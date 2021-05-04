// Code generated by mog. DO NOT EDIT.

package sourcepkg

import core "github.com/hashicorp/mog/internal/e2e/core"
// NodeToCore is a struct-to-struct conversion that translates FROM a protobuf
func NodeToCore(s Node) core.ClusterNode {
	var t core.ClusterNode
	t.ID = s.ID
	t.F1 = WorkloadToCore(s.F1)
	t.F2 = WorkloadPtrToCorePtr(s.F2)
	t.F3 = WorkloadPtrToCore(s.F3)
	t.F4 = WorkloadToCorePtr(s.F4)
	return t
}
// NodePtrToCore is a pointer-to-struct conversion that translates FROM a protobuf
func NodePtrToCore(s *Node) core.ClusterNode {
	var t core.ClusterNode
	if s == nil {
		return t
	}
	return NodeToCore(*s)
}
// NodeToCorePtr is a struct-to-pointer conversion that translates FROM a protobuf
func NodeToCorePtr(s Node) *core.ClusterNode {
	t := NodeToCore(s)
	return &t
}
// NodePtrToCorePtr is a pointer-to-pointer conversion that translates FROM a protobuf
func NodePtrToCorePtr(s *Node) *core.ClusterNode {
	if s == nil {
		return nil
	}
	t := NodeToCore(*s)
	return &t
}
// NewNodeFromCore is a struct-to-struct conversion that translates TO a protobuf
func NewNodeFromCore(t core.ClusterNode) Node {
	var s Node
	s.ID = t.ID
	s.F1 = NewWorkloadFromCore(t.F1)
	s.F2 = NewWorkloadPtrFromCorePtr(t.F2)
	s.F3 = NewWorkloadPtrFromCore(t.F3)
	s.F4 = NewWorkloadFromCorePtr(t.F4)
	return s
}
// NewNodeFromCorePtr is a pointer-to-struct conversion that translates TO a protobuf
func NewNodeFromCorePtr(t *core.ClusterNode) Node {
	var s Node
	if t == nil {
		return s
	}
	return NewNodeFromCore(*t)
}
// NewNodePtrFromCore is a struct-pointer conversion that translates TO a protobuf
func NewNodePtrFromCore(t core.ClusterNode) *Node {
	s := NewNodeFromCore(t)
	return &s
}
// NewNodePtrFromCorePtr is a pointer-to-pointer conversion that translates TO a protobuf
func NewNodePtrFromCorePtr(t *core.ClusterNode) *Node {
	if t == nil {
		return nil
	}
	s := NewNodeFromCore(*t)
	return &s
}
// WorkloadToCore is a struct-to-struct conversion that translates FROM a protobuf
func WorkloadToCore(s Workload) core.Workload {
	var t core.Workload
	t.ID = s.ID
	return t
}
// WorkloadPtrToCore is a pointer-to-struct conversion that translates FROM a protobuf
func WorkloadPtrToCore(s *Workload) core.Workload {
	var t core.Workload
	if s == nil {
		return t
	}
	return WorkloadToCore(*s)
}
// WorkloadToCorePtr is a struct-to-pointer conversion that translates FROM a protobuf
func WorkloadToCorePtr(s Workload) *core.Workload {
	t := WorkloadToCore(s)
	return &t
}
// WorkloadPtrToCorePtr is a pointer-to-pointer conversion that translates FROM a protobuf
func WorkloadPtrToCorePtr(s *Workload) *core.Workload {
	if s == nil {
		return nil
	}
	t := WorkloadToCore(*s)
	return &t
}
// NewWorkloadFromCore is a struct-to-struct conversion that translates TO a protobuf
func NewWorkloadFromCore(t core.Workload) Workload {
	var s Workload
	s.ID = t.ID
	return s
}
// NewWorkloadFromCorePtr is a pointer-to-struct conversion that translates TO a protobuf
func NewWorkloadFromCorePtr(t *core.Workload) Workload {
	var s Workload
	if t == nil {
		return s
	}
	return NewWorkloadFromCore(*t)
}
// NewWorkloadPtrFromCore is a struct-pointer conversion that translates TO a protobuf
func NewWorkloadPtrFromCore(t core.Workload) *Workload {
	s := NewWorkloadFromCore(t)
	return &s
}
// NewWorkloadPtrFromCorePtr is a pointer-to-pointer conversion that translates TO a protobuf
func NewWorkloadPtrFromCorePtr(t *core.Workload) *Workload {
	if t == nil {
		return nil
	}
	s := NewWorkloadFromCore(*t)
	return &s
}