package presentation

import (
	"os"

	"kostjc/domain"
)


// Code generated by 1_codegen_test.go DO NOT EDIT.


func cmdRun(b *domain.Domain, action string, payload []byte) {
	switch action {
	case domain.GuestLoginAction:
		in := domain.GuestLoginIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestLogin(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.GuestRegisterAction:
		in := domain.GuestRegisterIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestRegister(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.GuestResendVerificationEmailAction:
		in := domain.GuestResendVerificationEmailIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestResendVerificationEmail(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.GuestResetPasswordAction:
		in := domain.GuestResetPasswordIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestResetPassword(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.GuestVerifyEmailAction:
		in := domain.GuestVerifyEmailIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestVerifyEmail(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserBookingAction:
		in := domain.UserBookingIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserBooking(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserBuildingAction:
		in := domain.UserBuildingIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserBuilding(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserFacilityAction:
		in := domain.UserFacilityIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserFacility(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserLocationAction:
		in := domain.UserLocationIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserLocation(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserLogoutAction:
		in := domain.UserLogoutIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserLogout(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserProfileAction:
		in := domain.UserProfileIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserProfile(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserSessionKillAction:
		in := domain.UserSessionKillIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserSessionKill(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserTenantsAction:
		in := domain.UserTenantsIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserTenants(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	}
}

// Code generated by 1_codegen_test.go DO NOT EDIT.
