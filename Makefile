build:
	go build -o ./out/mutants-analyser ./cmd/mutants-analyser-api

run: build
	./out/mutants-analyser
