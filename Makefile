OUT=./dist/aiservice
MAIN=./cmd/aiservice

all: test build
build:
	CGO_ENABLED=0 GOOS=linux go build \
		-a -installsuffix cgo -ldflags '-extldflags "-static"' \
		-o $(OUT) -v $(MAIN)
test:
	CGO_ENABLED=0 go test ./... -count=1 -covermode=atomic -v -tags=unit
run:
	CGO_ENABLED=0 go run $(MAIN)
