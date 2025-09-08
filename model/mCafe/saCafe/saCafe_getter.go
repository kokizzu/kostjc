package saCafe

import (
	"kostjc/model/zCrud"

	"github.com/kokizzu/gotro/L"
)

func (m MenuLogs) FindByPagination(in *zCrud.PagerIn, out *zCrud.PagerOut) (res []MenuLogs) {
	const comment = `-- MenuLogs) FindByPagination`

	validFields := MenuLogsFieldTypeMap
	whereAndSql := out.WhereAndSqlCh(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + m.SqlTableName() + whereAndSql + `
LIMIT 1`
	row := m.Adapter.QueryRow(queryCount)
	err := row.Err()
	if L.IsError(err, `FindByPagination.Adapter.QueryRow error: `+queryCount) {
		return
	}
	err = row.Scan(&out.Total)
	L.IsError(err, comment+`: error while scanning total`)
	out.CalculatePages(in.Page, in.PerPage, out.Total)

	orderBySql := out.OrderBySqlCh(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + m.SqlAllFields() + `
FROM ` + m.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	rows, err := m.Adapter.Query(queryRows)
	if L.IsError(err, `FindByPagination.Adapter.Query error: `+queryRows) {
		return
	}

	res, err = m.ScanRowsAllCols(rows, out.PerPage)
	if L.IsError(err, `FindByPagination.ScanRowsAllCols error: `+queryRows) {
		return
	}

	return
}

func (s SaleLogs) FindByPagination(in *zCrud.PagerIn, out *zCrud.PagerOut) (res []SaleLogs) {
	const comment = `-- SaleLogs) FindByPagination`

	validFields := SaleLogsFieldTypeMap
	whereAndSql := out.WhereAndSqlCh(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + s.SqlTableName() + whereAndSql + `
LIMIT 1`
	row := s.Adapter.QueryRow(queryCount)
	err := row.Err()
	if L.IsError(err, `FindByPagination.Adapter.QueryRow error: `+queryCount) {
		return
	}
	err = row.Scan(&out.Total)
	L.IsError(err, comment+`: error while scanning total`)
	out.CalculatePages(in.Page, in.PerPage, out.Total)

	orderBySql := out.OrderBySqlCh(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + s.SqlAllFields() + `
FROM ` + s.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	rows, err := s.Adapter.Query(queryRows)
	if L.IsError(err, `FindByPagination.Adapter.Query error: `+queryRows) {
		return
	}

	res, err = s.ScanRowsAllCols(rows, out.PerPage)
	if L.IsError(err, `FindByPagination.ScanRowsAllCols error: `+queryRows) {
		return
	}

	return
}

func (s BorrowedUtensilLogs) FindByPagination(in *zCrud.PagerIn, out *zCrud.PagerOut) (res []BorrowedUtensilLogs) {
	const comment = `-- BorrowedUtensilLogs) FindByPagination`

	validFields := BorrowedUtensilLogsFieldTypeMap
	whereAndSql := out.WhereAndSqlCh(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + s.SqlTableName() + whereAndSql + `
LIMIT 1`
	row := s.Adapter.QueryRow(queryCount)
	err := row.Err()
	if L.IsError(err, `FindByPagination.Adapter.QueryRow error: `+queryCount) {
		return
	}
	err = row.Scan(&out.Total)
	L.IsError(err, comment+`: error while scanning total`)
	out.CalculatePages(in.Page, in.PerPage, out.Total)

	orderBySql := out.OrderBySqlCh(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + s.SqlAllFields() + `
FROM ` + s.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	rows, err := s.Adapter.Query(queryRows)
	if L.IsError(err, `FindByPagination.Adapter.Query error: `+queryRows) {
		return
	}

	res, err = s.ScanRowsAllCols(rows, out.PerPage)
	if L.IsError(err, `FindByPagination.ScanRowsAllCols error: `+queryRows) {
		return
	}

	return
}
