.PHONY: all fmt tags doc

all:
	go install -v ./...
	gofmt -s -w -l .
	go install -v ./...
	smlchk -path="shanhu.io/tools"
	golint ./...
	gotags -R . > tags

rall:
	go build -a ./...

fmt:
	gofmt -s -w -l .

tags:
	gotags -R . > tags

test:
	go test ./...

testv:
	go test -v ./...

lc:
	wc -l `find . -name "*.go"`

doc:
	godoc -http=:8000

lint:
	golint ./...
