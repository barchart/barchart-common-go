.PHONY: $(MAKECMDGOALS)

test:
	go test ./...

test-v:
	go test -v ./...

tag:
	./tag.sh