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

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package ecskube

import (
	"fmt"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
)

var (
	InternalECSDeployment = builders.NewInternalResource(
		"ecsdeployments",
		"ECSDeployment",
		func() runtime.Object { return &ECSDeployment{} },
		func() runtime.Object { return &ECSDeploymentList{} },
	)
	InternalECSDeploymentStatus = builders.NewInternalResourceStatus(
		"ecsdeployments",
		"ECSDeploymentStatus",
		func() runtime.Object { return &ECSDeployment{} },
		func() runtime.Object { return &ECSDeploymentList{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("ecskube.ecskube").WithKinds(
		InternalECSDeployment,
		InternalECSDeploymentStatus,
	)

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ECSDeployment struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   ECSDeploymentSpec
	Status ECSDeploymentStatus
}

type ECSDeploymentSpec struct {
}

type ECSDeploymentStatus struct {
}

//
// ECSDeployment Functions and Structs
//
// +k8s:deepcopy-gen=false
type ECSDeploymentStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type ECSDeploymentStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ECSDeploymentList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []ECSDeployment
}

func (ECSDeployment) NewStatus() interface{} {
	return ECSDeploymentStatus{}
}

func (pc *ECSDeployment) GetStatus() interface{} {
	return pc.Status
}

func (pc *ECSDeployment) SetStatus(s interface{}) {
	pc.Status = s.(ECSDeploymentStatus)
}

func (pc *ECSDeployment) GetSpec() interface{} {
	return pc.Spec
}

func (pc *ECSDeployment) SetSpec(s interface{}) {
	pc.Spec = s.(ECSDeploymentSpec)
}

func (pc *ECSDeployment) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *ECSDeployment) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc ECSDeployment) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store ECSDeployment.
// +k8s:deepcopy-gen=false
type ECSDeploymentRegistry interface {
	ListECSDeployments(ctx request.Context, options *internalversion.ListOptions) (*ECSDeploymentList, error)
	GetECSDeployment(ctx request.Context, id string, options *metav1.GetOptions) (*ECSDeployment, error)
	CreateECSDeployment(ctx request.Context, id *ECSDeployment) (*ECSDeployment, error)
	UpdateECSDeployment(ctx request.Context, id *ECSDeployment) (*ECSDeployment, error)
	DeleteECSDeployment(ctx request.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewECSDeploymentRegistry(sp builders.StandardStorageProvider) ECSDeploymentRegistry {
	return &storageECSDeployment{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageECSDeployment struct {
	builders.StandardStorageProvider
}

func (s *storageECSDeployment) ListECSDeployments(ctx request.Context, options *internalversion.ListOptions) (*ECSDeploymentList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*ECSDeploymentList), err
}

func (s *storageECSDeployment) GetECSDeployment(ctx request.Context, id string, options *metav1.GetOptions) (*ECSDeployment, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*ECSDeployment), nil
}

func (s *storageECSDeployment) CreateECSDeployment(ctx request.Context, object *ECSDeployment) (*ECSDeployment, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, true)
	if err != nil {
		return nil, err
	}
	return obj.(*ECSDeployment), nil
}

func (s *storageECSDeployment) UpdateECSDeployment(ctx request.Context, object *ECSDeployment) (*ECSDeployment, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil)
	if err != nil {
		return nil, err
	}
	return obj.(*ECSDeployment), nil
}

func (s *storageECSDeployment) DeleteECSDeployment(ctx request.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil)
	return sync, err
}