package presentation

import (
	"bytes"
	stdjson "encoding/json"
	"fmt"
	"kostjc/conf"
	"kostjc/domain"
	"reflect"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
	"github.com/kokizzu/gotro/Z"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type WebServer struct {
	*domain.Domain
}

var requiredHeader = M.SS{
	//domain.SomeUrl: `X-CC-Webhook-Signature`,
}

// priority:
// 1. query string
// 2. body
// 3. params
func webApiParseInput(ctx *fiber.Ctx, reqCommon *domain.RequestCommon, in any, url string) error {
	body := ctx.Body()
	reqCommon.Action = url
	reqCommon.Debug = reqCommon.Debug || conf.IsDebug()
	path := S.LeftOf(url, `?`) // without API_PREFIX
	if header, ok := requiredHeader[path]; ok {
		reqCommon.Header = ctx.Get(header)
		reqCommon.RawBody = string(body)
	}
	//L.Print(ctx.OriginalURL())
	if err := ctx.QueryParser(in); L.IsError(err, `ctx.QueryParser failed: `+url) {
		return err
	}
	if len(body) > 0 {
		retry := true
		trimmedBody := bytes.TrimSpace(body)
		if len(trimmedBody) > 0 && (trimmedBody[0] == '{' || strings.HasPrefix(ctx.Get(`content-type`), `application/json`)) {
			if err := unmarshalWebJSON(trimmedBody, in); err == nil {
				retry = false
			}
		}
		// application/x-www-form-urlencoded
		// multipart/form-data
		if retry {
			if err := ctx.BodyParser(in); L.IsError(err, `ctx.BodyParser failed: `+url) {
				fmt.Println(color.RedString(X.ToJsonPretty(in)))
				_ = ctx.JSON(M.SX{
					`error`: err.Error(),
				})
				return err
			}
		}
		trimBody := S.Left(string(body), 1024)
		if reqCommon.Debug && reqCommon.RawBody == `` {
			reqCommon.RawBody = trimBody
		}
	}
	if conf.IsDebug() && reqCommon.Debug && reqCommon.RawBody != `` {
		// prevent too large dump because multipart/form-data is raw binary
		contentType := ctx.Get(`content-type`)
		if !S.StartsWith(contentType, `multipart/form-data`) {
			//log.Print(ctx.GetReqHeaders())
			log.Print(reqCommon.RawBody)
		}
	}
	reqCommon.FromFiberCtx(ctx, ctx.UserContext())
	return nil
}

func unmarshalWebJSON(body []byte, in any) error {
	err := json.Unmarshal(body, in)
	if err == nil {
		return nil
	}

	normalized, changed, normalizeErr := normalizeJSONStringTaggedNumbers(body, reflect.TypeOf(in))
	if normalizeErr != nil || !changed {
		return err
	}
	if retryErr := json.Unmarshal(normalized, in); retryErr != nil {
		return retryErr
	}
	return nil
}

func normalizeJSONStringTaggedNumbers(body []byte, typ reflect.Type) ([]byte, bool, error) {
	var value any
	decoder := stdjson.NewDecoder(bytes.NewReader(body))
	decoder.UseNumber()
	if err := decoder.Decode(&value); err != nil {
		return nil, false, err
	}

	if !normalizeJSONValueForType(value, indirectType(typ)) {
		return nil, false, nil
	}

	normalized, err := stdjson.Marshal(value)
	return normalized, true, err
}

func normalizeJSONValueForType(value any, typ reflect.Type) bool {
	if typ == nil {
		return false
	}
	typ = indirectType(typ)

	switch typ.Kind() {
	case reflect.Struct:
		object, ok := value.(map[string]any)
		if !ok {
			return false
		}
		return normalizeJSONObjectForStruct(object, typ)
	case reflect.Slice, reflect.Array:
		items, ok := value.([]any)
		if !ok {
			return false
		}
		changed := false
		elemType := indirectType(typ.Elem())
		for _, item := range items {
			if normalizeJSONValueForType(item, elemType) {
				changed = true
			}
		}
		return changed
	case reflect.Map:
		object, ok := value.(map[string]any)
		if !ok {
			return false
		}
		valueType := indirectType(typ.Elem())
		if valueType.Kind() == reflect.Interface {
			return false
		}
		changed := false
		for _, item := range object {
			if normalizeJSONValueForType(item, valueType) {
				changed = true
			}
		}
		return changed
	default:
		return false
	}
}

func normalizeJSONObjectForStruct(object map[string]any, typ reflect.Type) bool {
	changed := false
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.PkgPath != `` && !field.Anonymous {
			continue
		}

		name, stringify, skip := jsonFieldInfo(field)
		if skip {
			continue
		}
		if field.Anonymous && name == `` {
			if normalizeJSONValueForType(object, field.Type) {
				changed = true
			}
			continue
		}
		if name == `` {
			name = field.Name
		}

		actualName := name
		fieldValue, ok := object[actualName]
		if !ok {
			for candidate := range object {
				if strings.EqualFold(candidate, name) {
					actualName = candidate
					fieldValue = object[candidate]
					ok = true
					break
				}
			}
		}
		if !ok {
			continue
		}
		if stringify {
			if number, ok := fieldValue.(stdjson.Number); ok {
				object[actualName] = number.String()
				changed = true
				continue
			}
		}
		if normalizeJSONValueForType(fieldValue, field.Type) {
			changed = true
		}
	}
	return changed
}

