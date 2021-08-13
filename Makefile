SHELL = /bin/bash

SLSENV=SERVERLESS_PLATFORM_VENDOR=tencent

all: build docker

build:
	GOOS=linux go build -o main main.go

deploy:
	${SLSENV} sls deploy --stage test

docker:
	docker run -it --rm  --env-file ~/.env -v $(shell pwd):/srv airdb/scf

log:
	${SLSENV} sls logs --tail --stage test

logrelease:
	${SLSENV} sls logs --tail --stage release
