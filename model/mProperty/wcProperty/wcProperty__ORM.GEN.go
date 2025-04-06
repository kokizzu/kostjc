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

// BuildingsMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcProperty__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcProperty__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcProperty__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcProperty__ORM.GEN.go
type BuildingsMutator struct {
	rqProperty.Buildings
	mutations []A.X
	logs      []A.X
}

// NewBuildingsMutator create new ORM writer/command object
func NewBuildingsMutator(adapter *Tt.Adapter) (res *BuildingsMutator) {
	res = &BuildingsMutator{Buildings: rqProperty.Buildings{Adapter: adapter}}
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
//		A.X{`=`, 3, b.FacilitiesObj},
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

// SetFacilitiesObj create mutations, should not duplicate
func (b *BuildingsMutator) SetFacilitiesObj(val string) bool { //nolint:dupl false positive
	if val != b.FacilitiesObj {
		b.mutations = append(b.mutations, A.X{`=`, 3, val})
		b.logs = append(b.logs, A.X{`facilitiesObj`, b.FacilitiesObj, val})
		b.FacilitiesObj = val
		return true
	}
	return false
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
	if !excludeMap[`facilitiesObj`] && (forceMap[`facilitiesObj`] || from.FacilitiesObj != ``) {
		b.FacilitiesObj = S.Trim(from.FacilitiesObj)
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
//	})
//	return !L.IsError(err, `Rooms.DoUpsert failed: `+r.SpaceName()+ `\n%#v`, arr)
// }

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
func (r *RoomsMutator) SetFirstUseAt(val int64) bool { //nolint:dupl false positive
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
	if !excludeMap[`firstUseAt`] && (forceMap[`firstUseAt`] || from.FirstUseAt != 0) {
		r.FirstUseAt = from.FirstUseAt
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
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
