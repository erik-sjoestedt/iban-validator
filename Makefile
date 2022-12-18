build:
	docker build -t iban-validator .
run:
	docker run --rm -p 5000:5000 iban-validator
test:
	go test ./...
