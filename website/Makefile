build:
	hugo
	docker build -t reg.lab.st/echo .

push: build
	docker push reg.lab.st/echo

.PHONY: build push
