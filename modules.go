package workfx

import (
	"database/sql"

	"github.com/freerware/work"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Modules defines the various work unit Fx modules.
var Modules = struct {
	SQLUnit        fx.Option
	BestEffortUnit fx.Option
}{
	SQLUnit: fx.Options(
		fx.Provide(func(p SQLUnitParameters) SQLWorkUniterResult {
			params := work.SQLUnitParameters{
				UnitParameters: work.UnitParameters{
					Inserters: p.Inserters,
					Updaters:  p.Updaters,
					Deleters:  p.Deleters,
					Logger:    p.Logger,
				},
				ConnectionPool: p.DB,
			}
			return SQLWorkUniterResult{
				WorkUniter: work.NewSQLUniter(params),
			}
		})),

	BestEffortUnit: fx.Options(
		fx.Provide(func(p UnitParameters) BestEffortWorkUniterResult {
			params := work.UnitParameters{
				Inserters: p.Inserters,
				Updaters:  p.Updaters,
				Deleters:  p.Deleters,
				Logger:    p.Logger,
			}
			return BestEffortWorkUniterResult{
				WorkUniter: work.NewBestEffortUniter(params),
			}
		})),
}

// SQLUnitParameters encapsulates the various dependencies
// required to construct SQL work units.
type SQLUnitParameters struct {
	fx.In

	Inserters map[work.TypeName]work.Inserter
	Updaters  map[work.TypeName]work.Updater
	Deleters  map[work.TypeName]work.Deleter
	Logger    *zap.Logger
	DB        *sql.DB `name:"rwDB"`
}

// UnitParameters encapsulates the various dependencies
// required to contruct various work units, such as the
// best effort work unit.
type UnitParameters struct {
	fx.In

	Inserters map[work.TypeName]work.Inserter
	Updaters  map[work.TypeName]work.Updater
	Deleters  map[work.TypeName]work.Deleter
	Logger    *zap.Logger
}

// SQLWorkUniterResult defines the SQL work uniter to be
// provided to the dependency injection container.
type SQLWorkUniterResult struct {
	fx.Out

	WorkUniter work.Uniter `name:"sqlWorkUniter",group:"workUniter"`
}

// BestEffortWorkUniterResult defines the best effort uniter
// to be provided to the dependency injection container.
type BestEffortWorkUniterResult struct {
	fx.Out

	WorkUniter work.Uniter `name:"bestEffortWorkUniter",group:"workUniter"`
}
