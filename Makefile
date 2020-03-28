.PHONY: test
test:
	go test -cover ./...

.PHONY: build
build:
	go build -o sketch-game

.PHONY: dev
dev: build
	./sketch-game -openCors

.PHONY: dev-ui
dev-ui:
	cd ui; yarn start
