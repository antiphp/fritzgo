// Package cmd contains reusable, command-related helper functions.
package cmd

import (
	"github.com/hamba/cmd/v2"
	"github.com/hamba/logger/v2"
	lctx "github.com/hamba/logger/v2/ctx"
	"github.com/urfave/cli/v2"
)

// NewLogger returns a preconfigured logger.
func NewLogger(c *cli.Context, svc string) (*logger.Logger, func(), error) {
	logger.TimeFormat = logger.TimeFormatISO8601

	log, err := cmd.NewLogger(c)
	if err != nil {
		return nil, nil, err
	}
	cancel := log.WithTimestamp()

	log = log.With(lctx.Str("svc", svc))

	return log, cancel, nil
}
