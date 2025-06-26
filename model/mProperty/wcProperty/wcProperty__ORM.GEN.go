package wcProperty

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"kostjc/model/mProperty/rqProperty"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

// BookingsMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcProperty__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcProperty__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcProperty__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcProperty__ORM.GEN.go
type BookingsMutator struct {
	rqProperty.Bookings
	mutations []A.X
	logs      []A.X
}

// NewBookingsMutator create new ORM writer/command object
func NewBookingsMutator(adapter *Tt.Adapter) (res *BookingsMutator) {
	res = &BookingsMutator{Bookings: rqProperty.Bookings{Adapter: adapter}}
	res.ExtraTenants = []any{}
	return
}

// Logs get array of logs [field, old, new]
func (b *BookingsMutator) Logs() []A.X { //nolint:dupl false positive
	return b.logs
}

// HaveMutation check whether Set* methods ever called
func (b *BookingsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(b.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (b *BookingsMutator) ClearMutations() { //nolint:dupl false positive
	b.mutations = []A.X{}
	b.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (b *BookingsMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := b.Adapter.Update(b.SpaceName(), b.UniqueIndexId(), A.X{b.Id}, b.ToUpdateArray())
	return !L.IsError(err, `Bookings.DoOverwriteById failed: `+b.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (b *BookingsMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !b.HaveMutation() {
		return true
	}
	_, err := b.Adapter.Update(b.SpaceName(), b.UniqueIndexId(), A.X{b.Id}, b.mutations)
	return !L.IsError(err, `Bookings.DoUpdateById failed: `+b.SpaceName())
}

// DoDeletePermanentById permanent delete
func (b *BookingsMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := b.Adapter.Delete(b.SpaceName(), b.UniqueIndexId(), A.X{b.Id})
	return !L.IsError(err, `Bookings.DoDeletePermanentById failed: `+b.SpaceName())
}

// func (b *BookingsMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := b.ToArray()
//	_, err := b.Adapter.Upsert(b.SpaceName(), arr, A.X{
//		A.X{`=`, 0, b.Id},
//		A.X{`=`, 1, b.DateStart},
//		A.X{`=`, 2, b.DateEnd},
//		A.X{`=`, 3, b.BasePriceIDR},
//		A.X{`=`, 4, b.FacilitiesObj},
//		A.X{`=`, 5, b.TotalPriceIDR},
//		A.X{`=`, 6, b.PaidAt},
//		A.X{`=`, 7, b.TenantId},
//		A.X{`=`, 8, b.CreatedAt},
//		A.X{`=`, 9, b.CreatedBy},
//		A.X{`=`, 10, b.UpdatedAt},
//		A.X{`=`, 11, b.UpdatedBy},
//		A.X{`=`, 12, b.DeletedAt},
//		A.X{`=`, 13, b.DeletedBy},
//		A.X{`=`, 14, b.RestoredBy},
//		A.X{`=`, 15, b.ExtraTenants},
//		A.X{`=`, 16, b.RoomId},
//	})
//	return !L.IsError(err, `Bookings.DoUpsert failed: `+b.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (b *BookingsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := b.ToArray()
	row, err := b.Adapter.Insert(b.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			b.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Bookings.DoInsert failed: `+b.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (b *BookingsMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := b.ToArray()
	row, err := b.Adapter.Replace(b.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			b.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Bookings.DoUpsert failed: `+b.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (b *BookingsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != b.Id {
		b.mutations = append(b.mutations, A.X{`=`, 0, val})
		b.logs = append(b.logs, A.X{`id`, b.Id, val})
		b.Id = val
		return true
	}
	return false
}

// SetDateStart create mutations, should not duplicate
func (b *BookingsMutator) SetDateStart(val string) bool { //nolint:dupl false positive
	if val != b.DateStart {
		b.mutations = append(b.mutations, A.X{`=`, 1, val})
		b.logs = append(b.logs, A.X{`dateStart`, b.DateStart, val})
		b.DateStart = val
		return true
	}
	return false
}

// SetDateEnd create mutations, should not duplicate
func (b *BookingsMutator) SetDateEnd(val string) bool { //nolint:dupl false positive
	if val != b.DateEnd {
		b.mutations = append(b.mutations, A.X{`=`, 2, val})
		b.logs = append(b.logs, A.X{`dateEnd`, b.DateEnd, val})
		b.DateEnd = val
		return true
	}
	return false
}

// SetBasePriceIDR create mutations, should not duplicate
func (b *BookingsMutator) SetBasePriceIDR(val int64) bool { //nolint:dupl false positive
	if val != b.BasePriceIDR {
		b.mutations = append(b.mutations, A.X{`=`, 3, val})
		b.logs = append(b.logs, A.X{`basePriceIDR`, b.BasePriceIDR, val})
		b.BasePriceIDR = val
		return true
	}
	return false
}

// SetFacilitiesObj create mutations, should not duplicate
func (b *BookingsMutator) SetFacilitiesObj(val string) bool { //nolint:dupl false positive
	if val != b.FacilitiesObj {
		b.mutations = append(b.mutations, A.X{`=`, 4, val})
		b.logs = append(b.logs, A.X{`facilitiesObj`, b.FacilitiesObj, val})
		b.FacilitiesObj = val
		return true
	}
	return false
}

// SetTotalPriceIDR create mutations, should not duplicate
func (b *BookingsMutator) SetTotalPriceIDR(val int64) bool { //nolint:dupl false positive
	if val != b.TotalPriceIDR {
		b.mutations = append(b.mutations, A.X{`=`, 5, val})
		b.logs = append(b.logs, A.X{`totalPriceIDR`, b.TotalPriceIDR, val})
		b.TotalPriceIDR = val
		return true
	}
	return false
}

// SetPaidAt create mutations, should not duplicate
func (b *BookingsMutator) SetPaidAt(val string) bool { //nolint:dupl false positive
	if val != b.PaidAt {
		b.mutations = append(b.mutations, A.X{`=`, 6, val})
		b.logs = append(b.logs, A.X{`paidAt`, b.PaidAt, val})
		b.PaidAt = val
		return true
	}
	return false
}

// SetTenantId create mutations, should not duplicate
func (b *BookingsMutator) SetTenantId(val uint64) bool { //nolint:dupl false positive
	if val != b.TenantId {
		b.mutations = append(b.mutations, A.X{`=`, 7, val})
		b.logs = append(b.logs, A.X{`tenantId`, b.TenantId, val})
		b.TenantId = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (b *BookingsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != b.CreatedAt {
		b.mutations = append(b.mutations, A.X{`=`, 8, val})
		b.logs = append(b.logs, A.X{`createdAt`, b.CreatedAt, val})
		b.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (b *BookingsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != b.CreatedBy {
		b.mutations = append(b.mutations, A.X{`=`, 9, val})
		b.logs = append(b.logs, A.X{`createdBy`, b.CreatedBy, val})
		b.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (b *BookingsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != b.UpdatedAt {
		b.mutations = append(b.mutations, A.X{`=`, 10, val})
		b.logs = append(b.logs, A.X{`updatedAt`, b.UpdatedAt, val})
		b.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (b *BookingsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != b.UpdatedBy {
		b.mutations = append(b.mutations, A.X{`=`, 11, val})
		b.logs = append(b.logs, A.X{`updatedBy`, b.UpdatedBy, val})
		b.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (b *BookingsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != b.DeletedAt {
		b.mutations = append(b.mutations, A.X{`=`, 12, val})
		b.logs = append(b.logs, A.X{`deletedAt`, b.DeletedAt, val})
		b.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (b *BookingsMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != b.DeletedBy {
		b.mutations = append(b.mutations, A.X{`=`, 13, val})
		b.logs = append(b.logs, A.X{`deletedBy`, b.DeletedBy, val})
		b.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (b *BookingsMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != b.RestoredBy {
		b.mutations = append(b.mutations, A.X{`=`, 14, val})
		b.logs = append(b.logs, A.X{`restoredBy`, b.RestoredBy, val})
		b.RestoredBy = val
		return true
	}
	return false
}

// SetExtraTenants create mutations, should not duplicate
func (b *BookingsMutator) SetExtraTenants(val []any) bool { //nolint:dupl false positive
	b.mutations = append(b.mutations, A.X{`=`, 15, val})
	b.logs = append(b.logs, A.X{`extraTenants`, b.ExtraTenants, val})
	b.ExtraTenants = val
	return true
}

// SetRoomId create mutations, should not duplicate
func (b *BookingsMutator) SetRoomId(val uint64) bool { //nolint:dupl false positive
	if val != b.RoomId {
		b.mutations = append(b.mutations, A.X{`=`, 16, val})
		b.logs = append(b.logs, A.X{`roomId`, b.RoomId, val})
		b.RoomId = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (b *BookingsMutator) SetAll(from rqProperty.Bookings, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
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
	if !excludeMap[`dateStart`] && (forceMap[`dateStart`] || from.DateStart != ``) {
		b.DateStart = S.Trim(from.DateStart)
		changed = true
	}
	if !excludeMap[`dateEnd`] && (forceMap[`dateEnd`] || from.DateEnd != ``) {
		b.DateEnd = S.Trim(from.DateEnd)
		changed = true
	}
	if !excludeMap[`basePriceIDR`] && (forceMap[`basePriceIDR`] || from.BasePriceIDR != 0) {
		b.BasePriceIDR = from.BasePriceIDR
		changed = true
	}
	if !excludeMap[`facilitiesObj`] && (forceMap[`facilitiesObj`] || from.FacilitiesObj != ``) {
		b.FacilitiesObj = S.Trim(from.FacilitiesObj)
		changed = true
	}
	if !excludeMap[`totalPriceIDR`] && (forceMap[`totalPriceIDR`] || from.TotalPriceIDR != 0) {
		b.TotalPriceIDR = from.TotalPriceIDR
		changed = true
	}
	if !excludeMap[`paidAt`] && (forceMap[`paidAt`] || from.PaidAt != ``) {
		b.PaidAt = S.Trim(from.PaidAt)
		changed = true
	}
	if !excludeMap[`tenantId`] && (forceMap[`tenantId`] || from.TenantId != 0) {
		b.TenantId = from.TenantId
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
	if !excludeMap[`extraTenants`] && (forceMap[`extraTenants`] || from.ExtraTenants != nil) {
		b.ExtraTenants = from.ExtraTenants
		changed = true
	}
	if !excludeMap[`roomId`] && (forceMap[`roomId`] || from.RoomId != 0) {
		b.RoomId = from.RoomId
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// BuildingsMutator DAO writer/command struct
type BuildingsMutator struct {
	rqProperty.Buildings
	mutations []A.X
	logs      []A.X
}

// NewBuildingsMutator create new ORM writer/command object
func NewBuildingsMutator(adapter *Tt.Adapter) (res *BuildingsMutator) {
	res = &BuildingsMutator{Buildings: rqProperty.Buildings{Adapter: adapter}}
	res.Facilities = []any{}
	return
}

// Logs get array of logs [field, old, new]
func (b *BuildingsMutator) Logs() []A.X { //nolint:dupl false positive
	return b.logs
}

// HaveMutation check whether Set* methods ever called
func (b *BuildingsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(b.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (b *BuildingsMutator) ClearMutations() { //nolint:dupl false positive
	b.mutations = []A.X{}
	b.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (b *BuildingsMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := b.Adapter.Update(b.SpaceName(), b.UniqueIndexId(), A.X{b.Id}, b.ToUpdateArray())
	return !L.IsError(err, `Buildings.DoOverwriteById failed: `+b.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (b *BuildingsMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !b.HaveMutation() {
		return true
	}
	_, err := b.Adapter.Update(b.SpaceName(), b.UniqueIndexId(), A.X{b.Id}, b.mutations)
	return !L.IsError(err, `Buildings.DoUpdateById failed: `+b.SpaceName())
}

// DoDeletePermanentById permanent delete
func (b *BuildingsMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := b.Adapter.Delete(b.SpaceName(), b.UniqueIndexId(), A.X{b.Id})
	return !L.IsError(err, `Buildings.DoDeletePermanentById failed: `+b.SpaceName())
}

// func (b *BuildingsMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := b.ToArray()
//	_, err := b.Adapter.Upsert(b.SpaceName(), arr, A.X{
//		A.X{`=`, 0, b.Id},
//		A.X{`=`, 1, b.BuildingName},
//		A.X{`=`, 2, b.LocationId},
//		A.X{`=`, 3, b.Facilities},
//		A.X{`=`, 4, b.CreatedAt},
//		A.X{`=`, 5, b.CreatedBy},
//		A.X{`=`, 6, b.UpdatedAt},
//		A.X{`=`, 7, b.UpdatedBy},
//		A.X{`=`, 8, b.DeletedAt},
//		A.X{`=`, 9, b.DeletedBy},
//		A.X{`=`, 10, b.RestoredBy},
//	})
//	return !L.IsError(err, `Buildings.DoUpsert failed: `+b.SpaceName()+ `\n%#v`, arr)
// }

// DoOverwriteByBuildingName update all columns, error if not exists, not using mutations/Set*
func (b *BuildingsMutator) DoOverwriteByBuildingName() bool { //nolint:dupl false positive
	_, err := b.Adapter.Update(b.SpaceName(), b.UniqueIndexBuildingName(), A.X{b.BuildingName}, b.ToUpdateArray())
	return !L.IsError(err, `Buildings.DoOverwriteByBuildingName failed: `+b.SpaceName())
}

// DoUpdateByBuildingName update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (b *BuildingsMutator) DoUpdateByBuildingName() bool { //nolint:dupl false positive
	if !b.HaveMutation() {
		return true
	}
	_, err := b.Adapter.Update(b.SpaceName(), b.UniqueIndexBuildingName(), A.X{b.BuildingName}, b.mutations)
	return !L.IsError(err, `Buildings.DoUpdateByBuildingName failed: `+b.SpaceName())
}

// DoDeletePermanentByBuildingName permanent delete
func (b *BuildingsMutator) DoDeletePermanentByBuildingName() bool { //nolint:dupl false positive
	_, err := b.Adapter.Delete(b.SpaceName(), b.UniqueIndexBuildingName(), A.X{b.BuildingName})
	return !L.IsError(err, `Buildings.DoDeletePermanentByBuildingName failed: `+b.SpaceName())
}

// DoInsert insert, error if already exists
func (b *BuildingsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := b.ToArray()
	row, err := b.Adapter.Insert(b.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			b.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Buildings.DoInsert failed: `+b.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (b *BuildingsMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := b.ToArray()
	row, err := b.Adapter.Replace(b.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			b.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Buildings.DoUpsert failed: `+b.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (b *BuildingsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != b.Id {
		b.mutations = append(b.mutations, A.X{`=`, 0, val})
		b.logs = append(b.logs, A.X{`id`, b.Id, val})
		b.Id = val
		return true
	}
	return false
}

// SetBuildingName create mutations, should not duplicate
func (b *BuildingsMutator) SetBuildingName(val string) bool { //nolint:dupl false positive
	if val != b.BuildingName {
		b.mutations = append(b.mutations, A.X{`=`, 1, val})
		b.logs = append(b.logs, A.X{`buildingName`, b.BuildingName, val})
		b.BuildingName = val
		return true
	}
	return false
}

// SetLocationId create mutations, should not duplicate
func (b *BuildingsMutator) SetLocationId(val uint64) bool { //nolint:dupl false positive
	if val != b.LocationId {
		b.mutations = append(b.mutations, A.X{`=`, 2, val})
		b.logs = append(b.logs, A.X{`locationId`, b.LocationId, val})
		b.LocationId = val
		return true
	}
	return false
}

// SetFacilities create mutations, should not duplicate
func (b *BuildingsMutator) SetFacilities(val []any) bool { //nolint:dupl false positive
	b.mutations = append(b.mutations, A.X{`=`, 3, val})
	b.logs = append(b.logs, A.X{`facilities`, b.Facilities, val})
	b.Facilities = val
	return true
}

// SetCreatedAt create mutations, should not duplicate
func (b *BuildingsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != b.CreatedAt {
		b.mutations = append(b.mutations, A.X{`=`, 4, val})
		b.logs = append(b.logs, A.X{`createdAt`, b.CreatedAt, val})
		b.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (b *BuildingsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != b.CreatedBy {
		b.mutations = append(b.mutations, A.X{`=`, 5, val})
		b.logs = append(b.logs, A.X{`createdBy`, b.CreatedBy, val})
		b.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (b *BuildingsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != b.UpdatedAt {
		b.mutations = append(b.mutations, A.X{`=`, 6, val})
		b.logs = append(b.logs, A.X{`updatedAt`, b.UpdatedAt, val})
		b.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (b *BuildingsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != b.UpdatedBy {
		b.mutations = append(b.mutations, A.X{`=`, 7, val})
		b.logs = append(b.logs, A.X{`updatedBy`, b.UpdatedBy, val})
		b.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (b *BuildingsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != b.DeletedAt {
		b.mutations = append(b.mutations, A.X{`=`, 8, val})
		b.logs = append(b.logs, A.X{`deletedAt`, b.DeletedAt, val})
		b.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (b *BuildingsMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != b.DeletedBy {
		b.mutations = append(b.mutations, A.X{`=`, 9, val})
		b.logs = append(b.logs, A.X{`deletedBy`, b.DeletedBy, val})
		b.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (b *BuildingsMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != b.RestoredBy {
		b.mutations = append(b.mutations, A.X{`=`, 10, val})
		b.logs = append(b.logs, A.X{`restoredBy`, b.RestoredBy, val})
		b.RestoredBy = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (b *BuildingsMutator) SetAll(from rqProperty.Buildings, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
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
	if !excludeMap[`buildingName`] && (forceMap[`buildingName`] || from.BuildingName != ``) {
		b.BuildingName = S.Trim(from.BuildingName)
		changed = true
	}
	if !excludeMap[`locationId`] && (forceMap[`locationId`] || from.LocationId != 0) {
		b.LocationId = from.LocationId
		changed = true
	}
	if !excludeMap[`facilities`] && (forceMap[`facilities`] || from.Facilities != nil) {
		b.Facilities = from.Facilities
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

// FacilitiesMutator DAO writer/command struct
type FacilitiesMutator struct {
	rqProperty.Facilities
	mutations []A.X
	logs      []A.X
}

// NewFacilitiesMutator create new ORM writer/command object
func NewFacilitiesMutator(adapter *Tt.Adapter) (res *FacilitiesMutator) {
	res = &FacilitiesMutator{Facilities: rqProperty.Facilities{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (f *FacilitiesMutator) Logs() []A.X { //nolint:dupl false positive
	return f.logs
}

// HaveMutation check whether Set* methods ever called
func (f *FacilitiesMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(f.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (f *FacilitiesMutator) ClearMutations() { //nolint:dupl false positive
	f.mutations = []A.X{}
	f.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (f *FacilitiesMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := f.Adapter.Update(f.SpaceName(), f.UniqueIndexId(), A.X{f.Id}, f.ToUpdateArray())
	return !L.IsError(err, `Facilities.DoOverwriteById failed: `+f.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (f *FacilitiesMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !f.HaveMutation() {
		return true
	}
	_, err := f.Adapter.Update(f.SpaceName(), f.UniqueIndexId(), A.X{f.Id}, f.mutations)
	return !L.IsError(err, `Facilities.DoUpdateById failed: `+f.SpaceName())
}

// DoDeletePermanentById permanent delete
func (f *FacilitiesMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := f.Adapter.Delete(f.SpaceName(), f.UniqueIndexId(), A.X{f.Id})
	return !L.IsError(err, `Facilities.DoDeletePermanentById failed: `+f.SpaceName())
}

// func (f *FacilitiesMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := f.ToArray()
//	_, err := f.Adapter.Upsert(f.SpaceName(), arr, A.X{
//		A.X{`=`, 0, f.Id},
//		A.X{`=`, 1, f.FacilityName},
//		A.X{`=`, 2, f.ExtraChargeIDR},
//		A.X{`=`, 3, f.CreatedAt},
//		A.X{`=`, 4, f.CreatedBy},
//		A.X{`=`, 5, f.UpdatedAt},
//		A.X{`=`, 6, f.UpdatedBy},
//		A.X{`=`, 7, f.DeletedAt},
//		A.X{`=`, 8, f.DeletedBy},
//		A.X{`=`, 9, f.RestoredBy},
//		A.X{`=`, 10, f.FacilityType},
//		A.X{`=`, 11, f.DescriptionEN},
//	})
//	return !L.IsError(err, `Facilities.DoUpsert failed: `+f.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (f *FacilitiesMutator) DoInsert() bool { //nolint:dupl false positive
	arr := f.ToArray()
	row, err := f.Adapter.Insert(f.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			f.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Facilities.DoInsert failed: `+f.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (f *FacilitiesMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := f.ToArray()
	row, err := f.Adapter.Replace(f.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			f.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Facilities.DoUpsert failed: `+f.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (f *FacilitiesMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != f.Id {
		f.mutations = append(f.mutations, A.X{`=`, 0, val})
		f.logs = append(f.logs, A.X{`id`, f.Id, val})
		f.Id = val
		return true
	}
	return false
}

// SetFacilityName create mutations, should not duplicate
func (f *FacilitiesMutator) SetFacilityName(val string) bool { //nolint:dupl false positive
	if val != f.FacilityName {
		f.mutations = append(f.mutations, A.X{`=`, 1, val})
		f.logs = append(f.logs, A.X{`facilityName`, f.FacilityName, val})
		f.FacilityName = val
		return true
	}
	return false
}

// SetExtraChargeIDR create mutations, should not duplicate
func (f *FacilitiesMutator) SetExtraChargeIDR(val int64) bool { //nolint:dupl false positive
	if val != f.ExtraChargeIDR {
		f.mutations = append(f.mutations, A.X{`=`, 2, val})
		f.logs = append(f.logs, A.X{`extraChargeIDR`, f.ExtraChargeIDR, val})
		f.ExtraChargeIDR = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (f *FacilitiesMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != f.CreatedAt {
		f.mutations = append(f.mutations, A.X{`=`, 3, val})
		f.logs = append(f.logs, A.X{`createdAt`, f.CreatedAt, val})
		f.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (f *FacilitiesMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != f.CreatedBy {
		f.mutations = append(f.mutations, A.X{`=`, 4, val})
		f.logs = append(f.logs, A.X{`createdBy`, f.CreatedBy, val})
		f.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (f *FacilitiesMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != f.UpdatedAt {
		f.mutations = append(f.mutations, A.X{`=`, 5, val})
		f.logs = append(f.logs, A.X{`updatedAt`, f.UpdatedAt, val})
		f.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (f *FacilitiesMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != f.UpdatedBy {
		f.mutations = append(f.mutations, A.X{`=`, 6, val})
		f.logs = append(f.logs, A.X{`updatedBy`, f.UpdatedBy, val})
		f.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (f *FacilitiesMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != f.DeletedAt {
		f.mutations = append(f.mutations, A.X{`=`, 7, val})
		f.logs = append(f.logs, A.X{`deletedAt`, f.DeletedAt, val})
		f.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (f *FacilitiesMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != f.DeletedBy {
		f.mutations = append(f.mutations, A.X{`=`, 8, val})
		f.logs = append(f.logs, A.X{`deletedBy`, f.DeletedBy, val})
		f.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (f *FacilitiesMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != f.RestoredBy {
		f.mutations = append(f.mutations, A.X{`=`, 9, val})
		f.logs = append(f.logs, A.X{`restoredBy`, f.RestoredBy, val})
		f.RestoredBy = val
		return true
	}
	return false
}

// SetFacilityType create mutations, should not duplicate
func (f *FacilitiesMutator) SetFacilityType(val string) bool { //nolint:dupl false positive
	if val != f.FacilityType {
		f.mutations = append(f.mutations, A.X{`=`, 10, val})
		f.logs = append(f.logs, A.X{`facilityType`, f.FacilityType, val})
		f.FacilityType = val
		return true
	}
	return false
}

// SetDescriptionEN create mutations, should not duplicate
func (f *FacilitiesMutator) SetDescriptionEN(val string) bool { //nolint:dupl false positive
	if val != f.DescriptionEN {
		f.mutations = append(f.mutations, A.X{`=`, 11, val})
		f.logs = append(f.logs, A.X{`descriptionEN`, f.DescriptionEN, val})
		f.DescriptionEN = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (f *FacilitiesMutator) SetAll(from rqProperty.Facilities, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		f.Id = from.Id
		changed = true
	}
	if !excludeMap[`facilityName`] && (forceMap[`facilityName`] || from.FacilityName != ``) {
		f.FacilityName = S.Trim(from.FacilityName)
		changed = true
	}
	if !excludeMap[`extraChargeIDR`] && (forceMap[`extraChargeIDR`] || from.ExtraChargeIDR != 0) {
		f.ExtraChargeIDR = from.ExtraChargeIDR
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		f.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		f.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		f.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		f.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		f.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		f.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		f.RestoredBy = from.RestoredBy
		changed = true
	}
	if !excludeMap[`facilityType`] && (forceMap[`facilityType`] || from.FacilityType != ``) {
		f.FacilityType = S.Trim(from.FacilityType)
		changed = true
	}
	if !excludeMap[`descriptionEN`] && (forceMap[`descriptionEN`] || from.DescriptionEN != ``) {
		f.DescriptionEN = S.Trim(from.DescriptionEN)
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// LocationsMutator DAO writer/command struct
type LocationsMutator struct {
	rqProperty.Locations
	mutations []A.X
	logs      []A.X
}

// NewLocationsMutator create new ORM writer/command object
func NewLocationsMutator(adapter *Tt.Adapter) (res *LocationsMutator) {
	res = &LocationsMutator{Locations: rqProperty.Locations{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (l *LocationsMutator) Logs() []A.X { //nolint:dupl false positive
	return l.logs
}

// HaveMutation check whether Set* methods ever called
func (l *LocationsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(l.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (l *LocationsMutator) ClearMutations() { //nolint:dupl false positive
	l.mutations = []A.X{}
	l.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (l *LocationsMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := l.Adapter.Update(l.SpaceName(), l.UniqueIndexId(), A.X{l.Id}, l.ToUpdateArray())
	return !L.IsError(err, `Locations.DoOverwriteById failed: `+l.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (l *LocationsMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !l.HaveMutation() {
		return true
	}
	_, err := l.Adapter.Update(l.SpaceName(), l.UniqueIndexId(), A.X{l.Id}, l.mutations)
	return !L.IsError(err, `Locations.DoUpdateById failed: `+l.SpaceName())
}

// DoDeletePermanentById permanent delete
func (l *LocationsMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := l.Adapter.Delete(l.SpaceName(), l.UniqueIndexId(), A.X{l.Id})
	return !L.IsError(err, `Locations.DoDeletePermanentById failed: `+l.SpaceName())
}

// func (l *LocationsMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := l.ToArray()
//	_, err := l.Adapter.Upsert(l.SpaceName(), arr, A.X{
//		A.X{`=`, 0, l.Id},
//		A.X{`=`, 1, l.Name},
//		A.X{`=`, 2, l.Address},
//		A.X{`=`, 3, l.GmapLocation},
//		A.X{`=`, 4, l.CreatedAt},
//		A.X{`=`, 5, l.CreatedBy},
//		A.X{`=`, 6, l.UpdatedAt},
//		A.X{`=`, 7, l.UpdatedBy},
//		A.X{`=`, 8, l.DeletedAt},
//		A.X{`=`, 9, l.DeletedBy},
//		A.X{`=`, 10, l.RestoredBy},
//	})
//	return !L.IsError(err, `Locations.DoUpsert failed: `+l.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (l *LocationsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := l.ToArray()
	row, err := l.Adapter.Insert(l.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			l.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Locations.DoInsert failed: `+l.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (l *LocationsMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := l.ToArray()
	row, err := l.Adapter.Replace(l.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			l.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Locations.DoUpsert failed: `+l.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (l *LocationsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != l.Id {
		l.mutations = append(l.mutations, A.X{`=`, 0, val})
		l.logs = append(l.logs, A.X{`id`, l.Id, val})
		l.Id = val
		return true
	}
	return false
}

// SetName create mutations, should not duplicate
func (l *LocationsMutator) SetName(val string) bool { //nolint:dupl false positive
	if val != l.Name {
		l.mutations = append(l.mutations, A.X{`=`, 1, val})
		l.logs = append(l.logs, A.X{`name`, l.Name, val})
		l.Name = val
		return true
	}
	return false
}

// SetAddress create mutations, should not duplicate
func (l *LocationsMutator) SetAddress(val string) bool { //nolint:dupl false positive
	if val != l.Address {
		l.mutations = append(l.mutations, A.X{`=`, 2, val})
		l.logs = append(l.logs, A.X{`address`, l.Address, val})
		l.Address = val
		return true
	}
	return false
}

// SetGmapLocation create mutations, should not duplicate
func (l *LocationsMutator) SetGmapLocation(val string) bool { //nolint:dupl false positive
	if val != l.GmapLocation {
		l.mutations = append(l.mutations, A.X{`=`, 3, val})
		l.logs = append(l.logs, A.X{`gmapLocation`, l.GmapLocation, val})
		l.GmapLocation = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (l *LocationsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != l.CreatedAt {
		l.mutations = append(l.mutations, A.X{`=`, 4, val})
		l.logs = append(l.logs, A.X{`createdAt`, l.CreatedAt, val})
		l.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (l *LocationsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != l.CreatedBy {
		l.mutations = append(l.mutations, A.X{`=`, 5, val})
		l.logs = append(l.logs, A.X{`createdBy`, l.CreatedBy, val})
		l.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (l *LocationsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != l.UpdatedAt {
		l.mutations = append(l.mutations, A.X{`=`, 6, val})
		l.logs = append(l.logs, A.X{`updatedAt`, l.UpdatedAt, val})
		l.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (l *LocationsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != l.UpdatedBy {
		l.mutations = append(l.mutations, A.X{`=`, 7, val})
		l.logs = append(l.logs, A.X{`updatedBy`, l.UpdatedBy, val})
		l.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (l *LocationsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != l.DeletedAt {
		l.mutations = append(l.mutations, A.X{`=`, 8, val})
		l.logs = append(l.logs, A.X{`deletedAt`, l.DeletedAt, val})
		l.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (l *LocationsMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != l.DeletedBy {
		l.mutations = append(l.mutations, A.X{`=`, 9, val})
		l.logs = append(l.logs, A.X{`deletedBy`, l.DeletedBy, val})
		l.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (l *LocationsMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != l.RestoredBy {
		l.mutations = append(l.mutations, A.X{`=`, 10, val})
		l.logs = append(l.logs, A.X{`restoredBy`, l.RestoredBy, val})
		l.RestoredBy = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (l *LocationsMutator) SetAll(from rqProperty.Locations, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
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
	if !excludeMap[`name`] && (forceMap[`name`] || from.Name != ``) {
		l.Name = S.Trim(from.Name)
		changed = true
	}
	if !excludeMap[`address`] && (forceMap[`address`] || from.Address != ``) {
		l.Address = S.Trim(from.Address)
		changed = true
	}
	if !excludeMap[`gmapLocation`] && (forceMap[`gmapLocation`] || from.GmapLocation != ``) {
		l.GmapLocation = S.Trim(from.GmapLocation)
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
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// PaymentsMutator DAO writer/command struct
type PaymentsMutator struct {
	rqProperty.Payments
	mutations []A.X
	logs      []A.X
}

// NewPaymentsMutator create new ORM writer/command object
func NewPaymentsMutator(adapter *Tt.Adapter) (res *PaymentsMutator) {
	res = &PaymentsMutator{Payments: rqProperty.Payments{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (p *PaymentsMutator) Logs() []A.X { //nolint:dupl false positive
	return p.logs
}

// HaveMutation check whether Set* methods ever called
func (p *PaymentsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(p.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (p *PaymentsMutator) ClearMutations() { //nolint:dupl false positive
	p.mutations = []A.X{}
	p.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (p *PaymentsMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id}, p.ToUpdateArray())
	return !L.IsError(err, `Payments.DoOverwriteById failed: `+p.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (p *PaymentsMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !p.HaveMutation() {
		return true
	}
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id}, p.mutations)
	return !L.IsError(err, `Payments.DoUpdateById failed: `+p.SpaceName())
}

// DoDeletePermanentById permanent delete
func (p *PaymentsMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := p.Adapter.Delete(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id})
	return !L.IsError(err, `Payments.DoDeletePermanentById failed: `+p.SpaceName())
}

// func (p *PaymentsMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := p.ToArray()
//	_, err := p.Adapter.Upsert(p.SpaceName(), arr, A.X{
//		A.X{`=`, 0, p.Id},
//		A.X{`=`, 1, p.BookingId},
//		A.X{`=`, 2, p.PaymentAt},
//		A.X{`=`, 3, p.PaidIDR},
//		A.X{`=`, 4, p.PaymentMethod},
//		A.X{`=`, 5, p.PaymentStatus},
//		A.X{`=`, 6, p.Note},
//		A.X{`=`, 7, p.CreatedAt},
//		A.X{`=`, 8, p.CreatedBy},
//		A.X{`=`, 9, p.UpdatedAt},
//		A.X{`=`, 10, p.UpdatedBy},
//		A.X{`=`, 11, p.DeletedAt},
//		A.X{`=`, 12, p.DeletedBy},
//		A.X{`=`, 13, p.RestoredBy},
//	})
//	return !L.IsError(err, `Payments.DoUpsert failed: `+p.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (p *PaymentsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := p.ToArray()
	row, err := p.Adapter.Insert(p.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			p.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Payments.DoInsert failed: `+p.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (p *PaymentsMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := p.ToArray()
	row, err := p.Adapter.Replace(p.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			p.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Payments.DoUpsert failed: `+p.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (p *PaymentsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != p.Id {
		p.mutations = append(p.mutations, A.X{`=`, 0, val})
		p.logs = append(p.logs, A.X{`id`, p.Id, val})
		p.Id = val
		return true
	}
	return false
}

// SetBookingId create mutations, should not duplicate
func (p *PaymentsMutator) SetBookingId(val uint64) bool { //nolint:dupl false positive
	if val != p.BookingId {
		p.mutations = append(p.mutations, A.X{`=`, 1, val})
		p.logs = append(p.logs, A.X{`bookingId`, p.BookingId, val})
		p.BookingId = val
		return true
	}
	return false
}

// SetPaymentAt create mutations, should not duplicate
func (p *PaymentsMutator) SetPaymentAt(val string) bool { //nolint:dupl false positive
	if val != p.PaymentAt {
		p.mutations = append(p.mutations, A.X{`=`, 2, val})
		p.logs = append(p.logs, A.X{`paymentAt`, p.PaymentAt, val})
		p.PaymentAt = val
		return true
	}
	return false
}

// SetPaidIDR create mutations, should not duplicate
func (p *PaymentsMutator) SetPaidIDR(val int64) bool { //nolint:dupl false positive
	if val != p.PaidIDR {
		p.mutations = append(p.mutations, A.X{`=`, 3, val})
		p.logs = append(p.logs, A.X{`paidIDR`, p.PaidIDR, val})
		p.PaidIDR = val
		return true
	}
	return false
}

// SetPaymentMethod create mutations, should not duplicate
func (p *PaymentsMutator) SetPaymentMethod(val string) bool { //nolint:dupl false positive
	if val != p.PaymentMethod {
		p.mutations = append(p.mutations, A.X{`=`, 4, val})
		p.logs = append(p.logs, A.X{`paymentMethod`, p.PaymentMethod, val})
		p.PaymentMethod = val
		return true
	}
	return false
}

// SetPaymentStatus create mutations, should not duplicate
func (p *PaymentsMutator) SetPaymentStatus(val string) bool { //nolint:dupl false positive
	if val != p.PaymentStatus {
		p.mutations = append(p.mutations, A.X{`=`, 5, val})
		p.logs = append(p.logs, A.X{`paymentStatus`, p.PaymentStatus, val})
		p.PaymentStatus = val
		return true
	}
	return false
}

// SetNote create mutations, should not duplicate
func (p *PaymentsMutator) SetNote(val string) bool { //nolint:dupl false positive
	if val != p.Note {
		p.mutations = append(p.mutations, A.X{`=`, 6, val})
		p.logs = append(p.logs, A.X{`note`, p.Note, val})
		p.Note = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (p *PaymentsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.CreatedAt {
		p.mutations = append(p.mutations, A.X{`=`, 7, val})
		p.logs = append(p.logs, A.X{`createdAt`, p.CreatedAt, val})
		p.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (p *PaymentsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.CreatedBy {
		p.mutations = append(p.mutations, A.X{`=`, 8, val})
		p.logs = append(p.logs, A.X{`createdBy`, p.CreatedBy, val})
		p.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (p *PaymentsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.UpdatedAt {
		p.mutations = append(p.mutations, A.X{`=`, 9, val})
		p.logs = append(p.logs, A.X{`updatedAt`, p.UpdatedAt, val})
		p.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (p *PaymentsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.UpdatedBy {
		p.mutations = append(p.mutations, A.X{`=`, 10, val})
		p.logs = append(p.logs, A.X{`updatedBy`, p.UpdatedBy, val})
		p.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (p *PaymentsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != p.DeletedAt {
		p.mutations = append(p.mutations, A.X{`=`, 11, val})
		p.logs = append(p.logs, A.X{`deletedAt`, p.DeletedAt, val})
		p.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (p *PaymentsMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.DeletedBy {
		p.mutations = append(p.mutations, A.X{`=`, 12, val})
		p.logs = append(p.logs, A.X{`deletedBy`, p.DeletedBy, val})
		p.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (p *PaymentsMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != p.RestoredBy {
		p.mutations = append(p.mutations, A.X{`=`, 13, val})
		p.logs = append(p.logs, A.X{`restoredBy`, p.RestoredBy, val})
		p.RestoredBy = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (p *PaymentsMutator) SetAll(from rqProperty.Payments, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		p.Id = from.Id
		changed = true
	}
	if !excludeMap[`bookingId`] && (forceMap[`bookingId`] || from.BookingId != 0) {
		p.BookingId = from.BookingId
		changed = true
	}
	if !excludeMap[`paymentAt`] && (forceMap[`paymentAt`] || from.PaymentAt != ``) {
		p.PaymentAt = S.Trim(from.PaymentAt)
		changed = true
	}
	if !excludeMap[`paidIDR`] && (forceMap[`paidIDR`] || from.PaidIDR != 0) {
		p.PaidIDR = from.PaidIDR
		changed = true
	}
	if !excludeMap[`paymentMethod`] && (forceMap[`paymentMethod`] || from.PaymentMethod != ``) {
		p.PaymentMethod = S.Trim(from.PaymentMethod)
		changed = true
	}
	if !excludeMap[`paymentStatus`] && (forceMap[`paymentStatus`] || from.PaymentStatus != ``) {
		p.PaymentStatus = S.Trim(from.PaymentStatus)
		changed = true
	}
	if !excludeMap[`note`] && (forceMap[`note`] || from.Note != ``) {
		p.Note = S.Trim(from.Note)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		p.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		p.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		p.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		p.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		p.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		p.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		p.RestoredBy = from.RestoredBy
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// RoomsMutator DAO writer/command struct
type RoomsMutator struct {
	rqProperty.Rooms
	mutations []A.X
	logs      []A.X
}

// NewRoomsMutator create new ORM writer/command object
func NewRoomsMutator(adapter *Tt.Adapter) (res *RoomsMutator) {
	res = &RoomsMutator{Rooms: rqProperty.Rooms{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (r *RoomsMutator) Logs() []A.X { //nolint:dupl false positive
	return r.logs
}

// HaveMutation check whether Set* methods ever called
func (r *RoomsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(r.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (r *RoomsMutator) ClearMutations() { //nolint:dupl false positive
	r.mutations = []A.X{}
	r.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (r *RoomsMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := r.Adapter.Update(r.SpaceName(), r.UniqueIndexId(), A.X{r.Id}, r.ToUpdateArray())
	return !L.IsError(err, `Rooms.DoOverwriteById failed: `+r.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (r *RoomsMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !r.HaveMutation() {
		return true
	}
	_, err := r.Adapter.Update(r.SpaceName(), r.UniqueIndexId(), A.X{r.Id}, r.mutations)
	return !L.IsError(err, `Rooms.DoUpdateById failed: `+r.SpaceName())
}

// DoDeletePermanentById permanent delete
func (r *RoomsMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := r.Adapter.Delete(r.SpaceName(), r.UniqueIndexId(), A.X{r.Id})
	return !L.IsError(err, `Rooms.DoDeletePermanentById failed: `+r.SpaceName())
}

// func (r *RoomsMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := r.ToArray()
//	_, err := r.Adapter.Upsert(r.SpaceName(), arr, A.X{
//		A.X{`=`, 0, r.Id},
//		A.X{`=`, 1, r.RoomName},
//		A.X{`=`, 2, r.BasePriceIDR},
//		A.X{`=`, 3, r.CurrentTenantId},
//		A.X{`=`, 4, r.FirstUseAt},
//		A.X{`=`, 5, r.CreatedAt},
//		A.X{`=`, 6, r.CreatedBy},
//		A.X{`=`, 7, r.UpdatedAt},
//		A.X{`=`, 8, r.UpdatedBy},
//		A.X{`=`, 9, r.DeletedAt},
//		A.X{`=`, 10, r.DeletedBy},
//		A.X{`=`, 11, r.RestoredBy},
//		A.X{`=`, 12, r.BuildingId},
//		A.X{`=`, 13, r.LastUseAt},
//		A.X{`=`, 14, r.RoomSize},
//		A.X{`=`, 15, r.ImageUrl},
//	})
//	return !L.IsError(err, `Rooms.DoUpsert failed: `+r.SpaceName()+ `\n%#v`, arr)
// }

// DoOverwriteByRoomName update all columns, error if not exists, not using mutations/Set*
func (r *RoomsMutator) DoOverwriteByRoomName() bool { //nolint:dupl false positive
	_, err := r.Adapter.Update(r.SpaceName(), r.UniqueIndexRoomName(), A.X{r.RoomName}, r.ToUpdateArray())
	return !L.IsError(err, `Rooms.DoOverwriteByRoomName failed: `+r.SpaceName())
}

// DoUpdateByRoomName update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (r *RoomsMutator) DoUpdateByRoomName() bool { //nolint:dupl false positive
	if !r.HaveMutation() {
		return true
	}
	_, err := r.Adapter.Update(r.SpaceName(), r.UniqueIndexRoomName(), A.X{r.RoomName}, r.mutations)
	return !L.IsError(err, `Rooms.DoUpdateByRoomName failed: `+r.SpaceName())
}

// DoDeletePermanentByRoomName permanent delete
func (r *RoomsMutator) DoDeletePermanentByRoomName() bool { //nolint:dupl false positive
	_, err := r.Adapter.Delete(r.SpaceName(), r.UniqueIndexRoomName(), A.X{r.RoomName})
	return !L.IsError(err, `Rooms.DoDeletePermanentByRoomName failed: `+r.SpaceName())
}

// DoInsert insert, error if already exists
func (r *RoomsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := r.ToArray()
	row, err := r.Adapter.Insert(r.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			r.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Rooms.DoInsert failed: `+r.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (r *RoomsMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := r.ToArray()
	row, err := r.Adapter.Replace(r.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			r.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Rooms.DoUpsert failed: `+r.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (r *RoomsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != r.Id {
		r.mutations = append(r.mutations, A.X{`=`, 0, val})
		r.logs = append(r.logs, A.X{`id`, r.Id, val})
		r.Id = val
		return true
	}
	return false
}

// SetRoomName create mutations, should not duplicate
func (r *RoomsMutator) SetRoomName(val string) bool { //nolint:dupl false positive
	if val != r.RoomName {
		r.mutations = append(r.mutations, A.X{`=`, 1, val})
		r.logs = append(r.logs, A.X{`roomName`, r.RoomName, val})
		r.RoomName = val
		return true
	}
	return false
}

// SetBasePriceIDR create mutations, should not duplicate
func (r *RoomsMutator) SetBasePriceIDR(val int64) bool { //nolint:dupl false positive
	if val != r.BasePriceIDR {
		r.mutations = append(r.mutations, A.X{`=`, 2, val})
		r.logs = append(r.logs, A.X{`basePriceIDR`, r.BasePriceIDR, val})
		r.BasePriceIDR = val
		return true
	}
	return false
}

// SetCurrentTenantId create mutations, should not duplicate
func (r *RoomsMutator) SetCurrentTenantId(val uint64) bool { //nolint:dupl false positive
	if val != r.CurrentTenantId {
		r.mutations = append(r.mutations, A.X{`=`, 3, val})
		r.logs = append(r.logs, A.X{`currentTenantId`, r.CurrentTenantId, val})
		r.CurrentTenantId = val
		return true
	}
	return false
}

// SetFirstUseAt create mutations, should not duplicate
func (r *RoomsMutator) SetFirstUseAt(val string) bool { //nolint:dupl false positive
	if val != r.FirstUseAt {
		r.mutations = append(r.mutations, A.X{`=`, 4, val})
		r.logs = append(r.logs, A.X{`firstUseAt`, r.FirstUseAt, val})
		r.FirstUseAt = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (r *RoomsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != r.CreatedAt {
		r.mutations = append(r.mutations, A.X{`=`, 5, val})
		r.logs = append(r.logs, A.X{`createdAt`, r.CreatedAt, val})
		r.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (r *RoomsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != r.CreatedBy {
		r.mutations = append(r.mutations, A.X{`=`, 6, val})
		r.logs = append(r.logs, A.X{`createdBy`, r.CreatedBy, val})
		r.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (r *RoomsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != r.UpdatedAt {
		r.mutations = append(r.mutations, A.X{`=`, 7, val})
		r.logs = append(r.logs, A.X{`updatedAt`, r.UpdatedAt, val})
		r.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (r *RoomsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != r.UpdatedBy {
		r.mutations = append(r.mutations, A.X{`=`, 8, val})
		r.logs = append(r.logs, A.X{`updatedBy`, r.UpdatedBy, val})
		r.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (r *RoomsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != r.DeletedAt {
		r.mutations = append(r.mutations, A.X{`=`, 9, val})
		r.logs = append(r.logs, A.X{`deletedAt`, r.DeletedAt, val})
		r.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (r *RoomsMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != r.DeletedBy {
		r.mutations = append(r.mutations, A.X{`=`, 10, val})
		r.logs = append(r.logs, A.X{`deletedBy`, r.DeletedBy, val})
		r.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (r *RoomsMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != r.RestoredBy {
		r.mutations = append(r.mutations, A.X{`=`, 11, val})
		r.logs = append(r.logs, A.X{`restoredBy`, r.RestoredBy, val})
		r.RestoredBy = val
		return true
	}
	return false
}

// SetBuildingId create mutations, should not duplicate
func (r *RoomsMutator) SetBuildingId(val uint64) bool { //nolint:dupl false positive
	if val != r.BuildingId {
		r.mutations = append(r.mutations, A.X{`=`, 12, val})
		r.logs = append(r.logs, A.X{`buildingId`, r.BuildingId, val})
		r.BuildingId = val
		return true
	}
	return false
}

// SetLastUseAt create mutations, should not duplicate
func (r *RoomsMutator) SetLastUseAt(val string) bool { //nolint:dupl false positive
	if val != r.LastUseAt {
		r.mutations = append(r.mutations, A.X{`=`, 13, val})
		r.logs = append(r.logs, A.X{`lastUseAt`, r.LastUseAt, val})
		r.LastUseAt = val
		return true
	}
	return false
}

// SetRoomSize create mutations, should not duplicate
func (r *RoomsMutator) SetRoomSize(val string) bool { //nolint:dupl false positive
	if val != r.RoomSize {
		r.mutations = append(r.mutations, A.X{`=`, 14, val})
		r.logs = append(r.logs, A.X{`roomSize`, r.RoomSize, val})
		r.RoomSize = val
		return true
	}
	return false
}

// SetImageUrl create mutations, should not duplicate
func (r *RoomsMutator) SetImageUrl(val string) bool { //nolint:dupl false positive
	if val != r.ImageUrl {
		r.mutations = append(r.mutations, A.X{`=`, 15, val})
		r.logs = append(r.logs, A.X{`imageUrl`, r.ImageUrl, val})
		r.ImageUrl = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (r *RoomsMutator) SetAll(from rqProperty.Rooms, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		r.Id = from.Id
		changed = true
	}
	if !excludeMap[`roomName`] && (forceMap[`roomName`] || from.RoomName != ``) {
		r.RoomName = S.Trim(from.RoomName)
		changed = true
	}
	if !excludeMap[`basePriceIDR`] && (forceMap[`basePriceIDR`] || from.BasePriceIDR != 0) {
		r.BasePriceIDR = from.BasePriceIDR
		changed = true
	}
	if !excludeMap[`currentTenantId`] && (forceMap[`currentTenantId`] || from.CurrentTenantId != 0) {
		r.CurrentTenantId = from.CurrentTenantId
		changed = true
	}
	if !excludeMap[`firstUseAt`] && (forceMap[`firstUseAt`] || from.FirstUseAt != ``) {
		r.FirstUseAt = S.Trim(from.FirstUseAt)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		r.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		r.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		r.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		r.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		r.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		r.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		r.RestoredBy = from.RestoredBy
		changed = true
	}
	if !excludeMap[`buildingId`] && (forceMap[`buildingId`] || from.BuildingId != 0) {
		r.BuildingId = from.BuildingId
		changed = true
	}
	if !excludeMap[`lastUseAt`] && (forceMap[`lastUseAt`] || from.LastUseAt != ``) {
		r.LastUseAt = S.Trim(from.LastUseAt)
		changed = true
	}
	if !excludeMap[`roomSize`] && (forceMap[`roomSize`] || from.RoomSize != ``) {
		r.RoomSize = S.Trim(from.RoomSize)
		changed = true
	}
	if !excludeMap[`imageUrl`] && (forceMap[`imageUrl`] || from.ImageUrl != ``) {
		r.ImageUrl = S.Trim(from.ImageUrl)
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// StocksMutator DAO writer/command struct
type StocksMutator struct {
	rqProperty.Stocks
	mutations []A.X
	logs      []A.X
}

// NewStocksMutator create new ORM writer/command object
func NewStocksMutator(adapter *Tt.Adapter) (res *StocksMutator) {
	res = &StocksMutator{Stocks: rqProperty.Stocks{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (s *StocksMutator) Logs() []A.X { //nolint:dupl false positive
	return s.logs
}

// HaveMutation check whether Set* methods ever called
func (s *StocksMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(s.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (s *StocksMutator) ClearMutations() { //nolint:dupl false positive
	s.mutations = []A.X{}
	s.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (s *StocksMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := s.Adapter.Update(s.SpaceName(), s.UniqueIndexId(), A.X{s.Id}, s.ToUpdateArray())
	return !L.IsError(err, `Stocks.DoOverwriteById failed: `+s.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (s *StocksMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !s.HaveMutation() {
		return true
	}
	_, err := s.Adapter.Update(s.SpaceName(), s.UniqueIndexId(), A.X{s.Id}, s.mutations)
	return !L.IsError(err, `Stocks.DoUpdateById failed: `+s.SpaceName())
}

// DoDeletePermanentById permanent delete
func (s *StocksMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := s.Adapter.Delete(s.SpaceName(), s.UniqueIndexId(), A.X{s.Id})
	return !L.IsError(err, `Stocks.DoDeletePermanentById failed: `+s.SpaceName())
}

// func (s *StocksMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := s.ToArray()
//	_, err := s.Adapter.Upsert(s.SpaceName(), arr, A.X{
//		A.X{`=`, 0, s.Id},
//		A.X{`=`, 1, s.StockName},
//		A.X{`=`, 2, s.StockAddedAt},
//		A.X{`=`, 3, s.Quantity},
//		A.X{`=`, 4, s.PriceIDR},
//		A.X{`=`, 5, s.CreatedAt},
//		A.X{`=`, 6, s.CreatedBy},
//		A.X{`=`, 7, s.UpdatedAt},
//		A.X{`=`, 8, s.UpdatedBy},
//		A.X{`=`, 9, s.DeletedAt},
//		A.X{`=`, 10, s.DeletedBy},
//		A.X{`=`, 11, s.RestoredBy},
//	})
//	return !L.IsError(err, `Stocks.DoUpsert failed: `+s.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (s *StocksMutator) DoInsert() bool { //nolint:dupl false positive
	arr := s.ToArray()
	row, err := s.Adapter.Insert(s.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			s.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Stocks.DoInsert failed: `+s.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (s *StocksMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := s.ToArray()
	row, err := s.Adapter.Replace(s.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			s.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Stocks.DoUpsert failed: `+s.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (s *StocksMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != s.Id {
		s.mutations = append(s.mutations, A.X{`=`, 0, val})
		s.logs = append(s.logs, A.X{`id`, s.Id, val})
		s.Id = val
		return true
	}
	return false
}

// SetStockName create mutations, should not duplicate
func (s *StocksMutator) SetStockName(val string) bool { //nolint:dupl false positive
	if val != s.StockName {
		s.mutations = append(s.mutations, A.X{`=`, 1, val})
		s.logs = append(s.logs, A.X{`stockName`, s.StockName, val})
		s.StockName = val
		return true
	}
	return false
}

// SetStockAddedAt create mutations, should not duplicate
func (s *StocksMutator) SetStockAddedAt(val string) bool { //nolint:dupl false positive
	if val != s.StockAddedAt {
		s.mutations = append(s.mutations, A.X{`=`, 2, val})
		s.logs = append(s.logs, A.X{`stockAddedAt`, s.StockAddedAt, val})
		s.StockAddedAt = val
		return true
	}
	return false
}

// SetQuantity create mutations, should not duplicate
func (s *StocksMutator) SetQuantity(val int64) bool { //nolint:dupl false positive
	if val != s.Quantity {
		s.mutations = append(s.mutations, A.X{`=`, 3, val})
		s.logs = append(s.logs, A.X{`quantity`, s.Quantity, val})
		s.Quantity = val
		return true
	}
	return false
}

// SetPriceIDR create mutations, should not duplicate
func (s *StocksMutator) SetPriceIDR(val int64) bool { //nolint:dupl false positive
	if val != s.PriceIDR {
		s.mutations = append(s.mutations, A.X{`=`, 4, val})
		s.logs = append(s.logs, A.X{`priceIDR`, s.PriceIDR, val})
		s.PriceIDR = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (s *StocksMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != s.CreatedAt {
		s.mutations = append(s.mutations, A.X{`=`, 5, val})
		s.logs = append(s.logs, A.X{`createdAt`, s.CreatedAt, val})
		s.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (s *StocksMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != s.CreatedBy {
		s.mutations = append(s.mutations, A.X{`=`, 6, val})
		s.logs = append(s.logs, A.X{`createdBy`, s.CreatedBy, val})
		s.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (s *StocksMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != s.UpdatedAt {
		s.mutations = append(s.mutations, A.X{`=`, 7, val})
		s.logs = append(s.logs, A.X{`updatedAt`, s.UpdatedAt, val})
		s.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (s *StocksMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != s.UpdatedBy {
		s.mutations = append(s.mutations, A.X{`=`, 8, val})
		s.logs = append(s.logs, A.X{`updatedBy`, s.UpdatedBy, val})
		s.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (s *StocksMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != s.DeletedAt {
		s.mutations = append(s.mutations, A.X{`=`, 9, val})
		s.logs = append(s.logs, A.X{`deletedAt`, s.DeletedAt, val})
		s.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (s *StocksMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != s.DeletedBy {
		s.mutations = append(s.mutations, A.X{`=`, 10, val})
		s.logs = append(s.logs, A.X{`deletedBy`, s.DeletedBy, val})
		s.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (s *StocksMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != s.RestoredBy {
		s.mutations = append(s.mutations, A.X{`=`, 11, val})
		s.logs = append(s.logs, A.X{`restoredBy`, s.RestoredBy, val})
		s.RestoredBy = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (s *StocksMutator) SetAll(from rqProperty.Stocks, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
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
	if !excludeMap[`stockName`] && (forceMap[`stockName`] || from.StockName != ``) {
		s.StockName = S.Trim(from.StockName)
		changed = true
	}
	if !excludeMap[`stockAddedAt`] && (forceMap[`stockAddedAt`] || from.StockAddedAt != ``) {
		s.StockAddedAt = S.Trim(from.StockAddedAt)
		changed = true
	}
	if !excludeMap[`quantity`] && (forceMap[`quantity`] || from.Quantity != 0) {
		s.Quantity = from.Quantity
		changed = true
	}
	if !excludeMap[`priceIDR`] && (forceMap[`priceIDR`] || from.PriceIDR != 0) {
		s.PriceIDR = from.PriceIDR
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

// WifiDevicesMutator DAO writer/command struct
type WifiDevicesMutator struct {
	rqProperty.WifiDevices
	mutations []A.X
	logs      []A.X
}

// NewWifiDevicesMutator create new ORM writer/command object
func NewWifiDevicesMutator(adapter *Tt.Adapter) (res *WifiDevicesMutator) {
	res = &WifiDevicesMutator{WifiDevices: rqProperty.WifiDevices{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (w *WifiDevicesMutator) Logs() []A.X { //nolint:dupl false positive
	return w.logs
}

// HaveMutation check whether Set* methods ever called
func (w *WifiDevicesMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(w.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (w *WifiDevicesMutator) ClearMutations() { //nolint:dupl false positive
	w.mutations = []A.X{}
	w.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (w *WifiDevicesMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := w.Adapter.Update(w.SpaceName(), w.UniqueIndexId(), A.X{w.Id}, w.ToUpdateArray())
	return !L.IsError(err, `WifiDevices.DoOverwriteById failed: `+w.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (w *WifiDevicesMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !w.HaveMutation() {
		return true
	}
	_, err := w.Adapter.Update(w.SpaceName(), w.UniqueIndexId(), A.X{w.Id}, w.mutations)
	return !L.IsError(err, `WifiDevices.DoUpdateById failed: `+w.SpaceName())
}

// DoDeletePermanentById permanent delete
func (w *WifiDevicesMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := w.Adapter.Delete(w.SpaceName(), w.UniqueIndexId(), A.X{w.Id})
	return !L.IsError(err, `WifiDevices.DoDeletePermanentById failed: `+w.SpaceName())
}

// func (w *WifiDevicesMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := w.ToArray()
//	_, err := w.Adapter.Upsert(w.SpaceName(), arr, A.X{
//		A.X{`=`, 0, w.Id},
//		A.X{`=`, 1, w.StartAt},
//		A.X{`=`, 2, w.EndAt},
//		A.X{`=`, 3, w.PaidAt},
//		A.X{`=`, 4, w.PriceIDR},
//		A.X{`=`, 5, w.TenantId},
//		A.X{`=`, 6, w.MacAddress},
//		A.X{`=`, 7, w.RoomId},
//		A.X{`=`, 8, w.CreatedAt},
//		A.X{`=`, 9, w.CreatedBy},
//		A.X{`=`, 10, w.UpdatedAt},
//		A.X{`=`, 11, w.UpdatedBy},
//		A.X{`=`, 12, w.DeletedAt},
//		A.X{`=`, 13, w.DeletedBy},
//		A.X{`=`, 14, w.RestoredBy},
//	})
//	return !L.IsError(err, `WifiDevices.DoUpsert failed: `+w.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (w *WifiDevicesMutator) DoInsert() bool { //nolint:dupl false positive
	arr := w.ToArray()
	row, err := w.Adapter.Insert(w.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			w.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `WifiDevices.DoInsert failed: `+w.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (w *WifiDevicesMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := w.ToArray()
	row, err := w.Adapter.Replace(w.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			w.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `WifiDevices.DoUpsert failed: `+w.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (w *WifiDevicesMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != w.Id {
		w.mutations = append(w.mutations, A.X{`=`, 0, val})
		w.logs = append(w.logs, A.X{`id`, w.Id, val})
		w.Id = val
		return true
	}
	return false
}

// SetStartAt create mutations, should not duplicate
func (w *WifiDevicesMutator) SetStartAt(val int64) bool { //nolint:dupl false positive
	if val != w.StartAt {
		w.mutations = append(w.mutations, A.X{`=`, 1, val})
		w.logs = append(w.logs, A.X{`startAt`, w.StartAt, val})
		w.StartAt = val
		return true
	}
	return false
}

// SetEndAt create mutations, should not duplicate
func (w *WifiDevicesMutator) SetEndAt(val int64) bool { //nolint:dupl false positive
	if val != w.EndAt {
		w.mutations = append(w.mutations, A.X{`=`, 2, val})
		w.logs = append(w.logs, A.X{`endAt`, w.EndAt, val})
		w.EndAt = val
		return true
	}
	return false
}

// SetPaidAt create mutations, should not duplicate
func (w *WifiDevicesMutator) SetPaidAt(val int64) bool { //nolint:dupl false positive
	if val != w.PaidAt {
		w.mutations = append(w.mutations, A.X{`=`, 3, val})
		w.logs = append(w.logs, A.X{`paidAt`, w.PaidAt, val})
		w.PaidAt = val
		return true
	}
	return false
}

// SetPriceIDR create mutations, should not duplicate
func (w *WifiDevicesMutator) SetPriceIDR(val int64) bool { //nolint:dupl false positive
	if val != w.PriceIDR {
		w.mutations = append(w.mutations, A.X{`=`, 4, val})
		w.logs = append(w.logs, A.X{`priceIDR`, w.PriceIDR, val})
		w.PriceIDR = val
		return true
	}
	return false
}

// SetTenantId create mutations, should not duplicate
func (w *WifiDevicesMutator) SetTenantId(val uint64) bool { //nolint:dupl false positive
	if val != w.TenantId {
		w.mutations = append(w.mutations, A.X{`=`, 5, val})
		w.logs = append(w.logs, A.X{`tenantId`, w.TenantId, val})
		w.TenantId = val
		return true
	}
	return false
}

// SetMacAddress create mutations, should not duplicate
func (w *WifiDevicesMutator) SetMacAddress(val string) bool { //nolint:dupl false positive
	if val != w.MacAddress {
		w.mutations = append(w.mutations, A.X{`=`, 6, val})
		w.logs = append(w.logs, A.X{`macAddress`, w.MacAddress, val})
		w.MacAddress = val
		return true
	}
	return false
}

// SetRoomId create mutations, should not duplicate
func (w *WifiDevicesMutator) SetRoomId(val uint64) bool { //nolint:dupl false positive
	if val != w.RoomId {
		w.mutations = append(w.mutations, A.X{`=`, 7, val})
		w.logs = append(w.logs, A.X{`roomId`, w.RoomId, val})
		w.RoomId = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (w *WifiDevicesMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != w.CreatedAt {
		w.mutations = append(w.mutations, A.X{`=`, 8, val})
		w.logs = append(w.logs, A.X{`createdAt`, w.CreatedAt, val})
		w.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (w *WifiDevicesMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != w.CreatedBy {
		w.mutations = append(w.mutations, A.X{`=`, 9, val})
		w.logs = append(w.logs, A.X{`createdBy`, w.CreatedBy, val})
		w.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (w *WifiDevicesMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != w.UpdatedAt {
		w.mutations = append(w.mutations, A.X{`=`, 10, val})
		w.logs = append(w.logs, A.X{`updatedAt`, w.UpdatedAt, val})
		w.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (w *WifiDevicesMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != w.UpdatedBy {
		w.mutations = append(w.mutations, A.X{`=`, 11, val})
		w.logs = append(w.logs, A.X{`updatedBy`, w.UpdatedBy, val})
		w.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (w *WifiDevicesMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != w.DeletedAt {
		w.mutations = append(w.mutations, A.X{`=`, 12, val})
		w.logs = append(w.logs, A.X{`deletedAt`, w.DeletedAt, val})
		w.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (w *WifiDevicesMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != w.DeletedBy {
		w.mutations = append(w.mutations, A.X{`=`, 13, val})
		w.logs = append(w.logs, A.X{`deletedBy`, w.DeletedBy, val})
		w.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (w *WifiDevicesMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != w.RestoredBy {
		w.mutations = append(w.mutations, A.X{`=`, 14, val})
		w.logs = append(w.logs, A.X{`restoredBy`, w.RestoredBy, val})
		w.RestoredBy = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (w *WifiDevicesMutator) SetAll(from rqProperty.WifiDevices, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		w.Id = from.Id
		changed = true
	}
	if !excludeMap[`startAt`] && (forceMap[`startAt`] || from.StartAt != 0) {
		w.StartAt = from.StartAt
		changed = true
	}
	if !excludeMap[`endAt`] && (forceMap[`endAt`] || from.EndAt != 0) {
		w.EndAt = from.EndAt
		changed = true
	}
	if !excludeMap[`paidAt`] && (forceMap[`paidAt`] || from.PaidAt != 0) {
		w.PaidAt = from.PaidAt
		changed = true
	}
	if !excludeMap[`priceIDR`] && (forceMap[`priceIDR`] || from.PriceIDR != 0) {
		w.PriceIDR = from.PriceIDR
		changed = true
	}
	if !excludeMap[`tenantId`] && (forceMap[`tenantId`] || from.TenantId != 0) {
		w.TenantId = from.TenantId
		changed = true
	}
	if !excludeMap[`macAddress`] && (forceMap[`macAddress`] || from.MacAddress != ``) {
		w.MacAddress = S.Trim(from.MacAddress)
		changed = true
	}
	if !excludeMap[`roomId`] && (forceMap[`roomId`] || from.RoomId != 0) {
		w.RoomId = from.RoomId
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		w.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		w.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		w.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		w.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		w.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		w.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		w.RestoredBy = from.RestoredBy
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
