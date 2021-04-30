// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1beta1sets

import (
	networking_enterprise_mesh_gloo_solo_io_v1beta1 "github.com/solo-io/gloo-mesh/pkg/api/networking.enterprise.mesh.gloo.solo.io/v1beta1"

	"github.com/rotisserie/eris"
	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type WasmDeploymentSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment
	// Insert a resource into the set.
	Insert(wasmDeployment ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(wasmDeploymentSet WasmDeploymentSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(wasmDeployment ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(wasmDeployment ezkube.ResourceId)
	// Return the union with the provided set
	Union(set WasmDeploymentSet) WasmDeploymentSet
	// Return the difference with the provided set
	Difference(set WasmDeploymentSet) WasmDeploymentSet
	// Return the intersection with the provided set
	Intersection(set WasmDeploymentSet) WasmDeploymentSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another WasmDeploymentSet
	Delta(newSet WasmDeploymentSet) sksets.ResourceDelta
}

func makeGenericWasmDeploymentSet(wasmDeploymentList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range wasmDeploymentList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type wasmDeploymentSet struct {
	set sksets.ResourceSet
}

func NewWasmDeploymentSet(wasmDeploymentList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) WasmDeploymentSet {
	return &wasmDeploymentSet{set: makeGenericWasmDeploymentSet(wasmDeploymentList)}
}

func NewWasmDeploymentSetFromList(wasmDeploymentList *networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeploymentList) WasmDeploymentSet {
	list := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment, 0, len(wasmDeploymentList.Items))
	for idx := range wasmDeploymentList.Items {
		list = append(list, &wasmDeploymentList.Items[idx])
	}
	return &wasmDeploymentSet{set: makeGenericWasmDeploymentSet(list)}
}

func (s *wasmDeploymentSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *wasmDeploymentSet) List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment))
		})
	}

	objs := s.Generic().List(genericFilters...)
	wasmDeploymentList := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment, 0, len(objs))
	for _, obj := range objs {
		wasmDeploymentList = append(wasmDeploymentList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment))
	}
	return wasmDeploymentList
}

func (s *wasmDeploymentSet) UnsortedList(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment))
		})
	}

	var wasmDeploymentList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		wasmDeploymentList = append(wasmDeploymentList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment))
	}
	return wasmDeploymentList
}

func (s *wasmDeploymentSet) Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment)
	}
	return newMap
}

func (s *wasmDeploymentSet) Insert(
	wasmDeploymentList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range wasmDeploymentList {
		s.Generic().Insert(obj)
	}
}

func (s *wasmDeploymentSet) Has(wasmDeployment ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(wasmDeployment)
}

func (s *wasmDeploymentSet) Equal(
	wasmDeploymentSet WasmDeploymentSet,
) bool {
	if s == nil {
		return wasmDeploymentSet == nil
	}
	return s.Generic().Equal(wasmDeploymentSet.Generic())
}

func (s *wasmDeploymentSet) Delete(WasmDeployment ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(WasmDeployment)
}

func (s *wasmDeploymentSet) Union(set WasmDeploymentSet) WasmDeploymentSet {
	if s == nil {
		return set
	}
	return NewWasmDeploymentSet(append(s.List(), set.List()...)...)
}

func (s *wasmDeploymentSet) Difference(set WasmDeploymentSet) WasmDeploymentSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &wasmDeploymentSet{set: newSet}
}

func (s *wasmDeploymentSet) Intersection(set WasmDeploymentSet) WasmDeploymentSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var wasmDeploymentList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment
	for _, obj := range newSet.List() {
		wasmDeploymentList = append(wasmDeploymentList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment))
	}
	return NewWasmDeploymentSet(wasmDeploymentList...)
}

func (s *wasmDeploymentSet) Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find WasmDeployment %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.WasmDeployment), nil
}

