package rqAuth

import (
	"time"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
	"github.com/tarantool/go-tarantool"

	"kostjc/model/zCrud"
)

func (u *Users) CheckPassword(pass string) error {
	return S.CheckPassword(u.Password, pass)
}

func (s *Sessions) AllActiveSession(userId uint64, now int64) (res []*Sessions) {
	const comment = `-- Sessions) AllActiveSession`

	query := comment + `
SELECT ` + s.SqlSelectAllFields() + `
FROM ` + s.SqlTableName() + `
WHERE ` + s.SqlUserId() + ` = ` + I.UToS(userId) + `
	AND ` + s.SqlExpiredAt() + ` > ` + I.ToS(now)
	s.Adapter.QuerySql(query, func(row []any) {
		rec := &Sessions{}
		res = append(res, rec.FromArray(row))
	})
	return
}

func (u *Users) CountUserRegisterToday() (res int64) {
	const comment = `-- Users) CountUserRegisterToday`
	currentDate := time.Now()
	beginCurrentDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, currentDate.Location())
	endCurrentDate := beginCurrentDate.AddDate(0, 0, 1).Add(-time.Second)

	beginDateUnix := I.ToS(beginCurrentDate.Unix())
	endDateUnix := I.ToS(endCurrentDate.Unix())

	queryCountRegisteredToday := comment + `
SELECT COUNT(*) 
FROM ` + u.SqlTableName() + `
WHERE ` + u.SqlCreatedAt() + ` >= ` + beginDateUnix + ` 
AND ` + u.SqlCreatedAt() + ` <= ` + endDateUnix

	u.Adapter.QuerySql(queryCountRegisteredToday, func(row []any) {
		res = X.ToI(row[0])
	})

	return res
}

func (u *Users) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Users) FindByPagination`

	validFields := UsersFieldTypeMap
	whereAndSql := ``
	if in.Search != `` {
		whereAndSql = out.SearchBySqlTt(in.Search, in.SearchBy, validFields)
	} else {
		whereAndSql = out.WhereAndSqlTt(in.Filters, validFields)
	}

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + u.SqlTableName() + whereAndSql + `
LIMIT 1`
	u.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + u.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	u.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return
}

// EmailMapByIds if using this function, please make sure the string is not from frontend/user
// if from frontend, please do check for sql injection
func (u *Users) EmailMapByIds(keys string) map[string]string {
	const comment = `-- Users) EmailMapByIds`
	queryRows := comment + `
SELECT ` + u.SqlId() + `
	, ` + u.SqlEmail() + `
FROM ` + u.SqlTableName() + `
WHERE ` + u.SqlId() + ` IN (` + keys + `)`

	res := map[string]string{}
	u.Adapter.QuerySql(queryRows, func(row []any) {
		res[X.ToS(row[0])] = X.ToS(row[1])
	})
	return res
}

func (t *Tenants) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Locations) FindByPagination`

	validFields := TenantsFieldTypeMap
	whereAndSql := ``
	if in.Search != `` {
		whereAndSql = out.SearchBySqlTt(in.Search, in.SearchBy, validFields)
	} else {
		whereAndSql = out.WhereAndSqlTt(in.Filters, validFields)
	}

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + t.SqlTableName() + whereAndSql + `
LIMIT 1`
	t.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + t.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql

	t.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return
}

func (t *Tenants) FindTenantChoices() map[uint64]string {
	const comment = `-- Tenants) FindTenantChoices`

	queryRows := comment + `
SELECT ` + t.SqlId() + `, ` + t.SqlTenantName() + `, ` + t.SqlTelegramUsername() + `, ` + t.SqlWhatsappNumber() + ` FROM ` + t.SqlTableName() + `
ORDER BY ` + t.SqlTenantName() + ` ASC`

	out := make(map[uint64]string)
	t.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 4 {
			nameToShow := X.ToS(row[1])
			if X.ToS(row[2]) != "" {
				nameToShow += ` / ` + X.ToS(row[2])
			}

			if X.ToS(row[3]) != "" {
				nameToShow += ` / ` + X.ToS(row[3])
			}

			out[X.ToU(row[0])] = nameToShow
		}
	})

	return out
}

func (u *Users) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := u.Adapter.Select(u.SpaceName(), u.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (u *Users) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + u.SqlTableName() + `
	LIMIT 1`

	u.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (u *Users) Truncate() bool {
	return u.Adapter.ExecBoxSpace(u.SpaceName()+`:truncate`, A.X{})
}

func (t *Tenants) GetRows(offset, limit uint32) (res [][]any) {
	resp, err := t.Adapter.Select(t.SpaceName(), t.UniqueIndexId(), offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `failed to query property`) {
		return
	}

	res = resp.Tuples()

	return
}

func (t *Tenants) CountTotalAllRows() (total uint64) {
	queryCount := `
	SELECT COUNT(1)
	FROM ` + t.SqlTableName() + `
	LIMIT 1`

	t.Adapter.QuerySql(queryCount, func(row []any) {
		if len(row) >= 1 {
			total = X.ToU(row[0])
		}
	})

	return
}

func (t *Tenants) Truncate() bool {
	return t.Adapter.ExecBoxSpace(t.SpaceName()+`:truncate`, A.X{})
}

func (s *Sessions) Truncate() bool {
	return s.Adapter.ExecBoxSpace(s.SpaceName()+`:truncate`, A.X{})
}

func (u *Users) FindUserChoices() map[uint64]string {
	const comment = `-- Users) FindUserChoices`

	queryRows := comment + `
SELECT ` + u.SqlId() + `, ` + u.SqlFullName() + `, ` + u.SqlUserName() + `
FROM ` + u.SqlTableName() + `
ORDER BY ` + u.SqlId() + ` ASC`

	out := make(map[uint64]string)
	u.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) != 3 {
			return
		}

		nameToShow := X.ToS(row[1])
		if nameToShow == "" {
			nameToShow = X.ToS(row[2])
		}

		out[X.ToU(row[0])] = nameToShow
	})

	return out
}
