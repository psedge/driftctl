// GENERATED, DO NOT EDIT THIS FILE
package aws

const AwsIamRolePolicyResourceType = "aws_iam_role_policy"

type AwsIamRolePolicy struct {
	Id         string  `cty:"id" computed:"true"`
	Name       *string `cty:"name" computed:"true"`
	NamePrefix *string `cty:"name_prefix"`
	Policy     *string `cty:"policy" jsonstring:"true"`
	Role       *string `cty:"role"`
}

func (r *AwsIamRolePolicy) TerraformId() string {
	return r.Id
}

func (r *AwsIamRolePolicy) TerraformType() string {
	return AwsIamRolePolicyResourceType
}
