package window_test

import (
	"github.com/kevin-yuan/rw"
	"fmt"
)

// setMenu set the app main menu to m.
func setMenu(m rw.Menu, w rw.Window) {
	old := rw.OSX_MainManu()
	fmt.Printf("old=%#v\n", old)
	if old != nil {
		if old == m {
			return
		} else {
			old.Release()
		}
	}
	rw.OSX_SetMainMenu(m)
}
