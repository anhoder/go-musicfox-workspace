// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package backgroundtransfer

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
	"github.com/saltosystems/winrt-go/windows/foundation/collections"
)

const SignatureResponseInformation string = "rc(Windows.Networking.BackgroundTransfer.ResponseInformation;{f8bb9a12-f713-4792-8b68-d9d297f91d2e})"

type ResponseInformation struct {
	ole.IUnknown
}

func (impl *ResponseInformation) GetIsResumable() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiResponseInformation))
	defer itf.Release()
	v := (*iResponseInformation)(unsafe.Pointer(itf))
	return v.GetIsResumable()
}

func (impl *ResponseInformation) GetActualUri() (*foundation.Uri, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiResponseInformation))
	defer itf.Release()
	v := (*iResponseInformation)(unsafe.Pointer(itf))
	return v.GetActualUri()
}

func (impl *ResponseInformation) GetStatusCode() (uint32, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiResponseInformation))
	defer itf.Release()
	v := (*iResponseInformation)(unsafe.Pointer(itf))
	return v.GetStatusCode()
}

func (impl *ResponseInformation) GetHeaders() (*collections.IMapView, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiResponseInformation))
	defer itf.Release()
	v := (*iResponseInformation)(unsafe.Pointer(itf))
	return v.GetHeaders()
}

const GUIDiResponseInformation string = "f8bb9a12-f713-4792-8b68-d9d297f91d2e"
const SignatureiResponseInformation string = "{f8bb9a12-f713-4792-8b68-d9d297f91d2e}"

type iResponseInformation struct {
	ole.IInspectable
}

type iResponseInformationVtbl struct {
	ole.IInspectableVtbl

	GetIsResumable uintptr
	GetActualUri   uintptr
	GetStatusCode  uintptr
	GetHeaders     uintptr
}

func (v *iResponseInformation) VTable() *iResponseInformationVtbl {
	return (*iResponseInformationVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iResponseInformation) GetIsResumable() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetIsResumable,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iResponseInformation) GetActualUri() (*foundation.Uri, error) {
	var out *foundation.Uri
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetActualUri,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.Uri
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iResponseInformation) GetStatusCode() (uint32, error) {
	var out uint32
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetStatusCode,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint32
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iResponseInformation) GetHeaders() (*collections.IMapView, error) {
	var out *collections.IMapView
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetHeaders,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out collections.IMapView
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
