package app

type apiService interface {
	CreateDevice(instanceID int, device Device) (*Device, error)
}

type Device struct {
	DeviceType   DeviceType
	Name         DeviceName
	SerialNumber SerialNumber
}

type VPNService struct {
	api apiService
}

func NewVPNService(api apiService) VPNService {
	return VPNService{
		api: api,
	}
}

func (c *VPNService) CreateIcomOSDevice(instanceID int, name DeviceName, serialNumber SerialNumber) (*Device, error) {
	device := Device{
		DeviceTypeIcomOS,
		name,
		serialNumber,
	}

	return c.api.CreateDevice(instanceID, device)
}
