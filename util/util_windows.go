package util

import(
    "strings"
    "io"
	"fmt"
	"github.com/kevin-yuan/rw/native"
)

// Windows_ControlTitleWithMnemonic applys the mnemonic char to the title.
// Only available in Windows platform.
func Windows_ControlTitleWithMnemonic(title string, mnemonic rune) string {
	if mnemonic == '&' {
		panic ("& is not a valid mnemonic")
	}
	if mnemonic != 0 {
		if index := strings.IndexRune(strings.Replace(title, "&", "&&", -1), mnemonic); index != -1 {
			title = title[:index] + "&" + title[index:]
		} else {
			title += ("(&" + string(mnemonic) + ")")
		}
	}
	return title
}

func printHandle(w io.Writer, handle native.Handle) {
	fmt.Fprintf(w, "%#X", handle)
}