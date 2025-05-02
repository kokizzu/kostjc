package rqProperty

import (
	"kostjc/model/zCrud"
	"strconv"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
	"github.com/tarantool/go-tarantool"
)

func (l *Locations) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Locations) FindByPagination`

	validFields := LocationsFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

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

	return
}

func (l *Locations) FindAll() []Locations {
	const comment = `-- Locations) FindAll`

	queryRows := comment + `
SELECT ` + l.SqlSelectAllFields() + ` FROM ` + l.SqlTableName()

	var rows = []Locations{}
	l.Adapter.QuerySql(queryRows, func(row []any) {
		l.FromArray(row)
		rows = append(rows, *l)
	})

	return rows
}

func (f *Facilities) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Facilities) FindByPagination`

	validFields := LocationsFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

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

	return
}

func (b *Buildings) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Buildings) FindByPagination`

	validFields := LocationsFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

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

	return
}

func (l *Locations) FindLocationChoices() map[uint64]string {
	const comment = `-- Locations) FindLocationChoices`

	queryRows := comment + `
SELECT ` + l.SqlId() + `, ` + l.SqlName() + ` FROM ` + l.SqlTableName() + `
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

	validFields := LocationsFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

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
WHERE ` + f.SqlFacilityType() + ` = 'Building'`

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
WHERE ` + f.SqlFacilityType() + ` = 'Room'`

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

	validFields := LocationsFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

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

	return
}

func (b *Buildings) FindBuildingChoices() map[uint64]string {
	const comment = `-- Buildings) FindBuildingChoices`

	queryRows := comment + `
SELECT ` + b.SqlId() + `, ` + b.SqlBuildingName() + ` FROM ` + b.SqlTableName() + `
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
SELECT ` + b.SqlId() + `, ` + b.SqlTotalPriceIDR() + ` FROM ` + b.SqlTableName() + `
ORDER BY ` + b.SqlId() + ` ASC`

	out := make(map[uint64]string)
	b.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 2 {
			idWithPrice := `#` + X.ToS(row[0]) + ` - ` + formatCurrency(X.ToI(row[1]), `IDR`)
			out[X.ToU(row[0])] = idWithPrice
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
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

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

	return
}

func (r *Rooms) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Rooms) FindByPagination`

	validFields := RoomsFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

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
