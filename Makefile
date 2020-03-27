GIT_HOST=github.com
GIT_REPOSITORY=$(GIT_HOST)/kei2100/protoc-gen-marshal-zap
PLUGIN_DIR=plugin/protoc-gen-marshal-zap

.PHONY: proto
proto:
	@$(MAKE) marshal-zap.pb.go

%.pb.go: %.proto
	protoc --go_out=$(*D) $(*).proto
	mv $(*D)/$(GIT_REPOSITORY)/$(@) $(*D)
	rm -r $(*D)/$(GIT_HOST)

.PHONY:
install: proto
	cd $(PLUGIN_DIR) && go install
