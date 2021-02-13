GIT_VERSION=`git describe --tags`

.PHONY: cli
cli:
	go build --tags "json1" -o circuit-breaker ./cmd
	
.PHONY: contracts-all
contracts-all: build-contracts abigen

.PHONY: build-contracts
build-contracts:
	solc --bin --abi -o bin --overwrite contracts/BMath.sol
	solc --bin --abi -o bin --overwrite contracts/LogSwapTest.sol
	solc --optimize --bin --abi -o bin --overwrite contracts/SimpleMultiCall.sol

.PHONY: abigen
abigen:
	abigen --abi bin/BNum.abi --bin bin/BNum.bin --pkg bmat --out bindings/bmath/bindings.go
	abigen --abi bin/LogSwapTest.abi --bin bin/LogSwapTest.bin --pkg logswap --out bindings/logswap/bindings.go
	abigen --abi bin/SimpleMultiCall.abi --bin bin/SimpleMultiCall.bin --pkg multicall --out bindings/multicall/bindings.go

# TESTFLAGS="-v -cover" make test
.PHONY: test
test:
	go test --tags "json1" $(TESTFLAGS) ./...


.PHONY: test-ci
test-ci:
	# test once without race (and generate coverage profile)
	go test --tags "json1" -short -failfast -coverprofile=coverage.txt -cover ./...
	# test with race
	go test --tags "json1" -short -failfast -race ./...

.PHONY: docker-build
docker-build:
	docker build --build-arg VERSION=$(GIT_VERSION) -t bonedaddy/circuit-breaker:$(GIT_VERSION) .
	docker image tag bonedaddy/circuit-breaker:$(GIT_VERSION) bonedaddy/circuit-breaker:latest

.PHONY: release
release:
	.scripts/release.sh
