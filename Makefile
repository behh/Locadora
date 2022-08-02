.PHONY: all
all: build
FORCE: ;

SHELL  := env LIBRARY_ENV=$(LIBRARY_ENV) $(SHELL)
LIBRARY_ENV ?= prod

BIN_DIR = $(PWD)/bin

OS ?= $(shell uname)

.PHONY: build

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: dependencies build-api zip

build-api:	
ifeq ($(OS),Linux)
	GOOS=windows GOARCH=amd64 go build -tags $(LIBRARY_ENV) -o ./bin/Locadora.exe ./main.go
else
	go build -tags $(LIBRARY_ENV) -o ./bin/Locadora.exe ./main.go
endif 

zip:
ifeq ($(OS),Linux)
	zip -j $(ZIP_DIR)/Locadora.zip ./bin/Locadora.exe 
	zip -r $(ZIP_DIR)/Locadora.zip ./web/static ./web/template
endif

test:
	go test -tags dev ./...

build-mocks:	
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen@v1.6.0	
	@mockgen -source=usecase/carro/interface.go -destination=usecase/carro/mock/filme.go -package=mock
	@mockgen -source=usecase/logger/interface.go -destination=usecase/logger/mock/logger.go -package=mock