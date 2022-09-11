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

package integration

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/funlessdev/fl-cli/internal/command/admin"
	"github.com/funlessdev/fl-cli/pkg/deploy"
	"github.com/funlessdev/fl-cli/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestAdminDevRun(t *testing.T) {
	runIntegration := os.Getenv("INTEGRATION_TESTS")
	if runIntegration == "" {
		t.Skip("set INTEGRATION_TESTS (optionally with DOCKER_HOST) to run this test")
	}

	admCmd := admin.Admin{Dev: struct{}{}}

	coreName := "fl-core-test"
	workerName := "fl-worker-test"
	flNetName := "fl-net-test"
	flRuntimeName := "fl-runtime-net-test"
	localDeployer := deploy.NewLocalDeployer(coreName, workerName, flNetName, flRuntimeName)

	b := log.NewLoggerBuilder()
	var outbuf bytes.Buffer
	logger, _ := b.WithDebug(true).WithWriter(&outbuf).Build()

	ctx := context.Background()

	t.Run("should successfully deploy funless when no errors occurr", func(t *testing.T) {
		err := admCmd.Dev.Run(ctx, localDeployer, logger)

		assert.NoError(t, err)

		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.41"))
		assert.NoError(t, err)

		assertContainer(t, ctx, cli, coreName)
		assertContainer(t, ctx, cli, workerName)

		assertNetwork(t, ctx, cli, flNetName)
		assertNetwork(t, ctx, cli, flRuntimeName)

		_ = localDeployer.RemoveCoreContainer(ctx)
		_ = localDeployer.RemoveWorkerContainer(ctx)
		_ = localDeployer.RemoveFLNetworks(ctx)
	})

	t.Run("should successfully deploy without creating networks if they already exist", func(t *testing.T) {
		_ = localDeployer.SetupFLNetworks(ctx)

		err := admCmd.Dev.Run(ctx, localDeployer, logger)

		assert.NoError(t, err)

		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.41"))
		assert.NoError(t, err)

		assertContainer(t, ctx, cli, coreName)
		assertContainer(t, ctx, cli, workerName)

		assertNetwork(t, ctx, cli, flNetName)
		assertNetwork(t, ctx, cli, flRuntimeName)

		_ = localDeployer.RemoveCoreContainer(ctx)
		_ = localDeployer.RemoveWorkerContainer(ctx)
		_ = localDeployer.RemoveFLNetworks(ctx)
	})

	t.Run("should fail if core is already running", func(t *testing.T) {
		_ = localDeployer.SetupFLNetworks(ctx)
		_ = localDeployer.PullCoreImage(ctx)
		_ = localDeployer.StartCore(ctx)

		err := admCmd.Dev.Run(ctx, localDeployer, logger)

		assert.Error(t, err)

		_ = localDeployer.RemoveCoreContainer(ctx)
		_ = localDeployer.RemoveFLNetworks(ctx)
	})
}

func assertContainer(t *testing.T, ctx context.Context, cli *client.Client, name string) {
	t.Helper()
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{Key: "name", Value: name}),
	})

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(containers), 1)
}

func assertNetwork(t *testing.T, ctx context.Context, cli *client.Client, name string) {
	t.Helper()
	networks, err := cli.NetworkList(ctx, types.NetworkListOptions{
		Filters: filters.NewArgs(filters.KeyValuePair{Key: "name", Value: name}),
	})

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(networks), 1)
}