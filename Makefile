all: build run

build:
	docker-compose build

run:
	docker-compose up

stop:
	docker-compose down

test:
	HOST_DB=localhost \
    PORT_DB=5432 \
    USER_DB=postgres \
    PASSWORD_DB=gatomagico4444 \
    NAME_DB=postgres \
	go test ./pkg/...