package rqProperty

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"kostjc/model/zCrud"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
	"github.com/tarantool/go-tarantool"
)

func (l *Locations) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Locations) FindByPagination`

	validFields := LocationsFieldTypeMap
	whereAndSql := ``
	if in.Search != `` {
		whereAndSql = out.SearchBySqlTt(in.Search, in.SearchBy, validFields)
	} else {
		whereAndSql = out.WhereAndSqlTt(in.Filters, validFields)
	}

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + l.SqlTableName() + whereAndSql + `
LIMIT 1`
	l.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + l.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql

	l.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (l *Locations) FindAll() []Locations {
	const comment = `-- Locations) FindAll`

	queryRows := comment + `
SELECT ` + l.SqlSelectAllFields() + ` FROM ` + l.SqlTableName() + `
WHERE "deletedAt" = 0`

	var rows = []Locations{}
	l.Adapter.QuerySql(queryRows, func(row []any) {
		l.FromArray(row)
		rows = append(rows, *l)
	})

	return rows
}

func (f *Facilities) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Facilities) FindByPagination`

	validFields := FacilitiesFieldTypeMap
	whereAndSql := ``
	if in.Search != `` {
		whereAndSql = out.SearchBySqlTt(in.Search, in.SearchBy, validFields)
	} else {
		whereAndSql = out.WhereAndSqlTt(in.Filters, validFields)
	}

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + f.SqlTableName() + whereAndSql + `
LIMIT 1`
	f.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + f.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql

	f.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (b *Buildings) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Buildings) FindByPagination`

	validFields := BuildingsFieldTypeMap
	whereAndSql := ``
	if in.Search != `` {
		whereAndSql = out.SearchBySqlTt(in.Search, in.SearchBy, validFields)
	} else {
		whereAndSql = out.WhereAndSqlTt(in.Filters, validFields)
	}

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + b.SqlTableName() + whereAndSql + `
LIMIT 1`
	b.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + b.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql

	b.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (l *Locations) FindLocationChoices() map[uint64]string {
	const comment = `-- Locations) FindLocationChoices`

	queryRows := comment + `
SELECT ` + l.SqlId() + `, ` + l.SqlName() + ` FROM ` + l.SqlTableName() + `
WHERE "deletedAt" = 0
ORDER BY ` + l.SqlName() + ` ASC`

	out := make(map[uint64]string)
	l.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 2 {
			out[X.ToU(row[0])] = X.ToS(row[1])
		}
	})

	return out
}

func (b *Bookings) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Bookings) FindByPagination`

	validFields := BookingsFieldTypeMap
	whereAndSql := ``
	if in.Search != `` {
		whereAndSql = out.SearchBySqlTt(in.Search, in.SearchBy, validFields)
	} else {
		whereAndSql = out.WhereAndSqlTt(in.Filters, validFields)
	}

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + b.SqlTableName() + whereAndSql + `
LIMIT 1`
	b.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	if S.Contains(orderBySql, `ORDER BY "id"`) {
		orderBySql = S.Replace(orderBySql, `ORDER BY "id"`, `ORDER BY "bookings"."id"`)
	}

	if S.Contains(orderBySql, `ORDER BY "createdAt"`) {
		orderBySql = S.Replace(orderBySql, `ORDER BY "createdAt"`, `ORDER BY "bookings"."createdAt"`)
	}

	if S.Contains(orderBySql, `ORDER BY "updatedAt"`) {
		orderBySql = S.Replace(orderBySql, `ORDER BY "updatedAt"`, `ORDER BY "bookings"."updatedAt"`)
	}

	if S.Contains(orderBySql, `ORDER BY "deletedAt"`) {
		orderBySql = S.Replace(orderBySql, `ORDER BY "deletedAt"`, `ORDER BY "bookings"."deletedAt"`)
	}

	if S.Contains(orderBySql, `ORDER BY "extraTenants" DESC`) {
		orderBySql = S.Replace(orderBySql, `ORDER BY "extraTenants" DESC`, ``)
	}

	if S.Contains(orderBySql, `ORDER BY "extraTenants" ASC`) {
		orderBySql = S.Replace(orderBySql, `ORDER BY "extraTenants" ASC`, ``)
	}

	if S.Contains(orderBySql, `ORDER BY "extraTenants"`) {
		orderBySql = S.Replace(orderBySql, `ORDER BY "extraTenants"`, ``)
	}

	queryRows := comment + `
SELECT
	"bookings"."id",
	"bookings"."roomId",
	COALESCE(SUM("payments"."paidIDR"), 0) AS "totalPaidIDR",
	"bookings"."dateStart",
	"bookings"."dateEnd",
	"bookings"."tenantId",
	"bookings"."basePriceIDR",
	"bookings"."facilitiesObj",
	"bookings"."totalPriceIDR",
	"bookings"."paidAt",
	"bookings"."extraTenants",
	"bookings"."createdAt",
	"bookings"."updatedAt",
	"bookings"."deletedAt"
FROM ` + b.SqlTableName() + `
LEFT JOIN "payments" ON "payments"."bookingId" = "bookings"."id"
` + whereAndSql + `
GROUP BY "bookings"."id"
` + orderBySql + limitOffsetSql

	b.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (f *Facilities) FindAll() []Facilities {
	const comment = `-- Facilities) FindAll`

	queryRows := comment + `
SELECT + ` + f.SqlSelectAllFields() + ` FROM ` + f.SqlTableName()

	var rows = []Facilities{}
	f.Adapter.QuerySql(queryRows, func(row []any) {
		f.FromArray(row)
		rows = append(rows, *f)
	})

	return rows
}

func (f *Facilities) FindAllTypeBuilding() []Facilities {
	const comment = `-- Facilities) FindAllTypeBuilding`

	queryRows := comment + `
SELECT + ` + f.SqlSelectAllFields() + ` FROM ` + f.SqlTableName() + `
WHERE ` + f.SqlFacilityType() + ` = 'Building'
	AND "deletedAt" = 0`

	var rows = []Facilities{}
	f.Adapter.QuerySql(queryRows, func(row []any) {
		f.FromArray(row)
		rows = append(rows, *f)
	})

	return rows
}

func (f *Facilities) FindAllTypeRoom() []Facilities {
	const comment = `-- Facilities) FindAllTypeRoom`

	queryRows := comment + `
SELECT + ` + f.SqlSelectAllFields() + ` FROM ` + f.SqlTableName() + `
WHERE ` + f.SqlFacilityType() + ` = 'Room'
	AND "deletedAt" = 0`

	var rows = []Facilities{}
	f.Adapter.QuerySql(queryRows, func(row []any) {
		f.FromArray(row)
		rows = append(rows, *f)
	})

	return rows
}

