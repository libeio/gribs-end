
GO ?= go
GO-BUILD = go build
GO-GET = go get -u
GO-FMT = go fmt
GO-MOD = go mod
GO-TEST = go test -v

RM ?= rm -rf
MV ?= mv -f
CP ?= cp -f

GOFILES := $(shell find . -name "*.go")
EXEFILES := $(shell find . -type f -perm /u=x,g=x,o=x)
PWDGOFILES := $(shell find . -maxdepth 1 -name "*.go")
TESTFILES := $(shell find . -name "*_test.go")

MODULENAME = grids.com
