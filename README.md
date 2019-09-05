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

| Key | Value | Required |
| --- | --- | --- |
| PLACEHOLDER_PROTOCOL | `lambda` | yes |
| PLACEHOLDER_PLACEHOLDER_ENDPOINT | string | yes |
| PLACEHOLDER_NEXZEN_APIKEY | string | yes | 
| PLACEHOLDER_STATIC_PREFIX | string | no |
| PLACEHOLDER_NEXTZEN_STYLE_URL | string | no |
| PLACEHOLDER_PROXY_TILES | boolean | no |
| PLACEHOLDER_PROXY_TILES_DSN | string | no |

#### Nextzen style URLs in a Lambda context

#### Proxying tiles in a Lambda context

Assuming you're connecting to a Placeholder endpoint running in an ECS instance (using the `docker-placeholder` container) then you will have had to set up your Lambda function inside a VPC (see above). Once you've done that your Lambda function will no longer be able to reach the external internet without adding a NAT gateway to your VPC. If you haven't done that then, unless you've pre-seeded all the tiles you're going to request, the tile proxy layer won't work because the proxy will never be able to fetch tiles from Nextzen (aka: the external internet).

### ECS

There is also a handy [Dockerfile](Dockerfile) for running things in a container. The following assumes that you want to run the `placeholder-client-www` server in AWS as a Fargate deployment with application load balancer (ALB) in front of it.

```
/usr/local/bin/placeholder-client-www,-placeholder-endpoint,{PLACEHOLDER_ENDPOINT},-host,0.0.0.0,-nextzen-apikey,{NEXTZEN_APIKEY}
```

Note the `-host 0.0.0.0` part. This is important. Without it the health checks performed by ALB will always fail.

## See also

* https://github.com/sfomuseum/go-placeholder-client
* https://github.com/sfomuseum/docker-placeholder
* https://github.com/pelias/placeholder/
