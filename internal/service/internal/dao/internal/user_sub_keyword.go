// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2023-08-06 17:07:07
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserSubKeywordDao is the data access object for table user_sub_keyword.
type UserSubKeywordDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns UserSubKeywordColumns // columns contains all the column names of Table for convenient usage.
}

// UserSubKeywordColumns defines and stores column names for table user_sub_keyword.
type UserSubKeywordColumns struct {
	Id          string //
	UserId      string //
	Keyword     string //
	OrderByDate string //
	CreateTime  string //
	Lang        string // feed language
	Status      string //
}

//  userSubKeywordColumns holds the columns for table user_sub_keyword.
var userSubKeywordColumns = UserSubKeywordColumns{
	Id:          "id",
	UserId:      "user_id",
	Keyword:     "keyword",
	OrderByDate: "order_by_date",
	CreateTime:  "create_time",
	Lang:        "lang",
	Status:      "status",
}

// NewUserSubKeywordDao creates and returns a new DAO object for table data access.
func NewUserSubKeywordDao() *UserSubKeywordDao {
	return &UserSubKeywordDao{
		group:   "default",
		table:   "user_sub_keyword",
		columns: userSubKeywordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserSubKeywordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserSubKeywordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserSubKeywordDao) Columns() UserSubKeywordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserSubKeywordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserSubKeywordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserSubKeywordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
