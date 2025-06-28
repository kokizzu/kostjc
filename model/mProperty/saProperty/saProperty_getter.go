package saProperty

import (
	"kostjc/model/zCrud"

	"github.com/kokizzu/gotro/L"
)

func (l LocationLogs) FindByPagination(in *zCrud.PagerIn, out *zCrud.PagerOut) (res []LocationLogs) {
	const comment = `-- LocationLogs) FindByPagination`

	validFields := LocationLogsFieldTypeMap
	whereAndSql := out.WhereAndSqlCh(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + l.SqlTableName() + whereAndSql + `
LIMIT 1`
	row := l.Adapter.QueryRow(queryCount)
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
SELECT ` + l.SqlAllFields() + `
FROM ` + l.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	rows, err := l.Adapter.Query(queryRows)
	if L.IsError(err, `FindByPagination.Adapter.Query error: `+queryRows) {
		return
	}

	res, err = l.ScanRowsAllCols(rows, out.PerPage)
	if L.IsError(err, `FindByPagination.ScanRowsAllCols error: `+queryRows) {
		return
	}

	return
}

func (f FacilityLogs) FindByPagination(in *zCrud.PagerIn, out *zCrud.PagerOut) (res []FacilityLogs) {
	const comment = `-- FacilityLogs) FindByPagination`

	validFields := FacilityLogsFieldTypeMap
	whereAndSql := out.WhereAndSqlCh(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + f.SqlTableName() + whereAndSql + `
LIMIT 1`
	row := f.Adapter.QueryRow(queryCount)
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
SELECT ` + f.SqlAllFields() + `
FROM ` + f.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	rows, err := f.Adapter.Query(queryRows)
	if L.IsError(err, `FindByPagination.Adapter.Query error: `+queryRows) {
		return
	}

	res, err = f.ScanRowsAllCols(rows, out.PerPage)
	if L.IsError(err, `FindByPagination.ScanRowsAllCols error: `+queryRows) {
		return
	}

	return
}

func (b BuildingLogs) FindByPagination(in *zCrud.PagerIn, out *zCrud.PagerOut) (res []BuildingLogs) {
	const comment = `-- BuildingLogs) FindByPagination`

	validFields := BuildingLogsFieldTypeMap
	whereAndSql := out.WhereAndSqlCh(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + b.SqlTableName() + whereAndSql + `
LIMIT 1`
	row := b.Adapter.QueryRow(queryCount)
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
SELECT ` + b.SqlAllFields() + `
FROM ` + b.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	rows, err := b.Adapter.Query(queryRows)
	if L.IsError(err, `FindByPagination.Adapter.Query error: `+queryRows) {
		return
	}

	res, err = b.ScanRowsAllCols(rows, out.PerPage)
	if L.IsError(err, `FindByPagination.ScanRowsAllCols error: `+queryRows) {
		return
	}

	return
}

func (r RoomLogs) FindByPagination(in *zCrud.PagerIn, out *zCrud.PagerOut) (res []RoomLogs) {
	const comment = `-- RoomLogs) FindByPagination`

	validFields := RoomLogsFieldTypeMap
	whereAndSql := out.WhereAndSqlCh(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + r.SqlTableName() + whereAndSql + `
LIMIT 1`
	row := r.Adapter.QueryRow(queryCount)
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
SELECT ` + r.SqlAllFields() + `
FROM ` + r.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	rows, err := r.Adapter.Query(queryRows)
	if L.IsError(err, `FindByPagination.Adapter.Query error: `+queryRows) {
		return
	}

	res, err = r.ScanRowsAllCols(rows, out.PerPage)
	if L.IsError(err, `FindByPagination.ScanRowsAllCols error: `+queryRows) {
		return
	}

	return
}

func (b BookingLogs) FindByPagination(in *zCrud.PagerIn, out *zCrud.PagerOut) (res []BookingLogs) {
	const comment = `-- BookingLogs) FindByPagination`

	validFields := BookingLogsFieldTypeMap
	whereAndSql := out.WhereAndSqlCh(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + b.SqlTableName() + whereAndSql + `
LIMIT 1`
	row := b.Adapter.QueryRow(queryCount)
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
SELECT ` + b.SqlAllFields() + `
FROM ` + b.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	rows, err := b.Adapter.Query(queryRows)
	if L.IsError(err, `FindByPagination.Adapter.Query error: `+queryRows) {
		return
	}

	res, err = b.ScanRowsAllCols(rows, out.PerPage)
	if L.IsError(err, `FindByPagination.ScanRowsAllCols error: `+queryRows) {
		return
	}

	return
}

func (p PaymentLogs) FindByPagination(in *zCrud.PagerIn, out *zCrud.PagerOut) (res []PaymentLogs) {
	const comment = `-- PaymentLogs) FindByPagination`

	validFields := PaymentLogsFieldTypeMap
	whereAndSql := out.WhereAndSqlCh(in.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + p.SqlTableName() + whereAndSql + `
LIMIT 1`
	row := p.Adapter.QueryRow(queryCount)
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
SELECT ` + p.SqlAllFields() + `
FROM ` + p.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	rows, err := p.Adapter.Query(queryRows)
	if L.IsError(err, `FindByPagination.Adapter.Query error: `+queryRows) {
		return
	}

	res, err = p.ScanRowsAllCols(rows, out.PerPage)
	if L.IsError(err, `FindByPagination.ScanRowsAllCols error: `+queryRows) {
		return
	}

	return
}

func (s StockLogs) FindByPagination(in *zCrud.PagerIn, out *zCrud.PagerOut) (res []StockLogs) {
	const comment = `-- StockLogs) FindByPagination`

	validFields := StockLogsFieldTypeMap
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

func (s WifiDeviceLogs) FindByPagination(in *zCrud.PagerIn, out *zCrud.PagerOut) (res []WifiDeviceLogs) {
	const comment = `-- WifiDeviceLogs) FindByPagination`

	validFields := WifiDeviceLogsFieldTypeMap
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
