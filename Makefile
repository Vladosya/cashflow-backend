.PHONY: run
run: # билдить проект с помощью команды make
	go run ./cmd/cashflow

.DEFAULT_GOAL := run