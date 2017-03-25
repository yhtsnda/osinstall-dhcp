build:
	go build

run:
	go run

pack:
	go build
	rm -rf ./target
	mkdir ./target
	cp -R ./config ./target
	cp -R ./data ./target
	cp ./ocb-dhcp ./target

pack_linux:
	env GOARCH=amd64  GOOS=linux go build
	rm -rf ./target
	mkdir ./target
	cp -R ./config ./target
	cp -R ./data ./target
	cp ./ocb-dhcp ./target

pack_mac:
		env GOARCH=amd64  GOOS=darwin go build
		rm -rf ./target
		mkdir ./target
		cp -R ./config ./target
		cp -R ./data ./target
		cp ./ocb-dhcp ./target
