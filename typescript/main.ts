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

        const products = ["apple", "peaches", "grapes"];

        // define resources here
        products.forEach(product => this.AclAndGroup(product))
    }

    private AclAndGroup(name: string) {
        new Acl(this, `${name}-acl-id`, {
            resourceName: `${name}-acl`,
            resourceType: 'Topic',
            aclPrincipal: 'User:Hans',
            aclHost: '*',
            aclOperation: 'All',
            aclPermissionType: 'Allow',
        })

        new Topic(this, `${name}-topic-id`, {
            name: name,
            partitions: 5,
            replicationFactor: 3,
            config: {"segment.ms": "20000"},
        })
    }
}

const app = new App();
new MyStack(app, "cdktf-poc");
app.synth();
