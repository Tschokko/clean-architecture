package app

type InvalidDeviceNameError struct{}

func (e *InvalidDeviceNameError) Error() string {
	return "invalid device name"
}

type EmptyDeviceNameError struct{}

func (e *EmptyDeviceNameError) Error() string {
	return "empty device name"
}

type InvalidSerialNumberError struct{}

func (e *InvalidSerialNumberError) Error() string {
	return "invalid serial number"
}

type EmptySerialNumberError struct{}

func (e *EmptySerialNumberError) Error() string {
	return "empty serial number"
}

type APIError struct {
}

func (e *APIError) Error() string {
	return "something went wrong with the iCS-VPN API :-/"
}

type InstanceNotFoundError struct {
}

func (e *InstanceNotFoundError) Error() string {
	return "instance not found"
}

type DuplicateDeviceNameError struct{}

func (e *DuplicateDeviceNameError) Error() string {
	return "duplicate device name"
}

type DuplicateSerialNumberError struct{}

func (e *DuplicateSerialNumberError) Error() string {
	return "duplicate serial number"
}
