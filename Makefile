makefile_dir		:= $(abspath $(shell pwd))
template_dir		:= $(makefile_dir)/generator/templates
m					?= "updates"

.PHONY: list purge build install generate test commit tag publish

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

purge:
	find . -name '*.enum.go' -type f -delete

generate: purge
	go generate ./...
	cp $(template_dir)/integer.type.go.tpl $(template_dir)/float.type.go.tpl
	cp $(template_dir)/integer.persistence.go.tpl $(template_dir)/float.persistence.go.tpl

build: generate
	go mod tidy
	go build

install: generate
	go install

test: generate
	go test ./...

push:
	git add . || true
	git commit -m "$(m)" || true
	git pull origin master
	git push origin master

tag:
	git tag -a $(tag) -m "$(tag)"
	git push origin $(tag)

tag-list:
	git fetch --tags
	git tag -l | sort -V

publish: generate
	@if ack replace go.mod ;then echo 'Remove the "replace" line from the go.mod file'; exit 1; fi
	make commit m=$(m)
	make tag tag=$(m)
