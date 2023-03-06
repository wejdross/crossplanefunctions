dockerhub:
	docker build -t fnio -f Dockerfile .
	docker tag fnio wejdross/fnio
	docker push wejdross/fnio

reload:
	kubectl delete -f claim.yaml
	kubectl apply -f claim.yaml

all: dockerhub reload