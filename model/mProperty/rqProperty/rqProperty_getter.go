package rqProperty

import (
	"kostjc/model/zCrud"

	"github.com/kokizzu/gotro/X"
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