func (s *wasmDeploymentSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *wasmDeploymentSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *wasmDeploymentSet) Delta(newSet WasmDeploymentSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

type VirtualDestinationSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination
	// Insert a resource into the set.
	Insert(virtualDestination ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(virtualDestinationSet VirtualDestinationSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(virtualDestination ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(virtualDestination ezkube.ResourceId)
	// Return the union with the provided set
	Union(set VirtualDestinationSet) VirtualDestinationSet
	// Return the difference with the provided set
	Difference(set VirtualDestinationSet) VirtualDestinationSet
	// Return the intersection with the provided set
	Intersection(set VirtualDestinationSet) VirtualDestinationSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another VirtualDestinationSet
	Delta(newSet VirtualDestinationSet) sksets.ResourceDelta
}

func makeGenericVirtualDestinationSet(virtualDestinationList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range virtualDestinationList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type virtualDestinationSet struct {
	set sksets.ResourceSet
}

func NewVirtualDestinationSet(virtualDestinationList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) VirtualDestinationSet {
	return &virtualDestinationSet{set: makeGenericVirtualDestinationSet(virtualDestinationList)}
}

func NewVirtualDestinationSetFromList(virtualDestinationList *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestinationList) VirtualDestinationSet {
	list := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination, 0, len(virtualDestinationList.Items))
	for idx := range virtualDestinationList.Items {
		list = append(list, &virtualDestinationList.Items[idx])
	}
	return &virtualDestinationSet{set: makeGenericVirtualDestinationSet(list)}
}

func (s *virtualDestinationSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *virtualDestinationSet) List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination))
		})
	}

	objs := s.Generic().List(genericFilters...)
	virtualDestinationList := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination, 0, len(objs))
	for _, obj := range objs {
		virtualDestinationList = append(virtualDestinationList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination))
	}
	return virtualDestinationList
}

func (s *virtualDestinationSet) UnsortedList(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination))
		})
	}

	var virtualDestinationList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		virtualDestinationList = append(virtualDestinationList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination))
	}
	return virtualDestinationList
}

func (s *virtualDestinationSet) Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination)
	}
	return newMap
}

func (s *virtualDestinationSet) Insert(
	virtualDestinationList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range virtualDestinationList {
		s.Generic().Insert(obj)
	}
}

func (s *virtualDestinationSet) Has(virtualDestination ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(virtualDestination)
}

func (s *virtualDestinationSet) Equal(
	virtualDestinationSet VirtualDestinationSet,
) bool {
	if s == nil {
		return virtualDestinationSet == nil
	}
	return s.Generic().Equal(virtualDestinationSet.Generic())
}

func (s *virtualDestinationSet) Delete(VirtualDestination ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(VirtualDestination)
}

func (s *virtualDestinationSet) Union(set VirtualDestinationSet) VirtualDestinationSet {
	if s == nil {
		return set
	}
	return NewVirtualDestinationSet(append(s.List(), set.List()...)...)
}

func (s *virtualDestinationSet) Difference(set VirtualDestinationSet) VirtualDestinationSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &virtualDestinationSet{set: newSet}
}

func (s *virtualDestinationSet) Intersection(set VirtualDestinationSet) VirtualDestinationSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var virtualDestinationList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination
	for _, obj := range newSet.List() {
		virtualDestinationList = append(virtualDestinationList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination))
	}
	return NewVirtualDestinationSet(virtualDestinationList...)
}

func (s *virtualDestinationSet) Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find VirtualDestination %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualDestination), nil
}

