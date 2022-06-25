build-test-image:
	docker build -f build/Dockerfile-e2e-test -t form3-accountapi-client-test .

test: build-test-image
	docker compose -f ../docker-compose.yml up --detach
	docker compose logs -f -t form3-accountapi-client-test --no-color
	docker compose -f ../docker-compose.yml down --timeout=0 --volumes

unit-test:
	go test ./...

.PHONY: build-test-image test unit-test