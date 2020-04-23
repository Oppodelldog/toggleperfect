.DEFAULT_GOAL := build-all

build-swagger:
	# swagger is required to generate api code: https://goswagger.io/install.html

generate-assets: ## generates static assets
	# statics is required to generate assets: https://github.com/go-playground/statics
	# go get -u github.com/go-playground/statics
	cp swagger.yml apidocs/v1/swagger.yml
	statics -i=apidocs -o=apidocs/dist.go  -pkg=apidocs -group=Dist -ignore="(dist.go|apidocs.go|.gitignore)"

prepare: build-swagger

api-server:
	swagger generate server \
		-f ./swagger.yml \
 		-A timetoggle \
 		-s api/server \
 		-m api/model \
 		-a api \
 		--exclude-main

api-ui:
	swagger serve swagger.yml

build:
	go build -o ../../../bin/timetoggle.so -buildmode=plugin app/lib/main.go

build-all: prepare generate-assets build
