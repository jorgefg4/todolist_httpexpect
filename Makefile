all: build run

build:
	docker build -t todolist .

run:
	docker run --rm -p 8000:8000 --name todolist todolist

stop:
	docker stop todolist