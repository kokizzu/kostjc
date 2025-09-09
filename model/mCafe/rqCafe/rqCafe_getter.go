package rqCafe

import (
	"kostjc/model/zCrud"

	"github.com/kokizzu/gotro/X"
)

func (l *Menus) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Menus) FindByPagination`

	validFields := MenusFieldTypeMap
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

	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (l *Menus) FindAll() []Menus {
	const comment = `-- Menus) FindAll`

	queryRows := comment + `
SELECT ` + l.SqlSelectAllFields() + ` FROM ` + l.SqlTableName() + `
WHERE "deletedAt" = 0`

	var rows = []Menus{}
	l.Adapter.QuerySql(queryRows, func(row []any) {
		l.FromArray(row)
		rows = append(rows, *l)
	})

	return rows
}

func (f *Menus) FindMenus() map[uint64]string {
	const comment = `-- Menus) FindMenus`

	queryRows := comment + `
SELECT ` + f.SqlId() + `, ` + f.SqlName() + `, ` + f.SqlSalePriceIDR() + ` FROM ` + f.SqlTableName() + `
WHERE "deletedAt" = 0
ORDER BY ` + f.SqlId() + ` ASC`

	out := make(map[uint64]string)
	f.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 3 {
			menuName := X.ToS(row[1])
			price := X.ToS(row[2])
			out[X.ToU(row[0])] = menuName + " (" + price + ")"
		}
	})

	return out
}

func (f *Menus) FindMenusSalesChoices() map[uint64]string {
	const comment = `-- Menus) FindMenusSaleChoices`

	queryRows := comment + `
SELECT ` + f.SqlId() + `, ` + f.SqlName() + `, ` + f.SqlSalePriceIDR() + `, ` + f.SqlImageUrl() + ` FROM ` + f.SqlTableName() + `
WHERE "deletedAt" = 0
ORDER BY ` + f.SqlId() + ` ASC`

	out := make(map[uint64]string)
	f.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 4 {
			menuName := X.ToS(row[1])
			price := X.ToS(row[2])
			imageUrl := X.ToS(row[3])
			out[X.ToU(row[0])] = menuName + " (" + price + ")" + " (" + imageUrl + ")"
		}
	})

	return out
}

func (l *Sales) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Sales) FindByPagination`

	validFields := SalesFieldTypeMap
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

	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (l *Sales) FindAll() []Sales {
	const comment = `-- Sales) FindAll`

	queryRows := comment + `
SELECT ` + l.SqlSelectAllFields() + ` FROM ` + l.SqlTableName() + `
WHERE "deletedAt" = 0`

	var rows = []Sales{}
	l.Adapter.QuerySql(queryRows, func(row []any) {
		l.FromArray(row)
		rows = append(rows, *l)
	})

	return rows
}

func (l *BorrowedUtensils) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- BorrowedUtensils) FindByPagination`

	validFields := BorrowedUtensilsFieldTypeMap
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

	out.Order = in.Order
	out.Filters = in.Filters

	return
}

func (l *BorrowedUtensils) FindAll() []BorrowedUtensils {
	const comment = `-- BorrowedUtensils) FindAll`

	queryRows := comment + `
SELECT ` + l.SqlSelectAllFields() + ` FROM ` + l.SqlTableName() + `
WHERE "deletedAt" = 0`

	var rows = []BorrowedUtensils{}
	l.Adapter.QuerySql(queryRows, func(row []any) {
		l.FromArray(row)
		rows = append(rows, *l)
	})

	return rows
}

func (l *Laundry) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Laundry) FindByPagination`

	validFields := LaundryFieldTypeMap
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
