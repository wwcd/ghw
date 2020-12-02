// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package bios

import (
	"github.com/StackExchange/wmi"
)

const wqlBIOS = "SELECT SerialNumber, InstallDate, Manufacturer, Version FROM CIM_BIOSElement"

type win32BIOS struct {
	SerialNumber *string
	InstallDate  *string
	Manufacturer *string
	Version      *string
}

func (i *Info) load() error {
	// Getting data from WMI
	var win32BIOSDescriptions []win32BIOS
	if err := wmi.Query(wqlBIOS, &win32BIOSDescriptions); err != nil {
		return err
	}
	if len(win32BIOSDescriptions) > 0 {
		i.SerialNumber = *win32BIOSDescriptions[0].SerialNumber
		i.Vendor = *win32BIOSDescriptions[0].Manufacturer
		i.Version = *win32BIOSDescriptions[0].Version
		i.Date = *win32BIOSDescriptions[0].InstallDate
	}
	return nil
}