func (f *Facilities) FindFacilitiesChoices() map[uint64]string {
	const comment = `-- Locations) FindLocationChoices`

	queryRows := comment + `
SELECT ` + f.SqlId() + `, ` + f.SqlFacilityName() + `, ` + f.SqlFacilityType() + ` FROM ` + f.SqlTableName() + `
ORDER BY ` + f.SqlFacilityName() + ` ASC`

	out := make(map[uint64]string)
	f.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 3 {
			facilityName := X.ToS(row[1])
			facilityType := X.ToS(row[2])
			if facilityType != `` {
				facilityName += ` (` + facilityType + `)`
			}

			out[X.ToU(row[0])] = facilityName
		}
	})

	return out
}

func (f *Facilities) FindFacilitiesBuildingChoices() map[uint64]string {
	const comment = `-- Locations) FindFacilitiesBuildingChoices`

	queryRows := comment + `
SELECT ` + f.SqlId() + `, ` + f.SqlFacilityName() + ` FROM ` + f.SqlTableName() + `
WHERE ` + f.SqlFacilityType() + ` = 'Building'
	AND "deletedAt" = 0
ORDER BY ` + f.SqlFacilityName() + ` ASC`

	out := make(map[uint64]string)
	f.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 2 {
			out[X.ToU(row[0])] = X.ToS(row[1])
		}
	})

	return out
}

func (p *Payments) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Payments) FindByPagination`

	validFields := PaymentsFieldTypeMap
	whereAndSql := ``
	if in.Search != `` {
		whereAndSql = out.SearchBySqlTt(in.Search, in.SearchBy, validFields)
	} else {
		whereAndSql = out.WhereAndSqlTt(in.Filters, validFields)
	}

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + p.SqlTableName() + whereAndSql + `
LIMIT 1`
	p.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + p.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql

	p.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (b *Buildings) FindBuildingChoices() map[uint64]string {
	const comment = `-- Buildings) FindBuildingChoices`

	queryRows := comment + `
SELECT ` + b.SqlId() + `, ` + b.SqlBuildingName() + ` FROM ` + b.SqlTableName() + `
WHERE "deletedAt" = 0
ORDER BY ` + b.SqlId() + ` ASC`

	out := make(map[uint64]string)
	b.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 2 {
			out[X.ToU(row[0])] = X.ToS(row[1])
		}
	})

	return out
}

func (b *Bookings) FindBookingChoices() map[uint64]string {
	const comment = `-- Bookings) FindBookingChoices`

	queryRows := comment + `
SELECT
	"bookings"."id",
	"bookings"."totalPriceIDR",
	COALESCE("rooms"."roomName", ''),
	"bookings"."dateStart",
	COALESCE("tenants"."tenantName", '')
FROM "bookings"
LEFT JOIN "tenants" ON "bookings"."tenantId" = "tenants"."id"
LEFT JOIN "rooms" ON "bookings"."roomId" = "rooms"."id"
ORDER BY "bookings"."id" ASC`

	out := make(map[uint64]string)
	b.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 5 {
			id := `#` + X.ToS(row[0])
			totalPriceIdr := formatCurrency(X.ToI(row[1]), `IDR`)
			roomName := X.ToS(row[2])
			dateStart := X.ToS(row[3])
			tenantName := X.ToS(row[4])

			formattedBooking := fmt.Sprintf(
				"%s - %s - %s - %s - %s",
				id, totalPriceIdr, roomName, dateStart, tenantName,
			)
			out[X.ToU(row[0])] = formattedBooking
		}
	})

	return out
}

func formatCurrency(value int64, currency string) string {
	str := strconv.FormatInt(value, 10)

	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	var result []rune
	for i, r := range runes {
		if i > 0 && i%3 == 0 {
			result = append(result, ',')
		}
		result = append(result, r)
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return currency + ` ` + string(result)
}

func (s *Stocks) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Stocks) FindByPagination`

	validFields := LocationsFieldTypeMap
	whereAndSql := ``
	if in.Search != `` {
		whereAndSql = out.SearchBySqlTt(in.Search, in.SearchBy, validFields)
	} else {
		whereAndSql = out.WhereAndSqlTt(in.Filters, validFields)
	}

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + s.SqlTableName() + whereAndSql + `
LIMIT 1`
	s.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + s.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql

	s.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (r *Rooms) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Rooms) FindByPagination`

	validFields := RoomsFieldTypeMap
	whereAndSql := ``
	if in.Search != `` {
		whereAndSql = out.SearchBySqlTt(in.Search, in.SearchBy, validFields)
	} else {
		whereAndSql = out.WhereAndSqlTt(in.Filters, validFields)
	}

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + r.SqlTableName() + whereAndSql + `
LIMIT 1`
	r.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + r.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	r.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})
	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (b *Bookings) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (b *Bookings) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + b.SqlTableName() + `
	LIMIT 1`

	b.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (b *Bookings) Truncate() bool {
	return b.Adapter.ExecBoxSpace(b.SpaceName()+`:truncate`, A.X{})
}

func (b *Buildings) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (b *Buildings) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + b.SqlTableName() + `
	LIMIT 1`

	b.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (b *Buildings) Truncate() bool {
	return b.Adapter.ExecBoxSpace(b.SpaceName()+`:truncate`, A.X{})
}

func (b *Facilities) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (b *Facilities) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + b.SqlTableName() + `
	LIMIT 1`

	b.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (b *Facilities) Truncate() bool {
	return b.Adapter.ExecBoxSpace(b.SpaceName()+`:truncate`, A.X{})
}

func (b *Locations) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (b *Locations) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + b.SqlTableName() + `
	LIMIT 1`

	b.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (b *Locations) Truncate() bool {
	return b.Adapter.ExecBoxSpace(b.SpaceName()+`:truncate`, A.X{})
}

func (b *Payments) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (b *Payments) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + b.SqlTableName() + `
	LIMIT 1`

	b.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (b *Payments) Truncate() bool {
	return b.Adapter.ExecBoxSpace(b.SpaceName()+`:truncate`, A.X{})
}

func (b *Rooms) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (b *Rooms) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + b.SqlTableName() + `
	LIMIT 1`

	b.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (b *Rooms) Truncate() bool {
	return b.Adapter.ExecBoxSpace(b.SpaceName()+`:truncate`, A.X{})
}

