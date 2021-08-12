SHELL = /bin/bash

all: build docker

build:
	GOOS=linux go build -o main main.go

deploy:
	scf deploy -t template.yaml  -f

docker:
	docker run -it --rm  --env-file ~/.env -v $(shell pwd):/srv airdb/scf

log:
	docker run -it --rm  --env-file ~/.env -v $(shell pwd):/srv airdb/scf scf logs --name timer-mina
