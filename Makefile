GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GODEP=$(GOTEST) -i

GOPATH:=$(CURDIR)/_vendor:$(CURDIR)/

APPNAME = gocco
DISTDIR = dist

all: dist

clean:
	rm -rf bin
	rm -rf $(DISTDIR)

env:
	$(GO) env

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 $(GO) build 	-o bin/$(APPNAME)

dist: clean build
	mkdir $(DISTDIR)
	cp bin/$(APPNAME) $(DISTDIR)
	echo "Dist done"

test:
	go test -v $(PACKAGES)
