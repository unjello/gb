TARGET=gb

all: deps build

deps: godep
	@dep ensure

build: deps
	@go build -ldflags="-s -w" -o $(TARGET)

clean:
	@rm -rf $(TARGET)
	@rm -rf build

install: build
	@mv $(TARGET) $(GOPATH)/bin/

godep:
	@go get -u github.com/golang/dep/...
