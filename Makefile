makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list purge build install generate test commit tag publish

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

purge:
	find . -name '*.enum.go' -type f -delete

generate: purge
	go generate ./...

build: generate
	go build

install: generate
	go install

test: generate
	go test ./...

commit:
	git add . || true
	git commit -m "$(m)" || true
	git push origin master

tag:
	git tag -a $(tag) -m "$(tag)"
	git push origin $(tag)

publish: generate
	@if ack replace go.mod ;then echo 'Remove the "replace" line from the go.mod file'; exit 1; fi
	make commit m=$(m)
	make tag tag=$(m)