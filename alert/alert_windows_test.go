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
			SetMessage("Positive: n/a\nNegative: n/a\nNeutral: OK").
			SetInformativeMessage("Style: Informational").
			SetNeutralButton(alert.OK, func() { t.Logf("OK selected\n") }, true).
			Show(nil)

		alert.New().
			SetStyle(alert.Informational).
			SetTitle("Alert smaple").
			SetMessage("Positive: OK(Def)\nNegative: Cancel\nNeutral: n/a").
			SetInformativeMessage("Style: Informational").
			SetPositiveButton(alert.OK, func() { t.Logf("OK selected\n") }, true).
			SetNegativeButton(alert.Cancel, func() { t.Logf("Cancel selected\n") }, false).
			Show(nil)

		alert.New().
			SetStyle(alert.Warning).
			SetTitle("Alert smaple").
			SetMessage("Positive: OK\nNegative: Cancel(Def)\nNeutral: n/a").
			SetInformativeMessage("Style: Warning").
			SetPositiveButton(alert.OK, func() { t.Logf("OK selected\n") }, false).
			SetNegativeButton(alert.Cancel, func() { t.Logf("Cancel selected\n") }, true).
			Show(nil)

		alert.New().
			SetStyle(alert.Critical).
			SetTitle("Alert smaple").
			SetMessage("Positive: Yes\nNegative: No\nNeutral: Cancel(Def)").
			SetInformativeMessage("Style: Critical").
			SetPositiveButton(alert.Yes, func() { t.Logf("OK selected\n") }, false).
			SetNegativeButton(alert.No, func() { t.Logf("Cancel selected\n") }, false).
			SetNeutralButton(alert.Cancel, nil, true).
			Show(nil)

		alert.New().
			SetTitle("Alert smaple").
			SetMessage("Positive: Abort\nNegative: Retry\nNeutral: Ignore").
			SetInformativeMessage("Style: Default").
			SetPositiveButton(alert.Abort, func() { t.Logf("Abort selected\n") }, false).
			SetNegativeButton(alert.Retry, func() { t.Logf("Retry selected\n") }, false).
			SetNeutralButton(alert.Ignore, nil, false).
			Show(nil)

		alert.New().
			SetTitle("Alert smaple").
			SetMessage("Positive: Cancel\nNegative: Retry\nNeutral: n/a").
			SetInformativeMessage("Style: Default").
			SetPositiveButton(alert.Cancel, func() { t.Logf("Cancel selected\n") }, false).
			SetNegativeButton(alert.Retry, func() { t.Logf("Retry selected\n") }, false).
			Show(nil)
		rw.Exit()
	})
}
