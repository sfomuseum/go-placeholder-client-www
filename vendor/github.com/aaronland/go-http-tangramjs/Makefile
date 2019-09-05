CWD=$(shell pwd)

go-bindata:
	mkdir -p cmd/go-bindata
	mkdir -p cmd/go-bindata-assetfs
	curl -s -o cmd/go-bindata/main.go https://raw.githubusercontent.com/whosonfirst/go-bindata/master/cmd/go-bindata/main.go
	curl -s -o cmd/go-bindata-assetfs/main.go https://raw.githubusercontent.com/whosonfirst/go-bindata-assetfs/master/cmd/go-bindata-assetfs/main.go

bake:
	@make bake-templates
	@make bake-assets

bake-templates:
	go build -o bin/go-bindata cmd/go-bindata/main.go
	rm -rf templates/html/*~
	bin/go-bindata -pkg templates -o assets/templates/html.go templates/html

bake-assets:	
	go build -o bin/go-bindata cmd/go-bindata/main.go
	go build -o bin/go-bindata-assetfs cmd/go-bindata-assetfs/main.go
	rm -f static/*~ static/css/*~ static/tangram/*~ static/javascript/*~
	@PATH=$(PATH):$(CWD)/bin bin/go-bindata-assetfs -pkg tangramjs -o assets.go static static/javascript static/css static/tangram

debug:
	go run -mod vendor examples/map/main.go -templates 'templates/html/*.html'

tangram: 
	curl -s -o static/javascript/tangram.js https://www.nextzen.org/tangram/tangram.debug.js
	curl -s -o static/javascript/tangram.min.js https://www.nextzen.org/tangram/tangram.min.js

styles: refill walkabout

refill:
	curl -s -o static/tangram/refill-style.zip https://www.nextzen.org/carto/refill-style/refill-style.zip
	curl -s -o static/tangram/refill-style-themes-label.zip https://www.nextzen.org/carto/refill-style/themes/label-10.zip

walkabout:
	curl -s -o static/tangram/walkabout-style.zip https://www.nextzen.org/carto/refill-style/walkabout-style.zip
