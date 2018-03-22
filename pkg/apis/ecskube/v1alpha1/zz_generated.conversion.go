// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	ecskube "github.com/robinpercy/ecs-kube/pkg/apis/ecskube"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_ECSDeployment_To_ecskube_ECSDeployment,
		Convert_ecskube_ECSDeployment_To_v1alpha1_ECSDeployment,
		Convert_v1alpha1_ECSDeploymentList_To_ecskube_ECSDeploymentList,
		Convert_ecskube_ECSDeploymentList_To_v1alpha1_ECSDeploymentList,
		Convert_v1alpha1_ECSDeploymentSpec_To_ecskube_ECSDeploymentSpec,
		Convert_ecskube_ECSDeploymentSpec_To_v1alpha1_ECSDeploymentSpec,
		Convert_v1alpha1_ECSDeploymentStatus_To_ecskube_ECSDeploymentStatus,
		Convert_ecskube_ECSDeploymentStatus_To_v1alpha1_ECSDeploymentStatus,
		Convert_v1alpha1_ECSDeploymentStatusStrategy_To_ecskube_ECSDeploymentStatusStrategy,
		Convert_ecskube_ECSDeploymentStatusStrategy_To_v1alpha1_ECSDeploymentStatusStrategy,
		Convert_v1alpha1_ECSDeploymentStrategy_To_ecskube_ECSDeploymentStrategy,
		Convert_ecskube_ECSDeploymentStrategy_To_v1alpha1_ECSDeploymentStrategy,
	)
}

