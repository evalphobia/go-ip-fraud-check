.PHONY: init lint build build-macos build-linux build-all clean

GO111MODULE=on
LINT_OPT := -E gofmt \
            -E golint \
            -E govet \
            -E gosec \
            -E unused \
            -E gosimple \
            -E structcheck \
            -E varcheck \
            -E ineffassign \
            -E deadcode \
            -E typecheck \
            -E misspell \
            -E whitespace \
            -E errcheck \
            --exclude '(comment on exported (method|function|type|const|var)|should have( a package)? comment|comment should be of the form)' \
            --timeout 5m

init:
	go mod download

lint:
	@type golangci-lint > /dev/null || go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint $(LINT_OPT) run ./...

# build binary
build:
	go build -o bin/go-ip-fraud-check ./cmd

build-macos:
	@make _build BUILD_OS=darwin BUILD_ARCH=amd64

build-macos-m1:
	@make _build BUILD_OS=darwin BUILD_ARCH=arm64

build-linux:
	@make _build BUILD_OS=linux BUILD_ARCH=amd64

build-windows:
	@make _build BUILD_OS=windows BUILD_ARCH=amd64

_build:
	@mkdir -p bin/release
	$(eval BUILD_OUTPUT := go-ip-fraud-check_${BUILD_OS}_${BUILD_ARCH}${BUILD_ARM})
	GOOS=${BUILD_OS} \
	GOARCH=${BUILD_ARCH} \
	GOARM=${BUILD_ARM} \
	go build -o bin/${BUILD_OUTPUT} ./cmd
	@if [ "${USE_ARCHIVE}" = "1" ]; then \
		gzip -k -f bin/${BUILD_OUTPUT} ;\
		mv bin/${BUILD_OUTPUT}.gz bin/release/ ;\
	fi

build-all: clean
	@make build-macos build-macos-m1 build-linux build-windows USE_ARCHIVE=1

clean:
	rm -f bin/go-ip-fraud-check_*
	rm -f bin/release/*
