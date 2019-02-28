default: setup test

.PHONY: setup
setup:
	go get github.com/stretchr/testify

.PHONY: test
test:
	go test .
