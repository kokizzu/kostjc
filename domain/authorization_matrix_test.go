package domain

import (
	"reflect"
	"strings"
	"testing"

	"kostjc/model/zCrud"

	"github.com/stretchr/testify/require"
)

func TestDomainAuthorizationMatrix(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	domainType := reflect.TypeOf(d)
	domainValue := reflect.ValueOf(d)

	for idx := 0; idx < domainType.NumMethod(); idx++ {
		method := domainType.Method(idx)
		name := method.Name
		if !strings.HasPrefix(name, "Admin") && !strings.HasPrefix(name, "Staff") {
			continue
		}
		if method.Type.NumIn() != 2 || method.Type.NumOut() != 1 {
			continue
		}

		t.Run(name, func(t *testing.T) {
			action := testActionForMethod(name)

			if strings.HasPrefix(name, "Admin") {
				adminOut := callDomainMethod(t, domainValue, method, testAdminRequestCommon(action))
				require.NotEqual(t, 403, adminOut.StatusCode, "admin should not be blocked by role")
				require.NotEqual(t, ErrRoleNotAllowed, adminOut.Error, "admin should not hit role guard")

				staffOut := callDomainMethod(t, domainValue, method, testStaffRequestCommon(action))
				require.Equal(t, 403, staffOut.StatusCode, "staff should be blocked from admin handler")
				require.Contains(t, []string{ErrRoleNotAllowed, ErrSegmentNotAllowed}, staffOut.Error)

				userOut := callDomainMethod(t, domainValue, method, testUserRequestCommon(action))
				require.Equal(t, 403, userOut.StatusCode, "user should be blocked from admin handler")
				require.Contains(t, []string{ErrRoleNotAllowed, ErrSegmentNotAllowed}, userOut.Error)
				return
			}

			adminOut := callDomainMethod(t, domainValue, method, testAdminRequestCommon(action))
			require.NotEqual(t, 403, adminOut.StatusCode, "admin should be allowed on staff handler")
			require.NotEqual(t, ErrRoleNotAllowed, adminOut.Error, "admin should not hit role guard")

			staffOut := callDomainMethod(t, domainValue, method, testStaffRequestCommon(action))
			require.NotEqual(t, 403, staffOut.StatusCode, "staff should be allowed on staff handler")
			require.NotEqual(t, ErrRoleNotAllowed, staffOut.Error, "staff should not hit role guard")

			userOut := callDomainMethod(t, domainValue, method, testUserRequestCommon(action))
			require.Equal(t, 403, userOut.StatusCode, "user should be blocked from staff handler")
			require.Equal(t, ErrRoleNotAllowed, userOut.Error)
		})
	}
}

func callDomainMethod(t *testing.T, domainValue reflect.Value, method reflect.Method, rc RequestCommon) ResponseCommon {
	t.Helper()

	inArg := buildDomainInput(method.Type.In(1), rc)

	var results []reflect.Value
	require.NotPanics(t, func() {
		results = domainValue.MethodByName(method.Name).Call([]reflect.Value{inArg})
	})
	require.Len(t, results, 1)

	out := results[0]
	respField := out.FieldByName("ResponseCommon")
	require.True(t, respField.IsValid(), "missing ResponseCommon on %s", method.Name)

	resp, ok := respField.Interface().(ResponseCommon)
	require.True(t, ok, "invalid ResponseCommon type on %s", method.Name)
	return resp
}

func buildDomainInput(inType reflect.Type, rc RequestCommon) reflect.Value {
	inPtr := reflect.New(inType.Elem())
	inVal := inPtr.Elem()

	if field := inVal.FieldByName("RequestCommon"); field.IsValid() && field.CanSet() {
		field.Set(reflect.ValueOf(rc))
	}
	if field := inVal.FieldByName("Cmd"); field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
		field.SetString(zCrud.CmdList)
	}
	if field := inVal.FieldByName("WithMeta"); field.IsValid() && field.CanSet() && field.Kind() == reflect.Bool {
		field.SetBool(true)
	}
	if field := inVal.FieldByName("Pager"); field.IsValid() && field.CanSet() {
		field.Set(reflect.ValueOf(zCrud.PagerIn{Page: 1, PerPage: 10}))
	}
	if field := inVal.FieldByName("YearMonth"); field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
		field.SetString(testTime.Format("200601"))
	}
	if field := inVal.FieldByName("MonthStart"); field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
		field.SetString(testTime.Format("200601"))
	}
	if field := inVal.FieldByName("MonthEnd"); field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
		field.SetString(testTime.AddDate(0, 3, 0).Format("200601"))
	}

	return inPtr
}

func testActionForMethod(name string) string {
	switch {
	case strings.HasPrefix(name, "Admin"):
		return "admin/" + lowerFirst(strings.TrimPrefix(name, "Admin"))
	case strings.HasPrefix(name, "Staff"):
		return "staff/" + lowerFirst(strings.TrimPrefix(name, "Staff"))
	default:
		return lowerFirst(name)
	}
}

func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = []rune(strings.ToLower(string(runes[0])))[0]
	return string(runes)
}