func (b *Stocks) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (b *Stocks) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + b.SqlTableName() + `
	LIMIT 1`

	b.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (b *Stocks) Truncate() bool {
	return b.Adapter.ExecBoxSpace(b.SpaceName()+`:truncate`, A.X{})
}

func (r *Rooms) FindRoomChoices() map[uint64]string {
	const comment = `-- Rooms) FindRoomChoices`

	queryRows := comment + `
SELECT ` + r.SqlId() + `, ` + r.SqlRoomName() + ` FROM ` + r.SqlTableName() + `
ORDER BY ` + r.SqlRoomName() + ` ASC`

	out := make(map[uint64]string)
	r.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 2 {
			out[X.ToU(row[0])] = X.ToS(row[1])
		}
	})

	return out
}

func (r *Rooms) FindRoomNames() (out []string) {
	const comment = `-- Rooms) FindRoomNames`

	queryRows := comment + `
SELECT ` + r.SqlRoomName() + ` FROM ` + r.SqlTableName() + `
WHERE "deletedAt" = 0
ORDER BY ` + r.SqlRoomName() + ` ASC`

	r.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 1 {
			roomName := X.ToS(row[0])
			out = append(out, roomName)
		}
	})

	return
}

const DateFormatYYYYMM = "2006-01"

func isValidYearMonth(yearMonth string) bool {
	_, err := time.Parse(DateFormatYYYYMM, yearMonth)
	return err == nil
}

type BookingDetail struct {
	Id           uint64 `json:"id"`
	RoomId       uint64 `json:"roomId"`
	RoomName     string `json:"roomName"`
	TenantId     uint64 `json:"tenantId"`
	TenantName   string `json:"tenantName"`
	DateStart    string `json:"dateStart"`
	DateEnd      string `json:"dateEnd"`
	AmountPaid   int64  `json:"amountPaid"`
	TotalPrice   int64  `json:"totalPrice"`
	DeletedAt    int64  `json:"deletedAt"`
	IsNearEnding bool   `json:"isNearEnding"`
	IsExtended   bool   `json:"isExtended"`
	ExtraTenants []any  `json:"extraTenants"`
}

type RoomBooking struct {
	RoomName string          `json:"roomName"`
	Bookings []BookingDetail `json:"bookings"`
}

func getEndOfMonth(ym string) string {
	t, _ := time.Parse(DateFormatYYYYMM, ym)
	endOfMonth := t.AddDate(0, 1, 0).AddDate(0, 0, -1)
	return endOfMonth.Format(time.DateOnly)
}

func (b *Bookings) FindBookingsPerQuartal(monthStart, monthEnd string) (out []BookingDetail) {
	const comment = `-- Bookings) FindBookingsPerQuartal`

	if !(isValidYearMonth(monthStart) || isValidYearMonth(monthEnd)) {
		monthStart = time.Now().Format(DateFormatYYYYMM)
		monthEnd = time.Now().AddDate(0, 4, 0).Format(DateFormatYYYYMM)
	}

	monthEnd = getEndOfMonth(monthEnd)
	monthStart += "-01"

	queryRows := comment + `
SELECT 
	"bookings"."id" AS "id",
	"rooms"."id" AS "roomId",
  "rooms"."roomName",
	"bookings"."tenantId",
  COALESCE("tenants"."tenantName", '') AS "tenantName",
  COALESCE("bookings"."dateStart", '') AS "dateStart",
  COALESCE("bookings"."dateEnd", '') AS "dateEnd",
  COALESCE(
		SUM(CASE
			WHEN "payments"."deletedAt" = 0
			THEN "payments"."paidIDR"
			ELSE 0 END
		),
	0) AS "totalPaidIDR",
  "bookings"."totalPriceIDR",
	"bookings"."deletedAt",
	"bookings"."extraTenants"
FROM "bookings"
LEFT JOIN "tenants" ON "bookings"."tenantId" = "tenants"."id"
LEFT JOIN "rooms" ON "bookings"."roomId" = "rooms"."id"
LEFT JOIN "payments" ON "bookings"."id" = "payments"."bookingId"
WHERE
	(
    "bookings"."dateStart" BETWEEN ` + S.Z(monthStart) + ` AND ` + S.Z(monthEnd) + `
    OR
    "bookings"."dateEnd" BETWEEN ` + S.Z(monthStart) + ` AND ` + S.Z(monthEnd) + `
  )
GROUP BY
	"bookings"."id",
	"rooms"."roomName"
ORDER BY "rooms"."roomName" ASC`

	b.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 11 {
			id := X.ToU(row[0])
			roomId := X.ToU(row[1])
			roomName := X.ToS(row[2])
			tenantId := X.ToU(row[3])
			tenantName := X.ToS(row[4])
			dateStart := X.ToS(row[5])
			dateEnd := X.ToS(row[6])
			totalPaidIdr := X.ToI(row[7])
			totalPriceIdr := X.ToI(row[8])
			deletedAt := X.ToI(row[9])
			extraTenants := X.ToArr(row[10])

			out = append(out, BookingDetail{
				Id:           id,
				RoomId:       roomId,
				RoomName:     roomName,
				TenantId:     tenantId,
				TenantName:   tenantName,
				DateStart:    dateStart,
				DateEnd:      dateEnd,
				AmountPaid:   totalPaidIdr,
				TotalPrice:   totalPriceIdr,
				DeletedAt:    deletedAt,
				ExtraTenants: extraTenants,
			})
		}
	})

	out = updateFlags(out)

	return
}

func updateFlags(bookings []BookingDetail) []BookingDetail {
	now := time.Now()
	sevenDaysLater := now.AddDate(0, 0, 7)

	for i := range bookings {
		current := &bookings[i]
		if current.DeletedAt != 0 {
			continue
		}

		dateEnd, err := time.Parse(time.DateOnly, current.DateEnd)
		if err != nil {
			continue
		}

		// ============ IS EXTENDED ============
		for _, other := range bookings {
			if other.DeletedAt != 0 || other.Id == current.Id {
				continue
			}
			if other.RoomId == current.RoomId {
				otherStart, err := time.Parse(time.DateOnly, other.DateStart)
				if err != nil {
					continue
				}
				if otherStart.After(dateEnd) {
					current.IsExtended = true
					break
				}
			}
		}

		// ============ IS NEAR ENDING ============
		hasFutureBookingSameRoom := false
		for _, other := range bookings {
			if other.DeletedAt != 0 || other.Id == current.Id {
				continue
			}
			if other.RoomId == current.RoomId {
				otherStart, err := time.Parse(time.DateOnly, other.DateStart)
				if err != nil {
					continue
				}
				if otherStart.After(now) {
					hasFutureBookingSameRoom = true
					break
				}
			}
		}

		if !hasFutureBookingSameRoom {
			if dateEnd.Before(now) {
				current.IsNearEnding = true // Sudah habis dan tidak diperpanjang
			} else if dateEnd.Before(sevenDaysLater.Add(24 * time.Hour)) {
				current.IsNearEnding = true // Akan habis <= 7 hari
			}
		}
	}

	return bookings
}

