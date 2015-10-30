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

stayall:
	STAYPATH=`pwd`/stay-tests stayall

lint:
	golint ./...

symdep:
	symdep e8vm.io/tools/dagvis
	symdep e8vm.io/tools/godep
	symdep e8vm.io/tools/goview
	symdep e8vm.io/tools/goload
	symdep e8vm.io/tools/e8doc
	symdep e8vm.io/tools/e8dag

check: fmt all lint symdep
