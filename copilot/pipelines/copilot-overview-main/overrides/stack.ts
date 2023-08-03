import * as cdk from 'aws-cdk-lib';
import * as path from 'path';
import { aws_codepipeline as cp } from 'aws-cdk-lib';
import { aws_iam as iam } from 'aws-cdk-lib';

interface TransformedStackProps extends cdk.StackProps {
    readonly appName: string;
}

export class TransformedStack extends cdk.Stack {
    public readonly template: cdk.cloudformation_include.CfnInclude;
    public readonly appName: string;

    constructor (scope: cdk.App, id: string, props: TransformedStackProps) {
        super(scope, id, props);
        this.template = new cdk.cloudformation_include.CfnInclude(this, 'Template', {
            templateFile: path.join('.build', 'in.yml'),
        });
        this.appName = props.appName;
        this.transformPipeline();
        this.transformBuildProjectPolicy();
    }
    
    // TODO: implement me.
    transformPipeline() {
        const pipeline = this.template.getResource("Pipeline") as cp.CfnPipeline;
        pipeline
        throw new Error("not implemented");
    }
    
    // TODO: implement me.
    transformBuildProjectPolicy() {
        const buildProjectPolicy = this.template.getResource("BuildProjectPolicy") as iam.CfnPolicy;
        throw new Error("not implemented");
    }
    
}