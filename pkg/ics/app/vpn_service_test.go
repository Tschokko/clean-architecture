package app

import "testing"

type mockAPIService struct {
}

func (m *mockAPIService) CreateDevice(instanceID int, device Device) (*Device, error) {
	return &Device{
		device.DeviceType,
		device.Name,
		device.SerialNumber,
	}, nil
}

func Test_CreateIcomOSDevice_IsDeviceTypeIcomOS(t *testing.T) {
	service := NewVPNService(&mockAPIService{})
	deviceName, _ := NewDeviceName("test")
	serialNumber, _ := NewSerialNumber("test")
	device, _ := service.CreateIcomOSDevice(1, *deviceName, *serialNumber)
	if device.DeviceType != DeviceTypeIcomOS {
		t.Errorf("device type is not DeviceTypIcomOS")
	}
}
