package rqCafe

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"kostjc/model/mCafe"

	"github.com/tarantool/go-tarantool"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

// Menus DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqCafe__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqCafe__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqCafe__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqCafe__ORM.GEN.go
type Menus struct {
	Adapter      *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id           uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	Name         string      `json:"name" form:"name" query:"name" long:"name" msg:"name"`
	HppIDR       int64       `json:"hppIDR" form:"hppIDR" query:"hppIDR" long:"hppIDR" msg:"hppIDR"`
	SalePriceIDR int64       `json:"salePriceIDR" form:"salePriceIDR" query:"salePriceIDR" long:"salePriceIDR" msg:"salePriceIDR"`
	Detail       string      `json:"detail" form:"detail" query:"detail" long:"detail" msg:"detail"`
	ImageUrl     string      `json:"imageUrl" form:"imageUrl" query:"imageUrl" long:"imageUrl" msg:"imageUrl"`
	CreatedAt    int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy    uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt    int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy    uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt    int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy    uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy   uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
}

// NewMenus create new ORM reader/query object
func NewMenus(adapter *Tt.Adapter) *Menus {
	return &Menus{Adapter: adapter}
}

// SpaceName returns full package and table name
func (m *Menus) SpaceName() string { //nolint:dupl false positive
	return string(mCafe.TableMenus) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (m *Menus) SqlTableName() string { //nolint:dupl false positive
	return `"menus"`
}

func (m *Menus) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (m *Menus) FindById() bool { //nolint:dupl false positive
	res, err := m.Adapter.Select(m.SpaceName(), m.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{m.Id})
	if L.IsError(err, `Menus.FindById failed: `+m.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		m.FromArray(rows[0])
		return true
	}
	return false
}

// UniqueIndexName return unique index name
func (m *Menus) UniqueIndexName() string { //nolint:dupl false positive
	return `name`
}

// FindByName Find one by Name
func (m *Menus) FindByName() bool { //nolint:dupl false positive
	res, err := m.Adapter.Select(m.SpaceName(), m.UniqueIndexName(), 0, 1, tarantool.IterEq, A.X{m.Name})
	if L.IsError(err, `Menus.FindByName failed: `+m.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		m.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (m *Menus) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "name"
	, "hppIDR"
	, "salePriceIDR"
	, "detail"
	, "imageUrl"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (m *Menus) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "name"
	, "hppIDR"
	, "salePriceIDR"
	, "detail"
	, "imageUrl"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// ToUpdateArray generate slice of update command
func (m *Menus) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, m.Id},
		A.X{`=`, 1, m.Name},
		A.X{`=`, 2, m.HppIDR},
		A.X{`=`, 3, m.SalePriceIDR},
		A.X{`=`, 4, m.Detail},
		A.X{`=`, 5, m.ImageUrl},
		A.X{`=`, 6, m.CreatedAt},
		A.X{`=`, 7, m.CreatedBy},
		A.X{`=`, 8, m.UpdatedAt},
		A.X{`=`, 9, m.UpdatedBy},
		A.X{`=`, 10, m.DeletedAt},
		A.X{`=`, 11, m.DeletedBy},
		A.X{`=`, 12, m.RestoredBy},
	}
}

// IdxId return name of the index
func (m *Menus) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (m *Menus) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxName return name of the index
func (m *Menus) IdxName() int { //nolint:dupl false positive
	return 1
}

// SqlName return name of the column being indexed
func (m *Menus) SqlName() string { //nolint:dupl false positive
	return `"name"`
}

// IdxHppIDR return name of the index
func (m *Menus) IdxHppIDR() int { //nolint:dupl false positive
	return 2
}

// SqlHppIDR return name of the column being indexed
func (m *Menus) SqlHppIDR() string { //nolint:dupl false positive
	return `"hppIDR"`
}

// IdxSalePriceIDR return name of the index
func (m *Menus) IdxSalePriceIDR() int { //nolint:dupl false positive
	return 3
}

