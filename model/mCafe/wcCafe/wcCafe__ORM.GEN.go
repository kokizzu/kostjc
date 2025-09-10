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

// BorrowedUtensilsMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcCafe__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcCafe__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcCafe__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcCafe__ORM.GEN.go
type BorrowedUtensilsMutator struct {
	rqCafe.BorrowedUtensils
	mutations []A.X
	logs      []A.X
}

// NewBorrowedUtensilsMutator create new ORM writer/command object
func NewBorrowedUtensilsMutator(adapter *Tt.Adapter) (res *BorrowedUtensilsMutator) {
	res = &BorrowedUtensilsMutator{BorrowedUtensils: rqCafe.BorrowedUtensils{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (b *BorrowedUtensilsMutator) Logs() []A.X { //nolint:dupl false positive
	return b.logs
}

// HaveMutation check whether Set* methods ever called
func (b *BorrowedUtensilsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(b.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (b *BorrowedUtensilsMutator) ClearMutations() { //nolint:dupl false positive
	b.mutations = []A.X{}
	b.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (b *BorrowedUtensilsMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := b.Adapter.Update(b.SpaceName(), b.UniqueIndexId(), A.X{b.Id}, b.ToUpdateArray())
	return !L.IsError(err, `BorrowedUtensils.DoOverwriteById failed: `+b.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (b *BorrowedUtensilsMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !b.HaveMutation() {
		return true
	}
	_, err := b.Adapter.Update(b.SpaceName(), b.UniqueIndexId(), A.X{b.Id}, b.mutations)
	return !L.IsError(err, `BorrowedUtensils.DoUpdateById failed: `+b.SpaceName())
}

// DoDeletePermanentById permanent delete
func (b *BorrowedUtensilsMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := b.Adapter.Delete(b.SpaceName(), b.UniqueIndexId(), A.X{b.Id})
	return !L.IsError(err, `BorrowedUtensils.DoDeletePermanentById failed: `+b.SpaceName())
}

// func (b *BorrowedUtensilsMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := b.ToArray()
//	_, err := b.Adapter.Upsert(b.SpaceName(), arr, A.X{
//		A.X{`=`, 0, b.Id},
//		A.X{`=`, 1, b.Customer},
//		A.X{`=`, 2, b.Items},
//		A.X{`=`, 3, b.Qty},
//		A.X{`=`, 4, b.Status},
//		A.X{`=`, 5, b.BorrowedAt},
//		A.X{`=`, 6, b.CreatedAt},
//		A.X{`=`, 7, b.CreatedBy},
//		A.X{`=`, 8, b.UpdatedAt},
//		A.X{`=`, 9, b.UpdatedBy},
//		A.X{`=`, 10, b.DeletedAt},
//		A.X{`=`, 11, b.DeletedBy},
//		A.X{`=`, 12, b.RestoredBy},
//	})
//	return !L.IsError(err, `BorrowedUtensils.DoUpsert failed: `+b.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (b *BorrowedUtensilsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := b.ToArray()
	row, err := b.Adapter.Insert(b.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			b.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `BorrowedUtensils.DoInsert failed: `+b.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (b *BorrowedUtensilsMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := b.ToArray()
	row, err := b.Adapter.Replace(b.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			b.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `BorrowedUtensils.DoUpsert failed: `+b.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (b *BorrowedUtensilsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != b.Id {
		b.mutations = append(b.mutations, A.X{`=`, 0, val})
		b.logs = append(b.logs, A.X{`id`, b.Id, val})
		b.Id = val
		return true
	}
	return false
}

// SetCustomer create mutations, should not duplicate
func (b *BorrowedUtensilsMutator) SetCustomer(val string) bool { //nolint:dupl false positive
	if val != b.Customer {
		b.mutations = append(b.mutations, A.X{`=`, 1, val})
		b.logs = append(b.logs, A.X{`customer`, b.Customer, val})
		b.Customer = val
		return true
	}
	return false
}

// SetItems create mutations, should not duplicate
func (b *BorrowedUtensilsMutator) SetItems(val string) bool { //nolint:dupl false positive
	if val != b.Items {
		b.mutations = append(b.mutations, A.X{`=`, 2, val})
		b.logs = append(b.logs, A.X{`items`, b.Items, val})
		b.Items = val
		return true
	}
	return false
}

// SetQty create mutations, should not duplicate
func (b *BorrowedUtensilsMutator) SetQty(val int64) bool { //nolint:dupl false positive
	if val != b.Qty {
		b.mutations = append(b.mutations, A.X{`=`, 3, val})
		b.logs = append(b.logs, A.X{`qty`, b.Qty, val})
		b.Qty = val
		return true
	}
	return false
}

// SetStatus create mutations, should not duplicate
func (b *BorrowedUtensilsMutator) SetStatus(val string) bool { //nolint:dupl false positive
	if val != b.Status {
		b.mutations = append(b.mutations, A.X{`=`, 4, val})
		b.logs = append(b.logs, A.X{`status`, b.Status, val})
		b.Status = val
		return true
	}
	return false
}

// SetBorrowedAt create mutations, should not duplicate
func (b *BorrowedUtensilsMutator) SetBorrowedAt(val string) bool { //nolint:dupl false positive
	if val != b.BorrowedAt {
		b.mutations = append(b.mutations, A.X{`=`, 5, val})
		b.logs = append(b.logs, A.X{`borrowedAt`, b.BorrowedAt, val})
		b.BorrowedAt = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (b *BorrowedUtensilsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != b.CreatedAt {
		b.mutations = append(b.mutations, A.X{`=`, 6, val})
		b.logs = append(b.logs, A.X{`createdAt`, b.CreatedAt, val})
		b.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (b *BorrowedUtensilsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != b.CreatedBy {
		b.mutations = append(b.mutations, A.X{`=`, 7, val})
		b.logs = append(b.logs, A.X{`createdBy`, b.CreatedBy, val})
		b.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (b *BorrowedUtensilsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != b.UpdatedAt {
		b.mutations = append(b.mutations, A.X{`=`, 8, val})
		b.logs = append(b.logs, A.X{`updatedAt`, b.UpdatedAt, val})
		b.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (b *BorrowedUtensilsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != b.UpdatedBy {
		b.mutations = append(b.mutations, A.X{`=`, 9, val})
		b.logs = append(b.logs, A.X{`updatedBy`, b.UpdatedBy, val})
		b.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (b *BorrowedUtensilsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != b.DeletedAt {
		b.mutations = append(b.mutations, A.X{`=`, 10, val})
		b.logs = append(b.logs, A.X{`deletedAt`, b.DeletedAt, val})
		b.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (b *BorrowedUtensilsMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != b.DeletedBy {
		b.mutations = append(b.mutations, A.X{`=`, 11, val})
		b.logs = append(b.logs, A.X{`deletedBy`, b.DeletedBy, val})
		b.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (b *BorrowedUtensilsMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != b.RestoredBy {
		b.mutations = append(b.mutations, A.X{`=`, 12, val})
		b.logs = append(b.logs, A.X{`restoredBy`, b.RestoredBy, val})
		b.RestoredBy = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (b *BorrowedUtensilsMutator) SetAll(from rqCafe.BorrowedUtensils, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		b.Id = from.Id
		changed = true
	}
	if !excludeMap[`customer`] && (forceMap[`customer`] || from.Customer != ``) {
		b.Customer = S.Trim(from.Customer)
		changed = true
	}
	if !excludeMap[`items`] && (forceMap[`items`] || from.Items != ``) {
		b.Items = S.Trim(from.Items)
		changed = true
	}
	if !excludeMap[`qty`] && (forceMap[`qty`] || from.Qty != 0) {
		b.Qty = from.Qty
		changed = true
	}
	if !excludeMap[`status`] && (forceMap[`status`] || from.Status != ``) {
		b.Status = S.Trim(from.Status)
		changed = true
	}
	if !excludeMap[`borrowedAt`] && (forceMap[`borrowedAt`] || from.BorrowedAt != ``) {
		b.BorrowedAt = S.Trim(from.BorrowedAt)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		b.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		b.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		b.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		b.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		b.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		b.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		b.RestoredBy = from.RestoredBy
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// LaundryMutator DAO writer/command struct
type LaundryMutator struct {
	rqCafe.Laundry
	mutations []A.X
	logs      []A.X
}

// NewLaundryMutator create new ORM writer/command object
func NewLaundryMutator(adapter *Tt.Adapter) (res *LaundryMutator) {
	res = &LaundryMutator{Laundry: rqCafe.Laundry{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (l *LaundryMutator) Logs() []A.X { //nolint:dupl false positive
	return l.logs
}

// HaveMutation check whether Set* methods ever called
func (l *LaundryMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(l.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (l *LaundryMutator) ClearMutations() { //nolint:dupl false positive
	l.mutations = []A.X{}
	l.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (l *LaundryMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := l.Adapter.Update(l.SpaceName(), l.UniqueIndexId(), A.X{l.Id}, l.ToUpdateArray())
	return !L.IsError(err, `Laundry.DoOverwriteById failed: `+l.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (l *LaundryMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !l.HaveMutation() {
		return true
	}
	_, err := l.Adapter.Update(l.SpaceName(), l.UniqueIndexId(), A.X{l.Id}, l.mutations)
	return !L.IsError(err, `Laundry.DoUpdateById failed: `+l.SpaceName())
}

// DoDeletePermanentById permanent delete
func (l *LaundryMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := l.Adapter.Delete(l.SpaceName(), l.UniqueIndexId(), A.X{l.Id})
	return !L.IsError(err, `Laundry.DoDeletePermanentById failed: `+l.SpaceName())
}

// func (l *LaundryMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := l.ToArray()
//	_, err := l.Adapter.Upsert(l.SpaceName(), arr, A.X{
//		A.X{`=`, 0, l.Id},
//		A.X{`=`, 1, l.Customer},
//		A.X{`=`, 2, l.Items},
//		A.X{`=`, 3, l.Status},
//		A.X{`=`, 4, l.Note},
//		A.X{`=`, 5, l.LaundryAt},
//		A.X{`=`, 6, l.CreatedAt},
//		A.X{`=`, 7, l.CreatedBy},
//		A.X{`=`, 8, l.UpdatedAt},
//		A.X{`=`, 9, l.UpdatedBy},
//		A.X{`=`, 10, l.DeletedAt},
//		A.X{`=`, 11, l.DeletedBy},
//		A.X{`=`, 12, l.RestoredBy},
//		A.X{`=`, 13, l.PaidAt},
//		A.X{`=`, 14, l.PaidBy},
//		A.X{`=`, 15, l.WashAt},
//		A.X{`=`, 16, l.WashBy},
//		A.X{`=`, 17, l.DryAt},
//		A.X{`=`, 18, l.DryBy},
//		A.X{`=`, 19, l.FoldAt},
//		A.X{`=`, 20, l.FoldBy},
//		A.X{`=`, 21, l.NotifyAt},
//		A.X{`=`, 22, l.NotifyBy},
//		A.X{`=`, 23, l.GivenAt},
//		A.X{`=`, 24, l.GivenBys},
//		A.X{`=`, 25, l.Weight},
//	})
//	return !L.IsError(err, `Laundry.DoUpsert failed: `+l.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (l *LaundryMutator) DoInsert() bool { //nolint:dupl false positive
	arr := l.ToArray()
	row, err := l.Adapter.Insert(l.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			l.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Laundry.DoInsert failed: `+l.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (l *LaundryMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := l.ToArray()
	row, err := l.Adapter.Replace(l.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			l.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Laundry.DoUpsert failed: `+l.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (l *LaundryMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != l.Id {
		l.mutations = append(l.mutations, A.X{`=`, 0, val})
		l.logs = append(l.logs, A.X{`id`, l.Id, val})
		l.Id = val
		return true
	}
	return false
}

// SetCustomer create mutations, should not duplicate
func (l *LaundryMutator) SetCustomer(val string) bool { //nolint:dupl false positive
	if val != l.Customer {
		l.mutations = append(l.mutations, A.X{`=`, 1, val})
		l.logs = append(l.logs, A.X{`customer`, l.Customer, val})
		l.Customer = val
		return true
	}
	return false
}

// SetItems create mutations, should not duplicate
func (l *LaundryMutator) SetItems(val string) bool { //nolint:dupl false positive
	if val != l.Items {
		l.mutations = append(l.mutations, A.X{`=`, 2, val})
		l.logs = append(l.logs, A.X{`items`, l.Items, val})
		l.Items = val
		return true
	}
	return false
}

// SetStatus create mutations, should not duplicate
func (l *LaundryMutator) SetStatus(val string) bool { //nolint:dupl false positive
	if val != l.Status {
		l.mutations = append(l.mutations, A.X{`=`, 3, val})
		l.logs = append(l.logs, A.X{`status`, l.Status, val})
		l.Status = val
		return true
	}
	return false
}

// SetNote create mutations, should not duplicate
func (l *LaundryMutator) SetNote(val string) bool { //nolint:dupl false positive
	if val != l.Note {
		l.mutations = append(l.mutations, A.X{`=`, 4, val})
		l.logs = append(l.logs, A.X{`note`, l.Note, val})
		l.Note = val
		return true
	}
	return false
}

// SetLaundryAt create mutations, should not duplicate
func (l *LaundryMutator) SetLaundryAt(val string) bool { //nolint:dupl false positive
	if val != l.LaundryAt {
		l.mutations = append(l.mutations, A.X{`=`, 5, val})
		l.logs = append(l.logs, A.X{`laundryAt`, l.LaundryAt, val})
		l.LaundryAt = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (l *LaundryMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != l.CreatedAt {
		l.mutations = append(l.mutations, A.X{`=`, 6, val})
		l.logs = append(l.logs, A.X{`createdAt`, l.CreatedAt, val})
		l.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (l *LaundryMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != l.CreatedBy {
		l.mutations = append(l.mutations, A.X{`=`, 7, val})
		l.logs = append(l.logs, A.X{`createdBy`, l.CreatedBy, val})
		l.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (l *LaundryMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != l.UpdatedAt {
		l.mutations = append(l.mutations, A.X{`=`, 8, val})
		l.logs = append(l.logs, A.X{`updatedAt`, l.UpdatedAt, val})
		l.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (l *LaundryMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != l.UpdatedBy {
		l.mutations = append(l.mutations, A.X{`=`, 9, val})
		l.logs = append(l.logs, A.X{`updatedBy`, l.UpdatedBy, val})
		l.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (l *LaundryMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != l.DeletedAt {
		l.mutations = append(l.mutations, A.X{`=`, 10, val})
		l.logs = append(l.logs, A.X{`deletedAt`, l.DeletedAt, val})
		l.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (l *LaundryMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != l.DeletedBy {
		l.mutations = append(l.mutations, A.X{`=`, 11, val})
		l.logs = append(l.logs, A.X{`deletedBy`, l.DeletedBy, val})
		l.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (l *LaundryMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != l.RestoredBy {
		l.mutations = append(l.mutations, A.X{`=`, 12, val})
		l.logs = append(l.logs, A.X{`restoredBy`, l.RestoredBy, val})
		l.RestoredBy = val
		return true
	}
	return false
}

// SetPaidAt create mutations, should not duplicate
func (l *LaundryMutator) SetPaidAt(val string) bool { //nolint:dupl false positive
	if val != l.PaidAt {
		l.mutations = append(l.mutations, A.X{`=`, 13, val})
		l.logs = append(l.logs, A.X{`paidAt`, l.PaidAt, val})
		l.PaidAt = val
		return true
	}
	return false
}

// SetPaidBy create mutations, should not duplicate
func (l *LaundryMutator) SetPaidBy(val uint64) bool { //nolint:dupl false positive
	if val != l.PaidBy {
		l.mutations = append(l.mutations, A.X{`=`, 14, val})
		l.logs = append(l.logs, A.X{`paidBy`, l.PaidBy, val})
		l.PaidBy = val
		return true
	}
	return false
}

// SetWashAt create mutations, should not duplicate
func (l *LaundryMutator) SetWashAt(val string) bool { //nolint:dupl false positive
	if val != l.WashAt {
		l.mutations = append(l.mutations, A.X{`=`, 15, val})
		l.logs = append(l.logs, A.X{`washAt`, l.WashAt, val})
		l.WashAt = val
		return true
	}
	return false
}

// SetWashBy create mutations, should not duplicate
func (l *LaundryMutator) SetWashBy(val uint64) bool { //nolint:dupl false positive
	if val != l.WashBy {
		l.mutations = append(l.mutations, A.X{`=`, 16, val})
		l.logs = append(l.logs, A.X{`washBy`, l.WashBy, val})
		l.WashBy = val
		return true
	}
	return false
}

// SetDryAt create mutations, should not duplicate
func (l *LaundryMutator) SetDryAt(val string) bool { //nolint:dupl false positive
	if val != l.DryAt {
		l.mutations = append(l.mutations, A.X{`=`, 17, val})
		l.logs = append(l.logs, A.X{`dryAt`, l.DryAt, val})
		l.DryAt = val
		return true
	}
	return false
}

// SetDryBy create mutations, should not duplicate
func (l *LaundryMutator) SetDryBy(val uint64) bool { //nolint:dupl false positive
	if val != l.DryBy {
		l.mutations = append(l.mutations, A.X{`=`, 18, val})
		l.logs = append(l.logs, A.X{`dryBy`, l.DryBy, val})
		l.DryBy = val
		return true
	}
	return false
}

// SetFoldAt create mutations, should not duplicate
func (l *LaundryMutator) SetFoldAt(val string) bool { //nolint:dupl false positive
	if val != l.FoldAt {
		l.mutations = append(l.mutations, A.X{`=`, 19, val})
		l.logs = append(l.logs, A.X{`foldAt`, l.FoldAt, val})
		l.FoldAt = val
		return true
	}
	return false
}

// SetFoldBy create mutations, should not duplicate
func (l *LaundryMutator) SetFoldBy(val uint64) bool { //nolint:dupl false positive
	if val != l.FoldBy {
		l.mutations = append(l.mutations, A.X{`=`, 20, val})
		l.logs = append(l.logs, A.X{`foldBy`, l.FoldBy, val})
		l.FoldBy = val
		return true
	}
	return false
}

// SetNotifyAt create mutations, should not duplicate
func (l *LaundryMutator) SetNotifyAt(val string) bool { //nolint:dupl false positive
	if val != l.NotifyAt {
		l.mutations = append(l.mutations, A.X{`=`, 21, val})
		l.logs = append(l.logs, A.X{`notifyAt`, l.NotifyAt, val})
		l.NotifyAt = val
		return true
	}
	return false
}

// SetNotifyBy create mutations, should not duplicate
func (l *LaundryMutator) SetNotifyBy(val uint64) bool { //nolint:dupl false positive
	if val != l.NotifyBy {
		l.mutations = append(l.mutations, A.X{`=`, 22, val})
		l.logs = append(l.logs, A.X{`notifyBy`, l.NotifyBy, val})
		l.NotifyBy = val
		return true
	}
	return false
}

// SetGivenAt create mutations, should not duplicate
func (l *LaundryMutator) SetGivenAt(val string) bool { //nolint:dupl false positive
	if val != l.GivenAt {
		l.mutations = append(l.mutations, A.X{`=`, 23, val})
		l.logs = append(l.logs, A.X{`givenAt`, l.GivenAt, val})
		l.GivenAt = val
		return true
	}
	return false
}

// SetGivenBys create mutations, should not duplicate
func (l *LaundryMutator) SetGivenBys(val string) bool { //nolint:dupl false positive
	if val != l.GivenBys {
		l.mutations = append(l.mutations, A.X{`=`, 24, val})
		l.logs = append(l.logs, A.X{`givenBys`, l.GivenBys, val})
		l.GivenBys = val
		return true
	}
	return false
}

// SetWeight create mutations, should not duplicate
func (l *LaundryMutator) SetWeight(val float64) bool { //nolint:dupl false positive
	if val != l.Weight {
		l.mutations = append(l.mutations, A.X{`=`, 25, val})
		l.logs = append(l.logs, A.X{`weight`, l.Weight, val})
		l.Weight = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (l *LaundryMutator) SetAll(from rqCafe.Laundry, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		l.Id = from.Id
		changed = true
	}
	if !excludeMap[`customer`] && (forceMap[`customer`] || from.Customer != ``) {
		l.Customer = S.Trim(from.Customer)
		changed = true
	}
	if !excludeMap[`items`] && (forceMap[`items`] || from.Items != ``) {
		l.Items = S.Trim(from.Items)
		changed = true
	}
	if !excludeMap[`status`] && (forceMap[`status`] || from.Status != ``) {
		l.Status = S.Trim(from.Status)
		changed = true
	}
	if !excludeMap[`note`] && (forceMap[`note`] || from.Note != ``) {
		l.Note = S.Trim(from.Note)
		changed = true
	}
	if !excludeMap[`laundryAt`] && (forceMap[`laundryAt`] || from.LaundryAt != ``) {
		l.LaundryAt = S.Trim(from.LaundryAt)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		l.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		l.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		l.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		l.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		l.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		l.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		l.RestoredBy = from.RestoredBy
		changed = true
	}
	if !excludeMap[`paidAt`] && (forceMap[`paidAt`] || from.PaidAt != ``) {
		l.PaidAt = S.Trim(from.PaidAt)
		changed = true
	}
	if !excludeMap[`paidBy`] && (forceMap[`paidBy`] || from.PaidBy != 0) {
		l.PaidBy = from.PaidBy
		changed = true
	}
	if !excludeMap[`washAt`] && (forceMap[`washAt`] || from.WashAt != ``) {
		l.WashAt = S.Trim(from.WashAt)
		changed = true
	}
	if !excludeMap[`washBy`] && (forceMap[`washBy`] || from.WashBy != 0) {
		l.WashBy = from.WashBy
		changed = true
	}
	if !excludeMap[`dryAt`] && (forceMap[`dryAt`] || from.DryAt != ``) {
		l.DryAt = S.Trim(from.DryAt)
		changed = true
	}
	if !excludeMap[`dryBy`] && (forceMap[`dryBy`] || from.DryBy != 0) {
		l.DryBy = from.DryBy
		changed = true
	}
	if !excludeMap[`foldAt`] && (forceMap[`foldAt`] || from.FoldAt != ``) {
		l.FoldAt = S.Trim(from.FoldAt)
		changed = true
	}
	if !excludeMap[`foldBy`] && (forceMap[`foldBy`] || from.FoldBy != 0) {
		l.FoldBy = from.FoldBy
		changed = true
	}
	if !excludeMap[`notifyAt`] && (forceMap[`notifyAt`] || from.NotifyAt != ``) {
		l.NotifyAt = S.Trim(from.NotifyAt)
		changed = true
	}
	if !excludeMap[`notifyBy`] && (forceMap[`notifyBy`] || from.NotifyBy != 0) {
		l.NotifyBy = from.NotifyBy
		changed = true
	}
	if !excludeMap[`givenAt`] && (forceMap[`givenAt`] || from.GivenAt != ``) {
		l.GivenAt = S.Trim(from.GivenAt)
		changed = true
	}
	if !excludeMap[`givenBys`] && (forceMap[`givenBys`] || from.GivenBys != ``) {
		l.GivenBys = S.Trim(from.GivenBys)
		changed = true
	}
	if !excludeMap[`weight`] && (forceMap[`weight`] || from.Weight != 0) {
		l.Weight = from.Weight
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// MenusMutator DAO writer/command struct
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
	res.MenuIds = []any{}
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
//		A.X{`=`, 4, s.MenuIds},
//		A.X{`=`, 5, s.QrisIDR},
//		A.X{`=`, 6, s.CashIDR},
//		A.X{`=`, 7, s.DebtIDR},
//		A.X{`=`, 8, s.TopupIDR},
//		A.X{`=`, 9, s.TotalPriceIDR},
//		A.X{`=`, 10, s.SalesDate},
//		A.X{`=`, 11, s.PaidAt},
//		A.X{`=`, 12, s.Note},
//		A.X{`=`, 13, s.Donation},
//		A.X{`=`, 14, s.TransferIDR},
//		A.X{`=`, 15, s.PaymentMethod},
//		A.X{`=`, 16, s.PaymentStatus},
//		A.X{`=`, 17, s.CreatedAt},
//		A.X{`=`, 18, s.CreatedBy},
//		A.X{`=`, 19, s.UpdatedAt},
//		A.X{`=`, 20, s.UpdatedBy},
//		A.X{`=`, 21, s.DeletedAt},
//		A.X{`=`, 22, s.DeletedBy},
//		A.X{`=`, 23, s.RestoredBy},
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

// SetMenuIds create mutations, should not duplicate
func (s *SalesMutator) SetMenuIds(val []any) bool { //nolint:dupl false positive
	s.mutations = append(s.mutations, A.X{`=`, 4, val})
	s.logs = append(s.logs, A.X{`menuIds`, s.MenuIds, val})
	s.MenuIds = val
	return true
}

// SetQrisIDR create mutations, should not duplicate
func (s *SalesMutator) SetQrisIDR(val int64) bool { //nolint:dupl false positive
	if val != s.QrisIDR {
		s.mutations = append(s.mutations, A.X{`=`, 5, val})
		s.logs = append(s.logs, A.X{`qrisIDR`, s.QrisIDR, val})
		s.QrisIDR = val
		return true
	}
	return false
}

// SetCashIDR create mutations, should not duplicate
func (s *SalesMutator) SetCashIDR(val int64) bool { //nolint:dupl false positive
	if val != s.CashIDR {
		s.mutations = append(s.mutations, A.X{`=`, 6, val})
		s.logs = append(s.logs, A.X{`cashIDR`, s.CashIDR, val})
		s.CashIDR = val
		return true
	}
	return false
}

// SetDebtIDR create mutations, should not duplicate
func (s *SalesMutator) SetDebtIDR(val int64) bool { //nolint:dupl false positive
	if val != s.DebtIDR {
		s.mutations = append(s.mutations, A.X{`=`, 7, val})
		s.logs = append(s.logs, A.X{`debtIDR`, s.DebtIDR, val})
		s.DebtIDR = val
		return true
	}
	return false
}

// SetTopupIDR create mutations, should not duplicate
func (s *SalesMutator) SetTopupIDR(val int64) bool { //nolint:dupl false positive
	if val != s.TopupIDR {
		s.mutations = append(s.mutations, A.X{`=`, 8, val})
		s.logs = append(s.logs, A.X{`topupIDR`, s.TopupIDR, val})
		s.TopupIDR = val
		return true
	}
	return false
}

// SetTotalPriceIDR create mutations, should not duplicate
func (s *SalesMutator) SetTotalPriceIDR(val int64) bool { //nolint:dupl false positive
	if val != s.TotalPriceIDR {
		s.mutations = append(s.mutations, A.X{`=`, 9, val})
		s.logs = append(s.logs, A.X{`totalPriceIDR`, s.TotalPriceIDR, val})
		s.TotalPriceIDR = val
		return true
	}
	return false
}

// SetSalesDate create mutations, should not duplicate
func (s *SalesMutator) SetSalesDate(val string) bool { //nolint:dupl false positive
	if val != s.SalesDate {
		s.mutations = append(s.mutations, A.X{`=`, 10, val})
		s.logs = append(s.logs, A.X{`salesDate`, s.SalesDate, val})
		s.SalesDate = val
		return true
	}
	return false
}

// SetPaidAt create mutations, should not duplicate
func (s *SalesMutator) SetPaidAt(val string) bool { //nolint:dupl false positive
	if val != s.PaidAt {
		s.mutations = append(s.mutations, A.X{`=`, 11, val})
		s.logs = append(s.logs, A.X{`paidAt`, s.PaidAt, val})
		s.PaidAt = val
		return true
	}
	return false
}

// SetNote create mutations, should not duplicate
func (s *SalesMutator) SetNote(val string) bool { //nolint:dupl false positive
	if val != s.Note {
		s.mutations = append(s.mutations, A.X{`=`, 12, val})
		s.logs = append(s.logs, A.X{`note`, s.Note, val})
		s.Note = val
		return true
	}
	return false
}

// SetDonation create mutations, should not duplicate
func (s *SalesMutator) SetDonation(val int64) bool { //nolint:dupl false positive
	if val != s.Donation {
		s.mutations = append(s.mutations, A.X{`=`, 13, val})
		s.logs = append(s.logs, A.X{`donation`, s.Donation, val})
		s.Donation = val
		return true
	}
	return false
}

// SetTransferIDR create mutations, should not duplicate
func (s *SalesMutator) SetTransferIDR(val int64) bool { //nolint:dupl false positive
	if val != s.TransferIDR {
		s.mutations = append(s.mutations, A.X{`=`, 14, val})
		s.logs = append(s.logs, A.X{`transferIDR`, s.TransferIDR, val})
		s.TransferIDR = val
		return true
	}
	return false
}

// SetPaymentMethod create mutations, should not duplicate
func (s *SalesMutator) SetPaymentMethod(val string) bool { //nolint:dupl false positive
	if val != s.PaymentMethod {
		s.mutations = append(s.mutations, A.X{`=`, 15, val})
		s.logs = append(s.logs, A.X{`paymentMethod`, s.PaymentMethod, val})
		s.PaymentMethod = val
		return true
	}
	return false
}

// SetPaymentStatus create mutations, should not duplicate
func (s *SalesMutator) SetPaymentStatus(val string) bool { //nolint:dupl false positive
	if val != s.PaymentStatus {
		s.mutations = append(s.mutations, A.X{`=`, 16, val})
		s.logs = append(s.logs, A.X{`paymentStatus`, s.PaymentStatus, val})
		s.PaymentStatus = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (s *SalesMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != s.CreatedAt {
		s.mutations = append(s.mutations, A.X{`=`, 17, val})
		s.logs = append(s.logs, A.X{`createdAt`, s.CreatedAt, val})
		s.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (s *SalesMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != s.CreatedBy {
		s.mutations = append(s.mutations, A.X{`=`, 18, val})
		s.logs = append(s.logs, A.X{`createdBy`, s.CreatedBy, val})
		s.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (s *SalesMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != s.UpdatedAt {
		s.mutations = append(s.mutations, A.X{`=`, 19, val})
		s.logs = append(s.logs, A.X{`updatedAt`, s.UpdatedAt, val})
		s.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (s *SalesMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != s.UpdatedBy {
		s.mutations = append(s.mutations, A.X{`=`, 20, val})
		s.logs = append(s.logs, A.X{`updatedBy`, s.UpdatedBy, val})
		s.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (s *SalesMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != s.DeletedAt {
		s.mutations = append(s.mutations, A.X{`=`, 21, val})
		s.logs = append(s.logs, A.X{`deletedAt`, s.DeletedAt, val})
		s.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (s *SalesMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != s.DeletedBy {
		s.mutations = append(s.mutations, A.X{`=`, 22, val})
		s.logs = append(s.logs, A.X{`deletedBy`, s.DeletedBy, val})
		s.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (s *SalesMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != s.RestoredBy {
		s.mutations = append(s.mutations, A.X{`=`, 23, val})
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
	if !excludeMap[`menuIds`] && (forceMap[`menuIds`] || from.MenuIds != nil) {
		s.MenuIds = from.MenuIds
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
	if !excludeMap[`donation`] && (forceMap[`donation`] || from.Donation != 0) {
		s.Donation = from.Donation
		changed = true
	}
	if !excludeMap[`transferIDR`] && (forceMap[`transferIDR`] || from.TransferIDR != 0) {
		s.TransferIDR = from.TransferIDR
		changed = true
	}
	if !excludeMap[`paymentMethod`] && (forceMap[`paymentMethod`] || from.PaymentMethod != ``) {
		s.PaymentMethod = S.Trim(from.PaymentMethod)
		changed = true
	}
	if !excludeMap[`paymentStatus`] && (forceMap[`paymentStatus`] || from.PaymentStatus != ``) {
		s.PaymentStatus = S.Trim(from.PaymentStatus)
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
