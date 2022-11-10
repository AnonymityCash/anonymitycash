ifndef GOOS
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
	GOOS := darwin
else ifeq ($(UNAME_S),Linux)
	GOOS := linux
else
$(error "$$GOOS is not defined. If you are using Windows, try to re-make using 'GOOS=windows make ...' ")
endif
endif

PACKAGES    := $(shell go list ./... | grep -v '/vendor/' | grep -v '/crypto/ed25519/chainkd' | grep -v '/mining/tensority')
PACKAGES += 'github.com/anonymitycash/anonymitycash/mining/tensority/go_algorithm'

BUILD_FLAGS := -ldflags "-X github.com/anonymitycash/anonymitycash/version.GitCommit=`git rev-parse HEAD`"

MINER_BINARY32 := miner-$(GOOS)_386
MINER_BINARY64 := miner-$(GOOS)_amd64

ANONYMITYCASHD_BINARY32 := anonymitycashd-$(GOOS)_386
ANONYMITYCASHD_BINARY64 := anonymitycashd-$(GOOS)_amd64

ANONYMITYCASHCLI_BINARY32 := anonymitycashcli-$(GOOS)_386
ANONYMITYCASHCLI_BINARY64 := anonymitycashcli-$(GOOS)_amd64

VERSION := $(shell awk -F= '/Version =/ {print $$2}' version/version.go | tr -d "\" ")

MINER_RELEASE32 := miner-$(VERSION)-$(GOOS)_386
MINER_RELEASE64 := miner-$(VERSION)-$(GOOS)_amd64

ANONYMITYCASHD_RELEASE32 := anonymitycashd-$(VERSION)-$(GOOS)_386
ANONYMITYCASHD_RELEASE64 := anonymitycashd-$(VERSION)-$(GOOS)_amd64

ANONYMITYCASHCLI_RELEASE32 := anonymitycashcli-$(VERSION)-$(GOOS)_386
ANONYMITYCASHCLI_RELEASE64 := anonymitycashcli-$(VERSION)-$(GOOS)_amd64

ANONYMITYCASH_RELEASE32 := anonymitycash-$(VERSION)-$(GOOS)_386
ANONYMITYCASH_RELEASE64 := anonymitycash-$(VERSION)-$(GOOS)_amd64

all: test target release-all install

anonymitycashd:
	@echo "Building anonymitycashd to cmd/anonymitycashd/anonymitycashd"
	@go build $(BUILD_FLAGS) -o cmd/anonymitycashd/anonymitycashd cmd/anonymitycashd/main.go

anonymitycashd-simd:
	@echo "Building SIMD version anonymitycashd to cmd/anonymitycashd/anonymitycashd"
	@cd mining/tensority/cgo_algorithm/lib/ && make
	@go build -tags="simd" $(BUILD_FLAGS) -o cmd/anonymitycashd/anonymitycashd cmd/anonymitycashd/main.go

anonymitycashcli:
	@echo "Building anonymitycashcli to cmd/anonymitycashcli/anonymitycashcli"
	@go build $(BUILD_FLAGS) -o cmd/anonymitycashcli/anonymitycashcli cmd/anonymitycashcli/main.go

install:
	@echo "Installing anonymitycashd and anonymitycashcli to $(GOPATH)/bin"
	@go install ./cmd/anonymitycashd
	@go install ./cmd/anonymitycashcli

target:
	mkdir -p $@

binary: target/$(ANONYMITYCASHD_BINARY32) target/$(ANONYMITYCASHD_BINARY64) target/$(ANONYMITYCASHCLI_BINARY32) target/$(ANONYMITYCASHCLI_BINARY64) target/$(MINER_BINARY32) target/$(MINER_BINARY64)

ifeq ($(GOOS),windows)
release: binary
	cd target && cp -f $(MINER_BINARY32) $(MINER_BINARY32).exe
	cd target && cp -f $(ANONYMITYCASHD_BINARY32) $(ANONYMITYCASHD_BINARY32).exe
	cd target && cp -f $(ANONYMITYCASHCLI_BINARY32) $(ANONYMITYCASHCLI_BINARY32).exe
	cd target && md5sum $(MINER_BINARY32).exe $(ANONYMITYCASHD_BINARY32).exe $(ANONYMITYCASHCLI_BINARY32).exe >$(ANONYMITYCASH_RELEASE32).md5
	cd target && zip $(ANONYMITYCASH_RELEASE32).zip $(MINER_BINARY32).exe $(ANONYMITYCASHD_BINARY32).exe $(ANONYMITYCASHCLI_BINARY32).exe $(ANONYMITYCASH_RELEASE32).md5
	cd target && rm -f $(MINER_BINARY32) $(ANONYMITYCASHD_BINARY32) $(ANONYMITYCASHCLI_BINARY32) $(MINER_BINARY32).exe $(ANONYMITYCASHD_BINARY32).exe $(ANONYMITYCASHCLI_BINARY32).exe $(ANONYMITYCASH_RELEASE32).md5
	cd target && cp -f $(MINER_BINARY64) $(MINER_BINARY64).exe
	cd target && cp -f $(ANONYMITYCASHD_BINARY64) $(ANONYMITYCASHD_BINARY64).exe
	cd target && cp -f $(ANONYMITYCASHCLI_BINARY64) $(ANONYMITYCASHCLI_BINARY64).exe
	cd target && md5sum $(MINER_BINARY64).exe $(ANONYMITYCASHD_BINARY64).exe $(ANONYMITYCASHCLI_BINARY64).exe >$(ANONYMITYCASH_RELEASE64).md5
	cd target && zip $(ANONYMITYCASH_RELEASE64).zip $(MINER_BINARY64).exe $(ANONYMITYCASHD_BINARY64).exe $(ANONYMITYCASHCLI_BINARY64).exe $(ANONYMITYCASH_RELEASE64).md5
	cd target && rm -f $(MINER_BINARY64) $(ANONYMITYCASHD_BINARY64) $(ANONYMITYCASHCLI_BINARY64) $(MINER_BINARY64).exe $(ANONYMITYCASHD_BINARY64).exe $(ANONYMITYCASHCLI_BINARY64).exe $(ANONYMITYCASH_RELEASE64).md5
