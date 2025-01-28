// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newResourceGroupPoliciesExclusive,
			TypeName: "aws_iam_group_policies_exclusive",
			Name:     "Group Policies Exclusive",
		},
		{
			Factory:  newResourceGroupPolicyAttachmentsExclusive,
			TypeName: "aws_iam_group_policy_attachments_exclusive",
			Name:     "Group Policy Attachments Exclusive",
		},
		{
			Factory:  newOrganizationsFeaturesResource,
			TypeName: "aws_iam_organizations_features",
			Name:     "Organizations Features",
		},
		{
			Factory:  newResourceRolePoliciesExclusive,
			TypeName: "aws_iam_role_policies_exclusive",
			Name:     "Role Policies Exclusive",
		},
		{
			Factory:  newResourceRolePolicyAttachmentsExclusive,
			TypeName: "aws_iam_role_policy_attachments_exclusive",
			Name:     "Role Policy Attachments Exclusive",
		},
		{
			Factory:  newResourceUserPoliciesExclusive,
			TypeName: "aws_iam_user_policies_exclusive",
			Name:     "User Policies Exclusive",
		},
		{
			Factory:  newResourceUserPolicyAttachmentsExclusive,
			TypeName: "aws_iam_user_policy_attachments_exclusive",
			Name:     "User Policy Attachments Exclusive",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceAccessKeys,
			TypeName: "aws_iam_access_keys",
			Name:     "Access Keys",
		},
		{
			Factory:  dataSourceAccountAlias,
			TypeName: "aws_iam_account_alias",
			Name:     "Account Alias",
		},
		{
			Factory:  dataSourceGroup,
			TypeName: "aws_iam_group",
			Name:     "Group",
		},
		{
			Factory:  dataSourceInstanceProfile,
			TypeName: "aws_iam_instance_profile",
			Name:     "Instance Profile",
		},
		{
			Factory:  dataSourceInstanceProfiles,
			TypeName: "aws_iam_instance_profiles",
			Name:     "Instance Profiles",
		},
		{
			Factory:  dataSourceOpenIDConnectProvider,
			TypeName: "aws_iam_openid_connect_provider",
			Name:     "OIDC Provider",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourcePolicy,
			TypeName: "aws_iam_policy",
			Name:     "Policy",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourcePolicyDocument,
			TypeName: "aws_iam_policy_document",
			Name:     "Policy Document",
		},
		{
			Factory:  dataSourcePrincipalPolicySimulation,
			TypeName: "aws_iam_principal_policy_simulation",
			Name:     "Principal Policy Simulation",
		},
		{
			Factory:  dataSourceRole,
			TypeName: "aws_iam_role",
			Name:     "Role",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceRoles,
			TypeName: "aws_iam_roles",
			Name:     "Roles",
		},
		{
			Factory:  dataSourceSAMLProvider,
			TypeName: "aws_iam_saml_provider",
			Name:     "SAML Provider",
		},
		{
			Factory:  dataSourceServerCertificate,
			TypeName: "aws_iam_server_certificate",
			Name:     "Server Certificate",
		},
		{
			Factory:  dataSourceSessionContext,
			TypeName: "aws_iam_session_context",
			Name:     "Session Context",
		},
		{
			Factory:  dataSourceUser,
			TypeName: "aws_iam_user",
			Name:     "User",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceUserSSHKey,
			TypeName: "aws_iam_user_ssh_key",
			Name:     "User SSH Key",
		},
		{
			Factory:  dataSourceUsers,
			TypeName: "aws_iam_users",
			Name:     "Users",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAccessKey,
			TypeName: "aws_iam_access_key",
			Name:     "Access Key",
		},
		{
			Factory:  resourceAccountAlias,
			TypeName: "aws_iam_account_alias",
			Name:     "Account Alias",
		},
		{
			Factory:  resourceAccountPasswordPolicy,
			TypeName: "aws_iam_account_password_policy",
			Name:     "Account Password Policy",
		},
		{
			Factory:  resourceGroup,
			TypeName: "aws_iam_group",
			Name:     "Group",
		},
		{
			Factory:  resourceGroupMembership,
			TypeName: "aws_iam_group_membership",
			Name:     "Group Membership",
		},
		{
			Factory:  resourceGroupPolicy,
			TypeName: "aws_iam_group_policy",
			Name:     "Group Policy",
		},
		{
			Factory:  resourceGroupPolicyAttachment,
			TypeName: "aws_iam_group_policy_attachment",
			Name:     "Group Policy Attachment",
		},
		{
			Factory:  resourceInstanceProfile,
			TypeName: "aws_iam_instance_profile",
			Name:     "Instance Profile",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "InstanceProfile",
			},
		},
		{
			Factory:  resourceOpenIDConnectProvider,
			TypeName: "aws_iam_openid_connect_provider",
			Name:     "OIDC Provider",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
				ResourceType:        "OIDCProvider",
			},
		},
		{
			Factory:  resourcePolicy,
			TypeName: "aws_iam_policy",
			Name:     "Policy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
				ResourceType:        "Policy",
			},
		},
		{
			Factory:  resourcePolicyAttachment,
			TypeName: "aws_iam_policy_attachment",
			Name:     "Policy Attachment",
		},
		{
			Factory:  resourceRole,
			TypeName: "aws_iam_role",
			Name:     "Role",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrName,
				ResourceType:        "Role",
			},
		},
		{
			Factory:  resourceRolePolicy,
			TypeName: "aws_iam_role_policy",
			Name:     "Role Policy",
		},
		{
			Factory:  resourceRolePolicyAttachment,
			TypeName: "aws_iam_role_policy_attachment",
			Name:     "Role Policy Attachment",
		},
		{
			Factory:  resourceSAMLProvider,
			TypeName: "aws_iam_saml_provider",
			Name:     "SAML Provider",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "SAMLProvider",
			},
		},
		{
			Factory:  resourceSecurityTokenServicePreferences,
			TypeName: "aws_iam_security_token_service_preferences",
			Name:     "Security Token Service Preferences",
		},
		{
			Factory:  resourceServerCertificate,
			TypeName: "aws_iam_server_certificate",
			Name:     "Server Certificate",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrName,
				ResourceType:        "ServerCertificate",
			},
		},
		{
			Factory:  resourceServiceLinkedRole,
			TypeName: "aws_iam_service_linked_role",
			Name:     "Service Linked Role",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "ServiceLinkedRole",
			},
		},
		{
			Factory:  resourceServiceSpecificCredential,
			TypeName: "aws_iam_service_specific_credential",
			Name:     "Service Specific Credential",
		},
		{
			Factory:  resourceSigningCertificate,
			TypeName: "aws_iam_signing_certificate",
			Name:     "Signing Certificate",
		},
		{
			Factory:  resourceUser,
			TypeName: "aws_iam_user",
			Name:     "User",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrName,
				ResourceType:        "User",
			},
		},
		{
			Factory:  resourceUserGroupMembership,
			TypeName: "aws_iam_user_group_membership",
			Name:     "User Group Membership",
		},
		{
			Factory:  resourceUserLoginProfile,
			TypeName: "aws_iam_user_login_profile",
			Name:     "User Login Profile",
		},
		{
			Factory:  resourceUserPolicy,
			TypeName: "aws_iam_user_policy",
			Name:     "User Policy",
		},
		{
			Factory:  resourceUserPolicyAttachment,
			TypeName: "aws_iam_user_policy_attachment",
			Name:     "User Policy Attachment",
		},
		{
			Factory:  resourceUserSSHKey,
			TypeName: "aws_iam_user_ssh_key",
			Name:     "User SSH Key",
		},
		{
			Factory:  resourceVirtualMFADevice,
			TypeName: "aws_iam_virtual_mfa_device",
			Name:     "Virtual MFA Device",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "VirtualMFADevice",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.IAM
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*iam.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*iam.Options){
		iam.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return iam.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*iam.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*iam.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *iam.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*iam.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
