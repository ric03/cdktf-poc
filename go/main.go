package main

import (
	"cdk.tf/go/stack/generated/Mongey/kafka/acl"
	"cdk.tf/go/stack/generated/Mongey/kafka/provider"
	"cdk.tf/go/stack/generated/Mongey/kafka/topic"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	provider.NewKafkaProvider(stack, jsii.String("kafka"), &provider.KafkaProviderConfig{
		BootstrapServers:    &[]*string{jsii.String("localhost:1234"), jsii.String("kafka-broker:1234")},
		Alias:               nil,
		CaCert:              nil,
		CaCertFile:          nil,
		ClientCert:          nil,
		ClientCertFile:      nil,
		ClientKey:           nil,
		ClientKeyFile:       nil,
		ClientKeyPassphrase: nil,
		SaslMechanism:       nil,
		SaslPassword:        nil,
		SaslUsername:        nil,
		SkipTlsVerify:       nil,
		Timeout:             nil,
		TlsEnabled:          nil,
	})

	// The code that defines your stack goes here
	products := []string{"apple", "peaches", "grapes"}

	for _, product := range products {
		NewAclAndTopic(stack, product)
	}

	return stack
}

func NewAclAndTopic(stack cdktf.TerraformStack, name string) {
	topic.NewTopic(stack, jsii.String(name), &topic.TopicConfig{
		Name:              jsii.String(name),
		Partitions:        jsii.Number[int](5),
		ReplicationFactor: jsii.Number[int](3),
		Config:            &map[string]*string{"segment.ms": jsii.String("20000")},
		Id:                nil,
	})

	acl.NewAcl(stack, jsii.String(name+"acl"), &acl.AclConfig{
		AclHost:           jsii.String("*"),
		AclOperation:      jsii.String("All"),
		AclPermissionType: jsii.String("Allow"),
		AclPrincipal:      jsii.String("User:Hans"),
		ResourceName:      jsii.String(name),
		ResourceType:      jsii.String("Topic"),
		Id:                jsii.String(name + "acl-id"),
	})
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "go")

	app.Synth()
}
