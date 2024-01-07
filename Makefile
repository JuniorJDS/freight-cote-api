run:
	go run api/main.go

mongo-local:
	docker-compose -f docker-compose-run-local.yaml stop
	docker-compose -f docker-compose-run-local.yaml rm -f
	docker-compose -f docker-compose-run-local.yaml up mongo-db

integration-tests:
	docker-compose -f docker-compose-tests.yaml stop
	docker-compose -f docker-compose-tests.yaml rm -f
	docker-compose -f docker-compose-tests.yaml build
	docker-compose -f docker-compose-tests.yaml up --exit-code-from app-test
