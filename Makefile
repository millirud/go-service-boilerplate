
linter-golangci: ### check by golangci linter
	golangci-lint run

linter-hadolint: ### check by hadolint linter
	hadolint Dockerfile

linter: ### check by all linters
	golangci-lint run && hadolint Dockerfile