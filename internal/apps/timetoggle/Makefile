.DEFAULT_GOAL := build

build-swagger:
	# for api build you will need swagger: https://goswagger.io/install.html

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

build: prepare
	go build -o ../../../bin/timetoggle.so -buildmode=plugin app/lib/main.go