func parseMonth(monthStr string) (time.Time, error) {
	return time.Parse("2006-01", monthStr)
}

func monthsInRange(start, end time.Time) []time.Time {
	var months []time.Time
	for t := start; !t.After(end); t = t.AddDate(0, 1, 0) {
		months = append(months, t)
	}
	return months
}

func GroupBookingsToQuarter(bookings []BookingDetail, monthStartStr, monthEndStr string) ([]RoomBooking, error) {
	monthStart, err := parseMonth(monthStartStr)
	if err != nil {
		return nil, err
	}
	monthEnd, err := parseMonth(monthEndStr)
	if err != nil {
		return nil, err
	}

	monthSlots := monthsInRange(monthStart, monthEnd)
	roomMap := make(map[string][]BookingDetail)

	for _, b := range bookings {
		startDate, err1 := time.Parse(time.DateOnly, b.DateStart)
		endDate, err2 := time.Parse(time.DateOnly, b.DateEnd)
		if err1 != nil || err2 != nil {
			continue
		}

		for i, m := range monthSlots {
			monthStart := m
			monthEnd := m.AddDate(0, 1, 0).Add(-time.Nanosecond) // end of month
			if (startDate.Before(monthEnd) || startDate.Equal(monthEnd)) &&
				(endDate.After(monthStart) || endDate.Equal(monthStart)) {
				if _, exists := roomMap[b.RoomName]; !exists {
					roomMap[b.RoomName] = make([]BookingDetail, len(monthSlots))
				}
				roomMap[b.RoomName][i] = b
			}
		}
	}

	// Build result
	var result []RoomBooking
	for room, bookings := range roomMap {
		result = append(result, RoomBooking{
			RoomName: room,
			Bookings: bookings,
		})
	}
	return result, nil
}

func (p *Payments) FindByBookingId(bookingId uint64) (rows []Payments) {
	const comment = `-- Payments) FindByBookingId`

	query := comment + `
SELECT ` + p.SqlSelectAllFields() + `
FROM ` + p.SqlTableName() + `
WHERE ` + p.SqlBookingId() + ` = ` + I.UToS(bookingId)

	p.Adapter.QuerySql(query, func(row []any) {
		p.FromArray(row)
		rows = append(rows, *p)
	})

	return
}

type RoomMissingTenantData struct {
	RoomId                 uint64 `json:"roomId"`
	RoomName               string `json:"roomName"`
	TenantId               uint64 `json:"tenantId"`
	TenantName             string `json:"tenantName"`
	TenantTelegramUsername string `json:"tenantTelegramUsername"`
	TenantWhatsappNumber   string `json:"tenantWhatsappNumber"`
	TenantWaAddedAt        bool   `json:"tenantWaAddedAt"`
	TenantTeleAddedAt      bool   `json:"tenantTeleAddedAt"`
	LastUseAt              string `json:"lastUseAt"`
}

func (r *Rooms) FindMissingTenantsData(yearMonth string) (out []RoomMissingTenantData) {
	const comment = `-- Rooms) FindMissingTenantsData`

	var startDate string // YYYY-MM-DD
	var endDate string   // YYYY-MM-DD

	year, month, isValid := isValidMonthYear(yearMonth)
	if !isValid {
		now := time.Now()
		startDate = startOfYearMonth(now.Year(), now.Month())
		endDate = endOfYearMonth(now.Year(), now.Month())
	} else {
		startDate = startOfYearMonth(year, month)
		endDate = endOfYearMonth(year, month)
	}

	queryRows := comment + `
SELECT 
  "rooms"."id",
  "rooms"."roomName",
  "tenants"."id",
  "tenants"."tenantName",
  "tenants"."telegramUsername",
  "tenants"."whatsappNumber",
  "tenants"."waAddedAt",
  "tenants"."teleAddedAt",
	"rooms"."lastUseAt"
FROM "rooms"
LEFT JOIN "tenants"
	ON "rooms"."currentTenantId" = "tenants"."id"
WHERE
	"rooms"."deletedAt" = 0
	AND "rooms"."lastUseAt" >= ` + S.Z(startDate) + `
	AND "rooms"."lastUseAt" <= ` + S.Z(endDate) + `
ORDER BY "rooms"."updatedAt"`

	r.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) != 9 {
			return
		}
		out = append(out, RoomMissingTenantData{
			RoomId:                 X.ToU(row[0]),
			RoomName:               X.ToS(row[1]),
			TenantId:               X.ToU(row[2]),
			TenantName:             X.ToS(row[3]),
			TenantTelegramUsername: X.ToS(row[4]),
			TenantWhatsappNumber:   X.ToS(row[5]),
			TenantWaAddedAt:        X.ToBool(row[6]),
			TenantTeleAddedAt:      X.ToBool(row[7]),
			LastUseAt:              X.ToS(row[8]),
		})
	})

	return
}

type TenantNearbyBirthday struct {
	TenantId   uint64 `json:"tenantId"`
	TenantName string `json:"tenantName"`
	BirthDay   string `json:"birthDay"`
	Age        int    `json:"age"`
	RoomName   string `json:"roomName"`
}

