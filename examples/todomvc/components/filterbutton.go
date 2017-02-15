package components

import (
	"github.com/denkhaus/vecty"
	"github.com/denkhaus/vecty/elem"
	"github.com/denkhaus/vecty/event"
	"github.com/denkhaus/vecty/examples/todomvc/actions"
	"github.com/denkhaus/vecty/examples/todomvc/dispatcher"
	"github.com/denkhaus/vecty/examples/todomvc/store"
	"github.com/denkhaus/vecty/examples/todomvc/store/model"
	"github.com/denkhaus/vecty/prop"
)

type FilterButton struct {
	vecty.Core

	Label  string
	Filter model.FilterState
}

func (b *FilterButton) onClick(event *vecty.Event) {
	dispatcher.Dispatch(&actions.SetFilter{
		Filter: b.Filter,
	})
}

func (b *FilterButton) Render() *vecty.HTML {
	return elem.ListItem(
		elem.Anchor(
			vecty.If(store.Filter == b.Filter, prop.Class("selected")),
			prop.Href("#"),
			event.Click(b.onClick).PreventDefault(),

			vecty.Text(b.Label),
		),
	)
}
