
linter-golangci: ### check by golangci linter
	golangci-lint run

linter-hadolint: ### check by hadolint linter
	hadolint Dockerfile

linter: ### check by all linters
	golangci-lint run && hadolint Dockerfile

swag-init: ### init swagger docs
	swag init -q --parseDependency --parseInternal -g cmd/app/main.go