func (r *Rooms) FindTenantsNearbyBirthdays() []TenantNearbyBirthday {
	const comment = `-- Rooms) FindTenantsNearbyBirthdays`

	queryRows := comment + `
SELECT
	"tenants"."id",
  "tenants"."tenantName",
  "tenants"."ktpDateBirth",
  "rooms"."roomName"
FROM "rooms"
LEFT JOIN "tenants" ON "rooms"."currentTenantId" = "tenants"."id"
WHERE "rooms"."deletedAt" = 0`

	mapTenants := make(map[any]bool)
	resp := []TenantNearbyBirthday{}
	r.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) != 4 {
			return
		}
		tenantId := X.ToU(row[0])
		if mapTenants[tenantId] {
			return
		}
		birthDay := X.ToS(row[2])
		resp = append(resp, TenantNearbyBirthday{
			TenantId:   tenantId,
			TenantName: X.ToS(row[1]),
			BirthDay:   birthDay,
			RoomName:   X.ToS(row[3]),
			Age:        getAgeByBirthday(birthDay),
		})

		mapTenants[tenantId] = true
	})

	now := time.Now()
	sevenDaysAgo := now.AddDate(0, 0, -7).Format("01-02")
	sevenDaysLater := now.AddDate(0, 0, 7).Format("01-02")

	out := []TenantNearbyBirthday{}
	for _, v := range resp {
		birthDate, err := time.Parse(time.DateOnly, v.BirthDay)
		if err != nil {
			continue
		}
		birthMMDD := birthDate.Format("01-02")

		if birthMMDD >= sevenDaysAgo && birthMMDD <= sevenDaysLater {
			out = append(out, v)
		}
	}

	sort.Slice(out, func(i, j int) bool {
		ti, _ := time.Parse(time.DateOnly, out[i].BirthDay)
		tj, _ := time.Parse(time.DateOnly, out[j].BirthDay)

		mi, di := ti.Month(), ti.Day()
		mj, dj := tj.Month(), tj.Day()

		if mi == mj {
			return di < dj
		}

		return mi < mj
	})

	return out
}

func getAgeByBirthday(birthday string) int {
	birthDate, err := time.Parse(time.DateOnly, birthday)
	if err != nil {
		return 0
	}

	now := time.Now()
	age := now.Year() - birthDate.Year()
	if now.YearDay() < birthDate.YearDay() {
		age--
	}

	return age
}

type AvailableRoom struct {
	RoomName       string `json:"roomName"`
	AvailableAt    string `json:"availableAt"`
	IsAvailableNow bool   `json:"isAvailableNow"`
	LastTenant     string `json:"lastTenant"`
}

func (r *Rooms) FindAvailableRooms() (out []AvailableRoom) {
	const comment = `-- Rooms) FindAvailableRooms`

	now := time.Now()
	dateTimeNow := now.Format(time.DateOnly)
	sevenDaysLater := now.AddDate(0, 0, 7).Format(time.DateOnly)

	queryRows := comment + `
WITH "all_rooms" AS (
	SELECT
		"rooms"."roomName" AS "roomName",
		"bookings"."dateEnd" AS "dateEnd",
		MAX("bookings"."dateEnd") AS "maxDateEnd",
		"rooms"."lastUseAt" AS "lastUseAt",
		"rooms"."currentTenantId" AS "currentTenantId",
		"rooms"."deletedAt" AS "deletedAt",
		"tenants"."tenantName" AS "tenantName"
	FROM "rooms"
	LEFT JOIN "bookings" ON "rooms"."id" = "bookings"."roomId"
	LEFT JOIN "tenants" ON "rooms"."currentTenantId" = "tenants"."id"
	GROUP BY "rooms"."roomName"
)

SELECT
	"roomName",
	"dateEnd",
	CASE
		WHEN "dateEnd" < '` + dateTimeNow + `'
		THEN 'TRUE'
	ELSE 'FALSE' END AS "isAvailable",
	"tenantName"
FROM "all_rooms"
WHERE
	(
		"dateEnd" <= '` + sevenDaysLater + `'
		OR "dateEnd" IS NULL
		OR "currentTenantId" = 0
	)
	AND "deletedAt" = 0
ORDER BY "dateEnd" ASC`

	r.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 4 {
			roomName := X.ToS(row[0])
			availableAt := X.ToS(row[1])
			isAvailableNow := X.ToBool(row[2])
			tenantName := X.ToS(row[3])
			out = append(out, AvailableRoom{
				RoomName:       roomName,
				AvailableAt:    availableAt,
				IsAvailableNow: isAvailableNow,
				LastTenant:     tenantName,
			})
		}
	})

	return
}

type UnpaidBookingTenant struct {
	TenantName string `json:"tenantName"`
	RoomName   string `json:"roomName"`
	TotalPaid  int64  `json:"totalPaid"`
	TotalPrice int64  `json:"totalPrice"`
	DateStart  string `json:"dateStart"`
	DateEnd    string `json:"dateEnd"`
}

func (b *Bookings) FindUnpaidBookingTenants() (out []UnpaidBookingTenant) {
	const comment = `-- Bookings) FindUnpaidBookingTenants`

	queryRows := comment + `
SELECT
	t."id",
	t."tenantName",
	r."roomName",
	COALESCE(SUM(p."paidIDR"), 0) AS totalPaidIDR,
	b."totalPriceIDR",
	b."dateStart",
	b."dateEnd"
FROM "bookings" b
LEFT JOIN "tenants" t ON b."tenantId" = t."id"
LEFT JOIN "rooms" r ON b."roomId" = r."id"
LEFT JOIN "payments" p ON b."id" = p."bookingId"
WHERE
	b."deletedAt" = 0
	AND t."deletedAt" = 0
	AND r."deletedAt" = 0
GROUP BY
	b."id", t."id", t."tenantName", r."roomName"
HAVING totalPaidIDR <> b."totalPriceIDR"
ORDER BY t."tenantName" ASC`

	b.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) != 7 {
			return
		}

		tenantName := X.ToS(row[1])
		roomName := X.ToS(row[2])
		totalPaid := X.ToI(row[3])
		totalPrice := X.ToI(row[4])
		dateStart := X.ToS(row[5])
		dateEnd := X.ToS(row[6])

		out = append(out, UnpaidBookingTenant{
			TenantName: tenantName,
			RoomName:   roomName,
			TotalPaid:  totalPaid,
			TotalPrice: totalPrice,
			DateStart:  dateStart,
			DateEnd:    dateEnd,
		})
	})

	sort.Slice(out, func(i, j int) bool {
		t1, _ := time.Parse(time.DateOnly, out[i].DateStart)
		t2, _ := time.Parse(time.DateOnly, out[j].DateStart)
		return t1.Before(t2)
	})

	return
}

type RevenueReport struct {
	DateStart   string `json:"dateStart"`
	BookingId   uint64 `json:"bookingId"`
	RevenueIDR  int64  `json:"revenueIDR"`
	DonationIDR int64  `json:"donationIDR"`
}

