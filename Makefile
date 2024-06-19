#Makefile

run:
	@go run main.go

test:
	@go test ./parser ./evaluator ./ast ./lexer ./object
