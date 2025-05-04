package mAuth

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

const (
	Id         = `id`
	CreatedAt  = `createdAt`
	CreatedBy  = `createdBy`
	UpdatedAt  = `updatedAt`
	UpdatedBy  = `updatedBy`
	DeletedAt  = `deletedAt`
	DeletedBy  = `deletedBy`
	RestoredBy = `restoredBy`
)

const (
	RoleUser  = `User`
	RoleAdmin = `Admin`
)

func IsValidRole(role string) bool {
	return role == RoleUser || role == RoleAdmin
}

const (
	TableUsers Tt.TableName = `users`

	Email         = `email`
	Password      = `password`
	PasswordSetAt = `passwordSetAt`
	SecretCode    = `secretCode`
	SecretCodeAt  = `secretCodeAt`
	VerifiedAt    = `verifiedAt`
	LastLoginAt   = `lastLoginAt`
	FullName      = `fullName`
	UserName      = `userName`
	Role          = `role`
)

const (
	TableSessions Tt.TableName = `sessions`

	SessionToken = `sessionToken`
	UserId       = `userId`
	ExpiredAt    = `expiredAt`
	Device       = `device`
	LoginAt      = `loginAt`
	LoginIPs     = `loginIPs`
)

const (
	TableTenants Tt.TableName = `tenants`

	TenantName       = `tenantName`
	KtpRegion        = `ktpRegion`
	KtpNumber        = `ktpNumber`
	KtpName          = `ktpName`
	KtpPlaceBirth    = `ktpPlaceBirth`
	KtpDateBirth     = `ktpDateBirth`
	KtpGender        = `ktpGender`
	KtpAddress       = `ktpAddress`
	KtpRtRw          = `ktpRtRw`
	KtpKelurahanDesa = `ktpKelurahanDesa`
	KtpKecamatan     = `ktpKecamatan`
	KtpReligion      = `ktpReligion`
	KtpMaritalStatus = `ktpMaritalStatus`
	KtpCitizenship   = `ktpCitizenship`
	KtpOccupation    = `ktpOccupation`
	TelegramUsername = `telegramUsername`
	WhatsappNumber   = `whatsappNumber`
)

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableUsers: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{Email, Tt.String},
			{Password, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{PasswordSetAt, Tt.Integer},
			{SecretCode, Tt.String},
			{SecretCodeAt, Tt.Integer},
			{VerifiedAt, Tt.Integer},
			{LastLoginAt, Tt.Integer},
			{FullName, Tt.String},
			{UserName, Tt.String},
			{Role, Tt.String},
		},
		AutoIncrementId:  true,
		Unique1:          Email,
		Unique2:          UserName, // after migration setting default usernames code done
		AutoCensorFields: []string{Password, SecretCode, SecretCodeAt},
		Engine:           Tt.Memtx,
	},
	TableSessions: {
		Fields: []Tt.Field{
			{SessionToken, Tt.String},
			{UserId, Tt.Unsigned},
			{ExpiredAt, Tt.Integer},
			{Device, Tt.String},
			{LoginAt, Tt.Integer},
			{LoginIPs, Tt.String},
		},
		Unique1: SessionToken,
		Engine:  Tt.Memtx,
	},
	TableTenants: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{TenantName, Tt.String},
			{KtpRegion, Tt.String},
			{KtpNumber, Tt.String},
			{KtpName, Tt.String},
			{KtpPlaceBirth, Tt.String},
			{KtpDateBirth, Tt.String},
			{KtpGender, Tt.String},
			{KtpAddress, Tt.String},
			{KtpRtRw, Tt.String},
			{KtpKelurahanDesa, Tt.String},
			{KtpKecamatan, Tt.String},
			{KtpReligion, Tt.String},
			{KtpMaritalStatus, Tt.String},
			{KtpCitizenship, Tt.String},
			{TelegramUsername, Tt.String},
			{WhatsappNumber, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
			{KtpOccupation, Tt.String},
		},
		AutoIncrementId: true,
		Engine:          Tt.Memtx,
		Unique1:         KtpNumber,
	},
}

const (
	TableActionLogs Ch.TableName = `actionLogs`

	RequestId  = `requestId`
	Error      = `error`
	ActorId    = `actorId`
	IpAddr4    = `ipAddr4`
	IpAddr6    = `ipAddr6`
	UserAgent  = `userAgent`
	Action     = `action`
	Traces     = `traces`
	StatusCode = `statusCode`
	Lat        = `lat`
	Long       = `long`

	Latency = `latency` // in seconds

	RefId = `refId`
)

var ClickhouseTables = map[Ch.TableName]*Ch.TableProp{
	TableActionLogs: {
		Fields: []Ch.Field{
			{CreatedAt, Ch.DateTime},
			{RequestId, Ch.String},
			{ActorId, Ch.UInt64},
			{Action, Ch.String},
			{StatusCode, Ch.Int16},
			{Traces, Ch.String},
			{Error, Ch.String},
			{IpAddr4, Ch.IPv4},
			{IpAddr6, Ch.IPv6},
			{UserAgent, Ch.String},
			{Lat, Ch.Float64},
			{Long, Ch.Float64},
			{Latency, Ch.Float64},
			{RefId, Ch.UInt64},
		},
		Orders: []string{CreatedAt, RequestId, ActorId, Action},
	},
}
