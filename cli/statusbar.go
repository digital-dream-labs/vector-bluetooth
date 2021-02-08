package cli

import (
	"github.com/schollz/progressbar/v3"
)

func (c *conf) renderLogStatus() {
	var bar *progressbar.ProgressBar
	var pc uint32
	for {
		r := <-c.status
		if bar == nil {
			bar = progressbar.Default(int64(r.LogStatus.PacketTotal), "log progress")
		}

		tc := r.LogStatus.PacketNumber - pc
		_ = bar.Add(int(tc))
		pc = r.LogStatus.PacketNumber
	}
}

func (c *conf) renderOTAStatus() {
	var bar *progressbar.ProgressBar
	var pc uint32
	for {
		r := <-c.status

		if r.OTAStatus.PacketTotal != 0 {
			if bar == nil {
				bar = progressbar.Default(int64(r.OTAStatus.PacketTotal), "OTA progress")
			}

			if r.OTAStatus.PacketNumber != pc {
				tc := r.OTAStatus.PacketNumber - pc
				_ = bar.Add(int(tc))
			}
			pc = r.OTAStatus.PacketNumber
		}
	}
}
