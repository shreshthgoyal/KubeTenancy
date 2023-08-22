// +build !ignore_autogenerated

/*
Copyright 2020 DevSpace Technologies Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	tenancy "github.com/loft-sh/kiosk/pkg/apis/tenancy"
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Account)(nil), (*tenancy.Account)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Account_To_tenancy_Account(a.(*Account), b.(*tenancy.Account), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*tenancy.Account)(nil), (*Account)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_tenancy_Account_To_v1alpha1_Account(a.(*tenancy.Account), b.(*Account), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AccountList)(nil), (*tenancy.AccountList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AccountList_To_tenancy_AccountList(a.(*AccountList), b.(*tenancy.AccountList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*tenancy.AccountList)(nil), (*AccountList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_tenancy_AccountList_To_v1alpha1_AccountList(a.(*tenancy.AccountList), b.(*AccountList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AccountSpec)(nil), (*tenancy.AccountSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AccountSpec_To_tenancy_AccountSpec(a.(*AccountSpec), b.(*tenancy.AccountSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*tenancy.AccountSpec)(nil), (*AccountSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_tenancy_AccountSpec_To_v1alpha1_AccountSpec(a.(*tenancy.AccountSpec), b.(*AccountSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AccountStatus)(nil), (*tenancy.AccountStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AccountStatus_To_tenancy_AccountStatus(a.(*AccountStatus), b.(*tenancy.AccountStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*tenancy.AccountStatus)(nil), (*AccountStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_tenancy_AccountStatus_To_v1alpha1_AccountStatus(a.(*tenancy.AccountStatus), b.(*AccountStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Space)(nil), (*tenancy.Space)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Space_To_tenancy_Space(a.(*Space), b.(*tenancy.Space), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*tenancy.Space)(nil), (*Space)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_tenancy_Space_To_v1alpha1_Space(a.(*tenancy.Space), b.(*Space), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SpaceList)(nil), (*tenancy.SpaceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SpaceList_To_tenancy_SpaceList(a.(*SpaceList), b.(*tenancy.SpaceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*tenancy.SpaceList)(nil), (*SpaceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_tenancy_SpaceList_To_v1alpha1_SpaceList(a.(*tenancy.SpaceList), b.(*SpaceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SpaceSpec)(nil), (*tenancy.SpaceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SpaceSpec_To_tenancy_SpaceSpec(a.(*SpaceSpec), b.(*tenancy.SpaceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*tenancy.SpaceSpec)(nil), (*SpaceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_tenancy_SpaceSpec_To_v1alpha1_SpaceSpec(a.(*tenancy.SpaceSpec), b.(*SpaceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SpaceStatus)(nil), (*tenancy.SpaceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SpaceStatus_To_tenancy_SpaceStatus(a.(*SpaceStatus), b.(*tenancy.SpaceStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*tenancy.SpaceStatus)(nil), (*SpaceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_tenancy_SpaceStatus_To_v1alpha1_SpaceStatus(a.(*tenancy.SpaceStatus), b.(*SpaceStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Account_To_tenancy_Account(in *Account, out *tenancy.Account, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_AccountSpec_To_tenancy_AccountSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_AccountStatus_To_tenancy_AccountStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Account_To_tenancy_Account is an autogenerated conversion function.
func Convert_v1alpha1_Account_To_tenancy_Account(in *Account, out *tenancy.Account, s conversion.Scope) error {
	return autoConvert_v1alpha1_Account_To_tenancy_Account(in, out, s)
}

func autoConvert_tenancy_Account_To_v1alpha1_Account(in *tenancy.Account, out *Account, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_tenancy_AccountSpec_To_v1alpha1_AccountSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_tenancy_AccountStatus_To_v1alpha1_AccountStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_tenancy_Account_To_v1alpha1_Account is an autogenerated conversion function.
func Convert_tenancy_Account_To_v1alpha1_Account(in *tenancy.Account, out *Account, s conversion.Scope) error {
	return autoConvert_tenancy_Account_To_v1alpha1_Account(in, out, s)
}

func autoConvert_v1alpha1_AccountList_To_tenancy_AccountList(in *AccountList, out *tenancy.AccountList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]tenancy.Account)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_AccountList_To_tenancy_AccountList is an autogenerated conversion function.
func Convert_v1alpha1_AccountList_To_tenancy_AccountList(in *AccountList, out *tenancy.AccountList, s conversion.Scope) error {
	return autoConvert_v1alpha1_AccountList_To_tenancy_AccountList(in, out, s)
}

func autoConvert_tenancy_AccountList_To_v1alpha1_AccountList(in *tenancy.AccountList, out *AccountList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Account)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_tenancy_AccountList_To_v1alpha1_AccountList is an autogenerated conversion function.
func Convert_tenancy_AccountList_To_v1alpha1_AccountList(in *tenancy.AccountList, out *AccountList, s conversion.Scope) error {
	return autoConvert_tenancy_AccountList_To_v1alpha1_AccountList(in, out, s)
}

func autoConvert_v1alpha1_AccountSpec_To_tenancy_AccountSpec(in *AccountSpec, out *tenancy.AccountSpec, s conversion.Scope) error {
	out.AccountSpec = in.AccountSpec
	return nil
}

// Convert_v1alpha1_AccountSpec_To_tenancy_AccountSpec is an autogenerated conversion function.
func Convert_v1alpha1_AccountSpec_To_tenancy_AccountSpec(in *AccountSpec, out *tenancy.AccountSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_AccountSpec_To_tenancy_AccountSpec(in, out, s)
}

func autoConvert_tenancy_AccountSpec_To_v1alpha1_AccountSpec(in *tenancy.AccountSpec, out *AccountSpec, s conversion.Scope) error {
	out.AccountSpec = in.AccountSpec
	return nil
}

// Convert_tenancy_AccountSpec_To_v1alpha1_AccountSpec is an autogenerated conversion function.
func Convert_tenancy_AccountSpec_To_v1alpha1_AccountSpec(in *tenancy.AccountSpec, out *AccountSpec, s conversion.Scope) error {
	return autoConvert_tenancy_AccountSpec_To_v1alpha1_AccountSpec(in, out, s)
}

func autoConvert_v1alpha1_AccountStatus_To_tenancy_AccountStatus(in *AccountStatus, out *tenancy.AccountStatus, s conversion.Scope) error {
	out.AccountStatus = in.AccountStatus
	return nil
}

// Convert_v1alpha1_AccountStatus_To_tenancy_AccountStatus is an autogenerated conversion function.
func Convert_v1alpha1_AccountStatus_To_tenancy_AccountStatus(in *AccountStatus, out *tenancy.AccountStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_AccountStatus_To_tenancy_AccountStatus(in, out, s)
}

func autoConvert_tenancy_AccountStatus_To_v1alpha1_AccountStatus(in *tenancy.AccountStatus, out *AccountStatus, s conversion.Scope) error {
	out.AccountStatus = in.AccountStatus
	return nil
}

// Convert_tenancy_AccountStatus_To_v1alpha1_AccountStatus is an autogenerated conversion function.
func Convert_tenancy_AccountStatus_To_v1alpha1_AccountStatus(in *tenancy.AccountStatus, out *AccountStatus, s conversion.Scope) error {
	return autoConvert_tenancy_AccountStatus_To_v1alpha1_AccountStatus(in, out, s)
}

func autoConvert_v1alpha1_Space_To_tenancy_Space(in *Space, out *tenancy.Space, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_SpaceSpec_To_tenancy_SpaceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_SpaceStatus_To_tenancy_SpaceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Space_To_tenancy_Space is an autogenerated conversion function.
func Convert_v1alpha1_Space_To_tenancy_Space(in *Space, out *tenancy.Space, s conversion.Scope) error {
	return autoConvert_v1alpha1_Space_To_tenancy_Space(in, out, s)
}

func autoConvert_tenancy_Space_To_v1alpha1_Space(in *tenancy.Space, out *Space, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_tenancy_SpaceSpec_To_v1alpha1_SpaceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_tenancy_SpaceStatus_To_v1alpha1_SpaceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_tenancy_Space_To_v1alpha1_Space is an autogenerated conversion function.
func Convert_tenancy_Space_To_v1alpha1_Space(in *tenancy.Space, out *Space, s conversion.Scope) error {
	return autoConvert_tenancy_Space_To_v1alpha1_Space(in, out, s)
}

func autoConvert_v1alpha1_SpaceList_To_tenancy_SpaceList(in *SpaceList, out *tenancy.SpaceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]tenancy.Space)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_SpaceList_To_tenancy_SpaceList is an autogenerated conversion function.
func Convert_v1alpha1_SpaceList_To_tenancy_SpaceList(in *SpaceList, out *tenancy.SpaceList, s conversion.Scope) error {
	return autoConvert_v1alpha1_SpaceList_To_tenancy_SpaceList(in, out, s)
}

func autoConvert_tenancy_SpaceList_To_v1alpha1_SpaceList(in *tenancy.SpaceList, out *SpaceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Space)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_tenancy_SpaceList_To_v1alpha1_SpaceList is an autogenerated conversion function.
func Convert_tenancy_SpaceList_To_v1alpha1_SpaceList(in *tenancy.SpaceList, out *SpaceList, s conversion.Scope) error {
	return autoConvert_tenancy_SpaceList_To_v1alpha1_SpaceList(in, out, s)
}

func autoConvert_v1alpha1_SpaceSpec_To_tenancy_SpaceSpec(in *SpaceSpec, out *tenancy.SpaceSpec, s conversion.Scope) error {
	out.Account = in.Account
	out.Finalizers = *(*[]v1.FinalizerName)(unsafe.Pointer(&in.Finalizers))
	return nil
}

// Convert_v1alpha1_SpaceSpec_To_tenancy_SpaceSpec is an autogenerated conversion function.
func Convert_v1alpha1_SpaceSpec_To_tenancy_SpaceSpec(in *SpaceSpec, out *tenancy.SpaceSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_SpaceSpec_To_tenancy_SpaceSpec(in, out, s)
}

func autoConvert_tenancy_SpaceSpec_To_v1alpha1_SpaceSpec(in *tenancy.SpaceSpec, out *SpaceSpec, s conversion.Scope) error {
	out.Account = in.Account
	out.Finalizers = *(*[]v1.FinalizerName)(unsafe.Pointer(&in.Finalizers))
	return nil
}

// Convert_tenancy_SpaceSpec_To_v1alpha1_SpaceSpec is an autogenerated conversion function.
func Convert_tenancy_SpaceSpec_To_v1alpha1_SpaceSpec(in *tenancy.SpaceSpec, out *SpaceSpec, s conversion.Scope) error {
	return autoConvert_tenancy_SpaceSpec_To_v1alpha1_SpaceSpec(in, out, s)
}

func autoConvert_v1alpha1_SpaceStatus_To_tenancy_SpaceStatus(in *SpaceStatus, out *tenancy.SpaceStatus, s conversion.Scope) error {
	out.Phase = v1.NamespacePhase(in.Phase)
	return nil
}

// Convert_v1alpha1_SpaceStatus_To_tenancy_SpaceStatus is an autogenerated conversion function.
func Convert_v1alpha1_SpaceStatus_To_tenancy_SpaceStatus(in *SpaceStatus, out *tenancy.SpaceStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_SpaceStatus_To_tenancy_SpaceStatus(in, out, s)
}

func autoConvert_tenancy_SpaceStatus_To_v1alpha1_SpaceStatus(in *tenancy.SpaceStatus, out *SpaceStatus, s conversion.Scope) error {
	out.Phase = v1.NamespacePhase(in.Phase)
	return nil
}

// Convert_tenancy_SpaceStatus_To_v1alpha1_SpaceStatus is an autogenerated conversion function.
func Convert_tenancy_SpaceStatus_To_v1alpha1_SpaceStatus(in *tenancy.SpaceStatus, out *SpaceStatus, s conversion.Scope) error {
	return autoConvert_tenancy_SpaceStatus_To_v1alpha1_SpaceStatus(in, out, s)
}
