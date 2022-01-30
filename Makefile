deploy: download-frontend
	docker-compose up -d

deploy-test: download-frontend
	docker-compose up

build: build-frontend build-backend

build-frontend:
	docker-compose -f docker-compose-helpers.yml run build-frontend && docker-compose -f docker-compose-helpers.yml down

build-backend:
	docker-compose -f docker-compose-helpers.yml run build-backend && docker-compose -f docker-compose-helpers.yml down

download-frontend:
	rm -rf ./frontend/dist ./frontend/dist.zip
	curl -s https://api.github.com/repos/HuakunShen/wol-web/releases/latest | grep "browser_download_url.*dist\.zip" | cut -d '"' -f 4 | wget -i -
	mv dist.zip frontend
	unzip ./frontend/dist.zip -d ./frontend

dev-backend:
	docker-compose -f docker-compose-helpers.yml run dev-backend

dev-frontend:
	docker-compose -f docker-compose-helpers.yml run dev-frontend

run-db:
	docker-compose run db

buildx:
	docker buildx build --push \
		--platform linux/arm64/v8,linux/arm/v6,linux/arm/v7,linux/amd64 \
		-t huakunshen/wol:latest .

clean:
	docker-compose down
	docker-compose -f docker-compose-helpers.yml down
	docker volume rm wol-web_wol-web-db