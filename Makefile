run_gateway:
		go run gateway/cmd/main.go

build_image:
		docker compose -f docker-compose.yml up -d --build

run_image:
		docker compose -f docker-compose.yml up -d