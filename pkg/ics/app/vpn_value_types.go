package app

import "strings"

type DeviceType int64

const (
	DeviceTypeIcomOS DeviceType = iota
)

type DeviceName struct {
	deviceName string
}

func NewDeviceName(deviceName string) (*DeviceName, error) {
	if strings.TrimSpace(deviceName) == "" {
		return nil, &EmptyDeviceNameError{}
	}
	if len(deviceName) > 100 {
		return nil, &InvalidDeviceNameError{}
	}
	return &DeviceName{deviceName: deviceName}, nil
}

type SerialNumber struct {
	serialNumber string
}

func NewSerialNumber(serialNumber string) (*SerialNumber, error) {
	if strings.TrimSpace(serialNumber) == "" {
		return nil, &EmptySerialNumberError{}
	}
	if len(serialNumber) > 100 {
		return nil, &InvalidSerialNumberError{}
	}
	return &SerialNumber{serialNumber: serialNumber}, nil
}
