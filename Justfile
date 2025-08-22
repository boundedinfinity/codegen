m   := "updates"

list:
    just --list

purge:
	find . -name '*.enum.go' -type f -delete

generate: purge
	go generate ./...

build: generate
	go mod tidy
	go build

install: generate
	go install

test: generate
	go test ./...

git-push:
    git add . || true
    git commit -m "{{ m }}" || true
    git push origin master

tag:
	git tag -a $(tag) -m "$(tag)"
	git push origin $(tag)

tag-list:
	git fetch --tags
	git tag -l | sort -V
