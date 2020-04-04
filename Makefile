GIT_HOST = github.com
GIT_REPOSITORY = $(GIT_HOST)/kei2100/protoc-gen-marshal-zap
PLUGIN_DIR = plugin/protoc-gen-marshal-zap
PROTOC_OPT ?=

.PHONY: proto
proto:
	@$(MAKE) marshal-zap.pb.go

%.pb.go: %.proto
	protoc $(PROTOC_OPT) --go_out=$(*D) $(*).proto
	mv $(*D)/$(GIT_REPOSITORY)/$(*).*.go $(*D)
	rm -r $(*D)/$(GIT_HOST)

.PHONY:
install: proto
	cd $(PLUGIN_DIR) && go install

.PHONY: test
test: install test.proto
	go test -v $$(go list ./test/...)

.PHONY: test.proto
test.proto:
	@PROTOC_OPT='-I. -Itest/types --marshal-zap_out=test/types' $(MAKE) test/types/types.pb.go

.PHONY: test.ci
test.ci:
	@which act > /dev/null 2>&1 || brew install nektos/tap/act
	act -P ubuntu-latest=nektos/act-environments-ubuntu:18.04 -s GITHUB_TOKEN=$${GITHUB_TOKEN}
