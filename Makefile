run:
	go run api/main.go

mongo-local:
	docker-compose -f docker-compose-run-local.yaml stop
	docker-compose -f docker-compose-run-local.yaml rm -f
	docker-compose -f docker-compose-run-local.yaml up mongo-db