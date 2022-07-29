# docker commands
build_docker:
	docker build -t cornerstone-issuer:0.0.5 .
	docker tag cornerstone-issuer:0.0.5 149875424875.dkr.ecr.af-south-1.amazonaws.com/cornerstone-issuer:0.0.5

push_docker:
	docker push 149875424875.dkr.ecr.af-south-1.amazonaws.com/cornerstone-issuer:0.0.5

# golang commands
fmt:
	go fmt ./...

lint: 
	golint ./...

test:
	go test -v -cover ./...

.PHONY: build_docker push_docker fmt lint test