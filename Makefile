default: build

test:
	go test ./...

build:
	go build

install: build
	mkdir -p ~/.tflint.d/plugins
	mv ./security-smells-ruleset ~/.tflint.d/plugins
