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

package service

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSignerMasking(t *testing.T) {
	cfg := &Signer{KeyFile: "/foo.pem", KeyPassword: "secret"}
	bs, err := json.Marshal(cfg)
	require.NoError(t, err)
	require.JSONEq(t, string(bs), `{"KeyFile":"/foo.pem","KeyPassword":"s****t"}`)
}
