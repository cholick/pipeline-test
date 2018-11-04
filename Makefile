default: all

GO_PACKAGES = $$(go list ./... | grep -v vendor)
GO_FILES = $$(find . -name "*.go" | grep -v vendor | uniq)

linux:
	GOOS=linux GOARCH=amd64 go build -o test.linux ./main.go

mac:
	GOOS=darwin GOARCH=amd64 go build -o test.darwin ./main.go

unit-test:
	@go test ${GO_PACKAGES}

vet:
	@go vet ${GO_PACKAGES}

fmt:
	gofmt -s -l -w $(GO_FILES)

test: fmt unit-test vet

build: linux mac

all: test build
