test:
	docker compose -f ../docker-compose.yml up --build --force-recreate form3-accountapi-client-test --attach form3-accountapi-client-test
	docker compose -f ../docker-compose.yml down --timeout=0 --volumes

unit-test:
	go test ./...

.PHONY: build-test-image test unit-test