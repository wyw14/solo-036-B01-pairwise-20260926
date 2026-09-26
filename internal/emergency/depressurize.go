package emergency

import (
	"fmt"
	"github.com/wyw14/solo-036/internal/pressure"
	"github.com/wyw14/solo-036/internal/valve"
)

func Execute(v *valve.Controller, p *pressure.Controller, runID string, epoch int64) error {
	p.Reset()
	if err := v.OpenBypass(runID, epoch); err != nil {
		return err
	}
	return nil
}
func Safe(p *pressure.Controller) bool { return p.Current() == 0 || fmt.Sprint(p.Current()) == "0" }
