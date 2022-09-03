// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
package admin

type Admin struct {
	Dev dev `cmd:"" help:"deploy locally with 1 core and 1 worker docker containers"`

	Init  coreInit `cmd:"" help:"todo init subcommand help"`
	Join  join     `cmd:"" help:"todo join subcommand help"`
	Token token    `cmd:"" help:"todo token subcommand help"`

	Reset reset `cmd:"" help:"removes the deployment of local containers"`
}