// SqlSalePriceIDR return name of the column being indexed
func (m *Menus) SqlSalePriceIDR() string { //nolint:dupl false positive
	return `"salePriceIDR"`
}

// IdxDetail return name of the index
func (m *Menus) IdxDetail() int { //nolint:dupl false positive
	return 4
}

// SqlDetail return name of the column being indexed
func (m *Menus) SqlDetail() string { //nolint:dupl false positive
	return `"detail"`
}

// IdxImageUrl return name of the index
func (m *Menus) IdxImageUrl() int { //nolint:dupl false positive
	return 5
}

// SqlImageUrl return name of the column being indexed
func (m *Menus) SqlImageUrl() string { //nolint:dupl false positive
	return `"imageUrl"`
}

// IdxCreatedAt return name of the index
func (m *Menus) IdxCreatedAt() int { //nolint:dupl false positive
	return 6
}

// SqlCreatedAt return name of the column being indexed
func (m *Menus) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (m *Menus) IdxCreatedBy() int { //nolint:dupl false positive
	return 7
}

// SqlCreatedBy return name of the column being indexed
func (m *Menus) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (m *Menus) IdxUpdatedAt() int { //nolint:dupl false positive
	return 8
}

// SqlUpdatedAt return name of the column being indexed
func (m *Menus) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (m *Menus) IdxUpdatedBy() int { //nolint:dupl false positive
	return 9
}

// SqlUpdatedBy return name of the column being indexed
func (m *Menus) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (m *Menus) IdxDeletedAt() int { //nolint:dupl false positive
	return 10
}

// SqlDeletedAt return name of the column being indexed
func (m *Menus) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (m *Menus) IdxDeletedBy() int { //nolint:dupl false positive
	return 11
}

// SqlDeletedBy return name of the column being indexed
func (m *Menus) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (m *Menus) IdxRestoredBy() int { //nolint:dupl false positive
	return 12
}

// SqlRestoredBy return name of the column being indexed
func (m *Menus) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// ToArray receiver fields to slice
func (m *Menus) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if m.Id != 0 {
		id = m.Id
	}
	return A.X{
		id,
		m.Name,         // 1
		m.HppIDR,       // 2
		m.SalePriceIDR, // 3
		m.Detail,       // 4
		m.ImageUrl,     // 5
		m.CreatedAt,    // 6
		m.CreatedBy,    // 7
		m.UpdatedAt,    // 8
		m.UpdatedBy,    // 9
		m.DeletedAt,    // 10
		m.DeletedBy,    // 11
		m.RestoredBy,   // 12
	}
}

// FromArray convert slice to receiver fields
func (m *Menus) FromArray(a A.X) *Menus { //nolint:dupl false positive
	m.Id = X.ToU(a[0])
	m.Name = X.ToS(a[1])
	m.HppIDR = X.ToI(a[2])
	m.SalePriceIDR = X.ToI(a[3])
	m.Detail = X.ToS(a[4])
	m.ImageUrl = X.ToS(a[5])
	m.CreatedAt = X.ToI(a[6])
	m.CreatedBy = X.ToU(a[7])
	m.UpdatedAt = X.ToI(a[8])
	m.UpdatedBy = X.ToU(a[9])
	m.DeletedAt = X.ToI(a[10])
	m.DeletedBy = X.ToU(a[11])
	m.RestoredBy = X.ToU(a[12])
	return m
}

