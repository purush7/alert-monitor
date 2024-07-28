.PHONY: docker-remove-server build-server run-server all
.SILENT: docker-remove-server build-server run-server all

local:
	docker-compose up -d
 
docker-remove-server:
	cont=$(shell docker ps -a  | grep alert-monitor | awk '{print $$1 }'); \
	if [ "$$cont" ]; then docker stop alert-monitor && docker rm alert-monitor; \
	else echo "no container named alert-monitor found" ; \
	fi

build-server:
	docker build -f Dockerfile -t alert-monitor .

run-server: docker-remove-server
	docker run -d --restart unless-stopped --cap-add=SYS_PTRACE -p 3336:3333 --network alertnetwork -v ${HOME}/.tmp:/tmp --name alert-monitor alert-monitor

all: build-server run-server