all: build

build :
	go get -v
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -installsuffix netgo -installsuffix cgo -v -ldflags "-s -w -X main.pBuildTime=`date -u +%Y%m%d.%H%M%S`" .

test : build
	go test -v
	golint > lint.txt
	go tool vet -v . > vet.txt
	gocov test | gocov-xml > coverage.xml
	go test -bench=. -test.benchmem -v | gobench2plot > benchmarks.xml

prepare : build
	cp freeformatter Docker/freeformatter

docker-devel : prepare
	-@sudo docker rmi -f bayugyug/freeformatter 2>/dev/null || true
	cd Docker && sudo docker build --no-cache --rm -t bayugyug/freeformatter .

docker-wheezy: prepare
	cd Docker && sudo docker build --no-cache --rm -t bayugyug/freeformatter  -f  wheezy/Dockerfile .

docker-scratch: prepare
	cd Docker && sudo docker build --no-cache --rm -t bayugyug/freeformatter:scratch -f  scratch/Dockerfile .

docker-busybox: prepare
	cd Docker && sudo docker build --no-cache --rm -t bayugyug/freeformatter:busybox -f  busybox/Dockerfile .

docker-alpine: prepare
	cd Docker && sudo docker build --no-cache --rm -t bayugyug/freeformatter:alpine  -f  alpine/Dockerfile .

clean:
	rm -f freeformatter Docker/freeformatter
	rm -f benchmarks.xml coverage.xml vet.txt lint.txt

re: clean all

