.PHONY: test
test:
	go test -mod vendor -cover ./...

.PHONY: build
build:
	go build -mod vendor -o sketch-game

.PHONY: dev
dev: build
	./sketch-game -openCors

.PHONY: dev-ui
dev-ui:
	cd ui; yarn start