func (b *Bookings) FindRevenueReports(yearMonth string) (out []RevenueReport) {
	const comment = `-- Bookings) FindRevenueReport`

	if !isValidYearMonth(yearMonth) {
		yearMonth = time.Now().Format(DateFormatYYYYMM)
	}

	query := comment + `
SELECT
	"bookings"."dateStart",
	"bookings"."id" AS "bookingId",
	SUM(CASE 
		WHEN "paymentMethod" != 'Donation'
			THEN "payments"."paidIDR"
		ELSE 0
	END) AS "revenueIDR",
	SUM(CASE 
		WHEN "paymentMethod" = 'Donation'
			THEN "payments"."paidIDR"
		ELSE 0
	END) AS "donationIDR"
FROM "bookings"
LEFT JOIN "payments" ON "bookings"."id" = "payments"."bookingId"
WHERE
	"bookings"."deletedAt" = 0
	AND SUBSTR("bookings"."dateStart", 1, 7) = '` + yearMonth + `'
GROUP BY "bookings"."id"`

	b.Adapter.QuerySql(query, func(row []any) {
		if len(row) == 4 {
			dateStart := X.ToS(row[0])
			bookingId := X.ToU(row[1])
			revenueIDR := X.ToI(row[2])
			donationIDR := X.ToI(row[3])
			out = append(out, RevenueReport{
				DateStart:   dateStart,
				BookingId:   bookingId,
				RevenueIDR:  revenueIDR,
				DonationIDR: donationIDR,
			})
		}
	})

	return
}

func (w *WifiDevices) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- WifiDevices) FindByPagination`

	validFields := WifiDevicesFieldTypeMap
	whereAndSql := ``
	if in.Search != `` {
		whereAndSql = out.SearchBySqlTt(in.Search, in.SearchBy, validFields)
	} else {
		whereAndSql = out.WhereAndSqlTt(in.Filters, validFields)
	}

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + w.SqlTableName() + whereAndSql + `
LIMIT 1`
	w.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + w.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql

	w.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	out.Order = in.Order
	out.Filters = in.Filters

	return
}

type WifiDeviceReport struct {
	TenantId  uint64 `json:"tenantId"`
	RoomId    uint64 `json:"roomId"`
	StartAt   string `json:"startAt"`
	EndAt     string `json:"endAt"`
	PaidAt    string `json:"paidAt"`
	DeletedAt int64  `json:"deletedAt"`
}

func (w *WifiDevices) FindWifiDeviceReports(yearMonth string) (out []WifiDeviceReport) {
	const comment = `-- WifiDevices) FindWifiDeviceReports`

	if !isValidYearMonth(yearMonth) {
		yearMonth = time.Now().Format(DateFormatYYYYMM)
	}

	query := comment + `
SELECT
	"tenantId",
	"roomId",
	"startAt",
	"endAt",
	"paidAt",
	"deletedAt"
FROM "wifiDevices"
WHERE SUBSTR("endAt", 1, 7) = '` + yearMonth + `'
`

	w.Adapter.QuerySql(query, func(row []any) {
		if len(row) != 6 {
			return
		}
		tenantId := X.ToU(row[0])
		roomId := X.ToU(row[1])
		startAt := X.ToS(row[2])
		endAt := X.ToS(row[3])
		paidAt := X.ToS(row[4])
		deletedAt := X.ToI(row[5])
		out = append(out, WifiDeviceReport{
			TenantId:  tenantId,
			RoomId:    roomId,
			StartAt:   startAt,
			EndAt:     endAt,
			PaidAt:    paidAt,
			DeletedAt: deletedAt,
		})
	})

	return
}

type DoubleBookingReportData struct {
	RoomId     uint64 `json:"-"`
	RoomName   string `json:"-"`
	TenantId   uint64 `json:"tenantId"`
	TenantName string `json:"tenantName"`
	DateStart  string `json:"dateStart"`
	DateEnd    string `json:"dateEnd"`
}

type DoubleBookingReport struct {
	RoomId   uint64                    `json:"roomId"`
	RoomName string                    `json:"roomName"`
	Tenants  []DoubleBookingReportData `json:"tenants"`
}

func (b *Bookings) FindDoubleBookingReports() (out []DoubleBookingReport) {
	const comment = `-- Bookings) FindDoubleBookingReports`

	dtNow := time.Now().Format(time.DateOnly)
	queryRows := comment + `
WITH overlapping_groups AS (
  SELECT 
    b."id" AS booking_id,
    b."roomId" AS room_id,
    r."roomName" AS room_name,
    b."tenantId" AS tenant_id,
    t."tenantName" AS tenant_name,
    b."dateStart" AS date_start,
    b."dateEnd" AS date_end
  FROM "bookings" b
	LEFT JOIN "bookings" b2
		ON b."roomId" = b2."roomId"
		AND b."id" <> b2."id"
		AND (
			b."dateStart" <= b2."dateEnd" AND
			b."dateEnd" >= b2."dateStart"
		)
  LEFT JOIN "tenants" t ON b."tenantId" = t."id"
  LEFT JOIN "rooms" r ON b."roomId" = r."id"
	WHERE b."deletedAt" = 0
)

SELECT 
  room_id,
  room_name,
  tenant_id,
  tenant_name,
  date_start,
  date_end
