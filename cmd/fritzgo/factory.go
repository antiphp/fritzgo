package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/antiphp/fritzgo/internal/tablerend"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/maps"
)

var styles = map[string]table.Style{
	"default": table.StyleDefault,
	"bold":    table.StyleBold,
	"fancy":   table.StyleColoredBlackOnCyanWhite,
	"none":    table.StyleLight,
}

func newRenderer(c *cli.Context) (*tablerend.TableRend, error) {
	style, ok := styles[c.String(flagRendererTableStyle)]
	if !ok {
		keys := maps.Keys(styles)
		sort.Strings(keys)

		return nil, fmt.Errorf("unsupported table style '%s', supported ones are: '%s'", c.String(flagRendererTableStyle), strings.Join(keys, "', '")) //nolint:goerr113
	}

	return tablerend.New(os.Stdout, tablerend.WithStyle(style)), nil
}
