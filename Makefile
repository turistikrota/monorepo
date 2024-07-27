jwt-key:
	ssh-keygen -t rsa -b 4096 -m PEM -f ./config/turistikrota_jwtRS256.key

jwt-pub:
	openssl rsa -in ./config/turistikrota_jwtRS256.key -pubout -outform PEM -out ./config/turistikrota_jwtRS256.key.pub

jwt: jwt-key jwt-pub

env: 
	cp ./config/.env.example ./config/.env

temp:
	mkdir temp && mkdir temp/db && mkdir temp/kv && mkdir temp/grafana

compose:
	docker-compose -f ./config/docker-compose.yml up -d

compose-build:
	docker-compose -f ./config/docker-compose.yml up -d --build --remove-orphans

compose-down:
	docker-compose -f ./config/docker-compose.yml down

build-app:
	cd apps/api && docker build -t github.com/turistikrota/api:latest .

secret-register:
	docker secret create turistikrota_private_key ./config/turistikrota_jwtRS256.key
	docker secret create turistikrota_public_key ./config/turistikrota_jwtRS256.key.pub

network:
	docker network create --driver overlay --attachable turistikrota

start-app:
	docker service create --name turistikrota --publish 4000:4000 --secret turistikrota_private_key --secret turistikrota_public_key --replicas 3 --env-file ./config/.env --network turistikrota github.com/turistikrota/api:latest

stop-app:
	docker service rm turistikrota

clean:
	rm -rf temp
	rm -rf config/turistikrota_jwtRS256.key
	rm -rf config/turistikrota_jwtRS256.key.pub

clean-docker:
	docker service rm turistikrota
	docker secret rm turistikrota_private_key
	docker secret rm turistikrota_public_key
	docker network rm turistikrota
	docker rmi github.com/turistikrota/api:latest

clean-all: clean clean-docker

reqs: temp jwt-key jwt-pub env network compose secret-register

burn: temp jwt-key jwt-pub env network compose secret-register build-app start-app

stop: compose-down stop-app clean

reload: stop-app build-app start-app

.PHONY: jwt-key jwt-pub jwt temp compose compose-down build-app secret-register network start-app stop-app clean start stop burn