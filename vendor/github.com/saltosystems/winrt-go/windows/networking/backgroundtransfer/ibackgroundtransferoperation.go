// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package backgroundtransfer

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
	"github.com/saltosystems/winrt-go/windows/storage/streams"
)

const GUIDIBackgroundTransferOperation string = "ded06846-90ca-44fb-8fb1-124154c0d539"
const SignatureIBackgroundTransferOperation string = "{ded06846-90ca-44fb-8fb1-124154c0d539}"

type IBackgroundTransferOperation struct {
	ole.IInspectable
}

type IBackgroundTransferOperationVtbl struct {
	ole.IInspectableVtbl

	GetGuid                uintptr
	GetRequestedUri        uintptr
	GetMethod              uintptr
	GetGroup               uintptr
	GetCostPolicy          uintptr
	SetCostPolicy          uintptr
	GetResultStreamAt      uintptr
	GetResponseInformation uintptr
}

func (v *IBackgroundTransferOperation) VTable() *IBackgroundTransferOperationVtbl {
	return (*IBackgroundTransferOperationVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IBackgroundTransferOperation) GetGuid() (syscall.GUID, error) {
	var out syscall.GUID
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetGuid,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out syscall.GUID
	)

	if hr != 0 {
		return syscall.GUID{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *IBackgroundTransferOperation) GetRequestedUri() (*foundation.Uri, error) {
	var out *foundation.Uri
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetRequestedUri,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.Uri
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IBackgroundTransferOperation) GetMethod() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetMethod,
		uintptr(unsafe.Pointer(v)),        // this
		uintptr(unsafe.Pointer(&outHStr)), // out string
	)

	if hr != 0 {
		return "", ole.NewError(hr)
	}

	out := outHStr.String()
	ole.DeleteHString(outHStr)
	return out, nil
}

func (v *IBackgroundTransferOperation) GetGroup() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetGroup,
		uintptr(unsafe.Pointer(v)),        // this
		uintptr(unsafe.Pointer(&outHStr)), // out string
	)

	if hr != 0 {
		return "", ole.NewError(hr)
	}

	out := outHStr.String()
	ole.DeleteHString(outHStr)
	return out, nil
}

func (v *IBackgroundTransferOperation) GetCostPolicy() (BackgroundTransferCostPolicy, error) {
	var out BackgroundTransferCostPolicy
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetCostPolicy,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out BackgroundTransferCostPolicy
	)

	if hr != 0 {
		return BackgroundTransferCostPolicyDefault, ole.NewError(hr)
	}

	return out, nil
}

func (v *IBackgroundTransferOperation) SetCostPolicy(value BackgroundTransferCostPolicy) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetCostPolicy,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(value),             // in BackgroundTransferCostPolicy
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *IBackgroundTransferOperation) GetResultStreamAt(position uint64) (*streams.IInputStream, error) {
	var out *streams.IInputStream
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetResultStreamAt,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(position),             // in uint64
		uintptr(unsafe.Pointer(&out)), // out streams.IInputStream
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IBackgroundTransferOperation) GetResponseInformation() (*ResponseInformation, error) {
	var out *ResponseInformation
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetResponseInformation,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out ResponseInformation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
