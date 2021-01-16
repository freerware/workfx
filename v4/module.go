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

package workfx

import (
	"github.com/freerware/work/v4/unit"
	"go.uber.org/fx"
)

// Module is the Fx module necessary to enable an Fx application with work units.
var Module = fx.Option(
	fx.Provide(func(p Parameters) Result {
		return Result{Uniter: unit.NewUniter(p.Options...)}
	}),
)

// UnitParameters encapsulates the various dependencies required to contruct work units.
type Parameters struct {
	fx.In

	Options []unit.Option `group:"unitOptions"`
}

// Result defines the uniter to be provided to the Fx application.
type Result struct {
	fx.Out

	Uniter unit.Uniter `name:"uniter"`
}
