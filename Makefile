SHELL=/bin/bash

.PHONY: lint
lint:
	@pushd ./src && \
		# golangci-lint cache clean && \
		golangci-lint run --config ../.github/.golangci.yml && \
		golangci-lint run --disable-all --config ../.github/.golangci.yml -E bodyclose && \
		golangci-lint run --disable-all --config ../.github/.golangci.yml -E exhaustive && \
		golangci-lint run --disable-all --config ../.github/.golangci.yml -E forbidigo && \
		golangci-lint run --disable-all --config ../.github/.golangci.yml -E forcetypeassert && \
		golangci-lint run --disable-all --config ../.github/.golangci.yml -E noctx && \
		golangci-lint run --disable-all --config ../.github/.golangci.yml -E whitespace && \
		golangci-lint run --disable-all --config ../.github/.golangci.yml -E gocognit && \
		golangci-lint run --disable-all --config ../.github/.golangci.yml -E unconvert && \
		golangci-lint run --disable-all --config ../.github/.golangci.yml -E gomnd && \
		golangci-lint run --disable-all --config ../.github/.golangci.yml -E errorlint && \
		golangci-lint run --disable-all --config ../.github/.golangci.yml -E gocyclo && \
		golangci-lint run --disable-all --config ../.github/.golangci.yml -E goimports && \
		golangci-lint run --disable-all --config ../.github/.golangci.yml -E gofmt && \
		golangci-lint run --disable-all --config ../.github/.golangci.yml -E errcheck && \
		golangci-lint run --disable-all --config ../.github/.golangci.yml -E gosec && \
	popd

		# golangci-lint run --config ../.github/.golangci.yml && \

dev-docker-up:
	@docker compose -f docker/development/docker-compose.yml up -d

dev-docker-down:
	@docker compose -f docker/development/docker-compose.yml down

test-docker-up:
	@docker compose -f docker/test/docker-compose.yml up -d