FROM overlapping_groups
WHERE date_end >= ` + S.Z(dtNow) + `
ORDER BY room_name`

	rawResults := []DoubleBookingReportData{}
	b.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 6 {
			rawResults = append(rawResults, DoubleBookingReportData{
				RoomId:     X.ToU(row[0]),
				RoomName:   X.ToS(row[1]),
				TenantId:   X.ToU(row[2]),
				TenantName: X.ToS(row[3]),
				DateStart:  X.ToS(row[4]),
				DateEnd:    X.ToS(row[5]),
			})
		}
	})

	out = groupDoubleBookingByRoom(rawResults)

	return
}

func parseDateYYYYMMDD(dateStr string) time.Time {
	t, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		return time.Now()
	}

	return t
}

func isOverlapping(a, b DoubleBookingReportData) bool {
	startA := parseDateYYYYMMDD(a.DateStart)
	endA := parseDateYYYYMMDD(a.DateEnd)
	startB := parseDateYYYYMMDD(b.DateStart)
	endB := parseDateYYYYMMDD(b.DateEnd)

	return startA.Before(endB.Add(24*time.Hour)) && startB.Before(endA.Add(24*time.Hour))
}

func hasAnyOverlap(data []DoubleBookingReportData) bool {
	for i := range data {
		for j := i + 1; j < len(data); j++ {
			if isOverlapping(data[i], data[j]) {
				return true
			}
		}
	}
	return false
}

func groupDoubleBookingByRoom(data []DoubleBookingReportData) []DoubleBookingReport {
	groupMap := make(map[uint64]*DoubleBookingReport)

	for _, d := range data {
		if groupMap[d.RoomId] == nil {
			groupMap[d.RoomId] = &DoubleBookingReport{
				RoomId:   d.RoomId,
				RoomName: d.RoomName,
				Tenants:  []DoubleBookingReportData{},
			}
		}
		groupMap[d.RoomId].Tenants = append(groupMap[d.RoomId].Tenants, d)
	}

	var result = []DoubleBookingReport{}
	for _, group := range groupMap {
		result = append(result, *group)
	}
	groupMap = nil

	var filteredResult = []DoubleBookingReport{}
	for _, report := range result {
		totalTenants := len(report.Tenants)
		if totalTenants > 1 {
			if hasAnyOverlap(report.Tenants) {
				filteredResult = append(filteredResult, report)
			}
		}
	}
	result = nil

	var finalResult = []DoubleBookingReport{}
	for _, report := range filteredResult {
		tenantCount := make(map[uint64]int)
		duplicateFound := false

		for _, tenant := range report.Tenants {
			tenantCount[tenant.TenantId]++
			if tenantCount[tenant.TenantId] > 1 {
				duplicateFound = true
				break
			}
		}

		if !duplicateFound {
			finalResult = append(finalResult, report)
		}
	}

	return finalResult
}

type RoomBookingInconsistency struct {
	RoomId          uint64 `json:"roomId"`
	TenantId        uint64 `json:"tenantId"`
	CurrentTenantId uint64 `json:"currentTenantId"`
	DateEnd         string `json:"dateEnd"`
	IsInconsistent  bool   `json:"isInconsistent"`
}

func (r *Rooms) FindRoomBookingInconsistencies() (out []RoomBookingInconsistency) {
	const comment = `-- Rooms) FindRoomBookingInconsistencies`

	queryRows := comment + `
WITH last AS ( 
	SELECT "roomId", MAX("dateEnd") "dateEnd" 
	FROM "bookings" GROUP BY 1 
) 
SELECT
	b."roomId",
	b."tenantId",
	r."currentTenantId",
	b."dateEnd",
	b."tenantId" <> r."currentTenantId" 
FROM last l 
LEFT JOIN "bookings" b 
	ON l."roomId" = b."roomId"
	AND b."dateEnd" = l."dateEnd" 
LEFT JOIN "rooms" r 
	ON l."roomId" = r."id"
`

	r.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 5 {
			out = append(out, RoomBookingInconsistency{
				RoomId:          X.ToU(row[0]),
				TenantId:        X.ToU(row[1]),
				CurrentTenantId: X.ToU(row[2]),
				DateEnd:         X.ToS(row[3]),
				IsInconsistent:  X.ToBool(row[4]),
			})
		}
	})

	return
}

func (r *Rooms) FixRoomBookingInconsistencies() {
	const comment = `-- Rooms) FixRoomBookingInconsistencies`

	query := comment + `
WITH last AS ( 
	SELECT "roomId", MAX("dateEnd") "dateEnd" 
	FROM "bookings" GROUP BY 1 
), lastBooking AS ( 
	SELECT l."roomId", b."tenantId" 
	FROM last l 
	LEFT JOIN "bookings" b 
		ON b."roomId" = l."roomId" 
		AND b."dateEnd" = l."dateEnd" 
) 
UPDATE "rooms" 
SET "currentTenantId" = COALESCE((
	SELECT lb."tenantId" 
	FROM lastBooking lb 
	WHERE lb."roomId" = "rooms"."id" 
),0)
WHERE "rooms"."deletedAt" = 0`

	r.Adapter.ExecSql(query)
}

func endOfYearMonth(year int, month time.Month) string { // returns: YYYY-MM-DD
	now := time.Now()
	firstOfNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, now.Location())
	return firstOfNextMonth.AddDate(0, 0, -1).Format("2006-01-02")
}

func startOfYearMonth(year int, month time.Month) string { // returns: YYYY-MM-DD
	now := time.Now()
	firstOfNextMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	return firstOfNextMonth.Format("2006-01-02")
}

func isValidMonthYear(in string) (year int, month time.Month, isValid bool) {
	yearMonthArr := S.Split(in, `-`)

	if len(yearMonthArr) != 2 {
		return
	}

	if _, err := time.Parse("2006-01", in); err != nil {
		return
	}

	year = int(S.ToI(yearMonthArr[0]))
	monthInt := S.ToI(yearMonthArr[1])

	if year < 1900 || year > 2100 {
		return
	}

	if monthInt < 1 || monthInt > 12 {
		return
	}

	month = time.Month(monthInt)

	isValid = true
	return
}

type ChartRevenueReport struct {
	Date       string `json:"date"`
	RevenueIDR int64  `json:"revenueIDR"`
}

func (p *Payments) GetChartRevenueReports(yearMonth string) (out []ChartRevenueReport) {
	const comment = `-- Payments) GetChartRevenueReports`

	var startDate string // YYYY-MM-DD
	var endDate string   // YYYY-MM-DD

	year, month, isValid := isValidMonthYear(yearMonth)
	if !isValid {
		now := time.Now()
		startDate = startOfYearMonth(now.Year(), now.Month())
		endDate = endOfYearMonth(now.Year(), now.Month())
	} else {
		startDate = startOfYearMonth(year, month)
		endDate = endOfYearMonth(year, month)
	}

	query := comment + `
SELECT
	"paymentAt",
	SUM(CASE 
		WHEN "paymentMethod" != 'Donation'
			THEN "paidIDR"
		ELSE 0
	END)
FROM "payments"
WHERE
	"deletedAt" = 0
	AND "paymentAt" >= ` + S.Z(startDate) + `
	AND "paymentAt" <= ` + S.Z(endDate) + `
GROUP BY "paymentAt"`

	p.Adapter.QuerySql(query, func(row []any) {
		if len(row) == 2 {
			out = append(out, ChartRevenueReport{
				Date:       X.ToS(row[0]),
				RevenueIDR: X.ToI(row[1]),
			})
		}
	})

	return
}

type UpcomingTenant struct {
	PrevTenant string `json:"prevTenant"`
	NextTenant string `json:"nextTenant"`
	RoomName   string `json:"roomName"`
	DateStart  string `json:"dateStart"`
	DateEnd    string `json:"dateEnd"`
}

func (b *Bookings) GetUpcomingTenants() (out []UpcomingTenant) {
	const comment = `-- Bookings) GetUpcomingTenants`

	dateNow := time.Now().Format(time.DateOnly)
	query := comment + `
