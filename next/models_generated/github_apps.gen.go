// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models_generated

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/diggerhq/digger/next/model"
)

func newGithubApp(db *gorm.DB, opts ...gen.DOOption) githubApp {
	_githubApp := githubApp{}

	_githubApp.githubAppDo.UseDB(db, opts...)
	_githubApp.githubAppDo.UseModel(&model.GithubApp{})

	tableName := _githubApp.githubAppDo.TableName()
	_githubApp.ALL = field.NewAsterisk(tableName)
	_githubApp.ID = field.NewString(tableName, "id")
	_githubApp.CreatedAt = field.NewTime(tableName, "created_at")
	_githubApp.UpdatedAt = field.NewTime(tableName, "updated_at")
	_githubApp.DeletedAt = field.NewField(tableName, "deleted_at")
	_githubApp.GithubID = field.NewInt64(tableName, "github_id")
	_githubApp.Name = field.NewString(tableName, "name")
	_githubApp.GithubAppURL = field.NewString(tableName, "github_app_url")

	_githubApp.fillFieldMap()

	return _githubApp
}

type githubApp struct {
	githubAppDo

	ALL          field.Asterisk
	ID           field.String
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	GithubID     field.Int64
	Name         field.String
	GithubAppURL field.String

	fieldMap map[string]field.Expr
}

func (g githubApp) Table(newTableName string) *githubApp {
	g.githubAppDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g githubApp) As(alias string) *githubApp {
	g.githubAppDo.DO = *(g.githubAppDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *githubApp) updateTableName(table string) *githubApp {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewString(table, "id")
	g.CreatedAt = field.NewTime(table, "created_at")
	g.UpdatedAt = field.NewTime(table, "updated_at")
	g.DeletedAt = field.NewField(table, "deleted_at")
	g.GithubID = field.NewInt64(table, "github_id")
	g.Name = field.NewString(table, "name")
	g.GithubAppURL = field.NewString(table, "github_app_url")

	g.fillFieldMap()

	return g
}

func (g *githubApp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *githubApp) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 7)
	g.fieldMap["id"] = g.ID
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
	g.fieldMap["deleted_at"] = g.DeletedAt
	g.fieldMap["github_id"] = g.GithubID
	g.fieldMap["name"] = g.Name
	g.fieldMap["github_app_url"] = g.GithubAppURL
}

func (g githubApp) clone(db *gorm.DB) githubApp {
	g.githubAppDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g githubApp) replaceDB(db *gorm.DB) githubApp {
	g.githubAppDo.ReplaceDB(db)
	return g
}

type githubAppDo struct{ gen.DO }

type IGithubAppDo interface {
	gen.SubQuery
	Debug() IGithubAppDo
	WithContext(ctx context.Context) IGithubAppDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGithubAppDo
	WriteDB() IGithubAppDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGithubAppDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGithubAppDo
	Not(conds ...gen.Condition) IGithubAppDo
	Or(conds ...gen.Condition) IGithubAppDo
	Select(conds ...field.Expr) IGithubAppDo
	Where(conds ...gen.Condition) IGithubAppDo
	Order(conds ...field.Expr) IGithubAppDo
	Distinct(cols ...field.Expr) IGithubAppDo
	Omit(cols ...field.Expr) IGithubAppDo
	Join(table schema.Tabler, on ...field.Expr) IGithubAppDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGithubAppDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGithubAppDo
	Group(cols ...field.Expr) IGithubAppDo
	Having(conds ...gen.Condition) IGithubAppDo
	Limit(limit int) IGithubAppDo
	Offset(offset int) IGithubAppDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGithubAppDo
	Unscoped() IGithubAppDo
	Create(values ...*model.GithubApp) error
	CreateInBatches(values []*model.GithubApp, batchSize int) error
	Save(values ...*model.GithubApp) error
	First() (*model.GithubApp, error)
	Take() (*model.GithubApp, error)
	Last() (*model.GithubApp, error)
	Find() ([]*model.GithubApp, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GithubApp, err error)
	FindInBatches(result *[]*model.GithubApp, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GithubApp) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGithubAppDo
	Assign(attrs ...field.AssignExpr) IGithubAppDo
	Joins(fields ...field.RelationField) IGithubAppDo
	Preload(fields ...field.RelationField) IGithubAppDo
	FirstOrInit() (*model.GithubApp, error)
	FirstOrCreate() (*model.GithubApp, error)
	FindByPage(offset int, limit int) (result []*model.GithubApp, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGithubAppDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g githubAppDo) Debug() IGithubAppDo {
	return g.withDO(g.DO.Debug())
}

func (g githubAppDo) WithContext(ctx context.Context) IGithubAppDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g githubAppDo) ReadDB() IGithubAppDo {
	return g.Clauses(dbresolver.Read)
}

