// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package log

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {

	t.Run("NewLoggerBuilder returns new builder", func(t *testing.T) {
		builder := NewLoggerBuilder()
		assert.NotNil(t, builder)
	})

	t.Run("NewLoggerBuilder returns builder with default values", func(t *testing.T) {
		builder := NewLoggerBuilder()
		assert.False(t, builder.(*loggerBuilder).debug)
		assert.True(t, builder.(*loggerBuilder).spinCfg.SuffixAutoColon)
		assert.Equal(t, "✓", builder.(*loggerBuilder).spinCfg.StopCharacter)
		assert.Equal(t, []string{"fgGreen"}, builder.(*loggerBuilder).spinCfg.StopColors)
		assert.Equal(t, "✗", builder.(*loggerBuilder).spinCfg.StopFailCharacter)
		assert.Equal(t, []string{"fgRed"}, builder.(*loggerBuilder).spinCfg.StopFailColors)
	})

	t.Run("Build returns new logger", func(t *testing.T) {
		logger, err := NewLoggerBuilder().Build()
		assert.NoError(t, err)
		assert.NotNil(t, logger)
		assert.False(t, logger.(*FLoggerImpl).debug)
	})

	t.Run("WithDebug true activates the debug in logger", func(t *testing.T) {
		logger, err := NewLoggerBuilder().WithDebug(true).Build()
		assert.NoError(t, err)
		assert.NotNil(t, logger)
		assert.True(t, logger.(*FLoggerImpl).debug)
	})

	t.Run("SpinnerFrequency generates error if input is less or equal to 0", func(t *testing.T) {
		logger, err := NewLoggerBuilder().SpinnerFrequency(0 * time.Millisecond).Build()
		assert.Error(t, err)
		assert.Nil(t, logger)
	})

	t.Run("SpinnerFrequency appends err if err is found", func(t *testing.T) {
		builder := NewLoggerBuilder()
		builder.(*loggerBuilder).err = errors.New("test err")
		_, err := builder.SpinnerFrequency(0 * time.Millisecond).Build()
		assert.Error(t, err)
		assert.Equal(t, "spinner frequency must be greater than 0, test err", err.Error())
	})

	t.Run("SpinnerCharSet fails if input is out of range [0,90]", func(t *testing.T) {
		logger, err := NewLoggerBuilder().SpinnerCharSet(91).Build()
		assert.Error(t, err)
		assert.Nil(t, logger)
	})
}