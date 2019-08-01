package workfx

import (
	"database/sql"
	"testing"

	"github.com/freerware/work"
	"github.com/stretchr/testify/require"
	"github.com/uber-go/tally"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
)

type dataMapper struct{}

func (dm dataMapper) Insert(e ...interface{}) error { return nil }
func (dm dataMapper) Update(e ...interface{}) error { return nil }
func (dm dataMapper) Delete(e ...interface{}) error { return nil }

type sqlDataMapper struct{}

func (dm sqlDataMapper) Insert(tx *sql.Tx, e ...interface{}) error { return nil }
func (dm sqlDataMapper) Update(tx *sql.Tx, e ...interface{}) error { return nil }
func (dm sqlDataMapper) Delete(tx *sql.Tx, e ...interface{}) error { return nil }

type entity struct{}

type sqlUniterParameters struct {
	fx.In

	Uniter work.Uniter `name:"sqlWorkUniter"`
}

type bestEffortUniterParameters struct {
	fx.In

	Uniter work.Uniter `name:"bestEffortWorkUniter"`
}

func TestSQLUnitModule(t *testing.T) {

	//arrange.
	type result struct {
		fx.Out

		Mappers map[work.TypeName]work.SQLDataMapper
		DB      *sql.DB `name:"rwDB"`
		Logger  *zap.Logger
		Scope   tally.Scope
	}

	unitDeps := func() result {
		mappers := make(map[work.TypeName]work.SQLDataMapper)

		dm := sqlDataMapper{}
		t := work.TypeNameOf(entity{})
		mappers[t] = &dm
		var db *sql.DB
		l := zap.NewExample()
		s := tally.NewTestScope("test", map[string]string{})
		return result{
			Mappers: mappers,
			DB:      db,
			Logger:  l,
			Scope:   s,
		}
	}
	var uniter *work.Uniter

	//action.
	fxtest.New(
		t,
		fx.Provide(unitDeps),
		Modules.SQLUnit,
		fx.Invoke(func(p sqlUniterParameters) {
			uniter = &p.Uniter
		}),
	).RequireStart().RequireStop()

	//assert.
	require.NotNil(t, uniter)
}

func TestBestEffortUnit(t *testing.T) {

	//arrange.
	type result struct {
		fx.Out

		Mappers map[work.TypeName]work.DataMapper
		Logger  *zap.Logger
		Scope   tally.Scope
	}
	unitDeps := func() result {
		mappers := make(map[work.TypeName]work.DataMapper)

		dm := dataMapper{}
		t := work.TypeNameOf(entity{})
		mappers[t] = &dm
		l := zap.NewExample()
		s := tally.NewTestScope("test", map[string]string{})
		return result{
			Mappers: mappers,
			Logger:  l,
			Scope:   s,
		}
	}
	var uniter *work.Uniter

	//action.
	fxtest.New(
		t,
		fx.Provide(unitDeps),
		Modules.BestEffortUnit,
		fx.Invoke(func(p bestEffortUniterParameters) {
			uniter = &p.Uniter
		}),
	).RequireStart().RequireStop()

	//assert.
	require.NotNil(t, uniter)
}

func TestSQLUnitAndBestEffortUnit(t *testing.T) {

	//arrange.
	type sResult struct {
		fx.Out

		Mappers map[work.TypeName]work.SQLDataMapper
		DB      *sql.DB `name:"rwDB"`
		Logger  *zap.Logger
		Scope   tally.Scope
	}

	type bResult struct {
		fx.Out

		Mappers map[work.TypeName]work.DataMapper
	}

	unitDeps := func() (sResult, bResult) {
		sqlMappers := make(map[work.TypeName]work.SQLDataMapper)
		mappers := make(map[work.TypeName]work.DataMapper)

		sdm := sqlDataMapper{}
		dm := dataMapper{}
		t := work.TypeNameOf(entity{})
		sqlMappers[t] = &sdm
		mappers[t] = &dm
		var db *sql.DB
		l := zap.NewExample()
		s := tally.NewTestScope("test", map[string]string{})
		return sResult{Mappers: sqlMappers, DB: db, Logger: l, Scope: s},
			bResult{Mappers: mappers}
	}
	var sql *work.Uniter
	var bestEffort *work.Uniter

	//action.
	fxtest.New(
		t,
		fx.Provide(unitDeps),
		Modules.BestEffortUnit,
		Modules.SQLUnit,
		fx.Invoke(func(sp sqlUniterParameters, bep bestEffortUniterParameters) {
			sql = &sp.Uniter
			bestEffort = &bep.Uniter
		}),
	).RequireStart().RequireStop()

	//assert.
	require.NotNil(t, sql)
	require.NotNil(t, bestEffort)
}