func (s *virtualDestinationSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *virtualDestinationSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *virtualDestinationSet) Delta(newSet VirtualDestinationSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

type VirtualGatewaySet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway
	// Insert a resource into the set.
	Insert(virtualGateway ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(virtualGatewaySet VirtualGatewaySet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(virtualGateway ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(virtualGateway ezkube.ResourceId)
	// Return the union with the provided set
	Union(set VirtualGatewaySet) VirtualGatewaySet
	// Return the difference with the provided set
	Difference(set VirtualGatewaySet) VirtualGatewaySet
	// Return the intersection with the provided set
	Intersection(set VirtualGatewaySet) VirtualGatewaySet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another VirtualGatewaySet
	Delta(newSet VirtualGatewaySet) sksets.ResourceDelta
}

func makeGenericVirtualGatewaySet(virtualGatewayList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range virtualGatewayList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type virtualGatewaySet struct {
	set sksets.ResourceSet
}

func NewVirtualGatewaySet(virtualGatewayList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway) VirtualGatewaySet {
	return &virtualGatewaySet{set: makeGenericVirtualGatewaySet(virtualGatewayList)}
}

func NewVirtualGatewaySetFromList(virtualGatewayList *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGatewayList) VirtualGatewaySet {
	list := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway, 0, len(virtualGatewayList.Items))
	for idx := range virtualGatewayList.Items {
		list = append(list, &virtualGatewayList.Items[idx])
	}
	return &virtualGatewaySet{set: makeGenericVirtualGatewaySet(list)}
}

func (s *virtualGatewaySet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *virtualGatewaySet) List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway))
		})
	}

	objs := s.Generic().List(genericFilters...)
	virtualGatewayList := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway, 0, len(objs))
	for _, obj := range objs {
		virtualGatewayList = append(virtualGatewayList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway))
	}
	return virtualGatewayList
}

func (s *virtualGatewaySet) UnsortedList(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway))
		})
	}

	var virtualGatewayList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		virtualGatewayList = append(virtualGatewayList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway))
	}
	return virtualGatewayList
}

func (s *virtualGatewaySet) Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway)
	}
	return newMap
}

func (s *virtualGatewaySet) Insert(
	virtualGatewayList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range virtualGatewayList {
		s.Generic().Insert(obj)
	}
}

func (s *virtualGatewaySet) Has(virtualGateway ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(virtualGateway)
}

func (s *virtualGatewaySet) Equal(
	virtualGatewaySet VirtualGatewaySet,
) bool {
	if s == nil {
		return virtualGatewaySet == nil
	}
	return s.Generic().Equal(virtualGatewaySet.Generic())
}

func (s *virtualGatewaySet) Delete(VirtualGateway ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(VirtualGateway)
}

func (s *virtualGatewaySet) Union(set VirtualGatewaySet) VirtualGatewaySet {
	if s == nil {
		return set
	}
	return NewVirtualGatewaySet(append(s.List(), set.List()...)...)
}

func (s *virtualGatewaySet) Difference(set VirtualGatewaySet) VirtualGatewaySet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &virtualGatewaySet{set: newSet}
}

func (s *virtualGatewaySet) Intersection(set VirtualGatewaySet) VirtualGatewaySet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var virtualGatewayList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway
	for _, obj := range newSet.List() {
		virtualGatewayList = append(virtualGatewayList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway))
	}
	return NewVirtualGatewaySet(virtualGatewayList...)
}

func (s *virtualGatewaySet) Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find VirtualGateway %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualGateway), nil
}

func (s *virtualGatewaySet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *virtualGatewaySet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *virtualGatewaySet) Delta(newSet VirtualGatewaySet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

type VirtualHostSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost
	// Insert a resource into the set.
	Insert(virtualHost ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(virtualHostSet VirtualHostSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(virtualHost ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(virtualHost ezkube.ResourceId)
	// Return the union with the provided set
	Union(set VirtualHostSet) VirtualHostSet
	// Return the difference with the provided set
	Difference(set VirtualHostSet) VirtualHostSet
	// Return the intersection with the provided set
	Intersection(set VirtualHostSet) VirtualHostSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another VirtualHostSet
	Delta(newSet VirtualHostSet) sksets.ResourceDelta
}

func makeGenericVirtualHostSet(virtualHostList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range virtualHostList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type virtualHostSet struct {
	set sksets.ResourceSet
}

func NewVirtualHostSet(virtualHostList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost) VirtualHostSet {
	return &virtualHostSet{set: makeGenericVirtualHostSet(virtualHostList)}
}

func NewVirtualHostSetFromList(virtualHostList *networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHostList) VirtualHostSet {
	list := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost, 0, len(virtualHostList.Items))
	for idx := range virtualHostList.Items {
		list = append(list, &virtualHostList.Items[idx])
	}
	return &virtualHostSet{set: makeGenericVirtualHostSet(list)}
}

func (s *virtualHostSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *virtualHostSet) List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost))
		})
	}

	objs := s.Generic().List(genericFilters...)
	virtualHostList := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost, 0, len(objs))
	for _, obj := range objs {
		virtualHostList = append(virtualHostList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost))
	}
	return virtualHostList
}

