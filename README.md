# go-placeholder-client-www

Too soon. Move along.

## AWS

### Lambda

This assumes you are running `cmd/server/main.go` as a Lambda function and connecting to an instance of Placeholder running inside ECS Fargate instance.

#### Roles

You will need a `IAM` role with the following (AWS managed) policies:

* `AWSLambdaBasicExecutionRole`
* `AWSLambdaVPCAccessExecutionRole`

#### VPC

_Please write me._

#### Environment variables

| Key | Value | Notes |
| --- | --- | --- |
| PLACEHOLDER_PLACEHOLDER_ENDPOINT | ... |
| PLACEHOLDER_PROTOCOL | lambda |
| PLACEHOLDER_BOOTSTRAP_PREFIX | ... |
| PLACEHOLDER_NEXTZENJS_PREFIX | ... |

The `PLACEHOLDER_BOOTSTRAP_PREFIX` and `PLACEHOLDER_NEXTZENJS_PREFIX` environment variables are only necessary if access to the Lambda function is being served through an API Gateway endpoint. The value of both variables should the name of the deployment prefixed with a leading "/", for example `/placeholder`. This will cause the `go-http-bootstrap` and `go-http-nextzenjs` HTML rewrite handlers to prefix URLs to their respective resources.

## See also

* https://github.com/sfomuseum/go-placeholder-client
* https://github.com/sfomuseum/docker-placeholder
* https://docs.aws.amazon.com/lambda/latest/dg/vpc.html