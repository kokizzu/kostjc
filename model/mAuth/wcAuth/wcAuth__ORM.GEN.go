package wcAuth

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"kostjc/model/mAuth/rqAuth"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

// SessionsMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcAuth__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcAuth__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcAuth__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcAuth__ORM.GEN.go
type SessionsMutator struct {
	rqAuth.Sessions
	mutations []A.X
	logs      []A.X
}

// NewSessionsMutator create new ORM writer/command object
func NewSessionsMutator(adapter *Tt.Adapter) (res *SessionsMutator) {
	res = &SessionsMutator{Sessions: rqAuth.Sessions{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (s *SessionsMutator) Logs() []A.X { //nolint:dupl false positive
	return s.logs
}

// HaveMutation check whether Set* methods ever called
func (s *SessionsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(s.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (s *SessionsMutator) ClearMutations() { //nolint:dupl false positive
	s.mutations = []A.X{}
	s.logs = []A.X{}
}

// func (s *SessionsMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := s.ToArray()
//	_, err := s.Adapter.Upsert(s.SpaceName(), arr, A.X{
//		A.X{`=`, 0, s.SessionToken},
//		A.X{`=`, 1, s.UserId},
//		A.X{`=`, 2, s.ExpiredAt},
//		A.X{`=`, 3, s.Device},
//		A.X{`=`, 4, s.LoginAt},
//		A.X{`=`, 5, s.LoginIPs},
//	})
//	return !L.IsError(err, `Sessions.DoUpsert failed: `+s.SpaceName()+ `\n%#v`, arr)
// }

// DoOverwriteBySessionToken update all columns, error if not exists, not using mutations/Set*
func (s *SessionsMutator) DoOverwriteBySessionToken() bool { //nolint:dupl false positive
	_, err := s.Adapter.Update(s.SpaceName(), s.UniqueIndexSessionToken(), A.X{s.SessionToken}, s.ToUpdateArray())
	return !L.IsError(err, `Sessions.DoOverwriteBySessionToken failed: `+s.SpaceName())
}

// DoUpdateBySessionToken update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (s *SessionsMutator) DoUpdateBySessionToken() bool { //nolint:dupl false positive
	if !s.HaveMutation() {
		return true
	}
	_, err := s.Adapter.Update(s.SpaceName(), s.UniqueIndexSessionToken(), A.X{s.SessionToken}, s.mutations)
	return !L.IsError(err, `Sessions.DoUpdateBySessionToken failed: `+s.SpaceName())
}

// DoDeletePermanentBySessionToken permanent delete
func (s *SessionsMutator) DoDeletePermanentBySessionToken() bool { //nolint:dupl false positive
	_, err := s.Adapter.Delete(s.SpaceName(), s.UniqueIndexSessionToken(), A.X{s.SessionToken})
	return !L.IsError(err, `Sessions.DoDeletePermanentBySessionToken failed: `+s.SpaceName())
}

// DoInsert insert, error if already exists
func (s *SessionsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := s.ToArray()
	_, err := s.Adapter.Insert(s.SpaceName(), arr)
	return !L.IsError(err, `Sessions.DoInsert failed: `+s.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (s *SessionsMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := s.ToArray()
	_, err := s.Adapter.Replace(s.SpaceName(), arr)
	return !L.IsError(err, `Sessions.DoUpsert failed: `+s.SpaceName()+`\n%#v`, arr)
}

// SetSessionToken create mutations, should not duplicate
func (s *SessionsMutator) SetSessionToken(val string) bool { //nolint:dupl false positive
	if val != s.SessionToken {
		s.mutations = append(s.mutations, A.X{`=`, 0, val})
		s.logs = append(s.logs, A.X{`sessionToken`, s.SessionToken, val})
		s.SessionToken = val
		return true
	}
	return false
}

// SetUserId create mutations, should not duplicate
func (s *SessionsMutator) SetUserId(val uint64) bool { //nolint:dupl false positive
	if val != s.UserId {
		s.mutations = append(s.mutations, A.X{`=`, 1, val})
		s.logs = append(s.logs, A.X{`userId`, s.UserId, val})
		s.UserId = val
		return true
	}
	return false
}

// SetExpiredAt create mutations, should not duplicate
func (s *SessionsMutator) SetExpiredAt(val int64) bool { //nolint:dupl false positive
	if val != s.ExpiredAt {
		s.mutations = append(s.mutations, A.X{`=`, 2, val})
		s.logs = append(s.logs, A.X{`expiredAt`, s.ExpiredAt, val})
		s.ExpiredAt = val
		return true
	}
	return false
}

// SetDevice create mutations, should not duplicate
func (s *SessionsMutator) SetDevice(val string) bool { //nolint:dupl false positive
	if val != s.Device {
		s.mutations = append(s.mutations, A.X{`=`, 3, val})
		s.logs = append(s.logs, A.X{`device`, s.Device, val})
		s.Device = val
		return true
	}
	return false
}

// SetLoginAt create mutations, should not duplicate
func (s *SessionsMutator) SetLoginAt(val int64) bool { //nolint:dupl false positive
	if val != s.LoginAt {
		s.mutations = append(s.mutations, A.X{`=`, 4, val})
		s.logs = append(s.logs, A.X{`loginAt`, s.LoginAt, val})
		s.LoginAt = val
		return true
	}
	return false
}

// SetLoginIPs create mutations, should not duplicate
func (s *SessionsMutator) SetLoginIPs(val string) bool { //nolint:dupl false positive
	if val != s.LoginIPs {
		s.mutations = append(s.mutations, A.X{`=`, 5, val})
		s.logs = append(s.logs, A.X{`loginIPs`, s.LoginIPs, val})
		s.LoginIPs = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (s *SessionsMutator) SetAll(from rqAuth.Sessions, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`sessionToken`] && (forceMap[`sessionToken`] || from.SessionToken != ``) {
		s.SessionToken = S.Trim(from.SessionToken)
		changed = true
	}
	if !excludeMap[`userId`] && (forceMap[`userId`] || from.UserId != 0) {
		s.UserId = from.UserId
		changed = true
	}
	if !excludeMap[`expiredAt`] && (forceMap[`expiredAt`] || from.ExpiredAt != 0) {
		s.ExpiredAt = from.ExpiredAt
		changed = true
	}
	if !excludeMap[`device`] && (forceMap[`device`] || from.Device != ``) {
		s.Device = S.Trim(from.Device)
		changed = true
	}
	if !excludeMap[`loginAt`] && (forceMap[`loginAt`] || from.LoginAt != 0) {
		s.LoginAt = from.LoginAt
		changed = true
	}
	if !excludeMap[`loginIPs`] && (forceMap[`loginIPs`] || from.LoginIPs != ``) {
		s.LoginIPs = S.Trim(from.LoginIPs)
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// TenantsMutator DAO writer/command struct
type TenantsMutator struct {
	rqAuth.Tenants
	mutations []A.X
	logs      []A.X
}

// NewTenantsMutator create new ORM writer/command object
func NewTenantsMutator(adapter *Tt.Adapter) (res *TenantsMutator) {
	res = &TenantsMutator{Tenants: rqAuth.Tenants{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (t *TenantsMutator) Logs() []A.X { //nolint:dupl false positive
	return t.logs
}

// HaveMutation check whether Set* methods ever called
func (t *TenantsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(t.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (t *TenantsMutator) ClearMutations() { //nolint:dupl false positive
	t.mutations = []A.X{}
	t.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (t *TenantsMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id}, t.ToUpdateArray())
	return !L.IsError(err, `Tenants.DoOverwriteById failed: `+t.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (t *TenantsMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !t.HaveMutation() {
		return true
	}
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id}, t.mutations)
	return !L.IsError(err, `Tenants.DoUpdateById failed: `+t.SpaceName())
}

// DoDeletePermanentById permanent delete
func (t *TenantsMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := t.Adapter.Delete(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id})
	return !L.IsError(err, `Tenants.DoDeletePermanentById failed: `+t.SpaceName())
}

// func (t *TenantsMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := t.ToArray()
//	_, err := t.Adapter.Upsert(t.SpaceName(), arr, A.X{
//		A.X{`=`, 0, t.Id},
//		A.X{`=`, 1, t.TenantName},
//		A.X{`=`, 2, t.KtpRegion},
//		A.X{`=`, 3, t.KtpNumber},
//		A.X{`=`, 4, t.KtpName},
//		A.X{`=`, 5, t.KtpPlaceBirth},
//		A.X{`=`, 6, t.KtpDateBirth},
//		A.X{`=`, 7, t.KtpGender},
//		A.X{`=`, 8, t.KtpAddress},
//		A.X{`=`, 9, t.KtpRtRw},
//		A.X{`=`, 10, t.KtpKelurahanDesa},
//		A.X{`=`, 11, t.KtpKecamatan},
//		A.X{`=`, 12, t.KtpReligion},
//		A.X{`=`, 13, t.KtpMaritalStatus},
//		A.X{`=`, 14, t.KtpCitizenship},
//		A.X{`=`, 15, t.TelegramUsername},
//		A.X{`=`, 16, t.WhatsappNumber},
//		A.X{`=`, 17, t.CreatedAt},
//		A.X{`=`, 18, t.CreatedBy},
//		A.X{`=`, 19, t.UpdatedAt},
//		A.X{`=`, 20, t.UpdatedBy},
//		A.X{`=`, 21, t.DeletedAt},
//		A.X{`=`, 22, t.DeletedBy},
//		A.X{`=`, 23, t.RestoredBy},
//		A.X{`=`, 24, t.KtpOccupation},
//		A.X{`=`, 25, t.WaAddedAt},
//		A.X{`=`, 26, t.TeleAddedAt},
//	})
//	return !L.IsError(err, `Tenants.DoUpsert failed: `+t.SpaceName()+ `\n%#v`, arr)
// }

// DoOverwriteByKtpNumber update all columns, error if not exists, not using mutations/Set*
func (t *TenantsMutator) DoOverwriteByKtpNumber() bool { //nolint:dupl false positive
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexKtpNumber(), A.X{t.KtpNumber}, t.ToUpdateArray())
	return !L.IsError(err, `Tenants.DoOverwriteByKtpNumber failed: `+t.SpaceName())
}

// DoUpdateByKtpNumber update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (t *TenantsMutator) DoUpdateByKtpNumber() bool { //nolint:dupl false positive
	if !t.HaveMutation() {
		return true
	}
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexKtpNumber(), A.X{t.KtpNumber}, t.mutations)
	return !L.IsError(err, `Tenants.DoUpdateByKtpNumber failed: `+t.SpaceName())
}

// DoDeletePermanentByKtpNumber permanent delete
func (t *TenantsMutator) DoDeletePermanentByKtpNumber() bool { //nolint:dupl false positive
	_, err := t.Adapter.Delete(t.SpaceName(), t.UniqueIndexKtpNumber(), A.X{t.KtpNumber})
	return !L.IsError(err, `Tenants.DoDeletePermanentByKtpNumber failed: `+t.SpaceName())
}

// DoInsert insert, error if already exists
func (t *TenantsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := t.ToArray()
	row, err := t.Adapter.Insert(t.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			t.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Tenants.DoInsert failed: `+t.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (t *TenantsMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := t.ToArray()
	row, err := t.Adapter.Replace(t.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			t.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Tenants.DoUpsert failed: `+t.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (t *TenantsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != t.Id {
		t.mutations = append(t.mutations, A.X{`=`, 0, val})
		t.logs = append(t.logs, A.X{`id`, t.Id, val})
		t.Id = val
		return true
	}
	return false
}

// SetTenantName create mutations, should not duplicate
func (t *TenantsMutator) SetTenantName(val string) bool { //nolint:dupl false positive
	if val != t.TenantName {
		t.mutations = append(t.mutations, A.X{`=`, 1, val})
		t.logs = append(t.logs, A.X{`tenantName`, t.TenantName, val})
		t.TenantName = val
		return true
	}
	return false
}

// SetKtpRegion create mutations, should not duplicate
func (t *TenantsMutator) SetKtpRegion(val string) bool { //nolint:dupl false positive
	if val != t.KtpRegion {
		t.mutations = append(t.mutations, A.X{`=`, 2, val})
		t.logs = append(t.logs, A.X{`ktpRegion`, t.KtpRegion, val})
		t.KtpRegion = val
		return true
	}
	return false
}

// SetKtpNumber create mutations, should not duplicate
func (t *TenantsMutator) SetKtpNumber(val string) bool { //nolint:dupl false positive
	if val != t.KtpNumber {
		t.mutations = append(t.mutations, A.X{`=`, 3, val})
		t.logs = append(t.logs, A.X{`ktpNumber`, t.KtpNumber, val})
		t.KtpNumber = val
		return true
	}
	return false
}

// SetKtpName create mutations, should not duplicate
func (t *TenantsMutator) SetKtpName(val string) bool { //nolint:dupl false positive
	if val != t.KtpName {
		t.mutations = append(t.mutations, A.X{`=`, 4, val})
		t.logs = append(t.logs, A.X{`ktpName`, t.KtpName, val})
		t.KtpName = val
		return true
	}
	return false
}

// SetKtpPlaceBirth create mutations, should not duplicate
func (t *TenantsMutator) SetKtpPlaceBirth(val string) bool { //nolint:dupl false positive
	if val != t.KtpPlaceBirth {
		t.mutations = append(t.mutations, A.X{`=`, 5, val})
		t.logs = append(t.logs, A.X{`ktpPlaceBirth`, t.KtpPlaceBirth, val})
		t.KtpPlaceBirth = val
		return true
	}
	return false
}

// SetKtpDateBirth create mutations, should not duplicate
func (t *TenantsMutator) SetKtpDateBirth(val string) bool { //nolint:dupl false positive
	if val != t.KtpDateBirth {
		t.mutations = append(t.mutations, A.X{`=`, 6, val})
		t.logs = append(t.logs, A.X{`ktpDateBirth`, t.KtpDateBirth, val})
		t.KtpDateBirth = val
		return true
	}
	return false
}

// SetKtpGender create mutations, should not duplicate
func (t *TenantsMutator) SetKtpGender(val string) bool { //nolint:dupl false positive
	if val != t.KtpGender {
		t.mutations = append(t.mutations, A.X{`=`, 7, val})
		t.logs = append(t.logs, A.X{`ktpGender`, t.KtpGender, val})
		t.KtpGender = val
		return true
	}
	return false
}

// SetKtpAddress create mutations, should not duplicate
func (t *TenantsMutator) SetKtpAddress(val string) bool { //nolint:dupl false positive
	if val != t.KtpAddress {
		t.mutations = append(t.mutations, A.X{`=`, 8, val})
		t.logs = append(t.logs, A.X{`ktpAddress`, t.KtpAddress, val})
		t.KtpAddress = val
		return true
	}
	return false
}

// SetKtpRtRw create mutations, should not duplicate
func (t *TenantsMutator) SetKtpRtRw(val string) bool { //nolint:dupl false positive
	if val != t.KtpRtRw {
		t.mutations = append(t.mutations, A.X{`=`, 9, val})
		t.logs = append(t.logs, A.X{`ktpRtRw`, t.KtpRtRw, val})
		t.KtpRtRw = val
		return true
	}
	return false
}

// SetKtpKelurahanDesa create mutations, should not duplicate
func (t *TenantsMutator) SetKtpKelurahanDesa(val string) bool { //nolint:dupl false positive
	if val != t.KtpKelurahanDesa {
		t.mutations = append(t.mutations, A.X{`=`, 10, val})
		t.logs = append(t.logs, A.X{`ktpKelurahanDesa`, t.KtpKelurahanDesa, val})
		t.KtpKelurahanDesa = val
		return true
	}
	return false
}

// SetKtpKecamatan create mutations, should not duplicate
func (t *TenantsMutator) SetKtpKecamatan(val string) bool { //nolint:dupl false positive
	if val != t.KtpKecamatan {
		t.mutations = append(t.mutations, A.X{`=`, 11, val})
		t.logs = append(t.logs, A.X{`ktpKecamatan`, t.KtpKecamatan, val})
		t.KtpKecamatan = val
		return true
	}
	return false
}

// SetKtpReligion create mutations, should not duplicate
func (t *TenantsMutator) SetKtpReligion(val string) bool { //nolint:dupl false positive
	if val != t.KtpReligion {
		t.mutations = append(t.mutations, A.X{`=`, 12, val})
		t.logs = append(t.logs, A.X{`ktpReligion`, t.KtpReligion, val})
		t.KtpReligion = val
		return true
	}
	return false
}

// SetKtpMaritalStatus create mutations, should not duplicate
func (t *TenantsMutator) SetKtpMaritalStatus(val string) bool { //nolint:dupl false positive
	if val != t.KtpMaritalStatus {
		t.mutations = append(t.mutations, A.X{`=`, 13, val})
		t.logs = append(t.logs, A.X{`ktpMaritalStatus`, t.KtpMaritalStatus, val})
		t.KtpMaritalStatus = val
		return true
	}
	return false
}

// SetKtpCitizenship create mutations, should not duplicate
func (t *TenantsMutator) SetKtpCitizenship(val string) bool { //nolint:dupl false positive
	if val != t.KtpCitizenship {
		t.mutations = append(t.mutations, A.X{`=`, 14, val})
		t.logs = append(t.logs, A.X{`ktpCitizenship`, t.KtpCitizenship, val})
		t.KtpCitizenship = val
		return true
	}
	return false
}

// SetTelegramUsername create mutations, should not duplicate
func (t *TenantsMutator) SetTelegramUsername(val string) bool { //nolint:dupl false positive
	if val != t.TelegramUsername {
		t.mutations = append(t.mutations, A.X{`=`, 15, val})
		t.logs = append(t.logs, A.X{`telegramUsername`, t.TelegramUsername, val})
		t.TelegramUsername = val
		return true
	}
	return false
}

// SetWhatsappNumber create mutations, should not duplicate
func (t *TenantsMutator) SetWhatsappNumber(val string) bool { //nolint:dupl false positive
	if val != t.WhatsappNumber {
		t.mutations = append(t.mutations, A.X{`=`, 16, val})
		t.logs = append(t.logs, A.X{`whatsappNumber`, t.WhatsappNumber, val})
		t.WhatsappNumber = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (t *TenantsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != t.CreatedAt {
		t.mutations = append(t.mutations, A.X{`=`, 17, val})
		t.logs = append(t.logs, A.X{`createdAt`, t.CreatedAt, val})
		t.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (t *TenantsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.CreatedBy {
		t.mutations = append(t.mutations, A.X{`=`, 18, val})
		t.logs = append(t.logs, A.X{`createdBy`, t.CreatedBy, val})
		t.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (t *TenantsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != t.UpdatedAt {
		t.mutations = append(t.mutations, A.X{`=`, 19, val})
		t.logs = append(t.logs, A.X{`updatedAt`, t.UpdatedAt, val})
		t.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (t *TenantsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.UpdatedBy {
		t.mutations = append(t.mutations, A.X{`=`, 20, val})
		t.logs = append(t.logs, A.X{`updatedBy`, t.UpdatedBy, val})
		t.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (t *TenantsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != t.DeletedAt {
		t.mutations = append(t.mutations, A.X{`=`, 21, val})
		t.logs = append(t.logs, A.X{`deletedAt`, t.DeletedAt, val})
		t.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (t *TenantsMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.DeletedBy {
		t.mutations = append(t.mutations, A.X{`=`, 22, val})
		t.logs = append(t.logs, A.X{`deletedBy`, t.DeletedBy, val})
		t.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (t *TenantsMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != t.RestoredBy {
		t.mutations = append(t.mutations, A.X{`=`, 23, val})
		t.logs = append(t.logs, A.X{`restoredBy`, t.RestoredBy, val})
		t.RestoredBy = val
		return true
	}
	return false
}

// SetKtpOccupation create mutations, should not duplicate
func (t *TenantsMutator) SetKtpOccupation(val string) bool { //nolint:dupl false positive
	if val != t.KtpOccupation {
		t.mutations = append(t.mutations, A.X{`=`, 24, val})
		t.logs = append(t.logs, A.X{`ktpOccupation`, t.KtpOccupation, val})
		t.KtpOccupation = val
		return true
	}
	return false
}

// SetWaAddedAt create mutations, should not duplicate
func (t *TenantsMutator) SetWaAddedAt(val string) bool { //nolint:dupl false positive
	if val != t.WaAddedAt {
		t.mutations = append(t.mutations, A.X{`=`, 25, val})
		t.logs = append(t.logs, A.X{`waAddedAt`, t.WaAddedAt, val})
		t.WaAddedAt = val
		return true
	}
	return false
}

// SetTeleAddedAt create mutations, should not duplicate
func (t *TenantsMutator) SetTeleAddedAt(val string) bool { //nolint:dupl false positive
	if val != t.TeleAddedAt {
		t.mutations = append(t.mutations, A.X{`=`, 26, val})
		t.logs = append(t.logs, A.X{`teleAddedAt`, t.TeleAddedAt, val})
		t.TeleAddedAt = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (t *TenantsMutator) SetAll(from rqAuth.Tenants, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		t.Id = from.Id
		changed = true
	}
	if !excludeMap[`tenantName`] && (forceMap[`tenantName`] || from.TenantName != ``) {
		t.TenantName = S.Trim(from.TenantName)
		changed = true
	}
	if !excludeMap[`ktpRegion`] && (forceMap[`ktpRegion`] || from.KtpRegion != ``) {
		t.KtpRegion = S.Trim(from.KtpRegion)
		changed = true
	}
	if !excludeMap[`ktpNumber`] && (forceMap[`ktpNumber`] || from.KtpNumber != ``) {
		t.KtpNumber = S.Trim(from.KtpNumber)
		changed = true
	}
	if !excludeMap[`ktpName`] && (forceMap[`ktpName`] || from.KtpName != ``) {
		t.KtpName = S.Trim(from.KtpName)
		changed = true
	}
	if !excludeMap[`ktpPlaceBirth`] && (forceMap[`ktpPlaceBirth`] || from.KtpPlaceBirth != ``) {
		t.KtpPlaceBirth = S.Trim(from.KtpPlaceBirth)
		changed = true
	}
	if !excludeMap[`ktpDateBirth`] && (forceMap[`ktpDateBirth`] || from.KtpDateBirth != ``) {
		t.KtpDateBirth = S.Trim(from.KtpDateBirth)
		changed = true
	}
	if !excludeMap[`ktpGender`] && (forceMap[`ktpGender`] || from.KtpGender != ``) {
		t.KtpGender = S.Trim(from.KtpGender)
		changed = true
	}
	if !excludeMap[`ktpAddress`] && (forceMap[`ktpAddress`] || from.KtpAddress != ``) {
		t.KtpAddress = S.Trim(from.KtpAddress)
		changed = true
	}
	if !excludeMap[`ktpRtRw`] && (forceMap[`ktpRtRw`] || from.KtpRtRw != ``) {
		t.KtpRtRw = S.Trim(from.KtpRtRw)
		changed = true
	}
	if !excludeMap[`ktpKelurahanDesa`] && (forceMap[`ktpKelurahanDesa`] || from.KtpKelurahanDesa != ``) {
		t.KtpKelurahanDesa = S.Trim(from.KtpKelurahanDesa)
		changed = true
	}
	if !excludeMap[`ktpKecamatan`] && (forceMap[`ktpKecamatan`] || from.KtpKecamatan != ``) {
		t.KtpKecamatan = S.Trim(from.KtpKecamatan)
		changed = true
	}
	if !excludeMap[`ktpReligion`] && (forceMap[`ktpReligion`] || from.KtpReligion != ``) {
		t.KtpReligion = S.Trim(from.KtpReligion)
		changed = true
	}
	if !excludeMap[`ktpMaritalStatus`] && (forceMap[`ktpMaritalStatus`] || from.KtpMaritalStatus != ``) {
		t.KtpMaritalStatus = S.Trim(from.KtpMaritalStatus)
		changed = true
	}
	if !excludeMap[`ktpCitizenship`] && (forceMap[`ktpCitizenship`] || from.KtpCitizenship != ``) {
		t.KtpCitizenship = S.Trim(from.KtpCitizenship)
		changed = true
	}
	if !excludeMap[`telegramUsername`] && (forceMap[`telegramUsername`] || from.TelegramUsername != ``) {
		t.TelegramUsername = S.Trim(from.TelegramUsername)
		changed = true
	}
	if !excludeMap[`whatsappNumber`] && (forceMap[`whatsappNumber`] || from.WhatsappNumber != ``) {
		t.WhatsappNumber = S.Trim(from.WhatsappNumber)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		t.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		t.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		t.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		t.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		t.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		t.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		t.RestoredBy = from.RestoredBy
		changed = true
	}
	if !excludeMap[`ktpOccupation`] && (forceMap[`ktpOccupation`] || from.KtpOccupation != ``) {
		t.KtpOccupation = S.Trim(from.KtpOccupation)
		changed = true
	}
	if !excludeMap[`waAddedAt`] && (forceMap[`waAddedAt`] || from.WaAddedAt != ``) {
		t.WaAddedAt = S.Trim(from.WaAddedAt)
		changed = true
	}
	if !excludeMap[`teleAddedAt`] && (forceMap[`teleAddedAt`] || from.TeleAddedAt != ``) {
		t.TeleAddedAt = S.Trim(from.TeleAddedAt)
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// UsersMutator DAO writer/command struct
type UsersMutator struct {
	rqAuth.Users
	mutations []A.X
	logs      []A.X
}

// NewUsersMutator create new ORM writer/command object
func NewUsersMutator(adapter *Tt.Adapter) (res *UsersMutator) {
	res = &UsersMutator{Users: rqAuth.Users{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (u *UsersMutator) Logs() []A.X { //nolint:dupl false positive
	return u.logs
}

// HaveMutation check whether Set* methods ever called
func (u *UsersMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(u.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (u *UsersMutator) ClearMutations() { //nolint:dupl false positive
	u.mutations = []A.X{}
	u.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (u *UsersMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexId(), A.X{u.Id}, u.ToUpdateArray())
	return !L.IsError(err, `Users.DoOverwriteById failed: `+u.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (u *UsersMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !u.HaveMutation() {
		return true
	}
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexId(), A.X{u.Id}, u.mutations)
	return !L.IsError(err, `Users.DoUpdateById failed: `+u.SpaceName())
}

// DoDeletePermanentById permanent delete
func (u *UsersMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := u.Adapter.Delete(u.SpaceName(), u.UniqueIndexId(), A.X{u.Id})
	return !L.IsError(err, `Users.DoDeletePermanentById failed: `+u.SpaceName())
}

// func (u *UsersMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := u.ToArray()
//	_, err := u.Adapter.Upsert(u.SpaceName(), arr, A.X{
//		A.X{`=`, 0, u.Id},
//		A.X{`=`, 1, u.Email},
//		A.X{`=`, 2, u.Password},
//		A.X{`=`, 3, u.CreatedAt},
//		A.X{`=`, 4, u.CreatedBy},
//		A.X{`=`, 5, u.UpdatedAt},
//		A.X{`=`, 6, u.UpdatedBy},
//		A.X{`=`, 7, u.DeletedAt},
//		A.X{`=`, 8, u.PasswordSetAt},
//		A.X{`=`, 9, u.SecretCode},
//		A.X{`=`, 10, u.SecretCodeAt},
//		A.X{`=`, 11, u.VerifiedAt},
//		A.X{`=`, 12, u.LastLoginAt},
//		A.X{`=`, 13, u.FullName},
//		A.X{`=`, 14, u.UserName},
//		A.X{`=`, 15, u.Role},
//	})
//	return !L.IsError(err, `Users.DoUpsert failed: `+u.SpaceName()+ `\n%#v`, arr)
// }

// DoOverwriteByEmail update all columns, error if not exists, not using mutations/Set*
func (u *UsersMutator) DoOverwriteByEmail() bool { //nolint:dupl false positive
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexEmail(), A.X{u.Email}, u.ToUpdateArray())
	return !L.IsError(err, `Users.DoOverwriteByEmail failed: `+u.SpaceName())
}

// DoUpdateByEmail update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (u *UsersMutator) DoUpdateByEmail() bool { //nolint:dupl false positive
	if !u.HaveMutation() {
		return true
	}
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexEmail(), A.X{u.Email}, u.mutations)
	return !L.IsError(err, `Users.DoUpdateByEmail failed: `+u.SpaceName())
}

// DoDeletePermanentByEmail permanent delete
func (u *UsersMutator) DoDeletePermanentByEmail() bool { //nolint:dupl false positive
	_, err := u.Adapter.Delete(u.SpaceName(), u.UniqueIndexEmail(), A.X{u.Email})
	return !L.IsError(err, `Users.DoDeletePermanentByEmail failed: `+u.SpaceName())
}

// DoOverwriteByUserName update all columns, error if not exists, not using mutations/Set*
func (u *UsersMutator) DoOverwriteByUserName() bool { //nolint:dupl false positive
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexUserName(), A.X{u.UserName}, u.ToUpdateArray())
	return !L.IsError(err, `Users.DoOverwriteByUserName failed: `+u.SpaceName())
}

// DoUpdateByUserName update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (u *UsersMutator) DoUpdateByUserName() bool { //nolint:dupl false positive
	if !u.HaveMutation() {
		return true
	}
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexUserName(), A.X{u.UserName}, u.mutations)
	return !L.IsError(err, `Users.DoUpdateByUserName failed: `+u.SpaceName())
}

// DoDeletePermanentByUserName permanent delete
func (u *UsersMutator) DoDeletePermanentByUserName() bool { //nolint:dupl false positive
	_, err := u.Adapter.Delete(u.SpaceName(), u.UniqueIndexUserName(), A.X{u.UserName})
	return !L.IsError(err, `Users.DoDeletePermanentByUserName failed: `+u.SpaceName())
}

// DoInsert insert, error if already exists
func (u *UsersMutator) DoInsert() bool { //nolint:dupl false positive
	arr := u.ToArray()
	row, err := u.Adapter.Insert(u.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			u.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Users.DoInsert failed: `+u.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (u *UsersMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := u.ToArray()
	row, err := u.Adapter.Replace(u.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			u.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Users.DoUpsert failed: `+u.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (u *UsersMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != u.Id {
		u.mutations = append(u.mutations, A.X{`=`, 0, val})
		u.logs = append(u.logs, A.X{`id`, u.Id, val})
		u.Id = val
		return true
	}
	return false
}

// SetEmail create mutations, should not duplicate
func (u *UsersMutator) SetEmail(val string) bool { //nolint:dupl false positive
	if val != u.Email {
		u.mutations = append(u.mutations, A.X{`=`, 1, val})
		u.logs = append(u.logs, A.X{`email`, u.Email, val})
		u.Email = val
		return true
	}
	return false
}

// SetPassword create mutations, should not duplicate
func (u *UsersMutator) SetPassword(val string) bool { //nolint:dupl false positive
	if val != u.Password {
		u.mutations = append(u.mutations, A.X{`=`, 2, val})
		u.Password = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (u *UsersMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != u.CreatedAt {
		u.mutations = append(u.mutations, A.X{`=`, 3, val})
		u.logs = append(u.logs, A.X{`createdAt`, u.CreatedAt, val})
		u.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (u *UsersMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != u.CreatedBy {
		u.mutations = append(u.mutations, A.X{`=`, 4, val})
		u.logs = append(u.logs, A.X{`createdBy`, u.CreatedBy, val})
		u.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (u *UsersMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != u.UpdatedAt {
		u.mutations = append(u.mutations, A.X{`=`, 5, val})
		u.logs = append(u.logs, A.X{`updatedAt`, u.UpdatedAt, val})
		u.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (u *UsersMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != u.UpdatedBy {
		u.mutations = append(u.mutations, A.X{`=`, 6, val})
		u.logs = append(u.logs, A.X{`updatedBy`, u.UpdatedBy, val})
		u.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (u *UsersMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != u.DeletedAt {
		u.mutations = append(u.mutations, A.X{`=`, 7, val})
		u.logs = append(u.logs, A.X{`deletedAt`, u.DeletedAt, val})
		u.DeletedAt = val
		return true
	}
	return false
}

// SetPasswordSetAt create mutations, should not duplicate
func (u *UsersMutator) SetPasswordSetAt(val int64) bool { //nolint:dupl false positive
	if val != u.PasswordSetAt {
		u.mutations = append(u.mutations, A.X{`=`, 8, val})
		u.logs = append(u.logs, A.X{`passwordSetAt`, u.PasswordSetAt, val})
		u.PasswordSetAt = val
		return true
	}
	return false
}

// SetSecretCode create mutations, should not duplicate
func (u *UsersMutator) SetSecretCode(val string) bool { //nolint:dupl false positive
	if val != u.SecretCode {
		u.mutations = append(u.mutations, A.X{`=`, 9, val})
		u.SecretCode = val
		return true
	}
	return false
}

// SetSecretCodeAt create mutations, should not duplicate
func (u *UsersMutator) SetSecretCodeAt(val int64) bool { //nolint:dupl false positive
	if val != u.SecretCodeAt {
		u.mutations = append(u.mutations, A.X{`=`, 10, val})
		u.SecretCodeAt = val
		return true
	}
	return false
}

// SetVerifiedAt create mutations, should not duplicate
func (u *UsersMutator) SetVerifiedAt(val int64) bool { //nolint:dupl false positive
	if val != u.VerifiedAt {
		u.mutations = append(u.mutations, A.X{`=`, 11, val})
		u.logs = append(u.logs, A.X{`verifiedAt`, u.VerifiedAt, val})
		u.VerifiedAt = val
		return true
	}
	return false
}

// SetLastLoginAt create mutations, should not duplicate
func (u *UsersMutator) SetLastLoginAt(val int64) bool { //nolint:dupl false positive
	if val != u.LastLoginAt {
		u.mutations = append(u.mutations, A.X{`=`, 12, val})
		u.logs = append(u.logs, A.X{`lastLoginAt`, u.LastLoginAt, val})
		u.LastLoginAt = val
		return true
	}
	return false
}

// SetFullName create mutations, should not duplicate
func (u *UsersMutator) SetFullName(val string) bool { //nolint:dupl false positive
	if val != u.FullName {
		u.mutations = append(u.mutations, A.X{`=`, 13, val})
		u.logs = append(u.logs, A.X{`fullName`, u.FullName, val})
		u.FullName = val
		return true
	}
	return false
}

// SetUserName create mutations, should not duplicate
func (u *UsersMutator) SetUserName(val string) bool { //nolint:dupl false positive
	if val != u.UserName {
		u.mutations = append(u.mutations, A.X{`=`, 14, val})
		u.logs = append(u.logs, A.X{`userName`, u.UserName, val})
		u.UserName = val
		return true
	}
	return false
}

// SetRole create mutations, should not duplicate
func (u *UsersMutator) SetRole(val string) bool { //nolint:dupl false positive
	if val != u.Role {
		u.mutations = append(u.mutations, A.X{`=`, 15, val})
		u.logs = append(u.logs, A.X{`role`, u.Role, val})
		u.Role = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (u *UsersMutator) SetAll(from rqAuth.Users, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		u.Id = from.Id
		changed = true
	}
	if !excludeMap[`email`] && (forceMap[`email`] || from.Email != ``) {
		u.Email = S.Trim(from.Email)
		changed = true
	}
	if !excludeMap[`password`] && (forceMap[`password`] || from.Password != ``) {
		u.Password = S.Trim(from.Password)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		u.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		u.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		u.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		u.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		u.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`passwordSetAt`] && (forceMap[`passwordSetAt`] || from.PasswordSetAt != 0) {
		u.PasswordSetAt = from.PasswordSetAt
		changed = true
	}
	if !excludeMap[`secretCode`] && (forceMap[`secretCode`] || from.SecretCode != ``) {
		u.SecretCode = S.Trim(from.SecretCode)
		changed = true
	}
	if !excludeMap[`secretCodeAt`] && (forceMap[`secretCodeAt`] || from.SecretCodeAt != 0) {
		u.SecretCodeAt = from.SecretCodeAt
		changed = true
	}
	if !excludeMap[`verifiedAt`] && (forceMap[`verifiedAt`] || from.VerifiedAt != 0) {
		u.VerifiedAt = from.VerifiedAt
		changed = true
	}
	if !excludeMap[`lastLoginAt`] && (forceMap[`lastLoginAt`] || from.LastLoginAt != 0) {
		u.LastLoginAt = from.LastLoginAt
		changed = true
	}
	if !excludeMap[`fullName`] && (forceMap[`fullName`] || from.FullName != ``) {
		u.FullName = S.Trim(from.FullName)
		changed = true
	}
	if !excludeMap[`userName`] && (forceMap[`userName`] || from.UserName != ``) {
		u.UserName = S.Trim(from.UserName)
		changed = true
	}
	if !excludeMap[`role`] && (forceMap[`role`] || from.Role != ``) {
		u.Role = S.Trim(from.Role)
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
