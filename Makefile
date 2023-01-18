start-docker:
		cd docker && docker-compose up
stop-docker:
		cd docker && docker-compose stop
run-gateway:
		cd api-gateway && go run ./
run-service:
		cd custshopsvc && go run ./
