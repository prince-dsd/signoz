package sqlite

import (
	"fmt"

	basedao "github.com/SigNoz/signoz/pkg/query-service/dao"
	basedsql "github.com/SigNoz/signoz/pkg/query-service/dao/sqlite"
	baseint "github.com/SigNoz/signoz/pkg/query-service/interfaces"
	"github.com/SigNoz/signoz/pkg/sqlstore"
	"github.com/uptrace/bun"
)

type modelDao struct {
	*basedsql.ModelDaoSqlite
	flags baseint.FeatureLookup
}

// SetFlagProvider sets the feature lookup provider
func (m *modelDao) SetFlagProvider(flags baseint.FeatureLookup) {
	m.flags = flags
}

// CheckFeature confirms if a feature is available
func (m *modelDao) checkFeature(key string) error {
	if m.flags == nil {
		return fmt.Errorf("flag provider not set")
	}

	return m.flags.CheckFeature(key)
}

// InitDB creates and extends base model DB repository
func InitDB(sqlStore sqlstore.SQLStore) (*modelDao, error) {
	dao, err := basedsql.InitDB(sqlStore)
	if err != nil {
		return nil, err
	}
	// set package variable so dependent base methods (e.g. AuthCache)  will work
	basedao.SetDB(dao)
	m := &modelDao{ModelDaoSqlite: dao}
	return m, nil
}

func (m *modelDao) DB() *bun.DB {
	return m.ModelDaoSqlite.DB()
}
