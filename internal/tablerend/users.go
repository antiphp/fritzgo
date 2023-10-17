// Package tablerend renders fritz entities as tables.
package tablerend

import (
	"github.com/antiphp/fritzgo/pkg/fritztypes"
	"github.com/jedib0t/go-pretty/v6/table"
)

// ListUsers renders a user list.
func (r *TableRend) ListUsers(users []fritztypes.User) error {
	tw := r.newTable()
	tw.SetTitle("List users")
	tw.AppendHeader(table.Row{"User name", "Default"})

	for _, user := range users {
		tw.AppendRow(table.Row{user.Name, user.Default})
	}

	return r.render(tw)
}
