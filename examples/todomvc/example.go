package main

import (
	"encoding/json"
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/denkhaus/vecty"
	"github.com/denkhaus/vecty/examples/todomvc/actions"
	"github.com/denkhaus/vecty/examples/todomvc/components"
	"github.com/denkhaus/vecty/examples/todomvc/dispatcher"
	"github.com/denkhaus/vecty/examples/todomvc/store"
	"github.com/denkhaus/vecty/examples/todomvc/store/model"
)

func main() {
	attachLocalStorage()

	vecty.SetTitle("GopherJS • TodoMVC")
	vecty.AddStylesheet("node_modules/todomvc-common/base.css")
	vecty.AddStylesheet("node_modules/todomvc-app-css/index.css")
	p := &components.PageView{}
	store.Listeners.Add(p, func() {
		p.Items = store.Items
		vecty.Rerender(p)
	})
	vecty.RenderBody(p)
}

func attachLocalStorage() {
	store.Listeners.Add(nil, func() {
		data, err := json.Marshal(store.Items)
		if err != nil {
			fmt.Printf("failed to store items: %s\n", err)
		}
		js.Global.Get("localStorage").Set("items", string(data))
	})

	if data := js.Global.Get("localStorage").Get("items"); data != js.Undefined {
		var items []*model.Item
		if err := json.Unmarshal([]byte(data.String()), &items); err != nil {
			fmt.Printf("failed to load items: %s\n", err)
		}
		dispatcher.Dispatch(&actions.ReplaceItems{
			Items: items,
		})
	}
}
