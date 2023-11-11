
.PHONY:up

up:
	$(MAKE) service-start SERVICE=kdyr

	docker-compose up -d ui

.PHONY:build
build:
	docker-compose build

.PHONY:stop
stop:
	docker-compose stop

.PHONY:down
down:
	docker-compose down

.PHONY:service-start
service-start:
	docker-compose up -d ${SERVICE}-postgres
	./hack/scripts/wait-for-postgres.sh ${SERVICE}-postgres
	$(MAKE) migrations-up SERVICE=${SERVICE}
	docker-compose up -d ${SERVICE}

.PHONY:service-build
service-build:
	docker-compose build ${SERVICE}

.PHONY:service-rebuild
service-rebuild:
	docker-compose stop ${SERVICE}
	docker-compose build ${SERVICE}
	$(MAKE) service-start SERVICE=${SERVICE}

.PHONY:service-stop
service-stop:
	docker-compose stop ${SERVICE}

.PHONY:migrations-up
migrations-up:
	docker run --rm -v $(DIR)/pkg/${SERVICE}/repository/postgres/migrations:/migrations \
		--network kirleg-do-you-remember_default migrate/migrate -path=/migrations/ \
		-database postgres://dev:12345@${SERVICE}-postgres:5432/${SERVICE}?sslmode=disable up

.PHONY:migrations-down
migrations-down:
	docker run --rm -v $(DIR)/pkg/${SERVICE}/repository/postgres/migrations:/migrations \
		--network kirleg-do-you-remember_default migrate/migrate -path=/migrations/ \
		-database postgres://dev:12345@${SERVICE}-postgres:5432/${SERVICE}?sslmode=disable down -all