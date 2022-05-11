docker-down:
	docker-compose -f 'docker-compose.yml'  -p 'cornerstone_issuer' down
	
docker-start:
	docker-compose -f 'docker-compose.yml'  -p 'cornerstone_issuer' start

docker-stop:
	docker-compose -f 'docker-compose.yml'  -p 'cornerstone_issuer' stop

docker-up:
	docker-compose -f "docker-compose.yml" up -d --build --remove-orphans

fmt:
	go fmt ./...

lint: 
	golint ./...

run:
	go run cmd/main.go
	
test:
	go test -v -cover ./...

build_docker:
	#docker build -t georgelza/cornerstone-issuer:0.0.2 .
	docker build -t cornerstone-issuer .
	docker tag cornerstone-issuer:0.0.2 149875424875.dkr.ecr.af-south-1.amazonaws.com/cornerstone-issuer:0.0.2

push_docker:
	#docker push georgelza/cornerstone-issuer:0.0.2
	docker push 149875424875.dkr.ecr.af-south-1.amazonaws.com/cornerstone-issuer:0.0.2


.PHONY: docker-down docker-start docker-stop docker-up fmt lint run test build_docker push_docker