// FromUncensoredArray convert slice to receiver fields
func (m *Menus) FromUncensoredArray(a A.X) *Menus { //nolint:dupl false positive
	m.Id = X.ToU(a[0])
	m.Name = X.ToS(a[1])
	m.HppIDR = X.ToI(a[2])
	m.SalePriceIDR = X.ToI(a[3])
	m.Detail = X.ToS(a[4])
	m.ImageUrl = X.ToS(a[5])
	m.CreatedAt = X.ToI(a[6])
	m.CreatedBy = X.ToU(a[7])
	m.UpdatedAt = X.ToI(a[8])
	m.UpdatedBy = X.ToU(a[9])
	m.DeletedAt = X.ToI(a[10])
	m.DeletedBy = X.ToU(a[11])
	m.RestoredBy = X.ToU(a[12])
	return m
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (m *Menus) FindOffsetLimit(offset, limit uint32, idx string) []Menus { //nolint:dupl false positive
	var rows []Menus
	res, err := m.Adapter.Select(m.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Menus.FindOffsetLimit failed: `+m.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Menus{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (m *Menus) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := m.Adapter.Select(m.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Menus.FindOffsetLimit failed: `+m.SpaceName()) {
		return rows, Tt.QueryMetaFrom(res, err)
	}
	tuples := res.Tuples()
	rows = make([]A.X, len(tuples))
	for z, row := range tuples {
		rows[z] = row
	}
	return rows, Tt.QueryMetaFrom(res, nil)
}

// Total count number of rows
func (m *Menus) Total() int64 { //nolint:dupl false positive
	rows := m.Adapter.CallBoxSpace(m.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// MenusFieldTypeMap returns key value of field name and key
var MenusFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:           Tt.Unsigned,
	`name`:         Tt.String,
	`hppIDR`:       Tt.Integer,
	`salePriceIDR`: Tt.Integer,
	`detail`:       Tt.String,
	`imageUrl`:     Tt.String,
	`createdAt`:    Tt.Integer,
	`createdBy`:    Tt.Unsigned,
	`updatedAt`:    Tt.Integer,
	`updatedBy`:    Tt.Unsigned,
	`deletedAt`:    Tt.Integer,
	`deletedBy`:    Tt.Unsigned,
	`restoredBy`:   Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Sales DAO reader/query struct
type Sales struct {
	Adapter       *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id            uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	Cashier       string      `json:"cashier" form:"cashier" query:"cashier" long:"cashier" msg:"cashier"`
	Who           string      `json:"who" form:"who" query:"who" long:"who" msg:"who"`
	SalesDate     string      `json:"salesDate" form:"salesDate" query:"salesDate" long:"salesDate" msg:"salesDate"`
	PaidAt        string      `json:"paidAt" form:"paidAt" query:"paidAt" long:"paidAt" msg:"paidAt"`
	PaidWith      string      `json:"paidWith" form:"paidWith" query:"paidWith" long:"paidWith" msg:"paidWith"`
	Note          string      `json:"note" form:"note" query:"note" long:"note" msg:"note"`
	TotalPriceIDR int64       `json:"totalPriceIDR" form:"totalPriceIDR" query:"totalPriceIDR" long:"totalPriceIDR" msg:"totalPriceIDR"`
	CreatedAt     int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy     uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt     int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy     uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt     int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy     uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy    uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
}

// NewSales create new ORM reader/query object
func NewSales(adapter *Tt.Adapter) *Sales {
	return &Sales{Adapter: adapter}
}

// SpaceName returns full package and table name
func (s *Sales) SpaceName() string { //nolint:dupl false positive
	return string(mCafe.TableSales) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (s *Sales) SqlTableName() string { //nolint:dupl false positive
	return `"sales"`
}

func (s *Sales) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (s *Sales) FindById() bool { //nolint:dupl false positive
	res, err := s.Adapter.Select(s.SpaceName(), s.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{s.Id})
	if L.IsError(err, `Sales.FindById failed: `+s.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		s.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (s *Sales) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "cashier"
	, "who"
	, "salesDate"
	, "paidAt"
	, "paidWith"
	, "note"
	, "totalPriceIDR"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (s *Sales) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "cashier"
	, "who"
	, "salesDate"
	, "paidAt"
	, "paidWith"
	, "note"
	, "totalPriceIDR"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// ToUpdateArray generate slice of update command
func (s *Sales) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, s.Id},
		A.X{`=`, 1, s.Cashier},
		A.X{`=`, 2, s.Who},
		A.X{`=`, 3, s.SalesDate},
		A.X{`=`, 4, s.PaidAt},
		A.X{`=`, 5, s.PaidWith},
		A.X{`=`, 6, s.Note},
		A.X{`=`, 7, s.TotalPriceIDR},
		A.X{`=`, 8, s.CreatedAt},
		A.X{`=`, 9, s.CreatedBy},
		A.X{`=`, 10, s.UpdatedAt},
		A.X{`=`, 11, s.UpdatedBy},
		A.X{`=`, 12, s.DeletedAt},
		A.X{`=`, 13, s.DeletedBy},
		A.X{`=`, 14, s.RestoredBy},
	}
}

// IdxId return name of the index
func (s *Sales) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (s *Sales) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxCashier return name of the index
func (s *Sales) IdxCashier() int { //nolint:dupl false positive
	return 1
}

// SqlCashier return name of the column being indexed
func (s *Sales) SqlCashier() string { //nolint:dupl false positive
	return `"cashier"`
}

// IdxWho return name of the index
func (s *Sales) IdxWho() int { //nolint:dupl false positive
	return 2
}

// SqlWho return name of the column being indexed
func (s *Sales) SqlWho() string { //nolint:dupl false positive
	return `"who"`
}

// IdxSalesDate return name of the index
func (s *Sales) IdxSalesDate() int { //nolint:dupl false positive
	return 3
}

// SqlSalesDate return name of the column being indexed
func (s *Sales) SqlSalesDate() string { //nolint:dupl false positive
	return `"salesDate"`
}

// IdxPaidAt return name of the index
func (s *Sales) IdxPaidAt() int { //nolint:dupl false positive
	return 4
}

// SqlPaidAt return name of the column being indexed
func (s *Sales) SqlPaidAt() string { //nolint:dupl false positive
	return `"paidAt"`
}

// IdxPaidWith return name of the index
func (s *Sales) IdxPaidWith() int { //nolint:dupl false positive
	return 5
}

// SqlPaidWith return name of the column being indexed
func (s *Sales) SqlPaidWith() string { //nolint:dupl false positive
	return `"paidWith"`
}

// IdxNote return name of the index
func (s *Sales) IdxNote() int { //nolint:dupl false positive
	return 6
}

// SqlNote return name of the column being indexed
func (s *Sales) SqlNote() string { //nolint:dupl false positive
	return `"note"`
}

// IdxTotalPriceIDR return name of the index
func (s *Sales) IdxTotalPriceIDR() int { //nolint:dupl false positive
	return 7
}

// SqlTotalPriceIDR return name of the column being indexed
func (s *Sales) SqlTotalPriceIDR() string { //nolint:dupl false positive
	return `"totalPriceIDR"`
}

// IdxCreatedAt return name of the index
func (s *Sales) IdxCreatedAt() int { //nolint:dupl false positive
	return 8
}

// SqlCreatedAt return name of the column being indexed
func (s *Sales) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (s *Sales) IdxCreatedBy() int { //nolint:dupl false positive
	return 9
}

// SqlCreatedBy return name of the column being indexed
func (s *Sales) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (s *Sales) IdxUpdatedAt() int { //nolint:dupl false positive
	return 10
}

// SqlUpdatedAt return name of the column being indexed
func (s *Sales) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (s *Sales) IdxUpdatedBy() int { //nolint:dupl false positive
	return 11
}

// SqlUpdatedBy return name of the column being indexed
func (s *Sales) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (s *Sales) IdxDeletedAt() int { //nolint:dupl false positive
	return 12
}

// SqlDeletedAt return name of the column being indexed
func (s *Sales) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (s *Sales) IdxDeletedBy() int { //nolint:dupl false positive
	return 13
}

// SqlDeletedBy return name of the column being indexed
func (s *Sales) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (s *Sales) IdxRestoredBy() int { //nolint:dupl false positive
	return 14
}

// SqlRestoredBy return name of the column being indexed
func (s *Sales) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// ToArray receiver fields to slice
func (s *Sales) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if s.Id != 0 {
		id = s.Id
	}
	return A.X{
		id,
		s.Cashier,       // 1
		s.Who,           // 2
		s.SalesDate,     // 3
		s.PaidAt,        // 4
		s.PaidWith,      // 5
		s.Note,          // 6
		s.TotalPriceIDR, // 7
		s.CreatedAt,     // 8
		s.CreatedBy,     // 9
		s.UpdatedAt,     // 10
		s.UpdatedBy,     // 11
		s.DeletedAt,     // 12
		s.DeletedBy,     // 13
		s.RestoredBy,    // 14
	}
}

// FromArray convert slice to receiver fields
func (s *Sales) FromArray(a A.X) *Sales { //nolint:dupl false positive
	s.Id = X.ToU(a[0])
	s.Cashier = X.ToS(a[1])
	s.Who = X.ToS(a[2])
	s.SalesDate = X.ToS(a[3])
	s.PaidAt = X.ToS(a[4])
	s.PaidWith = X.ToS(a[5])
	s.Note = X.ToS(a[6])
	s.TotalPriceIDR = X.ToI(a[7])
	s.CreatedAt = X.ToI(a[8])
	s.CreatedBy = X.ToU(a[9])
	s.UpdatedAt = X.ToI(a[10])
	s.UpdatedBy = X.ToU(a[11])
	s.DeletedAt = X.ToI(a[12])
	s.DeletedBy = X.ToU(a[13])
	s.RestoredBy = X.ToU(a[14])
	return s
}

// FromUncensoredArray convert slice to receiver fields
func (s *Sales) FromUncensoredArray(a A.X) *Sales { //nolint:dupl false positive
	s.Id = X.ToU(a[0])
	s.Cashier = X.ToS(a[1])
	s.Who = X.ToS(a[2])
	s.SalesDate = X.ToS(a[3])
	s.PaidAt = X.ToS(a[4])
	s.PaidWith = X.ToS(a[5])
	s.Note = X.ToS(a[6])
	s.TotalPriceIDR = X.ToI(a[7])
	s.CreatedAt = X.ToI(a[8])
	s.CreatedBy = X.ToU(a[9])
	s.UpdatedAt = X.ToI(a[10])
	s.UpdatedBy = X.ToU(a[11])
	s.DeletedAt = X.ToI(a[12])
	s.DeletedBy = X.ToU(a[13])
	s.RestoredBy = X.ToU(a[14])
	return s
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (s *Sales) FindOffsetLimit(offset, limit uint32, idx string) []Sales { //nolint:dupl false positive
	var rows []Sales
	res, err := s.Adapter.Select(s.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Sales.FindOffsetLimit failed: `+s.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Sales{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (s *Sales) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := s.Adapter.Select(s.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Sales.FindOffsetLimit failed: `+s.SpaceName()) {
		return rows, Tt.QueryMetaFrom(res, err)
	}
	tuples := res.Tuples()
	rows = make([]A.X, len(tuples))
	for z, row := range tuples {
		rows[z] = row
	}
	return rows, Tt.QueryMetaFrom(res, nil)
}

// Total count number of rows
func (s *Sales) Total() int64 { //nolint:dupl false positive
	rows := s.Adapter.CallBoxSpace(s.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// SalesFieldTypeMap returns key value of field name and key
var SalesFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:            Tt.Unsigned,
	`cashier`:       Tt.String,
	`who`:           Tt.String,
	`salesDate`:     Tt.String,
	`paidAt`:        Tt.String,
	`paidWith`:      Tt.String,
	`note`:          Tt.String,
	`totalPriceIDR`: Tt.Integer,
	`createdAt`:     Tt.Integer,
	`createdBy`:     Tt.Unsigned,
	`updatedAt`:     Tt.Integer,
	`updatedBy`:     Tt.Unsigned,
	`deletedAt`:     Tt.Integer,
	`deletedBy`:     Tt.Unsigned,
	`restoredBy`:    Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
