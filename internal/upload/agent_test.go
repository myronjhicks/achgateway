// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package upload

import (
	"testing"
	"time"

	"github.com/moov-io/achgateway/internal/service"
	"github.com/moov-io/base/log"

	"github.com/stretchr/testify/require"
)

func TestAgent(t *testing.T) {
	cfg := service.UploadAgents{
		Agents: []service.UploadAgent{
			{
				ID:   "mock",
				Mock: &service.MockAgent{},
			},
		},
	}
	agent, err := New(log.NewTestLogger(), cfg, "mock")
	require.NoError(t, err)

	if aa, ok := agent.(*MockAgent); !ok {
		t.Errorf("unexpected agent: %#v", aa)
	}

	// setup a second (retrying) agent
	cfg.Retry = &service.UploadRetry{
		Interval:   1 * time.Second,
		MaxRetries: 3,
	}
	agent, err = New(log.NewTestLogger(), cfg, "mock")
	require.NoError(t, err)

	if aa, ok := agent.(*RetryAgent); !ok {
		t.Errorf("unexpected agent: %#v", agent)
	} else {
		if aa, ok := aa.underlying.(*MockAgent); !ok {
			t.Errorf("unexpected agent: %#v", aa)
		}
	}
}
