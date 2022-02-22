.PHONY: all docker push test

IMAGE := internal/twirp-php81-gen:build

all: docker

docker:
	docker build --rm -t $(IMAGE) .

push:
	docker push $(IMAGE)

test:
	docker run -it --rm $(IMAGE) sh
