.PHONY: all default clean check container-dev update-gomod

all: clean default

default:
	gofmt -s -w .
	go build -o bin/ulat main.go
	@echo "Sucesfully built ulat"

update-gomod:
	go get -t -v -d -u ./..
	go mod tidy

check: default
	go install -v -x github.com/tsenart/deadcode@latest
	go install -v -x honnef.co/go/tools/cmd/staticcheck@latest
#	go test -v ./...
	deadcode ./
	go vet ./...
	staticcheck -checks all,-ST1000 ./...

container-dev:
	rm -f artifacts
	ln -s ../apt-ostree-config artifacts
	docker build -t ulat .
	docker run -i -t --privileged --network host \
		-v $(PWD):/workspace \
		-v $(PWD)/artifacts:/artifacts \
		-v /srv:/srv ulat

clean:
	rm -rf bin
