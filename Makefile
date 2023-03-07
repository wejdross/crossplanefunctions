dockerhub:
	docker build -t fnio -f Dockerfile .
	docker tag fnio ghcr.io/wejdross/fnio
	docker push ghcr.io/wejdross/fnio

reload:
	kubectl delete -f claim.yaml
	kubectl apply -f claim.yaml

all: dockerhub reload