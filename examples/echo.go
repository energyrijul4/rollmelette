// Copyright (c) Gabriel de Quadros Ligneul
// SPDX-License-Identifier: Apache-2.0 (see LICENSE)

package main

import (
	"log/slog"

	"github.com/gligneul/rollmelette"
)

// EchoApplication is an application that emits a voucher, a notice, and a report for each advance
// input; and a report for each inspect input.
type EchoApplication struct{}

func (e *EchoApplication) Advance(
	env rollmelette.Env,
	metadata rollmelette.Metadata,
	payload []byte,
) error {
	slog.Info("received advance input")
	env.Voucher(metadata.MsgSender, payload)
	env.Notice(payload)
	env.Report(payload)
	return nil
}

func (e *EchoApplication) Inspect(env rollmelette.EnvInspector, payload []byte) error {
	slog.Info("received inspect input")
	env.Report(payload)
	return nil
}
