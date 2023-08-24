// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package admintest

import (
	"testing"

	"github.com/moov-io/base/admin"

	"github.com/stretchr/testify/require"
)

func Server(t *testing.T) *admin.Server {
	t.Helper()

	adminServer, err := admin.New(admin.Opts{
		Addr: ":0",
	})
	require.NoError(t, err)
	require.NotNil(t, adminServer)

	go adminServer.Listen()

	t.Cleanup(func() {
		adminServer.Shutdown()
	})

	return adminServer
}