func autoConvert_v1alpha1_ECSDeployment_To_ecskube_ECSDeployment(in *ECSDeployment, out *ecskube.ECSDeployment, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ECSDeploymentSpec_To_ecskube_ECSDeploymentSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ECSDeploymentStatus_To_ecskube_ECSDeploymentStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ECSDeployment_To_ecskube_ECSDeployment is an autogenerated conversion function.
func Convert_v1alpha1_ECSDeployment_To_ecskube_ECSDeployment(in *ECSDeployment, out *ecskube.ECSDeployment, s conversion.Scope) error {
	return autoConvert_v1alpha1_ECSDeployment_To_ecskube_ECSDeployment(in, out, s)
}

func autoConvert_ecskube_ECSDeployment_To_v1alpha1_ECSDeployment(in *ecskube.ECSDeployment, out *ECSDeployment, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_ecskube_ECSDeploymentSpec_To_v1alpha1_ECSDeploymentSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_ecskube_ECSDeploymentStatus_To_v1alpha1_ECSDeploymentStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_ecskube_ECSDeployment_To_v1alpha1_ECSDeployment is an autogenerated conversion function.
func Convert_ecskube_ECSDeployment_To_v1alpha1_ECSDeployment(in *ecskube.ECSDeployment, out *ECSDeployment, s conversion.Scope) error {
	return autoConvert_ecskube_ECSDeployment_To_v1alpha1_ECSDeployment(in, out, s)
}

func autoConvert_v1alpha1_ECSDeploymentList_To_ecskube_ECSDeploymentList(in *ECSDeploymentList, out *ecskube.ECSDeploymentList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ecskube.ECSDeployment)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_ECSDeploymentList_To_ecskube_ECSDeploymentList is an autogenerated conversion function.
func Convert_v1alpha1_ECSDeploymentList_To_ecskube_ECSDeploymentList(in *ECSDeploymentList, out *ecskube.ECSDeploymentList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ECSDeploymentList_To_ecskube_ECSDeploymentList(in, out, s)
}

func autoConvert_ecskube_ECSDeploymentList_To_v1alpha1_ECSDeploymentList(in *ecskube.ECSDeploymentList, out *ECSDeploymentList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ECSDeployment)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_ecskube_ECSDeploymentList_To_v1alpha1_ECSDeploymentList is an autogenerated conversion function.
func Convert_ecskube_ECSDeploymentList_To_v1alpha1_ECSDeploymentList(in *ecskube.ECSDeploymentList, out *ECSDeploymentList, s conversion.Scope) error {
	return autoConvert_ecskube_ECSDeploymentList_To_v1alpha1_ECSDeploymentList(in, out, s)
}

func autoConvert_v1alpha1_ECSDeploymentSpec_To_ecskube_ECSDeploymentSpec(in *ECSDeploymentSpec, out *ecskube.ECSDeploymentSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_ECSDeploymentSpec_To_ecskube_ECSDeploymentSpec is an autogenerated conversion function.
func Convert_v1alpha1_ECSDeploymentSpec_To_ecskube_ECSDeploymentSpec(in *ECSDeploymentSpec, out *ecskube.ECSDeploymentSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ECSDeploymentSpec_To_ecskube_ECSDeploymentSpec(in, out, s)
}

func autoConvert_ecskube_ECSDeploymentSpec_To_v1alpha1_ECSDeploymentSpec(in *ecskube.ECSDeploymentSpec, out *ECSDeploymentSpec, s conversion.Scope) error {
	return nil
}

// Convert_ecskube_ECSDeploymentSpec_To_v1alpha1_ECSDeploymentSpec is an autogenerated conversion function.
func Convert_ecskube_ECSDeploymentSpec_To_v1alpha1_ECSDeploymentSpec(in *ecskube.ECSDeploymentSpec, out *ECSDeploymentSpec, s conversion.Scope) error {
	return autoConvert_ecskube_ECSDeploymentSpec_To_v1alpha1_ECSDeploymentSpec(in, out, s)
}

func autoConvert_v1alpha1_ECSDeploymentStatus_To_ecskube_ECSDeploymentStatus(in *ECSDeploymentStatus, out *ecskube.ECSDeploymentStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_ECSDeploymentStatus_To_ecskube_ECSDeploymentStatus is an autogenerated conversion function.
func Convert_v1alpha1_ECSDeploymentStatus_To_ecskube_ECSDeploymentStatus(in *ECSDeploymentStatus, out *ecskube.ECSDeploymentStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ECSDeploymentStatus_To_ecskube_ECSDeploymentStatus(in, out, s)
}

func autoConvert_ecskube_ECSDeploymentStatus_To_v1alpha1_ECSDeploymentStatus(in *ecskube.ECSDeploymentStatus, out *ECSDeploymentStatus, s conversion.Scope) error {
	return nil
}

// Convert_ecskube_ECSDeploymentStatus_To_v1alpha1_ECSDeploymentStatus is an autogenerated conversion function.
func Convert_ecskube_ECSDeploymentStatus_To_v1alpha1_ECSDeploymentStatus(in *ecskube.ECSDeploymentStatus, out *ECSDeploymentStatus, s conversion.Scope) error {
	return autoConvert_ecskube_ECSDeploymentStatus_To_v1alpha1_ECSDeploymentStatus(in, out, s)
}

func autoConvert_v1alpha1_ECSDeploymentStatusStrategy_To_ecskube_ECSDeploymentStatusStrategy(in *ECSDeploymentStatusStrategy, out *ecskube.ECSDeploymentStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1alpha1_ECSDeploymentStatusStrategy_To_ecskube_ECSDeploymentStatusStrategy is an autogenerated conversion function.
func Convert_v1alpha1_ECSDeploymentStatusStrategy_To_ecskube_ECSDeploymentStatusStrategy(in *ECSDeploymentStatusStrategy, out *ecskube.ECSDeploymentStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_ECSDeploymentStatusStrategy_To_ecskube_ECSDeploymentStatusStrategy(in, out, s)
}

func autoConvert_ecskube_ECSDeploymentStatusStrategy_To_v1alpha1_ECSDeploymentStatusStrategy(in *ecskube.ECSDeploymentStatusStrategy, out *ECSDeploymentStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_ecskube_ECSDeploymentStatusStrategy_To_v1alpha1_ECSDeploymentStatusStrategy is an autogenerated conversion function.
func Convert_ecskube_ECSDeploymentStatusStrategy_To_v1alpha1_ECSDeploymentStatusStrategy(in *ecskube.ECSDeploymentStatusStrategy, out *ECSDeploymentStatusStrategy, s conversion.Scope) error {
	return autoConvert_ecskube_ECSDeploymentStatusStrategy_To_v1alpha1_ECSDeploymentStatusStrategy(in, out, s)
}

func autoConvert_v1alpha1_ECSDeploymentStrategy_To_ecskube_ECSDeploymentStrategy(in *ECSDeploymentStrategy, out *ecskube.ECSDeploymentStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1alpha1_ECSDeploymentStrategy_To_ecskube_ECSDeploymentStrategy is an autogenerated conversion function.
func Convert_v1alpha1_ECSDeploymentStrategy_To_ecskube_ECSDeploymentStrategy(in *ECSDeploymentStrategy, out *ecskube.ECSDeploymentStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_ECSDeploymentStrategy_To_ecskube_ECSDeploymentStrategy(in, out, s)
}

func autoConvert_ecskube_ECSDeploymentStrategy_To_v1alpha1_ECSDeploymentStrategy(in *ecskube.ECSDeploymentStrategy, out *ECSDeploymentStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_ecskube_ECSDeploymentStrategy_To_v1alpha1_ECSDeploymentStrategy is an autogenerated conversion function.
func Convert_ecskube_ECSDeploymentStrategy_To_v1alpha1_ECSDeploymentStrategy(in *ecskube.ECSDeploymentStrategy, out *ECSDeploymentStrategy, s conversion.Scope) error {
	return autoConvert_ecskube_ECSDeploymentStrategy_To_v1alpha1_ECSDeploymentStrategy(in, out, s)
}