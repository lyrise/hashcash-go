GOCMD=go
PACKAGE_NAME=hashcash

build:
	$(GOCMD) build -o ./build/$(PACKAGE_NAME) ./cmd/$(PACKAGE_NAME)/main.go
clean:
	$(GOCMD) clean
run:
	$(GOCMD) run ./cmd/$(PACKAGE_NAME)/main.go
