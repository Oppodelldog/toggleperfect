
build: build-plugins
	go build -o bin/toggleperfect cmd/main.go

build-plugins:
	cd internal/apps/timetoggle && make
	cd internal/apps/stocks && make
	cd internal/apps/mails && make

start: build
	nohup bin/toggleperfect > toggleperfect.log 2>&1 &

install-service:
	sudo rm /etc/systemd/system/toggleperfect.service
	sudo ln -s /home/pi/toggleperfect/toggleperfect.service /etc/systemd/system/toggleperfect.service
	sudo systemctl enable toggleperfect