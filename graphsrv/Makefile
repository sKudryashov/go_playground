build:
	docker-compose build --no-cache

up:
	docker-compose -f docker-compose.yml up --force-recreate

start:
	$(MAKE) build && $(MAKE) up

test:
	go test ./...

.PHONY: test
