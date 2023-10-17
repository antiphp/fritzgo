// Package tablerend renders tables.
package tablerend

import (
	"bytes"
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
)

// TableRend renders tables.
type TableRend struct {
	w io.Writer

	style table.Style
}

// OptFunc represents an optional configuration.
type OptFunc func(tr *TableRend)

// WithStyle applies optional table style.
func WithStyle(style table.Style) OptFunc {
	return func(tr *TableRend) {
		tr.style = style
	}
}

// New returns a new table renderer.
func New(w io.Writer, opts ...OptFunc) *TableRend {
	tr := TableRend{
		w:     w,
		style: table.StyleDefault,
	}

	for _, opt := range opts {
		opt(&tr)
	}

	return &tr
}

func (r *TableRend) newTable() *table.Table {
	tw := &table.Table{}
	tw.SetStyle(r.style)
	return tw
}

func (r *TableRend) render(tw *table.Table) error {
	b := &bytes.Buffer{}
	b.WriteString("\n")
	b.WriteString(tw.Render())
	b.WriteString("\n")

	_, err := r.w.Write(b.Bytes())
	return err
}
