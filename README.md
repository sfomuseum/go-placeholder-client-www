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

## See also

* https://github.com/sfomuseum/go-placeholder-client
* https://github.com/sfomuseum/docker-placeholder
* https://docs.aws.amazon.com/lambda/latest/dg/vpc.html