package presentation

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"

	"kostjc/domain"
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/zCrud"
)

func (w *WebServer) WebStatic(fw *fiber.App, d *domain.Domain) {

	fw.Get(`/`, func(ctx *fiber.Ctx) error {
		_, user, segments := userInfoFromContext(ctx, d)

		return views.RenderIndex(ctx, M.SX{
			`title`:    `KostJC`,
			`user`:     user,
			`segments`: segments,
		})
	})

	fw.Get(`/`+domain.UserLocationAction, func(ctx *fiber.Ctx) error {
		fmt.Println(`Here 1`)
		var in domain.UserLocationIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.UserLocationAction)
		fmt.Println(`Here 2`)
		if err != nil {
			fmt.Println(`Here 3`)
			return err
		}

		fmt.Println(`Here 4`)

		if notLogin(ctx, d, in.RequestCommon) {
			fmt.Println(`Here 5`)
			return ctx.Redirect(`/`, 302)
		}

		fmt.Println(`Here 6`)
		user, segments := userInfoFromRequest(in.RequestCommon, d)

		fmt.Println(`Here 7`)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.UserLocation(&in)

		fmt.Println(`Here 8`)
		return views.RenderLocation(ctx, M.SX{
			`title`:     `User Location`,
			`user`:      user,
			`segments`:  segments,
			`location`:  out.Location,
			`locations`: out.Locations,
			`fields`:    out.Meta.Fields,
			`pager`:     out.Pager,
		})
	})

	fw.Get(`/debug`, func(ctx *fiber.Ctx) error {
		return views.RenderDebug(ctx, M.SX{})
	})
}

func notLogin(ctx *fiber.Ctx, d *domain.Domain, in domain.RequestCommon) bool {
	var check domain.ResponseCommon
	sess := d.MustLogin(in, &check)
	if sess == nil {
		_ = views.RenderError(ctx, M.SX{
			`error`: check.Error,
		})
		return true
	}
	return false
}

func notAdmin(ctx *fiber.Ctx, d *domain.Domain, in domain.RequestCommon) bool {
	var check domain.ResponseCommon
	sess := d.MustAdmin(in, &check)
	if sess == nil {
		_ = views.RenderError(ctx, M.SX{
			`error`: check.Error,
		})
		return true
	}
	return false
}

func userInfoFromContext(c *fiber.Ctx, d *domain.Domain) (domain.UserProfileIn, *rqAuth.Users, M.SB) {
	var in domain.UserProfileIn
	err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserProfileAction)
	var user *rqAuth.Users
	segments := M.SB{}
	if err == nil {
		out := d.UserProfile(&in)
		user = out.User
		segments = out.Segments
	}
	return in, user, segments
}

func userInfoFromRequest(rc domain.RequestCommon, d *domain.Domain) (*rqAuth.Users, M.SB) {
	var user *rqAuth.Users
	var segments = M.SB{}
	out := d.UserProfile(&domain.UserProfileIn{
		RequestCommon: rc,
	})
	user = out.User
	segments = out.Segments
	return user, segments
}
