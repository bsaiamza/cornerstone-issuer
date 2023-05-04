# docker commands
build_docker:
	docker build -t cornerstone-issuer:latest .
	docker tag cornerstone-issuer:latest 149875424875.dkr.ecr.af-south-1.amazonaws.com/cornerstone-issuer:latest

push_docker:
	docker push 149875424875.dkr.ecr.af-south-1.amazonaws.com/cornerstone-issuer:latest

# golang commands
fmt:
	go fmt ./...

lint: 
	golint ./...

test:
	go test -v -cover ./...

.PHONY: build_docker push_docker fmt lint test