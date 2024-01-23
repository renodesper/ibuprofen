package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws-apigateway/sdk/v2/go/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// An execution role to use for the Lambda function
		policy, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				{
					"Action": "sts:AssumeRole",
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"Service": "lambda.amazonaws.com",
					},
				},
			},
		})
		if err != nil {
			return err
		}

		role, err := iam.NewRole(ctx, "ibuprofen_role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(policy),
			ManagedPolicyArns: pulumi.StringArray{
				iam.ManagedPolicyAWSLambdaBasicExecutionRole,
			},
		})
		if err != nil {
			return err
		}

		// A Lambda function to invoke
		fn, err := lambda.NewFunction(ctx, "ibuprofen_fn", &lambda.FunctionArgs{
			Role:    role.Arn,
			Runtime: pulumi.String("go1.x"),
			Handler: pulumi.String("main"),
			Code:    pulumi.NewFileArchive("./function/app.zip"),
		})
		if err != nil {
			return err
		}

		// A REST API to route requests to HTML content and the Lambda function
		method := apigateway.MethodGET
		api, err := apigateway.NewRestAPI(ctx, "ibuprofen", &apigateway.RestAPIArgs{
			Routes: []apigateway.RouteArgs{
				{Path: "/hello", Method: &method, EventHandler: fn},
			},
		})
		if err != nil {
			return err
		}

		// The URL at which the REST API will be served
		ctx.Export("url", api.Url)
		return nil
	})
}
