
build:
	go build -o bin/toggleperfect cmd/main.go

start: build
	nohup bin/toggleperfect > toggleperfect.log 2>&1 &
