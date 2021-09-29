// +build cortexm,!nxp,!qemu

package runtime

import (
	"device/arm"
)

var ResetOnAbort bool

func abort() {
	// optionally cause a system reset upon runtime._panic() or runtime.runtimePanic()
	if ResetOnAbort {
		arm.SystemReset()
	}
	// lock up forever
	for {
		arm.Asm("wfi")
	}
}
