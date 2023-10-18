# Steps to reproduce

## Requirements

NodeJS 16+, npm 16+

    npm install

## Generate the providers

    cdktf get

The HCL providers are converted into typescript code and saved to `./.gen/providers`.
Those are required for the actual code. If not present they can be added by running a cdktf cli command.

## Write Code

Edit `main.ts`.
Add your terraform provider and all the components.

Fix the import path of generated providers added by IntelliJ.
Add a leading `./` to the import path, eg:

    - import {Acl} from ".gen/providers/kafka/acl";
    + import {Acl} from "./.gen/providers/kafka/acl";

## Compile Typescript, to check for errors

    npm run build

## Generate the files

    cdktf synth

This will create `cdktf.out/stacks/cdktf-poc/cdk.tf.json`.
The generated json file is rather verbose, and not an actual hcl/tf file.

## Run terraform

    cd cdktf.out/stacks/cdktf-poc/
    terraform plan
