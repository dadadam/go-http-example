BINARY_NAME=example-app
OS_NAME := $(shell uname -s | tr A-Z a-z)


ifeq ($(OS),Windows_NT)
    CCFLAGS += -D WIN32
	OS_NAME := windows
    ifeq ($(PROCESSOR_ARCHITEW6432),AMD64)
        CCFLAGS += -D AMD64
    else
        ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
            CCFLAGS += -D AMD64

        endif
        ifeq ($(PROCESSOR_ARCHITECTURE),x86)
            CCFLAGS += -D IA32
        endif
    endif
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        CCFLAGS += -D LINUX
		OS_NAME := linux
    endif
    ifeq ($(UNAME_S),Darwin)
        CCFLAGS += -D OSX
		OS_NAME := darwin
    endif
    UNAME_P := $(shell uname -p)
    ifeq ($(UNAME_P),x86_64)
        CCFLAGS += -D AMD64
    endif
    ifneq ($(filter %86,$(UNAME_P)),)
        CCFLAGS += -D IA32
    endif
    ifneq ($(filter arm%,$(UNAME_P)),)
        CCFLAGS += -D ARM
    endif
endif


build:
	GOARCH=amd64 GOOS=${OS_NAME} go build -o ${BINARY_NAME}-${OS_NAME} cmd/example-app/main.go

run:
	go run ./cmd/example-app/main.go

clean:
	go clean
	rm ${BINARY_NAME}-${OS_NAME}