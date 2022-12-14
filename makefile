.DEFAULT_GOAL=build

.PHONY: all clean fmt lint vet build ls_bins ls_srcs

SRCS := $(wildcard *.go)
BINS := $(SRCS:%.c=% %.exe)

all: fmt vet lint build

build: vet
	@echo "Go building (${SRCS})"
	go build ${SRCS}

vet: fmt
	go vet ./...

fmt:
	go fmt ./...

lint: fmt
	golint ./...

clean:
	@echo "Cleaning (${BINS})"
	rm -rvf ${BINS}

ls_bins: 
	@echo "${BINS}"

ls_srcs: 
	@echo "${SRCS}"