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
	"database/sql"
	"time"

	"github.com/freerware/work/v4/unit"
	"github.com/uber-go/tally"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Option(
	fx.Provide(func(p Parameters) Result {
		return Result{Uniter: unit.NewUniter(p.AsOptions()...)}
	}),
)

// UnitParameters encapsulates the various dependencies
// required to contruct work units.
type Parameters struct {
	fx.In

	Mappers                      map[unit.TypeName]unit.DataMapper
	Logger                       *zap.Logger                       `optional:"true"`
	Scope                        tally.Scope                       `optional:"true"`
	DB                           *sql.DB                           `name:"rwDB",optional:"true"`
	Actions                      map[unit.ActionType][]unit.Action `optional:"true"`
	DisableDefaultLoggingActions bool                              `name:"disableDefaultUnitLoggingActions",optional:"true"`
	RetryAttempts                int                               `name:"unitRetryAttempts",optional:"true"`
	RetryDelay                   time.Duration                     `name:"unitRetryDelay",optional:"true"`
	RetryMaximumJitter           time.Duration                     `name:"unitRetryMaximumJitter",optional:"true"`
	RetryType                    unit.RetryType                    `optional:"true"`
}

func (p Parameters) appendOption(options []unit.Option, param interface{}, option unit.Option) []unit.Option {
	if param == nil {
		return options
	}
	append(options, option)
}

func (p Parameters) AsOptions() (opts []unit.Option) {
	opts = p.appendOption(opts, p.Mappers, unit.DataMappers(p.Mappers))
	opts = p.appendOption(opts, p.Logger, unit.Logger(p.Logger))
	opts = p.appendOption(opts, p.Scope, unit.Scope(p.Scope))
	opts = p.appendOption(opts, p.DB, unit.DB(p.DB))
	opts = p.appendOption(opts, p.Actions, unit.Actions(p.Actions))
	opts = p.appendOption(opts, p.DisableDefaultLoggingActions, unit.DisableDefaultLoggingActions(p.DisableDefaultLoggingActions))
	opts = p.appendOption(opts, p.RetryAttempts, unit.RetryAttempts(p.RetryAttempts))
	opts = p.appendOption(opts, p.RetryDelay, unit.RetryDelay(p.RetryDelay))
	opts = p.appendOption(opts, p.RetryMaximumJitter, unit.RetryMaximumJitter(p.RetryMaximumJitter))
	opts = p.appendOption(opts, p.RetryType, unit.RetryType(p.RetryType))
	return
}

// BestEffortWorkUniterResult defines the best effort uniter
// to be provided to the dependency injection container.
type Result struct {
	fx.Out

	Uniter unit.Uniter `group:"uniter"`
}
