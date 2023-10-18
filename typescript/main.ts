import {Construct} from "constructs";
import {App, TerraformStack} from "cdktf";
import {Topic} from "./.gen/providers/kafka/topic";
import {Acl} from "./.gen/providers/kafka/acl";
import {KafkaProvider} from "./.gen/providers/kafka/provider";

class MyStack extends TerraformStack {
    constructor(scope: Construct, id: string) {
        super(scope, id);

        new KafkaProvider(this, 'my-kakfa-provider-id', {
            bootstrapServers: [],
        })

        // define resources here
        new Topic(this, 'my-topic-id', {
            name: 'temperature',
            partitions: 5,
            replicationFactor: 3,
            config: {"segment.ms": "20000"},
        })

        new Acl(this, 'my-acl-id', {
            resourceName: 'temperature_acl',
            resourceType: 'Topic',
            aclPrincipal: 'User:Hans',
            aclHost: '*',
            aclOperation: 'All',
            aclPermissionType: 'Allow',
        })
    }
}

const app = new App();
new MyStack(app, "cdktf-poc");
app.synth();
