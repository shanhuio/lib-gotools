.PHONY: all fmt tags doc

all:
	go install -v ./...
	gofmt -s -w -l .
	go install -v ./...
	e8chk -path="e8vm.io/tools"
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

asmt:
	make -C asm/tests --no-print-directory

lint:
	golint ./...

check: fmt all lint
