/* Copyright 2019 Freerware
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package workfx_test

import (
	"context"
	"testing"

	"github.com/freerware/work/v4/unit"
	"github.com/freerware/workfx/v4"
	"github.com/stretchr/testify/require"
	"github.com/uber-go/tally"
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

		Options []unit.Option `group:"unitOptions"`
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
