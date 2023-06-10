#!/bin/bash
help:
	@echo "Wallet"
	@echo "https://github.com/piovani/wallet"
	@echo "-----------------------------------------------"
	@echo "COMMANDS:                                      "
	@echo "make help     # prints usage info              "
	@echo "make start    # start run API                  "
	@echo "make build    # build version from APP         "


start:
	go run main.go rest

build:
	go build -o service