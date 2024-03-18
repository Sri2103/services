
up:
	docker compose up -d

down:
	docker compose down
rm_docker_images:
	docker images | grep services|awk '{print $3}'|xargs docker rmi