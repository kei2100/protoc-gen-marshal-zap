GIT_HOST=github.com
GIT_REPOSITORY=$(GIT_HOST)/kei2100/protoc-gen-marshal-zap
PLUGIN_DIR=plugin/protoc-gen-marshal-zap

.PHONY: proto
proto:
	@$(MAKE) marshal-zap.pb.go

%.pb.go: %.proto
	cd $(*D)
	protoc --go_out=. $(*).proto

.PHONY:
install: proto
	cd $(PLUGIN_DIR) && go install