else
release: binary
	cd target && md5sum $(MINER_BINARY32) $(ANONYMITYCASHD_BINARY32) $(ANONYMITYCASHCLI_BINARY32) >$(ANONYMITYCASH_RELEASE32).md5
	cd target && tar -czf $(ANONYMITYCASH_RELEASE32).tgz $(MINER_BINARY32) $(ANONYMITYCASHD_BINARY32) $(ANONYMITYCASHCLI_BINARY32) $(ANONYMITYCASH_RELEASE32).md5
	cd target && rm -f $(MINER_BINARY32) $(ANONYMITYCASHD_BINARY32) $(ANONYMITYCASHCLI_BINARY32) $(ANONYMITYCASH_RELEASE32).md5
	cd target && md5sum $(MINER_BINARY64) $(ANONYMITYCASHD_BINARY64) $(ANONYMITYCASHCLI_BINARY64) >$(ANONYMITYCASH_RELEASE64).md5
	cd target && tar -czf $(ANONYMITYCASH_RELEASE64).tgz $(MINER_BINARY64) $(ANONYMITYCASHD_BINARY64) $(ANONYMITYCASHCLI_BINARY64) $(ANONYMITYCASH_RELEASE64).md5
	cd target && rm -f $(MINER_BINARY64) $(ANONYMITYCASHD_BINARY64) $(ANONYMITYCASHCLI_BINARY64) $(ANONYMITYCASH_RELEASE64).md5
endif

release-all: clean
	GOOS=darwin  make release
	GOOS=linux   make release
	GOOS=windows make release

clean:
	@echo "Cleaning binaries built..."
	@rm -rf cmd/anonymitycashd/anonymitycashd
	@rm -rf cmd/anonymitycashcli/anonymitycashcli
	@rm -rf cmd/miner/miner
	@rm -rf target
	@rm -rf $(GOPATH)/bin/anonymitycashd
	@rm -rf $(GOPATH)/bin/anonymitycashcli
	@echo "Cleaning temp test data..."
	@rm -rf test/pseudo_hsm*
	@rm -rf blockchain/pseudohsm/testdata/pseudo/
	@echo "Cleaning sm2 pem files..."
	@rm -rf crypto/sm2/*.pem
	@echo "Done."

target/$(ANONYMITYCASHD_BINARY32):
	CGO_ENABLED=0 GOARCH=386 go build $(BUILD_FLAGS) -o $@ cmd/anonymitycashd/main.go

target/$(ANONYMITYCASHD_BINARY64):
	CGO_ENABLED=0 GOARCH=amd64 go build $(BUILD_FLAGS) -o $@ cmd/anonymitycashd/main.go

target/$(ANONYMITYCASHCLI_BINARY32):
	CGO_ENABLED=0 GOARCH=386 go build $(BUILD_FLAGS) -o $@ cmd/anonymitycashcli/main.go

target/$(ANONYMITYCASHCLI_BINARY64):
	CGO_ENABLED=0 GOARCH=amd64 go build $(BUILD_FLAGS) -o $@ cmd/anonymitycashcli/main.go

target/$(MINER_BINARY32):
	CGO_ENABLED=0 GOARCH=386 go build $(BUILD_FLAGS) -o $@ cmd/miner/main.go

target/$(MINER_BINARY64):
	CGO_ENABLED=0 GOARCH=amd64 go build $(BUILD_FLAGS) -o $@ cmd/miner/main.go

test:
	@echo "====> Running go test"
	@go test -tags "network" $(PACKAGES)

benchmark:
	@go test -bench $(PACKAGES)

functional-tests:
	@go test -timeout=5m -tags="functional" ./test 

ci: test functional-tests

.PHONY: all target release-all clean test benchmark
