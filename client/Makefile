test:
	docker build -f build/Dockerfile-e2e-test -t form3-accountapi-client-test .
	docker run form3-accountapi-client-test

unit-test:
	go test ./...
