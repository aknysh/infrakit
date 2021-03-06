package group

import (
	"github.com/docker/infrakit/spi/group"
)

// PluginServer returns a RPCService that conforms to the net/rpc rpc call convention.
func PluginServer(p group.Plugin) RPCService {
	return &Group{plugin: p}
}

// Group the exported type needed to conform to json-rpc call convention
type Group struct {
	plugin group.Plugin
}

// CommitGroup is the rpc method to commit a group
func (p *Group) CommitGroup(req *CommitGroupRequest, resp *CommitGroupResponse) error {
	details, err := p.plugin.CommitGroup(req.Spec, req.Pretend)
	if err != nil {
		return err
	}
	resp.Details = details
	return nil
}

// FreeGroup is the rpc method to free a group
func (p *Group) FreeGroup(req *FreeGroupRequest, resp *FreeGroupResponse) error {
	err := p.plugin.FreeGroup(req.ID)
	if err != nil {
		return err
	}
	resp.OK = true
	return nil
}

// DescribeGroup is the rpc method to describe a group
func (p *Group) DescribeGroup(req *DescribeGroupRequest, resp *DescribeGroupResponse) error {
	desc, err := p.plugin.DescribeGroup(req.ID)
	if err != nil {
		return err
	}
	resp.Description = desc
	return nil
}

// DestroyGroup is the rpc method to destroy a group
func (p *Group) DestroyGroup(req *DestroyGroupRequest, resp *DestroyGroupResponse) error {
	err := p.plugin.DestroyGroup(req.ID)
	if err != nil {
		return err
	}
	resp.OK = true
	return nil
}

// InspectGroups is the rpc method to inspect groups
func (p *Group) InspectGroups(req *InspectGroupsRequest, resp *InspectGroupsResponse) error {
	groups, err := p.plugin.InspectGroups()
	if err != nil {
		return err
	}
	resp.Groups = groups
	return nil
}
