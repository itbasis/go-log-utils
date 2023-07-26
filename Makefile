go-dependencies:
	# https://asdf-vm.com/
	asdf install golang
	# https://github.com/securego/gosec
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	#
	go get -u -t -v ./... || :

go-all: go-dependencies
	go generate ./...
	go mod tidy || :
	go test -v -count=1 ./...
	gosec ./...
