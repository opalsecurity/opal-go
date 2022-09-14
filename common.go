/*
   Utils to facilitate easier usage of the SDK.
*/
package opal

func GroupToUpdateGroupInfo(group Group) (updateGroupInfo UpdateGroupInfo) {
	return UpdateGroupInfo{
		GroupId:                group.GroupId,
		Description:            group.Description,
		AdminOwnerId:           group.AdminOwnerId,
		MaxDuration:            group.MaxDuration,
		RequireManagerApproval: group.RequireManagerApproval,
		RequireSupportTicket:   group.RequireSupportTicket,
		Name:                   group.Name,
		FolderId:               group.FolderId,
		RequireMfaToApprove:    group.RequireMfaToApprove,
		AutoApproval:           group.AutoApproval,
		RequestTemplateId:      group.RequestTemplateId,
	}
}

func ResourceToUpdateResourceInfo(resource Resource) (updateResourceInfo UpdateResourceInfo) {
	return UpdateResourceInfo{
		ResourceId:             resource.ResourceId,
		Description:            resource.Description,
		AdminOwnerId:           resource.AdminOwnerId,
		MaxDuration:            resource.MaxDuration,
		RequireManagerApproval: resource.RequireManagerApproval,
		RequireSupportTicket:   resource.RequireSupportTicket,
		Name:                   resource.Name,
		FolderId:               resource.FolderId,
		RequireMfaToApprove:    resource.RequireMfaToApprove,
		AutoApproval:           resource.AutoApproval,
		RequestTemplateId:      resource.RequestTemplateId,
	}
}
