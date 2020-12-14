package middlewares

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/aws"
	"github.com/sirupsen/logrus"
)

// Remove default security group from remote resources
type VPCDefaultSecurityGroupSanitizer struct{}

func NewVPCDefaultSecurityGroupSanitizer() VPCDefaultSecurityGroupSanitizer {
	return VPCDefaultSecurityGroupSanitizer{}
}

func (m VPCDefaultSecurityGroupSanitizer) Execute(remoteResources, resourcesFromState *[]resource.Resource) error {
	newRemoteResources := make([]resource.Resource, 0)

	for _, remoteResource := range *remoteResources {
		existInState := false

		// Ignore all resources other than security group
		if remoteResource.TerraformType() != aws.AwsSecurityGroupResourceType {
			newRemoteResources = append(newRemoteResources, remoteResource)
			continue
		}

		securityGroup, _ := remoteResource.(*aws.AwsSecurityGroup)

		if !isDefaultSecurityGroup(securityGroup) {
			newRemoteResources = append(newRemoteResources, remoteResource)
			continue
		}

		for _, stateResource := range *resourcesFromState {
			if resource.IsSameResource(remoteResource, stateResource) {
				existInState = true
				break
			}
		}

		if existInState {
			newRemoteResources = append(newRemoteResources, remoteResource)
		}

		if !existInState {
			logrus.WithFields(logrus.Fields{
				"id":   remoteResource.TerraformId(),
				"type": remoteResource.TerraformType(),
			}).Debug("Ignoring default unmanaged security group")
		}
	}

	*remoteResources = newRemoteResources

	return nil
}

// Return true if the security group is the default one
func isDefaultSecurityGroup(securityGroup *aws.AwsSecurityGroup) bool {
	return *securityGroup.Name == "default"
}