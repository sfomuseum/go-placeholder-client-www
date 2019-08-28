# go-placeholder-client-www

![](docs/images/canada.png)

Too soon. Move along.

## AWS

![](docs/images/arch.jpg)

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
| PLACEHOLDER_PROTOCOL | `lambda` |
| PLACEHOLDER_PLACEHOLDER_ENDPOINT | ... |
| PLACEHOLDER_STATIC_PREFIX | ... |
| PLACEHOLDER_NEXZEN_APIKEY | ... |

_Please write about `PLACEHOLDER_STATIC_PREFIX`._

## See also

* https://github.com/sfomuseum/go-placeholder-client
* https://github.com/sfomuseum/docker-placeholder
* https://github.com/pelias/placeholder/
