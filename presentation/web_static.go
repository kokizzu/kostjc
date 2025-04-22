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
			`title`:    `KostJC | Home`,
			`user`:     user,
			`segments`: segments,
		})
	})

	fw.Get(`/user`, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		in.RequestCommon.Action = domain.UserSessionsActiveAction
		out := d.UserSessionsActive(&domain.UserSessionsActiveIn{
			RequestCommon: in.RequestCommon,
		})

		return views.RenderUser(ctx, M.SX{
			`title`:          `KostJC | User Profile`,
			`user`:           user,
			`segments`:       segments,
			`activeSessions`: out.SessionsActive,
		})
	})

	fw.Get(`/`+domain.AdminLocationAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminLocationIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminLocationAction)
		if err != nil {
			return err
		}

		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminLocation(&in)

		return views.RenderLocation(ctx, M.SX{
			`title`:     `KostJC | Location Management`,
			`user`:      user,
			`segments`:  segments,
			`location`:  out.Location,
			`locations`: out.Locations,
			`fields`:    out.Meta.Fields,
			`pager`:     out.Pager,
		})
	})

	fw.Get(`/`+domain.AdminFacilityAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminFacilityIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminFacilityAction)
		if err != nil {
			return err
		}

		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminFacility(&in)

		return views.RenderFacility(ctx, M.SX{
			`title`:      `KostJC | Facility Management`,
			`user`:       user,
			`segments`:   segments,
			`facility`:   out.Facility,
			`facilities`: out.Facilities,
			`fields`:     out.Meta.Fields,
			`pager`:      out.Pager,
		})
	})

	fw.Get(`/`+domain.AdminBuildingAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminBuildingIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminBuildingAction)
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
		out := d.AdminBuilding(&in)

		return views.RenderBuilding(ctx, M.SX{
			`title`:             `KostJC | Building Management`,
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

	fw.Get(`/`+domain.AdminTenantsAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminTenantsIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminTenantsAction)
		if err != nil {
			return err
		}

		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminTenants(&in)

		return views.RenderTenants(ctx, M.SX{
			`title`:    `KostJC | Tenant Management`,
			`user`:     user,
			`segments`: segments,
			`tenant`:   out.Tenant,
			`tenants`:  out.Tenants,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
		})
	})

	fw.Get(`/`+domain.AdminPaymentAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminPaymentIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminPaymentAction)
		if err != nil {
			return err
		}

		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		bk := rqProperty.NewBookings(d.PropOltp)
		bookings := bk.FindBookingChoices()

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminPayment(&in)

		return views.RenderPayment(ctx, M.SX{
			`title`:    `KostJC | Payment Management`,
			`user`:     user,
			`segments`: segments,
			`payment`:  out.Payment,
			`payments`: out.Payments,
			`bookings`: bookings,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
		})
	})

	fw.Get(`/`+domain.AdminBookingAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminBookingIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminBookingAction)
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
		out := d.AdminBooking(&in)

		return views.RenderBooking(ctx, M.SX{
			`title`:      `KostJC | Booking Management`,
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
