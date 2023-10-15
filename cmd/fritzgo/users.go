package main

import (
	"context"
	"fmt"

	"github.com/antiphp/fritzgo"
	"github.com/antiphp/fritzgo/cmd"
	lctx "github.com/hamba/logger/v2/ctx"
	"github.com/urfave/cli/v2"
)

func runUsersList(c *cli.Context) error {
	ctx, cancel := context.WithCancel(c.Context)
	defer cancel()

	log, logClose, err := cmd.NewLogger(c, "users.list")
	if err != nil {
		return fmt.Errorf("creating logger: %w", err)
	}
	defer logClose()

	fritz, err := fritzgo.New(c.String(flagFritzURL), c.String(flagFritzUser), c.String(flagFritzPass), log)
	if err != nil {
		log.Error("init fritz", lctx.Err(err))
		return errAlreadyLogged
	}

	if err := fritz.ListUsers(ctx); err != nil {
		log.Error("list users", lctx.Err(err))
		return errAlreadyLogged
	}

	return nil
}
