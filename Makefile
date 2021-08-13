SHELL = /bin/bash

all: build docker

build:
	GOOS=linux go build -o main main.go

deploy:
	sls deploy --stage test

docker:
	docker run -it --rm  --env-file ~/.env -v $(shell pwd):/srv airdb/scf

log:
	${SLSENV} sls logs --tail --stage test

logrelease:
	${SLSENV} sls logs --tail --stage release
