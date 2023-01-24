// Copyright 2022 Giuseppe De Palma, Matteo Trentin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mod

import (
	"context"

	"github.com/funlessdev/fl-cli/pkg/client"
	"github.com/funlessdev/fl-cli/pkg/log"
)

type List struct {
	Count bool `name:"count" short:"c" default:"false" help:"return number of results"`
}

func (l *List) Run(ctx context.Context, modHandler client.ModHandler, logger log.FLogger) error {
	res, err := modHandler.List(ctx)
	if err != nil {
		return extractError(err)
	}

	data := res.GetData()

	for _, v := range data {
		logger.Info(*v.Name)
	}

	if l.Count {
		logger.Infof("Count: %d\n", len(data))
	}
	return nil
}