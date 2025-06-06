package wcCafe

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"kostjc/model/mCafe/rqCafe"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

// MenusMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcCafe__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcCafe__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcCafe__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcCafe__ORM.GEN.go
type MenusMutator struct {
	rqCafe.Menus
	mutations []A.X
	logs      []A.X
}

// NewMenusMutator create new ORM writer/command object
func NewMenusMutator(adapter *Tt.Adapter) (res *MenusMutator) {
	res = &MenusMutator{Menus: rqCafe.Menus{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (m *MenusMutator) Logs() []A.X { //nolint:dupl false positive
	return m.logs
}

// HaveMutation check whether Set* methods ever called
func (m *MenusMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(m.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (m *MenusMutator) ClearMutations() { //nolint:dupl false positive
	m.mutations = []A.X{}
	m.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (m *MenusMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := m.Adapter.Update(m.SpaceName(), m.UniqueIndexId(), A.X{m.Id}, m.ToUpdateArray())
	return !L.IsError(err, `Menus.DoOverwriteById failed: `+m.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (m *MenusMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !m.HaveMutation() {
		return true
	}
	_, err := m.Adapter.Update(m.SpaceName(), m.UniqueIndexId(), A.X{m.Id}, m.mutations)
	return !L.IsError(err, `Menus.DoUpdateById failed: `+m.SpaceName())
}

// DoDeletePermanentById permanent delete
func (m *MenusMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := m.Adapter.Delete(m.SpaceName(), m.UniqueIndexId(), A.X{m.Id})
	return !L.IsError(err, `Menus.DoDeletePermanentById failed: `+m.SpaceName())
}

// func (m *MenusMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := m.ToArray()
//	_, err := m.Adapter.Upsert(m.SpaceName(), arr, A.X{
//		A.X{`=`, 0, m.Id},
//		A.X{`=`, 1, m.Name},
//		A.X{`=`, 2, m.HppIDR},
//		A.X{`=`, 3, m.SalePriceIDR},
//		A.X{`=`, 4, m.Detail},
//		A.X{`=`, 5, m.ImageUrl},
//		A.X{`=`, 6, m.CreatedAt},
//		A.X{`=`, 7, m.CreatedBy},
//		A.X{`=`, 8, m.UpdatedAt},
//		A.X{`=`, 9, m.UpdatedBy},
//		A.X{`=`, 10, m.DeletedAt},
//		A.X{`=`, 11, m.DeletedBy},
//		A.X{`=`, 12, m.RestoredBy},
//	})
//	return !L.IsError(err, `Menus.DoUpsert failed: `+m.SpaceName()+ `\n%#v`, arr)
// }

// DoOverwriteByName update all columns, error if not exists, not using mutations/Set*
func (m *MenusMutator) DoOverwriteByName() bool { //nolint:dupl false positive
	_, err := m.Adapter.Update(m.SpaceName(), m.UniqueIndexName(), A.X{m.Name}, m.ToUpdateArray())
	return !L.IsError(err, `Menus.DoOverwriteByName failed: `+m.SpaceName())
}

// DoUpdateByName update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (m *MenusMutator) DoUpdateByName() bool { //nolint:dupl false positive
	if !m.HaveMutation() {
		return true
	}
	_, err := m.Adapter.Update(m.SpaceName(), m.UniqueIndexName(), A.X{m.Name}, m.mutations)
	return !L.IsError(err, `Menus.DoUpdateByName failed: `+m.SpaceName())
}

// DoDeletePermanentByName permanent delete
func (m *MenusMutator) DoDeletePermanentByName() bool { //nolint:dupl false positive
	_, err := m.Adapter.Delete(m.SpaceName(), m.UniqueIndexName(), A.X{m.Name})
	return !L.IsError(err, `Menus.DoDeletePermanentByName failed: `+m.SpaceName())
}

// DoInsert insert, error if already exists
func (m *MenusMutator) DoInsert() bool { //nolint:dupl false positive
	arr := m.ToArray()
	row, err := m.Adapter.Insert(m.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			m.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Menus.DoInsert failed: `+m.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (m *MenusMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := m.ToArray()
	row, err := m.Adapter.Replace(m.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			m.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Menus.DoUpsert failed: `+m.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (m *MenusMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != m.Id {
		m.mutations = append(m.mutations, A.X{`=`, 0, val})
		m.logs = append(m.logs, A.X{`id`, m.Id, val})
		m.Id = val
		return true
	}
	return false
}

// SetName create mutations, should not duplicate
func (m *MenusMutator) SetName(val string) bool { //nolint:dupl false positive
	if val != m.Name {
		m.mutations = append(m.mutations, A.X{`=`, 1, val})
		m.logs = append(m.logs, A.X{`name`, m.Name, val})
		m.Name = val
		return true
	}
	return false
}

// SetHppIDR create mutations, should not duplicate
func (m *MenusMutator) SetHppIDR(val int64) bool { //nolint:dupl false positive
	if val != m.HppIDR {
		m.mutations = append(m.mutations, A.X{`=`, 2, val})
		m.logs = append(m.logs, A.X{`hppIDR`, m.HppIDR, val})
		m.HppIDR = val
		return true
	}
	return false
}

// SetSalePriceIDR create mutations, should not duplicate
func (m *MenusMutator) SetSalePriceIDR(val int64) bool { //nolint:dupl false positive
	if val != m.SalePriceIDR {
		m.mutations = append(m.mutations, A.X{`=`, 3, val})
		m.logs = append(m.logs, A.X{`salePriceIDR`, m.SalePriceIDR, val})
		m.SalePriceIDR = val
		return true
	}
	return false
}

// SetDetail create mutations, should not duplicate
func (m *MenusMutator) SetDetail(val string) bool { //nolint:dupl false positive
	if val != m.Detail {
		m.mutations = append(m.mutations, A.X{`=`, 4, val})
		m.logs = append(m.logs, A.X{`detail`, m.Detail, val})
		m.Detail = val
		return true
	}
	return false
}

// SetImageUrl create mutations, should not duplicate
func (m *MenusMutator) SetImageUrl(val string) bool { //nolint:dupl false positive
	if val != m.ImageUrl {
		m.mutations = append(m.mutations, A.X{`=`, 5, val})
		m.logs = append(m.logs, A.X{`imageUrl`, m.ImageUrl, val})
		m.ImageUrl = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (m *MenusMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != m.CreatedAt {
		m.mutations = append(m.mutations, A.X{`=`, 6, val})
		m.logs = append(m.logs, A.X{`createdAt`, m.CreatedAt, val})
		m.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (m *MenusMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != m.CreatedBy {
		m.mutations = append(m.mutations, A.X{`=`, 7, val})
		m.logs = append(m.logs, A.X{`createdBy`, m.CreatedBy, val})
		m.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (m *MenusMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != m.UpdatedAt {
		m.mutations = append(m.mutations, A.X{`=`, 8, val})
		m.logs = append(m.logs, A.X{`updatedAt`, m.UpdatedAt, val})
		m.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (m *MenusMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != m.UpdatedBy {
		m.mutations = append(m.mutations, A.X{`=`, 9, val})
		m.logs = append(m.logs, A.X{`updatedBy`, m.UpdatedBy, val})
		m.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (m *MenusMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != m.DeletedAt {
		m.mutations = append(m.mutations, A.X{`=`, 10, val})
		m.logs = append(m.logs, A.X{`deletedAt`, m.DeletedAt, val})
		m.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (m *MenusMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != m.DeletedBy {
		m.mutations = append(m.mutations, A.X{`=`, 11, val})
		m.logs = append(m.logs, A.X{`deletedBy`, m.DeletedBy, val})
		m.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (m *MenusMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != m.RestoredBy {
		m.mutations = append(m.mutations, A.X{`=`, 12, val})
		m.logs = append(m.logs, A.X{`restoredBy`, m.RestoredBy, val})
		m.RestoredBy = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (m *MenusMutator) SetAll(from rqCafe.Menus, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		m.Id = from.Id
		changed = true
	}
	if !excludeMap[`name`] && (forceMap[`name`] || from.Name != ``) {
		m.Name = S.Trim(from.Name)
		changed = true
	}
	if !excludeMap[`hppIDR`] && (forceMap[`hppIDR`] || from.HppIDR != 0) {
		m.HppIDR = from.HppIDR
		changed = true
	}
	if !excludeMap[`salePriceIDR`] && (forceMap[`salePriceIDR`] || from.SalePriceIDR != 0) {
		m.SalePriceIDR = from.SalePriceIDR
		changed = true
	}
	if !excludeMap[`detail`] && (forceMap[`detail`] || from.Detail != ``) {
		m.Detail = S.Trim(from.Detail)
		changed = true
	}
	if !excludeMap[`imageUrl`] && (forceMap[`imageUrl`] || from.ImageUrl != ``) {
		m.ImageUrl = S.Trim(from.ImageUrl)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		m.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		m.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		m.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		m.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		m.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		m.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		m.RestoredBy = from.RestoredBy
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// SalesMutator DAO writer/command struct
type SalesMutator struct {
	rqCafe.Sales
	mutations []A.X
	logs      []A.X
}

// NewSalesMutator create new ORM writer/command object
func NewSalesMutator(adapter *Tt.Adapter) (res *SalesMutator) {
	res = &SalesMutator{Sales: rqCafe.Sales{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (s *SalesMutator) Logs() []A.X { //nolint:dupl false positive
	return s.logs
}

// HaveMutation check whether Set* methods ever called
func (s *SalesMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(s.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (s *SalesMutator) ClearMutations() { //nolint:dupl false positive
	s.mutations = []A.X{}
	s.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (s *SalesMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := s.Adapter.Update(s.SpaceName(), s.UniqueIndexId(), A.X{s.Id}, s.ToUpdateArray())
	return !L.IsError(err, `Sales.DoOverwriteById failed: `+s.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (s *SalesMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !s.HaveMutation() {
		return true
	}
	_, err := s.Adapter.Update(s.SpaceName(), s.UniqueIndexId(), A.X{s.Id}, s.mutations)
	return !L.IsError(err, `Sales.DoUpdateById failed: `+s.SpaceName())
}

// DoDeletePermanentById permanent delete
func (s *SalesMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := s.Adapter.Delete(s.SpaceName(), s.UniqueIndexId(), A.X{s.Id})
	return !L.IsError(err, `Sales.DoDeletePermanentById failed: `+s.SpaceName())
}

// func (s *SalesMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := s.ToArray()
//	_, err := s.Adapter.Upsert(s.SpaceName(), arr, A.X{
//		A.X{`=`, 0, s.Id},
//		A.X{`=`, 1, s.Cashier},
//		A.X{`=`, 2, s.TenantId},
//		A.X{`=`, 3, s.BuyerName},
//		A.X{`=`, 4, s.SalesDate},
//		A.X{`=`, 5, s.PaidAt},
//		A.X{`=`, 6, s.Note},
//		A.X{`=`, 7, s.UtensilsTaken},
//		A.X{`=`, 8, s.UtensilsReturnedAt},
//		A.X{`=`, 9, s.QrisIDR},
//		A.X{`=`, 10, s.CashIDR},
//		A.X{`=`, 11, s.DebtIDR},
//		A.X{`=`, 12, s.TopupIDR},
//		A.X{`=`, 13, s.TotalPriceIDR},
//		A.X{`=`, 14, s.CreatedAt},
//		A.X{`=`, 15, s.CreatedBy},
//		A.X{`=`, 16, s.UpdatedAt},
//		A.X{`=`, 17, s.UpdatedBy},
//		A.X{`=`, 18, s.DeletedAt},
//		A.X{`=`, 19, s.DeletedBy},
//		A.X{`=`, 20, s.RestoredBy},
//	})
//	return !L.IsError(err, `Sales.DoUpsert failed: `+s.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (s *SalesMutator) DoInsert() bool { //nolint:dupl false positive
	arr := s.ToArray()
	row, err := s.Adapter.Insert(s.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			s.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Sales.DoInsert failed: `+s.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (s *SalesMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := s.ToArray()
	row, err := s.Adapter.Replace(s.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			s.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Sales.DoUpsert failed: `+s.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (s *SalesMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != s.Id {
		s.mutations = append(s.mutations, A.X{`=`, 0, val})
		s.logs = append(s.logs, A.X{`id`, s.Id, val})
		s.Id = val
		return true
	}
	return false
}

// SetCashier create mutations, should not duplicate
func (s *SalesMutator) SetCashier(val string) bool { //nolint:dupl false positive
	if val != s.Cashier {
		s.mutations = append(s.mutations, A.X{`=`, 1, val})
		s.logs = append(s.logs, A.X{`cashier`, s.Cashier, val})
		s.Cashier = val
		return true
	}
	return false
}

// SetTenantId create mutations, should not duplicate
func (s *SalesMutator) SetTenantId(val uint64) bool { //nolint:dupl false positive
	if val != s.TenantId {
		s.mutations = append(s.mutations, A.X{`=`, 2, val})
		s.logs = append(s.logs, A.X{`tenantId`, s.TenantId, val})
		s.TenantId = val
		return true
	}
	return false
}

// SetBuyerName create mutations, should not duplicate
func (s *SalesMutator) SetBuyerName(val string) bool { //nolint:dupl false positive
	if val != s.BuyerName {
		s.mutations = append(s.mutations, A.X{`=`, 3, val})
		s.logs = append(s.logs, A.X{`buyerName`, s.BuyerName, val})
		s.BuyerName = val
		return true
	}
	return false
}

// SetSalesDate create mutations, should not duplicate
func (s *SalesMutator) SetSalesDate(val string) bool { //nolint:dupl false positive
	if val != s.SalesDate {
		s.mutations = append(s.mutations, A.X{`=`, 4, val})
		s.logs = append(s.logs, A.X{`salesDate`, s.SalesDate, val})
		s.SalesDate = val
		return true
	}
	return false
}

// SetPaidAt create mutations, should not duplicate
func (s *SalesMutator) SetPaidAt(val string) bool { //nolint:dupl false positive
	if val != s.PaidAt {
		s.mutations = append(s.mutations, A.X{`=`, 5, val})
		s.logs = append(s.logs, A.X{`paidAt`, s.PaidAt, val})
		s.PaidAt = val
		return true
	}
	return false
}

// SetNote create mutations, should not duplicate
func (s *SalesMutator) SetNote(val string) bool { //nolint:dupl false positive
	if val != s.Note {
		s.mutations = append(s.mutations, A.X{`=`, 6, val})
		s.logs = append(s.logs, A.X{`note`, s.Note, val})
		s.Note = val
		return true
	}
	return false
}

// SetUtensilsTaken create mutations, should not duplicate
func (s *SalesMutator) SetUtensilsTaken(val string) bool { //nolint:dupl false positive
	if val != s.UtensilsTaken {
		s.mutations = append(s.mutations, A.X{`=`, 7, val})
		s.logs = append(s.logs, A.X{`utensilsTaken`, s.UtensilsTaken, val})
		s.UtensilsTaken = val
		return true
	}
	return false
}

// SetUtensilsReturnedAt create mutations, should not duplicate
func (s *SalesMutator) SetUtensilsReturnedAt(val string) bool { //nolint:dupl false positive
	if val != s.UtensilsReturnedAt {
		s.mutations = append(s.mutations, A.X{`=`, 8, val})
		s.logs = append(s.logs, A.X{`utensilsReturnedAt`, s.UtensilsReturnedAt, val})
		s.UtensilsReturnedAt = val
		return true
	}
	return false
}

// SetQrisIDR create mutations, should not duplicate
func (s *SalesMutator) SetQrisIDR(val int64) bool { //nolint:dupl false positive
	if val != s.QrisIDR {
		s.mutations = append(s.mutations, A.X{`=`, 9, val})
		s.logs = append(s.logs, A.X{`qrisIDR`, s.QrisIDR, val})
		s.QrisIDR = val
		return true
	}
	return false
}

// SetCashIDR create mutations, should not duplicate
func (s *SalesMutator) SetCashIDR(val int64) bool { //nolint:dupl false positive
	if val != s.CashIDR {
		s.mutations = append(s.mutations, A.X{`=`, 10, val})
		s.logs = append(s.logs, A.X{`cashIDR`, s.CashIDR, val})
		s.CashIDR = val
		return true
	}
	return false
}

// SetDebtIDR create mutations, should not duplicate
func (s *SalesMutator) SetDebtIDR(val int64) bool { //nolint:dupl false positive
	if val != s.DebtIDR {
		s.mutations = append(s.mutations, A.X{`=`, 11, val})
		s.logs = append(s.logs, A.X{`debtIDR`, s.DebtIDR, val})
		s.DebtIDR = val
		return true
	}
	return false
}

// SetTopupIDR create mutations, should not duplicate
func (s *SalesMutator) SetTopupIDR(val int64) bool { //nolint:dupl false positive
	if val != s.TopupIDR {
		s.mutations = append(s.mutations, A.X{`=`, 12, val})
		s.logs = append(s.logs, A.X{`topupIDR`, s.TopupIDR, val})
		s.TopupIDR = val
		return true
	}
	return false
}

// SetTotalPriceIDR create mutations, should not duplicate
func (s *SalesMutator) SetTotalPriceIDR(val int64) bool { //nolint:dupl false positive
	if val != s.TotalPriceIDR {
		s.mutations = append(s.mutations, A.X{`=`, 13, val})
		s.logs = append(s.logs, A.X{`totalPriceIDR`, s.TotalPriceIDR, val})
		s.TotalPriceIDR = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (s *SalesMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != s.CreatedAt {
		s.mutations = append(s.mutations, A.X{`=`, 14, val})
		s.logs = append(s.logs, A.X{`createdAt`, s.CreatedAt, val})
		s.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (s *SalesMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != s.CreatedBy {
		s.mutations = append(s.mutations, A.X{`=`, 15, val})
		s.logs = append(s.logs, A.X{`createdBy`, s.CreatedBy, val})
		s.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (s *SalesMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != s.UpdatedAt {
		s.mutations = append(s.mutations, A.X{`=`, 16, val})
		s.logs = append(s.logs, A.X{`updatedAt`, s.UpdatedAt, val})
		s.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (s *SalesMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != s.UpdatedBy {
		s.mutations = append(s.mutations, A.X{`=`, 17, val})
		s.logs = append(s.logs, A.X{`updatedBy`, s.UpdatedBy, val})
		s.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (s *SalesMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != s.DeletedAt {
		s.mutations = append(s.mutations, A.X{`=`, 18, val})
		s.logs = append(s.logs, A.X{`deletedAt`, s.DeletedAt, val})
		s.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (s *SalesMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != s.DeletedBy {
		s.mutations = append(s.mutations, A.X{`=`, 19, val})
		s.logs = append(s.logs, A.X{`deletedBy`, s.DeletedBy, val})
		s.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (s *SalesMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != s.RestoredBy {
		s.mutations = append(s.mutations, A.X{`=`, 20, val})
		s.logs = append(s.logs, A.X{`restoredBy`, s.RestoredBy, val})
		s.RestoredBy = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (s *SalesMutator) SetAll(from rqCafe.Sales, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		s.Id = from.Id
		changed = true
	}
	if !excludeMap[`cashier`] && (forceMap[`cashier`] || from.Cashier != ``) {
		s.Cashier = S.Trim(from.Cashier)
		changed = true
	}
	if !excludeMap[`tenantId`] && (forceMap[`tenantId`] || from.TenantId != 0) {
		s.TenantId = from.TenantId
		changed = true
	}
	if !excludeMap[`buyerName`] && (forceMap[`buyerName`] || from.BuyerName != ``) {
		s.BuyerName = S.Trim(from.BuyerName)
		changed = true
	}
	if !excludeMap[`salesDate`] && (forceMap[`salesDate`] || from.SalesDate != ``) {
		s.SalesDate = S.Trim(from.SalesDate)
		changed = true
	}
	if !excludeMap[`paidAt`] && (forceMap[`paidAt`] || from.PaidAt != ``) {
		s.PaidAt = S.Trim(from.PaidAt)
		changed = true
	}
	if !excludeMap[`note`] && (forceMap[`note`] || from.Note != ``) {
		s.Note = S.Trim(from.Note)
		changed = true
	}
	if !excludeMap[`utensilsTaken`] && (forceMap[`utensilsTaken`] || from.UtensilsTaken != ``) {
		s.UtensilsTaken = S.Trim(from.UtensilsTaken)
		changed = true
	}
	if !excludeMap[`utensilsReturnedAt`] && (forceMap[`utensilsReturnedAt`] || from.UtensilsReturnedAt != ``) {
		s.UtensilsReturnedAt = S.Trim(from.UtensilsReturnedAt)
		changed = true
	}
	if !excludeMap[`qrisIDR`] && (forceMap[`qrisIDR`] || from.QrisIDR != 0) {
		s.QrisIDR = from.QrisIDR
		changed = true
	}
	if !excludeMap[`cashIDR`] && (forceMap[`cashIDR`] || from.CashIDR != 0) {
		s.CashIDR = from.CashIDR
		changed = true
	}
	if !excludeMap[`debtIDR`] && (forceMap[`debtIDR`] || from.DebtIDR != 0) {
		s.DebtIDR = from.DebtIDR
		changed = true
	}
	if !excludeMap[`topupIDR`] && (forceMap[`topupIDR`] || from.TopupIDR != 0) {
		s.TopupIDR = from.TopupIDR
		changed = true
	}
	if !excludeMap[`totalPriceIDR`] && (forceMap[`totalPriceIDR`] || from.TotalPriceIDR != 0) {
		s.TotalPriceIDR = from.TotalPriceIDR
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		s.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		s.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		s.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		s.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		s.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		s.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		s.RestoredBy = from.RestoredBy
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
