.PHONY: fmt
fmt:
	go fmt ./...

# Test
.PHONY: clean-test-cache
clean-test-cache:
	go clean -testcache

.PHONY: test
test: fmt clean-test-cache 
	go test ./...

.PHONY: testv
testv: fmt clean-test-cache 
	go test -v ./...