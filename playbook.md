## build-tags

Intro on build
https://pkg.go.dev/net/http/pprof
pprof - security
https://go.dev/doc/go1.17#go-command (go:build)
go build/run/test
explicit tags
By platform
    go env GOOS GOARCH
    go tool dist list
    curl http://localhost:8080/by
https://github.com/etcd-io/bbolt

## cgo

main.go
Dockerfile
make build
make run
docker run --rm -it app/httpd /bin/sh
ls
ldd httpd (libc)
/lib64/ld-linux-x86-64.so.2: No such file or directory

https://pkg.go.dev/net#hdr-Name_Resolution

In Dockerfile
ENV CGO_ENABLED=0 (before RUN go build)
make build
make run

## cross-compile

main.go
go build
GOOS=linux go build (time)
https://goreleaser.com/
.goreleser.yaml
git tag v1.2.3
make
ls dist
tar tzf dist/quotes_1.2.3-SNAPSHOT-f5f461e_linux_arm64.tar.gz

## embed

db.go
    _ embed
main.go
    embed.FS
    redirect

start
    http://localhost:8080/docs
    make

## generate
quotes.txt
gen_quotes.go
gen.go
go generate
quotes.go
    DO NOT EDIT

## inject
main
Makefile
make
./inject -version
version=1.2.3 make
./inject -version

## race

main.go
go run .
curl http://localhost:8080
curl http://localhost:8080
https://github.com/rakyll/hey
hey -n 1000 http://localhost:8080
curl http://localhost:8080

go run -race .
hey -n 1000 http://localhost:8080

Fix with sync atomic
https://go.dev/doc/articles/race_detector#Runtime_Overheads