func (g githubAppDo) WriteDB() IGithubAppDo {
	return g.Clauses(dbresolver.Write)
}

func (g githubAppDo) Session(config *gorm.Session) IGithubAppDo {
	return g.withDO(g.DO.Session(config))
}

func (g githubAppDo) Clauses(conds ...clause.Expression) IGithubAppDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g githubAppDo) Returning(value interface{}, columns ...string) IGithubAppDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g githubAppDo) Not(conds ...gen.Condition) IGithubAppDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g githubAppDo) Or(conds ...gen.Condition) IGithubAppDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g githubAppDo) Select(conds ...field.Expr) IGithubAppDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g githubAppDo) Where(conds ...gen.Condition) IGithubAppDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g githubAppDo) Order(conds ...field.Expr) IGithubAppDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g githubAppDo) Distinct(cols ...field.Expr) IGithubAppDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g githubAppDo) Omit(cols ...field.Expr) IGithubAppDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g githubAppDo) Join(table schema.Tabler, on ...field.Expr) IGithubAppDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g githubAppDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGithubAppDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g githubAppDo) RightJoin(table schema.Tabler, on ...field.Expr) IGithubAppDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g githubAppDo) Group(cols ...field.Expr) IGithubAppDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g githubAppDo) Having(conds ...gen.Condition) IGithubAppDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g githubAppDo) Limit(limit int) IGithubAppDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g githubAppDo) Offset(offset int) IGithubAppDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g githubAppDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGithubAppDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g githubAppDo) Unscoped() IGithubAppDo {
	return g.withDO(g.DO.Unscoped())
}

func (g githubAppDo) Create(values ...*model.GithubApp) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g githubAppDo) CreateInBatches(values []*model.GithubApp, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g githubAppDo) Save(values ...*model.GithubApp) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g githubAppDo) First() (*model.GithubApp, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GithubApp), nil
	}
}

func (g githubAppDo) Take() (*model.GithubApp, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GithubApp), nil
	}
}

func (g githubAppDo) Last() (*model.GithubApp, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GithubApp), nil
	}
}

func (g githubAppDo) Find() ([]*model.GithubApp, error) {
	result, err := g.DO.Find()
	return result.([]*model.GithubApp), err
}

func (g githubAppDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GithubApp, err error) {
	buf := make([]*model.GithubApp, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g githubAppDo) FindInBatches(result *[]*model.GithubApp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g githubAppDo) Attrs(attrs ...field.AssignExpr) IGithubAppDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g githubAppDo) Assign(attrs ...field.AssignExpr) IGithubAppDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g githubAppDo) Joins(fields ...field.RelationField) IGithubAppDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g githubAppDo) Preload(fields ...field.RelationField) IGithubAppDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g githubAppDo) FirstOrInit() (*model.GithubApp, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GithubApp), nil
	}
}

func (g githubAppDo) FirstOrCreate() (*model.GithubApp, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GithubApp), nil
	}
}

func (g githubAppDo) FindByPage(offset int, limit int) (result []*model.GithubApp, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g githubAppDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g githubAppDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g githubAppDo) Delete(models ...*model.GithubApp) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *githubAppDo) withDO(do gen.Dao) *githubAppDo {
	g.DO = *do.(*gen.DO)
	return g
}
