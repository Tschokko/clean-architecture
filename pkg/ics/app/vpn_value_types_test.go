package app

import (
	"errors"
	"strings"
	"testing"
)

//
// Value Type DeviceName
//

func Test_NewDeviceName_NoErrorsIfValid(t *testing.T) {
	_, err := NewDeviceName("valid")
	if err != nil {
		t.Errorf("got an error: %s", err)
		return
	}
}

func Test_NewDeviceName_ErrorIfEmpty(t *testing.T) {
	_, err := NewDeviceName("")
	if err == nil {
		t.Errorf("no error returned")
		return
	}
	target := &EmptyDeviceNameError{}
	if !errors.As(err, &target) {
		t.Errorf("wrong error type: %s", err)
		return
	}
}

func Test_NewDeviceName_ErrorIfToLong(t *testing.T) {
	_, err := NewDeviceName(strings.Repeat("t", 101))
	if err == nil {
		t.Errorf("no error returned")
		return
	}
	target := &InvalidDeviceNameError{}
	if !errors.As(err, &target) {
		t.Errorf("wrong error type: %s", err)
		return
	}
}

//
// Value Type SerialNumber
//

func Test_NewSerialNumber_NoErrorsIfValid(t *testing.T) {
	_, err := NewSerialNumber("valid")
	if err != nil {
		t.Errorf("got an error: %s", err)
		return
	}
}

func Test_NewSerialNumber_ErrorIfEmpty(t *testing.T) {
	_, err := NewSerialNumber("")
	if err == nil {
		t.Errorf("no error returned")
		return
	}
	target := &EmptySerialNumberError{}
	if !errors.As(err, &target) {
		t.Errorf("wrong error type: %s", err)
		return
	}
}

func Test_NewSerialNumber_ErrorIfToLong(t *testing.T) {
	_, err := NewSerialNumber(strings.Repeat("t", 101))
	if err == nil {
		t.Errorf("no error returned")
		return
	}
	target := &InvalidSerialNumberError{}
	if !errors.As(err, &target) {
		t.Errorf("wrong error type: %s", err)
		return
	}
}
