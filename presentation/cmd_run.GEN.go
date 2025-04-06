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

	}
}

// Code generated by 1_codegen_test.go DO NOT EDIT.
