all:
	@cd $(GOPATH)/src; go install github.com/Cloud-Foundations/health-agent/cmd/*


TARBALL_TARGET = /tmp/$(LOGNAME)/health-agent.tar.gz

tarball:
	@cd $(GOPATH)/src; go install github.com/Cloud-Foundations/health-agent/cmd/health-agent
	@tar --owner=0 --group=0 -czf $(TARBALL_TARGET) \
	init.d/health-agent.* \
	scripts/install.lib \
	-C cmd/health-agent install \
	-C $(GOPATH) bin/health-agent


format:
	gofmt -s -w .

format-imports:
	goimports -w .


test:
	@find * -name '*_test.go' -printf 'github.com/Cloud-Foundations/health-agent/%h\n' \
	| sort -u | xargs -r go test
