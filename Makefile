.DEFAULT_GOAL := build-all

build-all: build build-plugins

build:
	go build -o bin/toggleperfect cmd/main.go

build-plugins:
	cd internal/apps/timetoggle && make
	cd internal/apps/stocks && make
	cd internal/apps/mails && make

start: build
	nohup bin/toggleperfect > toggleperfect.log 2>&1 &

install-service:
	sudo service toggleperfect stop || true
	sudo rm /etc/systemd/system/toggleperfect.service
	sudo ln -s /home/pi/toggleperfect/toggleperfect.service /etc/systemd/system/toggleperfect.service
	sudo systemctl enable toggleperfect
	sudo service toggleperfect start

setup: ## Install tools
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s v1.27.0
	mkdir .bin || true; mv bin/golangci-lint .bin/golangci-lint || true

lint: ## Run the linters
	golangci-lint run

fmt-all: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

fmt: ## gofmt and goimports all uncommited go files
	 git diff --name-only | grep .go | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

