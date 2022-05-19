/* 
    Utils to facilitate easier usage of the SDK.
*/
package opal

func GroupToUpdateGroupInfo(group Group) (updateGroupInfo UpdateGroupInfo) {
    return UpdateGroupInfo{
        GroupId: group.GroupId,
        Description: group.Description,
        OwnerTeamId: group.OwnerTeamId,
        Visibility: group.Visibility,
        MaxDuration: group.MaxDuration,
        RequireManagerApproval: group.RequireManagerApproval,
        RequireSupportTicket: group.RequireSupportTicket,
    }
}

func ResourceToUpdateResourceInfo(resource Resource) (updateResourceInfo UpdateResourceInfo) {
    return UpdateResourceInfo{
        ResourceId: resource.ResourceId,
        Description: resource.Description,
        OwnerTeamId: resource.OwnerTeamId,
        Visibility: resource.Visibility,
        MaxDuration: resource.MaxDuration,
        RequireManagerApproval: resource.RequireManagerApproval,
        RequireSupportTicket: resource.RequireSupportTicket,
    }
}
