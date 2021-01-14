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

type UniterParameters struct {
	fx.In

	Uniter unit.Uniter `name:"bestEffortWorkUniter"`
}

func TestModule(t *testing.T) {
	//arrange.
	type result struct {
		fx.Out

		Mappers map[unit.TypeName]unit.DataMapper
		Logger  *zap.Logger
		Scope   tally.Scope
	}
	unitDeps := func() result {
		mappers := make(map[unit.TypeName]unit.DataMapper)

		dm := dataMapper{}
		t := unit.TypeNameOf(entity{})
		mappers[t] = &dm
		l := zap.NewExample()
		s := tally.NewTestScope("test", map[string]string{})
		return result{
			Mappers: mappers,
			Logger:  l,
			Scope:   s,
		}
	}
	var uniter *unit.Uniter

	//action.
	fxtest.New(
		t,
		fx.Provide(unitDeps),
		workfx.Module,
		fx.Invoke(func(p UniterParameters) {
			uniter = &p.Uniter
		}),
	).RequireStart().RequireStop()

	//assert.
	require.NotNil(t, uniter)
}