SELECT
	r."roomName",
	(
		SELECT t2."tenantName"
		FROM "bookings" b2
		JOIN "tenants" t2 ON t2."id" = b2."tenantId"
		WHERE b2."roomId" = b."roomId"
			AND b2."tenantId" <> b."tenantId"
			AND b2."dateStart" < b."dateStart"
		ORDER BY b2."dateStart" DESC
		LIMIT 1
	) AS prev_tenant,
	t."tenantName" AS next_tenant,
	b."dateStart",
	b."dateEnd"
FROM "bookings" b
LEFT JOIN "rooms" r ON r."id" = b."roomId"
LEFT JOIN "tenants" t ON t."id" = b."tenantId"
WHERE b."dateStart" >= ` + S.Z(dateNow) + `
	AND b."deletedAt" = 0
	AND NOT EXISTS (
		SELECT 1
		FROM "bookings" prev
		WHERE prev."roomId" = b."roomId"
			AND prev."tenantId" = b."tenantId"
			AND prev."dateStart" < b."dateStart"
  )
ORDER BY b."dateStart" ASC`

	b.Adapter.QuerySql(query, func(row []any) {
		if len(row) != 5 {
			return
		}
		out = append(out, UpcomingTenant{
			RoomName:   X.ToS(row[0]),
			PrevTenant: X.ToS(row[1]),
			NextTenant: X.ToS(row[2]),
			DateStart:  X.ToS(row[3]),
			DateEnd:    X.ToS(row[4]),
		})
	})

	return
}

type PaymentOfBooking struct {
	BookingId uint64 `json:"bookingId"`
	PaymentId uint64 `json:"paymentId"`
	PaymentAt string `json:"paymentAt"`
	PaidIDR   int64  `json:"paidIDR"`
}

func (p *Payments) GetPaymentsByTenantId(tenantId uint64) (out []PaymentOfBooking) {
	const comment = `-- Payments) GetPaymentsByTenantId`

	query := comment + `
SELECT
	p."bookingId",
	p."id",
	p."paymentAt",
	p."paidIDR"
FROM "payments" p
LEFT JOIN "bookings" b
	ON b."id" = p."bookingId"
	AND b."tenantId" = ` + I.UToS(tenantId) + `
WHERE p."deletedAt" = 0
	AND b."deletedAt" = 0`

	p.Adapter.QuerySql(query, func(row []any) {
		if len(row) != 4 {
			return
		}

		out = append(out, PaymentOfBooking{
			BookingId: X.ToU(row[0]),
			PaymentId: X.ToU(row[1]),
			PaymentAt: X.ToS(row[2]),
			PaidIDR:   X.ToI(row[3]),
		})
	})

	return
}

type TenantBookingDetail struct {
	BookingId     uint64 `json:"bookingId"`
	TotalPaidIDR  int64  `json:"totalPaidIDR"`
	TotalPriceIDR int64  `json:"totalPriceIDR"`
	DateStart     string `json:"dateStart"`
	DateEnd       string `json:"dateEnd"`
	RoomName      string `json:"roomName"`
}

func (b *Bookings) GetBookingsByTenantId(tenantId uint64) (out []TenantBookingDetail) {
	const comment = `-- Bookings) GetBookingsByTenantId`

	query := comment + `
SELECT
	b."id",
	COALESCE(SUM(p."paidIDR"), 0) AS totalPaidIDR,
	b."totalPriceIDR",
	b."dateStart",
	b."dateEnd",
	r."roomName"
FROM "bookings" b
LEFT JOIN "payments" p ON b."id" = p."bookingId"
LEFT JOIN "rooms" r ON b."roomId" = r."id"
WHERE b."tenantId" = ` + I.UToS(tenantId) + `
GROUP BY b."id"
ORDER BY b."dateStart" ASC`

	b.Adapter.QuerySql(query, func(row []any) {
		if len(row) != 6 {
			return
		}

		out = append(out, TenantBookingDetail{
			BookingId:     X.ToU(row[0]),
			TotalPaidIDR:  X.ToI(row[1]),
			TotalPriceIDR: X.ToI(row[2]),
			DateStart:     X.ToS(row[3]),
			DateEnd:       X.ToS(row[4]),
			RoomName:      X.ToS(row[5]),
		})
	})

	return
}

func (w *WifiDevices) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := w.Adapter.Select(w.SpaceName(), w.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query wifi devices`) {
		return
	}

	res = resp.Tuples()

	return
}

func (w *WifiDevices) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + w.SqlTableName() + `
	LIMIT 1`

	w.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (w *WifiDevices) Truncate() bool {
	return w.Adapter.ExecBoxSpace(w.SpaceName()+`:truncate`, A.X{})
}

type PricePerDayReport struct {
	RoomName    string  `json:"roomName"`
	TenantName  string  `json:"tenantName"`
	DateStart   string  `json:"dateStart"`
	DateEnd     string  `json:"dateEnd"`
	PricePerDay float64 `json:"pricePerDay"`
	TotalPaid   int64   `json:"totalPaid"`
	TotalPrice  int64   `json:"totalPrice"`
}

func (b *Bookings) FindPricePerDayReport(yearMonth string) (out []PricePerDayReport) {
	const comment = `-- Bookings) FindPricePerDayReport`

	_, _, isValid := isValidMonthYear(yearMonth)
	if !isValid {
		now := time.Now()
		yearMonth = startOfYearMonth(now.Year(), now.Month())
	}

	query := comment + `
SELECT
	r."roomName",
	t."tenantName",
	b."dateStart",
	b."dateEnd",
	COALESCE(SUM(p."paidIDR"), 0) AS totalPaidIDR,
	b."totalPriceIDR"
FROM "bookings" b
LEFT JOIN "payments" p ON b."id" = p."bookingId"
LEFT JOIN "rooms" r ON b."roomId" = r."id"
LEFT JOIN "tenants" t ON r."currentTenantId" = t."id"
WHERE
	b."deletedAt" = 0
	AND SUBSTR(b."dateStart", 1, 7) = '` + yearMonth + `'
GROUP BY b."id"`

	b.Adapter.QuerySql(query, func(row []any) {
		if len(row) != 6 {
			return
		}

		out = append(out, PricePerDayReport{
			RoomName:   X.ToS(row[0]),
			TenantName: X.ToS(row[1]),
			DateStart:  X.ToS(row[2]),
			DateEnd:    X.ToS(row[3]),
			TotalPaid:  X.ToI(row[4]),
			TotalPrice: X.ToI(row[5]),
		})
	})

	return
}
