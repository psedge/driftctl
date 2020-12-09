// GENERATED, DO NOT EDIT THIS FILE
package aws

const AwsDbSubnetGroupResourceType = "aws_db_subnet_group"

type AwsDbSubnetGroup struct {
	Arn         *string           `cty:"arn" computed:"true"`
	Description *string           `cty:"description"`
	Id          string            `cty:"id" computed:"true"`
	Name        *string           `cty:"name" computed:"true"`
	NamePrefix  *string           `cty:"name_prefix" computed:"true"`
	SubnetIds   []string          `cty:"subnet_ids"`
	Tags        map[string]string `cty:"tags"`
}

func (r *AwsDbSubnetGroup) TerraformId() string {
	return r.Id
}

func (r *AwsDbSubnetGroup) TerraformType() string {
	return AwsDbSubnetGroupResourceType
}
