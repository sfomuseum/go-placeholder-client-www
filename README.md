# go-placeholder-client-www

![](docs/images/canada.png)

Too soon. Move along.

## Tools

### server

```
$> server -h
Usage of server:
  -host string
    	The host to listen for requests on. (default "localhost")
  -nextzen-apikey string
    	A valid Nextzen API key
  -nextzen-style-url string
    	... (default "/tangram/refill-style.zip")
  -nextzen-tile-url string
    	... (default "https://{s}.tile.nextzen.org/tilezen/vector/v1/512/all/{z}/{x}/{y}.mvt")
  -placeholder-endpoint string
    	The address of the Placeholder endpoint to query. (default "http://localhost:3000")
  -port int
    	The port to listen for requests on. (default 8080)
  -protocol string
    	The protocol for placeholder-client server to listen on. Valid protocols are: http, lambda. (default "http")
  -proxy-tiles
    	...
  -proxy-tiles-dsn string
    	... (default "cache=gocache")
  -proxy-tiles-timeout int
    	The maximum number of seconds to allow for fetching a tile from the proxy. (default 30)
  -proxy-tiles-url string
    	... (default "/tiles/")
  -static-prefix string
    	Prepend this prefix to URLs for static assets.
  -templates string
    	An optional string for local templates. This is anything that can be read by the 'templates.ParseGlob' method.
```

For example

```
go run -mod vendor cmd/server/main.go \
	-nextzen-apikey {NEXTZEN_APIKEY} \
	-proxy-tiles
	-proxy-tiles-dsn 'cache=blob blob=s3://{S3_BUCKET}?region={S3_REGION}&credentials={AWS_CREDENTIALS'
```

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

Assuming you're accessing this Lambda function through an API Gateway endpoint you will need to specify an externally hosted Tilezen style for the `PLACEHOLDER_NEXTZEN_STYLE_URL` environment variable. Something like: `https://www.nextzen.org/carto/refill-style/refill-style.zip`

The reasons for this is that by default API Gateway returns data as plain or base-64 encoded text. It is possible to configure API Gateway to return binary data but this is dependent on the client _requesting_ the data to pass along an HTTP `Accept` header. In this case that means Tangram.js needs to request style bundles as `Accept: application/zip` which it doesn't as of this writing.

Pending an update to Tangram.js there are some possible alternatives (to allow using local Tangram styles) but they are still being investigated.

#### Proxying tiles in a Lambda context

Assuming you're connecting to a Placeholder endpoint running in an ECS instance (using the `docker-placeholder` container) then you will have had to set up your Lambda function inside a VPC (see above). Once you've done that your Lambda function will no longer be able to reach the external internet without adding a NAT gateway to your VPC. If you haven't done that then, unless you've pre-seeded all the tiles you're going to request, the tile proxy layer won't work because the proxy will never be able to fetch tiles from S3 let alone Nextzen (aka: the external internet).

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
