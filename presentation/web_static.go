package presentation

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"

	"kostjc/domain"
	"kostjc/model/mAuth"
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty"
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

	fw.Get(`/`+domain.UserReportAction, func(ctx *fiber.Ctx) error {
		var in domain.UserReportIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.UserReportAction)
		if err != nil {
			return err
		}

		if notLogin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		in.MonthStart = time.Now().Format(rqProperty.DateFormatYYYYMM)
		in.MonthEnd = time.Now().AddDate(0, 3, 0).Format(rqProperty.DateFormatYYYYMM)
		out := d.UserReport(&in)

		return views.RenderUserReport(ctx, M.SX{
			`title`:              `KostJC | User Report`,
			`user`:               user,
			`segments`:           segments,
			`bookingsPerQuartal`: out.Bookings,
			`roomNames`:          out.RoomNames,
		})
	})

	fw.Get(`/`+domain.StaffBookingAction, func(ctx *fiber.Ctx) error {
		var in domain.StaffBookingIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.StaffBookingAction)
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
		out := d.StaffBooking(&in)

		return views.RenderStaffBooking(ctx, M.SX{
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

	fw.Get(`/`+domain.AdminUsersManagementAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminUsersManagementIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminUsersManagementAction)
		if err != nil {
			return err
		}

		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminUsersManagement(&in)

		return views.RenderAdminUsersManagement(ctx, M.SX{
			`title`:    `KostJC | Users Management`,
			`user`:     user,
			`segments`: segments,
			`users`:    out.Users,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
		})
	})

	fw.Get(`/`+domain.AdminLocationAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminLocationIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminLocationAction)
		if err != nil {
			return err
		}

		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminLocation(&in)

		return views.RenderAdminLocation(ctx, M.SX{
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

		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminFacility(&in)

		return views.RenderAdminFacility(ctx, M.SX{
			`title`:      `KostJC | Facility Management`,
			`user`:       user,
			`segments`:   segments,
			`facility`:   out.Facility,
			`facilities`: out.Facilities,
			`fields`:     out.Meta.Fields,
			`pager`:      out.Pager,
		})
	})

	fw.Get(`/`+domain.AdminRoomAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminRoomIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminRoomAction)
		if err != nil {
			return err
		}

		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		tenant := rqAuth.NewTenants(d.AuthOltp)
		tenants := tenant.FindTenantChoices()

		building := rqProperty.NewBuildings(d.PropOltp)
		buildings := building.FindBuildingChoices()

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminRoom(&in)

		return views.RenderAdminRoom(ctx, M.SX{
			`title`:     `KostJC | Room Management`,
			`user`:      user,
			`segments`:  segments,
			`room`:      out.Room,
			`rooms`:     out.Rooms,
			`tenants`:   tenants,
			`buildings`: buildings,
			`fields`:    out.Meta.Fields,
			`pager`:     out.Pager,
		})
	})

	fw.Get(`/`+domain.AdminBuildingAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminBuildingIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminBuildingAction)
		if err != nil {
			return err
		}

		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		loc := rqProperty.NewLocations(d.PropOltp)
		locations := loc.FindLocationChoices()

		fac := rqProperty.NewFacilities(d.PropOltp)
		facilities := fac.FindAllTypeBuilding()
		facilitiesChoices := fac.FindFacilitiesBuildingChoices()

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminBuilding(&in)

		return views.RenderAdminBuilding(ctx, M.SX{
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

		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminTenants(&in)

		return views.RenderAdminTenants(ctx, M.SX{
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

		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		bk := rqProperty.NewBookings(d.PropOltp)
		bookings := bk.FindBookingChoices()

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminPayment(&in)

		return views.RenderAdminPayment(ctx, M.SX{
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

		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		tenant := rqAuth.NewTenants(d.AuthOltp)
		tenants := tenant.FindTenantChoices()

		fac := rqProperty.NewFacilities(d.PropOltp)
		facilities := fac.FindAllTypeRoom()

		room := rqProperty.NewRooms(d.PropOltp)
		rooms := room.FindRoomChoices()

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		in.Pager.Order = []string{`+` + mProperty.DateStart}
		out := d.AdminBooking(&in)

		return views.RenderAdminBooking(ctx, M.SX{
			`title`:      `KostJC | Booking Management`,
			`user`:       user,
			`segments`:   segments,
			`booking`:    out.Booking,
			`bookings`:   out.Bookings,
			`tenants`:    tenants,
			`facilities`: facilities,
			`rooms`:      rooms,
			`fields`:     out.Meta.Fields,
			`pager`:      out.Pager,
		})
	})

	fw.Get(`/`+domain.AdminStockAction, func(ctx *fiber.Ctx) error {
		var in domain.AdminStockIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.AdminStockAction)
		if err != nil {
			return err
		}

		if notAdmin(ctx, d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.AdminStock(&in)

		return views.RenderAdminStock(ctx, M.SX{
			`title`:    `KostJC | Stock Management`,
			`user`:     user,
			`segments`: segments,
			`stock`:    out.Stock,
			`stocks`:   out.Stocks,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
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

	if sess.Role != mAuth.RoleAdmin {
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
