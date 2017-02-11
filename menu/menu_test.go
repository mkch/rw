package menu_test

import (
	"fmt"
	"github.com/mkch/rw"
	"github.com/mkch/rw/event"
	"github.com/mkch/rw/menu"
	"github.com/mkch/rw/window"
	"runtime"
	"sort"
	"strconv"
	"testing"
)

func TestItemOperations(t *testing.T) {
	rw.Debug = true
	rw.Run(func() { itemOpTestStartup(t) })
}

type byItemTitle struct {
	rw.Menu
}

func (b byItemTitle) Len() int {
	return b.Menu.ItemCount()
}

func (b byItemTitle) Swap(i, j int) {
	var lower, upper int
	if i > j {
		lower, upper = j, i
	} else if i < j {
		lower, upper = i, j
	} else {
		return
	}
	m := b.Menu
	itemLower := m.Item(lower)
	itemUpper := m.Item(upper)
	m.RemoveItem(itemLower)
	m.RemoveItem(itemUpper)
	// Insert the (swapped) lower item first to maintain the right insertion positions.
	// e.g. For i == 2, j == 5,
	//   insert(item5, 2) // item5 at index 2
	//   insert(item2, 5) // item5 at index 2(unaffected), and item2 at index 5
	// but,
	//   insert(item2, 5) // item2 at index 5
	//   insert(item5, 2) // item5 at index 2, item2 at index 6(affected!)
	m.InsertItem(itemUpper, lower)
	m.InsertItem(itemLower, upper)
}

func (b byItemTitle) Less(i, j int) bool {
	m := b.Menu
	vi, _ := strconv.Atoi(m.Item(i).Title())
	vj, _ := strconv.Atoi(m.Item(j).Title())
	return vi < vj
}

func itemOpTestStartup(t *testing.T) {
	w1 := window.New()
	w1.SetFrame(rw.Rect{100, 200, 200, 300})
	w1.ShowActive()
	w1.OnClose().SetListener(func() { rw.Exit() })
	w1.SetTitle("W1")
	/*menu1 := */ itemOpTest(t, w1)
	//rw.SetMainMenu(menu1)

	w2 := window.New()
	w2.ShowActive()
	w2.SetTitle("W2")
	if runtime.GOOS != "darwin" {
		itemOpTest(t, w2)
	}

}

func itemOpTest(t *testing.T, win rw.Window) rw.Menu {
	m1 := menu.NewBuilder().
		BeginItem("3").
		SetKeyboardShortcut(rw.DefaultModifierKey, 'A').
		SetOnClickListener(
			func(event event.Event) bool {
				item := event.Sender().(rw.MenuItem)
				fmt.Printf("Sender=%v\n", event.Sender())
				item.SetChecked(!item.Checked())
				return true
			}).
		End().
		BeginItem("6").End().
		Build()
	m2 := menu.NewBuilder().
		BeginItem("5").End().
		BeginItem("2").End().
		BeginItem("1").End().
		BeginItem("4").End().
		Build()

	var m rw.Menu
	op := menu.NewItemBuilder(nil, nil).
		SetTitle("Sort").
		SetKeyboardShortcut(rw.DefaultModifierKey, 'R').
		SetOnClickListener(
			func(event event.Event) bool {
				item := event.Sender().(rw.MenuItem)
				fmt.Printf("Sort on %v\n", item.Menu().Opener().Menu().Window())
				if item.Enabled() == false {
					t.Fatalf("Listener of disabled menu item is called")
				}
				for _, item := range m1.Items() {
					m2.AddItem(item)
				}
				sort.Sort(byItemTitle{m2})
				if m2.Item(0).Title() != "1" ||
					m2.Item(1).Title() != "2" ||
					m2.Item(2).Title() != "3" ||
					m2.Item(3).Title() != "4" ||
					m2.Item(4).Title() != "5" {
					t.Errorf("Sort menu item error")
				}

				m.Item(2).SetTitle("Sorted")
				m.Item(1).SetTitle("EMPTY")
				sender := event.Sender().(rw.MenuItem)
				fmt.Printf("sender=%v\n", sender)
				sender.SetEnabled(false)
				if sender.Enabled() != false {
					t.Errorf("Menu item is not disabled")
				}
				return true
			}).Build()
	x := menu.NewItemBuilder(nil, nil).
		SetTitle("Exit").
		SetMnemonic('x').
		SetOnClickListener(func(event event.Event) bool {
			// windows only
			//event.Sender().(rw.MenuItem).Menu().Opener().Menu().Window().Close()
			win.Close()
			return true
		}).Build()

	m =
		menu.NewBuilder().
			BeginItem("Menu").
			BeginSubmenu().
			AddItem(op).
			AddSeparator().
			AddItem(x).
			End().
			End().
			BeginItem("Unsorted1").
			SetSubmenu(m1).
			End().
			BeginItem("Unsorted2").
			SetSubmenu(m2).
			End().
			Build()

	setMenu(win, m)
	return m
}