func (s *virtualHostSet) UnsortedList(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost))
		})
	}

	var virtualHostList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		virtualHostList = append(virtualHostList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost))
	}
	return virtualHostList
}

func (s *virtualHostSet) Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost)
	}
	return newMap
}

func (s *virtualHostSet) Insert(
	virtualHostList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range virtualHostList {
		s.Generic().Insert(obj)
	}
}

func (s *virtualHostSet) Has(virtualHost ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(virtualHost)
}

func (s *virtualHostSet) Equal(
	virtualHostSet VirtualHostSet,
) bool {
	if s == nil {
		return virtualHostSet == nil
	}
	return s.Generic().Equal(virtualHostSet.Generic())
}

func (s *virtualHostSet) Delete(VirtualHost ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(VirtualHost)
}

func (s *virtualHostSet) Union(set VirtualHostSet) VirtualHostSet {
	if s == nil {
		return set
	}
	return NewVirtualHostSet(append(s.List(), set.List()...)...)
}

func (s *virtualHostSet) Difference(set VirtualHostSet) VirtualHostSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &virtualHostSet{set: newSet}
}

func (s *virtualHostSet) Intersection(set VirtualHostSet) VirtualHostSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var virtualHostList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost
	for _, obj := range newSet.List() {
		virtualHostList = append(virtualHostList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost))
	}
	return NewVirtualHostSet(virtualHostList...)
}

func (s *virtualHostSet) Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find VirtualHost %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.VirtualHost), nil
}

func (s *virtualHostSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *virtualHostSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *virtualHostSet) Delta(newSet VirtualHostSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}

type RouteTableSet interface {
	// Get the set stored keys
	Keys() sets.String
	// List of resources stored in the set. Pass an optional filter function to filter on the list.
	List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable
	// Unsorted list of resources stored in the set. Pass an optional filter function to filter on the list.
	UnsortedList(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable
	// Return the Set as a map of key to resource.
	Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable
	// Insert a resource into the set.
	Insert(routeTable ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable)
	// Compare the equality of the keys in two sets (not the resources themselves)
	Equal(routeTableSet RouteTableSet) bool
	// Check if the set contains a key matching the resource (not the resource itself)
	Has(routeTable ezkube.ResourceId) bool
	// Delete the key matching the resource
	Delete(routeTable ezkube.ResourceId)
	// Return the union with the provided set
	Union(set RouteTableSet) RouteTableSet
	// Return the difference with the provided set
	Difference(set RouteTableSet) RouteTableSet
	// Return the intersection with the provided set
	Intersection(set RouteTableSet) RouteTableSet
	// Find the resource with the given ID
	Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable, error)
	// Get the length of the set
	Length() int
	// returns the generic implementation of the set
	Generic() sksets.ResourceSet
	// returns the delta between this and and another RouteTableSet
	Delta(newSet RouteTableSet) sksets.ResourceDelta
}

func makeGenericRouteTableSet(routeTableList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range routeTableList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type routeTableSet struct {
	set sksets.ResourceSet
}

func NewRouteTableSet(routeTableList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable) RouteTableSet {
	return &routeTableSet{set: makeGenericRouteTableSet(routeTableList)}
}

func NewRouteTableSetFromList(routeTableList *networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTableList) RouteTableSet {
	list := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable, 0, len(routeTableList.Items))
	for idx := range routeTableList.Items {
		list = append(list, &routeTableList.Items[idx])
	}
	return &routeTableSet{set: makeGenericRouteTableSet(list)}
}

