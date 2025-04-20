package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"

	"kostjc/domain"
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty/rqProperty"
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
		var in domain.UserLocationIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.UserLocationAction)
		if err != nil {
			return err
		}

		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.UserLocation(&in)

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

	fw.Get(`/`+domain.UserFacilityAction, func(ctx *fiber.Ctx) error {
		var in domain.UserFacilityIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.UserFacilityAction)
		if err != nil {
			return err
		}

		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.UserFacility(&in)

		return views.RenderFacility(ctx, M.SX{
			`title`:      `User Facility`,
			`user`:       user,
			`segments`:   segments,
			`facility`:   out.Facility,
			`facilities`: out.Facilities,
			`fields`:     out.Meta.Fields,
			`pager`:      out.Pager,
		})
	})

	fw.Get(`/`+domain.UserBuildingAction, func(ctx *fiber.Ctx) error {
		var in domain.UserBuildingIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.UserBuildingAction)
		if err != nil {
			return err
		}

		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		loc := rqProperty.NewLocations(d.PropOltp)
		locations := loc.FindLocationChoices()

		fac := rqProperty.NewFacilities(d.PropOltp)
		facilities := fac.FindAll()
		facilitiesChoices := fac.FindFacilitiesChoices()

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.UserBuilding(&in)

		return views.RenderBuilding(ctx, M.SX{
			`title`:             `User Building`,
			`user`:              user,
			`segments`:          segments,
			`building`:          out.Building,
			`buildings`:         out.Buildings,
			`locations`:         locations,
			`facilities`:        facilities,
			`facilitiesChoices`: facilitiesChoices,
			`fields`:            out.Meta.Fields,
			`pager`:             out.Pager,
		})
	})

	fw.Get(`/`+domain.UserTenantsAction, func(ctx *fiber.Ctx) error {
		var in domain.UserTenantsIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.UserTenantsAction)
		if err != nil {
			return err
		}

		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.UserTenants(&in)

		return views.RenderTenantsManagement(ctx, M.SX{
			`title`:    `User Tenants`,
			`user`:     user,
			`segments`: segments,
			`tenant`:   out.Tenant,
			`tenants`:  out.Tenants,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
		})
	})

	fw.Get(`/`+domain.UserBookingAction, func(ctx *fiber.Ctx) error {
		var in domain.UserBookingIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.UserBookingAction)
		if err != nil {
			return err
		}

		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		tenant := rqAuth.NewTenants(d.AuthOltp)
		tenants := tenant.FindTenantChoices()

		fac := rqProperty.NewFacilities(d.PropOltp)
		facilities := fac.FindAll()

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.UserBooking(&in)

		return views.RenderBooking(ctx, M.SX{
			`title`:      `User Booking`,
			`user`:       user,
			`segments`:   segments,
			`booking`:    out.Booking,
			`bookings`:   out.Bookings,
			`tenants`:    tenants,
			`facilities`: facilities,
			`fields`:     out.Meta.Fields,
			`pager`:      out.Pager,
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
