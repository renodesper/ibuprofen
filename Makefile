BREWPATH=/opt/homebrew
GOCMD=go
GOFUMPT=gofumpt
GOFUMPTPATH := $(GOPATH)/bin/gofumpt
GOFUMPTBREWPATH := $(BREWPATH)/bin/gofumpt

all: help

.PHONY: format
format:
	test -s ${GOFUMPTBREWPATH} || test -s ${GOFUMPTPATH} || $(GOCMD) install mvdan.cc/gofumpt@latest
	$(GOFUMPT) -l -w .

.PHONY: build
build:
	cd function && \
		rm -f app.zip main && \
		GOOS=linux GOARCH=amd64 $(GOCMD) build -o main . && \
		zip app.zip main

.PHONY: deploy
deploy: build
	pulumi up

.PHONY: clean
clean:
	cd function && \
		rm -f app.zip main

.PHONY: help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<arg>${RESET}'
	@echo ''
	@echo 'Arguments:'
	@echo "  ${YELLOW}help       ${RESET} ${GREEN}Show this help message${RESET}"
	@echo "  ${YELLOW}format     ${RESET} ${GREEN}Format '*.go' files with gofumpt${RESET}"
	@echo "  ${YELLOW}build      ${RESET} ${GREEN}Build the lambda code${RESET}"
	@echo "  ${YELLOW}deploy     ${RESET} ${GREEN}Deploy lambda function to AWS${RESET}"
	@echo "  ${YELLOW}clean      ${RESET} ${GREEN}Remove generated files${RESET}"
