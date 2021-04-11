build-frontend:
	docker-compose -f docker-compose-helpers.yml run build-frontend

deploy-test:
	docker-compose up

deploy:
	docker-compose up -d