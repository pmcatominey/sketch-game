.PHONY: tools
tools:
	GO111MODULE=off go get github.com/rakyll/statik

.PHONY: test
test:
	go test -cover ./...

.PHONY: build
build:
	go build -o sketch-game cmd/main.go

.PHONY: dev
dev: build
	./sketch-game -openCors

.PHONY: dev-ui
dev-ui:
	cd ui; yarn start

.PHONY: dist
dist:
	cd ui; npm run build
	statik -src ui/dist -dest pkg -p ui
	go build -o sketch-game cmd/main.go