func jsonFieldInfo(field reflect.StructField) (name string, stringify bool, skip bool) {
	tag := field.Tag.Get(`json`)
	if tag == `-` {
		return ``, false, true
	}
	if tag == `` {
		return ``, false, false
	}
	parts := strings.Split(tag, `,`)
	name = parts[0]
	for _, option := range parts[1:] {
		if option == `string` {
			stringify = true
			break
		}
	}
	return name, stringify, false
}

func indirectType(typ reflect.Type) reflect.Type {
	for typ != nil && typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ
}

func (w *WebServer) Start(log *zerolog.Logger) {
	fw := fiber.New(fiber.Config{
		AppName:     conf.PROJECT_NAME,
		ProxyHeader: `X-Real-IP`,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// check if actionLogs are there, if error, then you need to run migration: go run main.go migrate
	w.InsertActionLog(&domain.RequestCommon{
		UserAgent: "server",
		IpAddress: "127.0.0.1",
		Action:    "server/start",
	}, &domain.ResponseCommon{})

	// load svelte templates
	views = &Views{}
	views.LoadAll()

	fw.Use(limiter.New(limiter.Config{
		Max:               300,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
	if conf.IsDebug() { // TODO: use faster logger for production
		fw.Use(logger.New())
	} else { // prevent panic on production
		copy := recover.ConfigDefault
		copy.EnableStackTrace = true
		fw.Use(recover.New(copy))
	}

	// assign static routes (GET)
	w.WebStatic(fw, w.Domain)

	// API routes (POST)
	ApiRoutes(fw, w.Domain)

	//zImport.GoogleSheetCountryDataToJson("1TmAjrclFHUwDA1487ifQjX4FzYt9y7eJ0gwyxtwZMJU")

	// finally serve static files from svelte directory if any
	fw.Static(`/`, `./svelte`, fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        false,
		CacheDuration: 5 * time.Second,
		MaxAge:        3600,
	})

	log.Err(fw.Listen(w.Domain.WebCfg.ListenAddr()))
}

type Views struct {
	cache map[string]*Z.TemplateChain
}

var views *Views

func (v *Views) LoadAll() {
	if v.cache == nil {
		v.cache = map[string]*Z.TemplateChain{}
	}
	debug := conf.IsDebug()
	const svelteDir = `svelte/`
	var err error
	for svelte, html := range viewList {
		v.cache[svelte], err = Z.ParseFile(debug, debug, svelteDir+html)
		L.PanicIf(err, `failed to parse `+html)
	}
}
