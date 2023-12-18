default:
    @just --list

# build tinyurl binary
build:
    go build -o gourl ./cmd/gourl

# update go packages
update:
    @cd ./cmd/gourl && go get -u

# run golangci-lint
lint *flags:
    golangci-lint run -c .golangci.yml {{ flags }}
