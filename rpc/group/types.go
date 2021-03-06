package group

import (
	"github.com/docker/infrakit/spi/group"
)

// CommitGroupRequest is the rpc wrapper for input to commit a group
type CommitGroupRequest struct {
	Spec    group.Spec
	Pretend bool
}

// CommitGroupResponse is the rpc wrapper for the results to commit a group
type CommitGroupResponse struct {
	Details string
}

// FreeGroupRequest is the rpc wrapper for input to free a group
type FreeGroupRequest struct {
	ID group.ID
}

// FreeGroupResponse is the rpc wrapper for the results to free a group
type FreeGroupResponse struct {
	OK bool
}

// DescribeGroupRequest is the rpc wrapper for the input to inspect a group
type DescribeGroupRequest struct {
	ID group.ID
}

// DescribeGroupResponse is the rpc wrapper for the results from inspecting a group
type DescribeGroupResponse struct {
	Description group.Description
}

// DestroyGroupRequest is the rpc wrapper for the input to destroy a group
type DestroyGroupRequest struct {
	ID group.ID
}

// DestroyGroupResponse is the rpc wrapper for the output from destroying a group
type DestroyGroupResponse struct {
	OK bool
}

// InspectGroupsRequest is the rpc wrapper for the input to destroy a group
type InspectGroupsRequest struct {
	ID group.ID
}

// InspectGroupsResponse is the rpc wrapper for the output from destroying a group
type InspectGroupsResponse struct {
	Groups []group.Spec
}

// RPCService is the interface for exposing the group plugin as a RPC service. It conforms to the call conventions
// defined in net/rpc
type RPCService interface {
	CommitGroup(req *CommitGroupRequest, resp *CommitGroupResponse) error
	FreeGroup(req *FreeGroupRequest, resp *FreeGroupResponse) error
	DescribeGroup(req *DescribeGroupRequest, resp *DescribeGroupResponse) error
	DestroyGroup(req *DestroyGroupRequest, resp *DestroyGroupResponse) error
	InspectGroups(req *InspectGroupsRequest, resp *InspectGroupsResponse) error
}
