package tablerend

import (
	"github.com/antiphp/fritzgo/pkg/fritztypes"
	"github.com/jedib0t/go-pretty/v6/table"
)

// Info renders a fritz info.
func (r *TableRend) Info(info fritztypes.Info) error {
	tw := r.newTable()
	tw.SetTitle("Info")

	tw.AppendRows([]table.Row{
		{"Name", info.Name},
		{"Version", info.Version},
		{"Mac address", info.Mac},
		{"URL", info.URL},
	})

	return r.render(tw)
}
