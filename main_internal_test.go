// Copyright (c) 2020 Hervé Gouchet. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package main

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/matryer/is"
	errup "github.com/rvflash/goup/internal/errors"
	"github.com/rvflash/goup/pkg/goup"
)

func TestPatterns(t *testing.T) {
	var (
		are = is.New(t)
		dt  = map[string]struct {
			in  []string
			out string
		}{
			"Default": {},
			"Ok":      {in: []string{"a", "b", "", "d", " e "}, out: "a,b,d,e"},
		}
	)
	for name, tt := range dt {
		tt := tt
		t.Run(name, func(t *testing.T) {
			are.Equal(tt.out, patterns(tt.in...)) // mismatch result
		})
	}
}

func TestRun(t *testing.T) {
	var (
		stderr = new(strings.Builder)
		stdout = new(strings.Builder)
		are    = is.New(t)
		dt     = map[string]struct {
			ctx    context.Context
			cnf    goup.Config
			args   []string
			stderr io.Writer
			stdout io.Writer
			err    error
		}{
			"default":      {err: errup.ErrMissing},
			"no context":   {stdout: stdout, stderr: stderr, err: errup.ErrMod},
			"context only": {ctx: context.Background(), stdout: stdout, stderr: stderr, err: errup.ErrMod},
			"ok": {
				ctx:    context.Background(),
				cnf:    goup.Config{PrintVersion: true},
				stdout: stdout,
				stderr: stderr,
			},
		}
	)
	for name, tt := range dt {
		tt := tt
		t.Run(name, func(t *testing.T) {
			err := run(tt.ctx, tt.cnf, tt.args, tt.stdout, tt.stderr)
			are.True(errors.Is(err, tt.err)) // mismatch error
		})
	}
}
