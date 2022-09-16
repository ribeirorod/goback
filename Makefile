VERSION ?= 0.1.0-dev

GO_LDFLAGS ?= -s -w -X internal.Version=${VERSION}
GO_BUILDTAGS ?=

DESTDIR ?= ./bin

.PHONY: all
all: build

.PHONY: build ## Build the the compose cli-plugin
build:
	CGO_ENABLED=0 GO111MODULE=on go build -trimpath -tags "$(GO_BUILDTAGS)" -ldflags "$(GO_LDFLAGS)" -o "$(DESTDIR)/app" ./cmd
