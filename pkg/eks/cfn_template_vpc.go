package eks

import (
	"net"

	"github.com/awslabs/goformation/cloudformation"
)

func newParameter(t *cloudformation.Template, name, valueType, defaultValue string) interface{} {
	p := map[string]string{"Type": valueType}
	if defaultValue != "" {
		p["Default"] = defaultValue
	}
	t.Parameters[name] = p
	return newRef(name)
}

func newStringParameter(t *cloudformation.Template, name, defaultValue string) interface{} {
	return newParameter(t, name, "String", defaultValue)
}

func newRef(refName string) interface{} {
	return map[string]string{"Ref": refName}
}

func newSub(sub string) interface{} {
	return map[string]string{"Sub": sub}
}

func addResourcesForVPC(t *cloudformation.Template, stackName string, globalCIDR *net.IPNet, subnets map[string]*net.IPNet) {
	t.Resources["VPC"] = &cloudformation.UntypedAWSEC2VPC{
		CidrBlock:          globalCIDR.String(),
		EnableDnsSupport:   true,
		EnableDnsHostnames: true,
		Tags: []cloudformation.Tag{{
			Key:   "Name",
			Value: stackName + ".VPC",
		}},
	}
	refVPC := newRef("VPC")

	t.Resources["InternetGateway"] = &cloudformation.AWSEC2InternetGateway{}
	refIG := newRef("InternetGateway")
	t.Resources["VPCGatewayAttachment"] = &cloudformation.UntypedAWSEC2VPCGatewayAttachment{
		InternetGatewayId: refIG,
		VpcId:             refVPC,
	}

	t.Resources["RouteTable"] = &cloudformation.UntypedAWSEC2RouteTable{
		VpcId: refVPC,
		Tags: []cloudformation.Tag{
			{Key: "Name", Value: "Public Subnets"},
			{Key: "Network", Value: "Public"},
		},
	}
	refRT := newRef("RouteTable")

	t.Resources["Route"] = &cloudformation.UntypedAWSEC2Route{
		RouteTableId:         refRT,
		DestinationCidrBlock: "0.0.0.0/0",
		GatewayId:            newRef("InternetGateway"),
	}

	for az, subnet := range subnets {
		name := "Subnet_" + az
		t.Resources[name] = &cloudformation.UntypedAWSEC2Subnet{
			AvailabilityZone: az,
			CidrBlock:        subnet.String(),
			VpcId:            refVPC,
			Tags: []cloudformation.Tag{{
				Key:   "Name",
				Value: stackName + "." + name,
			}},
		}
		t.Resources["RouteTableAssociation_"+az] = &cloudformation.UntypedAWSEC2SubnetRouteTableAssociation{
			SubnetId:     newRef(name),
			RouteTableId: refRT,
		}
	}

	t.Resources["ControlPlaneSecurityGroup"] = &cloudformation.UntypedAWSEC2SecurityGroup{
		GroupDescription: "Cluster communication with worker nodes",
		VpcId:            refVPC,
	}
}
