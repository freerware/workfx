package workfx_test

import (
	"context"
	"testing"

	"github.com/freerware/work/v4/unit"
	"github.com/stretchr/testify/require"
	"github.com/uber-go/tally"
	"github.com/workfx/v4"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
)

type dataMapper struct{}

func (dm dataMapper) Insert(ctx context.Context, mCtx unit.MapperContext, e ...interface{}) error {
	return nil
}
func (dm dataMapper) Update(ctx context.Context, mCtx unit.MapperContext, e ...interface{}) error {
	return nil
}
func (dm dataMapper) Delete(ctx context.Context, mCtx unit.MapperContext, e ...interface{}) error {
	return nil
}

type entity struct{}

type Parameters struct {
	fx.In

	Uniter unit.Uniter `name:"uniter"`
}

func TestModule(t *testing.T) {
	//arrange.
	type result struct {
		fx.Out

		Options []unit.Option `name:"unitOptions"`
	}

	unitDeps := func() result {
		mappers := make(map[unit.TypeName]unit.DataMapper)
		dm := dataMapper{}
		t := unit.TypeNameOf(entity{})
		mappers[t] = &dm
		l := zap.NewExample()
		s := tally.NewTestScope("test", map[string]string{})
		return result{
			Options: []unit.Option{
				unit.DataMappers(mappers),
				unit.Logger(l),
				unit.Scope(s),
			},
		}
	}
	var uniter *unit.Uniter

	//action.
	fxtest.New(
		t,
		fx.Provide(unitDeps),
		workfx.Module,
		fx.Invoke(func(p Parameters) {
			uniter = &p.Uniter
		}),
	).RequireStart().RequireStop()

	//assert.
	require.NotNil(t, uniter)
}
