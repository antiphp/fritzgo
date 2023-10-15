// Package render renders fritz entities.
package render

import (
	"io"

	"github.com/antiphp/fritzgo/pkg/fritztypes"
	"github.com/jedib0t/go-pretty/v6/table"
)

func UsersList(w io.Writer, users []fritztypes.User) error {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"User name", "Is default?"})

	for _, user := range users {
		tw.AppendRow(table.Row{Blue.Render(user.Name), user.Default})
	}

	_, err := w.Write([]byte(tw.Render()))
	return err
}