func (s *routeTableSet) Keys() sets.String {
	if s == nil {
		return sets.String{}
	}
	return s.Generic().Keys()
}

func (s *routeTableSet) List(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable))
		})
	}

	objs := s.Generic().List(genericFilters...)
	routeTableList := make([]*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable, 0, len(objs))
	for _, obj := range objs {
		routeTableList = append(routeTableList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable))
	}
	return routeTableList
}

func (s *routeTableSet) UnsortedList(filterResource ...func(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable) bool) []*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable {
	if s == nil {
		return nil
	}
	var genericFilters []func(ezkube.ResourceId) bool
	for _, filter := range filterResource {
		genericFilters = append(genericFilters, func(obj ezkube.ResourceId) bool {
			return filter(obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable))
		})
	}

	var routeTableList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable
	for _, obj := range s.Generic().UnsortedList(genericFilters...) {
		routeTableList = append(routeTableList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable))
	}
	return routeTableList
}

func (s *routeTableSet) Map() map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable {
	if s == nil {
		return nil
	}

	newMap := map[string]*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable{}
	for k, v := range s.Generic().Map() {
		newMap[k] = v.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable)
	}
	return newMap
}

func (s *routeTableSet) Insert(
	routeTableList ...*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable,
) {
	if s == nil {
		panic("cannot insert into nil set")
	}

	for _, obj := range routeTableList {
		s.Generic().Insert(obj)
	}
}

func (s *routeTableSet) Has(routeTable ezkube.ResourceId) bool {
	if s == nil {
		return false
	}
	return s.Generic().Has(routeTable)
}

func (s *routeTableSet) Equal(
	routeTableSet RouteTableSet,
) bool {
	if s == nil {
		return routeTableSet == nil
	}
	return s.Generic().Equal(routeTableSet.Generic())
}

func (s *routeTableSet) Delete(RouteTable ezkube.ResourceId) {
	if s == nil {
		return
	}
	s.Generic().Delete(RouteTable)
}

func (s *routeTableSet) Union(set RouteTableSet) RouteTableSet {
	if s == nil {
		return set
	}
	return NewRouteTableSet(append(s.List(), set.List()...)...)
}

func (s *routeTableSet) Difference(set RouteTableSet) RouteTableSet {
	if s == nil {
		return set
	}
	newSet := s.Generic().Difference(set.Generic())
	return &routeTableSet{set: newSet}
}

func (s *routeTableSet) Intersection(set RouteTableSet) RouteTableSet {
	if s == nil {
		return nil
	}
	newSet := s.Generic().Intersection(set.Generic())
	var routeTableList []*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable
	for _, obj := range newSet.List() {
		routeTableList = append(routeTableList, obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable))
	}
	return NewRouteTableSet(routeTableList...)
}

func (s *routeTableSet) Find(id ezkube.ResourceId) (*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable, error) {
	if s == nil {
		return nil, eris.Errorf("empty set, cannot find RouteTable %v", sksets.Key(id))
	}
	obj, err := s.Generic().Find(&networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*networking_enterprise_mesh_gloo_solo_io_v1beta1.RouteTable), nil
}

func (s *routeTableSet) Length() int {
	if s == nil {
		return 0
	}
	return s.Generic().Length()
}

func (s *routeTableSet) Generic() sksets.ResourceSet {
	if s == nil {
		return nil
	}
	return s.set
}

func (s *routeTableSet) Delta(newSet RouteTableSet) sksets.ResourceDelta {
	if s == nil {
		return sksets.ResourceDelta{
			Inserted: newSet.Generic(),
		}
	}
	return s.Generic().Delta(newSet.Generic())
}
