run:
	go run api/main.go

mongo-local:
	docker-compose -f docker-compose-run-local-for-tests.yaml stop
	docker-compose -f docker-compose-run-local-for-tests.yaml rm -f
	docker-compose -f docker-compose-run-local-for-tests.yaml up mongo-db

run-local-environment:
	docker-compose -f docker-compose-run-local.yaml down
	docker-compose -f docker-compose-run-local.yaml stop
	docker-compose -f docker-compose-run-local.yaml build --force
	docker-compose -f docker-compose-run-local.yaml up

integration-tests:
	docker-compose -f docker-compose-tests.yaml stop
	docker-compose -f docker-compose-tests.yaml rm -f
	docker-compose -f docker-compose-tests.yaml build
	docker-compose -f docker-compose-tests.yaml up --exit-code-from app-test
