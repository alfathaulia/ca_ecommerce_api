test:
	go test -v -cover ./...

mock:
	cd domain && mockery --all --keeptree 

run:
	go run app/main.go

.PHONY: test mock run