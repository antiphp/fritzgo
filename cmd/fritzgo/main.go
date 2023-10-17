// Package main runs the agent.
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/ettle/strcase"
	"github.com/hamba/cmd/v2"
	"github.com/urfave/cli/v2"
)

const (
	flagFritzURL           = "fritz.url"
	flagFritzUser          = "fritz.user"
	flagFritzPass          = "fritz.pass"
	flagRendererTableStyle = "render.table.style"

	catFritz = "FRITZ!Box:"

	catStyle = "Rendering style:"
)

var (
	buildVersion   = "<unknown>"
	buildTimestamp = "0"
	buildTime      = time.Unix(func() int64 { n, _ := strconv.Atoi(buildTimestamp); return int64(n) }(), 0)
)

var flags = cmd.Flags{
	&cli.StringFlag{
		Name:     flagFritzURL,
		Usage:    "The FRITZ!Box address.",
		Category: catFritz,
		Value:    "http://fritz.box",
		EnvVars:  []string{strcase.ToSNAKE(flagFritzURL)},
	},
	&cli.StringFlag{
		Name:     flagRendererTableStyle,
		Usage:    "The rendering table style.",
		Category: catStyle,
		Value:    "default",
		EnvVars:  []string{strcase.ToSNAKE(flagRendererTableStyle)},
	},
}.Merge(cmd.LogFlags)

func main() {
	os.Exit(mainWithExitCode())
}

// errAlreadyLogged represents an error in the run function and allows main to determine whether the error has already been
// logged (and its content can be ignored) or not. This is due to the fact that the logger is created inside run.
var errAlreadyLogged = errors.New("run error")

func mainWithExitCode() int {
	app := cli.NewApp()
	app.Name = "FritzGo"
	app.Version = buildVersion + " @ " + buildTime.Format(time.RFC3339)
	app.Usage = "CLI tool to access FRITZ!Box data"
	app.Flags = flags
	app.Commands = []*cli.Command{
		{
			Name:   "info",
			Action: run(info),
		},
		{
			Name: "users",
			Subcommands: []*cli.Command{
				{
					Name:   "list",
					Action: run(usersList),
				},
			},
		},
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	if err := app.RunContext(ctx, os.Args); err != nil {
		switch { //nolint:gocritic // Be flexible.
		case !errors.Is(err, errAlreadyLogged):
			_, _ = fmt.Fprintf(os.Stderr, "Error: %v", err)
		}
		return 1
	}
	return 0
}
