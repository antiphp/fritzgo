package main

import (
	"fmt"

	"github.com/antiphp/fritzgo"
	"github.com/antiphp/fritzgo/cmd"
	"github.com/antiphp/fritzgo/pkg/fritzclient"
	lctx "github.com/hamba/logger/v2/ctx"
	"github.com/urfave/cli/v2"
)

func run(runner func(*cli.Context, *fritzgo.FritzGo) error) func(*cli.Context) error {
	return func(c *cli.Context) error {
		log, logClose, err := cmd.NewLogger(c, "fritzgo")
		if err != nil {
			return fmt.Errorf("creating logger: %w", err)
		}
		defer logClose()

		client, err := fritzclient.New(c.String(flagFritzURL), c.String(flagFritzUser), c.String(flagFritzPass))
		if err != nil {
			log.Error("Could not create client", lctx.Err(err))
			return errAlreadyLogged
		}

		renderer, err := newRenderer(c)
		if err != nil {
			log.Error("Could not create renderer", lctx.Err(err))
			return errAlreadyLogged
		}

		return runner(c, fritzgo.New(client, renderer, log))
	}
}

func info(c *cli.Context, fritz *fritzgo.FritzGo) error {
	if err := fritz.Info(c.Context); err != nil {
		return fmt.Errorf("info: %w", err)
	}
	return nil
}

func usersList(c *cli.Context, fritz *fritzgo.FritzGo) error {
	if err := fritz.ListUsers(c.Context); err != nil {
		return fmt.Errorf("list users: %w", err)
	}
	return nil
}
