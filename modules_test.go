package workfx

import (
	"database/sql"
	"testing"

	"github.com/freerware/work"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

type dataMapper struct{}

func (dm dataMapper) Insert(e ...interface{}) error { return nil }
func (dm dataMapper) Update(e ...interface{}) error { return nil }
func (dm dataMapper) Delete(e ...interface{}) error { return nil }

type entity struct{}

type result struct {
	fx.Out

	Inserters map[work.TypeName]work.Inserter
	Updaters  map[work.TypeName]work.Updater
	Deleters  map[work.TypeName]work.Deleter
	DB        *sql.DB `name:"rwDB"`
}

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
	unitDeps := func() result {
		inserters := make(map[work.TypeName]work.Inserter)
		updaters := make(map[work.TypeName]work.Updater)
		deleters := make(map[work.TypeName]work.Deleter)

		dm := dataMapper{}
		t := work.TypeNameOf(entity{})
		inserters[t], updaters[t], deleters[t] = &dm, &dm, &dm
		var db *sql.DB
		return result{
			Inserters: inserters,
			Updaters:  updaters,
			Deleters:  deleters,
			DB:        db,
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
	unitDeps := func() result {
		inserters := make(map[work.TypeName]work.Inserter)
		updaters := make(map[work.TypeName]work.Updater)
		deleters := make(map[work.TypeName]work.Deleter)

		dm := dataMapper{}
		t := work.TypeNameOf(entity{})
		inserters[t], updaters[t], deleters[t] = &dm, &dm, &dm
		var db *sql.DB
		return result{
			Inserters: inserters,
			Updaters:  updaters,
			Deleters:  deleters,
			DB:        db,
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
	unitDeps := func() result {
		inserters := make(map[work.TypeName]work.Inserter)
		updaters := make(map[work.TypeName]work.Updater)
		deleters := make(map[work.TypeName]work.Deleter)

		dm := dataMapper{}
		t := work.TypeNameOf(entity{})
		inserters[t], updaters[t], deleters[t] = &dm, &dm, &dm
		var db *sql.DB
		return result{
			Inserters: inserters,
			Updaters:  updaters,
			Deleters:  deleters,
			DB:        db,
		}
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
