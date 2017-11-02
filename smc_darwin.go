package smc

/*
#cgo LDFLAGS: -framework IOKit
#include <IOKit/IOKitLib.h>
#include "smc.h"
*/
import "C"

/*
Temperature sensors
-------------------
TB0T
TC0D
TC0P
TM0P
TN0P
Th0H
Ts0P
TN1P
Th1H
*/
const KEY_CPU_TEMP string = "TC0P"
const KEY_FAN0_SPEED string = "F0Ac"
const KEY_FAN1_SPEED string = "F1Ac"

func Open() {
	C.SMCOpen()
}

func GetTemperature(key string) float64 {
	if key == "" {
		key = KEY_CPU_TEMP
	}
	return float64(C.SMCGetTemperature(C.CString(key)))
}

func GetFanSpeed(key string) float64 {
	if key == "" {
		key = KEY_FAN0_SPEED
	}
	return float64(C.SMCGetFanSpeed(C.CString(key)))
}

func Close() {
	C.SMCClose()
}
