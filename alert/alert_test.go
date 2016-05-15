package alert_test

import (
	"github.com/mkch/rw"
	"github.com/mkch/rw/alert"
	"testing"
)

func TestAlert(t *testing.T) {
	rw.Run(func() {
		alert.New().
			SetStyle(alert.Informational).
			SetTitle("Alert smaple").
			SetMessage("Positive: OK(Def)\nNegative: Cancel\nNeutral: Empty").
			SetInformativeMessage("Style: Informational").
			SetPositiveButton("OK", func() { t.Logf("OK selected\n") }, true).
			SetNegativeButton("Cancel", func() { t.Logf("Cancel selected\n") }, false).
			SetNeutralButton("Empty", nil, false).
			Show(nil)

		alert.New().
			SetStyle(alert.Warning).
			SetTitle("Alert smaple").
			SetMessage("Positive: OK\nNegative: Cancel(Def)\nNeutral: Empty").
			SetInformativeMessage("Style: Warning").
			SetPositiveButton("OK", func() { t.Logf("OK selected\n") }, false).
			SetNegativeButton("Cancel", func() { t.Logf("Cancel selected\n") }, true).
			SetNeutralButton("Empty", nil, false).
			Show(nil)

		alert.New().
			SetStyle(alert.Critical).
			SetTitle("Alert smaple").
			SetMessage("Positive: OK\nNegative: Cancel\nNeutral: Empty(Def)").
			SetInformativeMessage("Style: Critical").
			SetPositiveButton("OK", func() { t.Logf("OK selected\n") }, false).
			SetNegativeButton("Cancel", func() { t.Logf("Cancel selected\n") }, false).
			SetNeutralButton("Empty", nil, true).
			Show(nil)

		rw.Exit()
	})
}
