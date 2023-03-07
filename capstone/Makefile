createdb:
	createdb goal_tracker

dropdb:
	dropdb goal_tracker

mcreate:
	migrate create -ext sql -dir db/migrations -seq

migrateup:
	migrate -path db/migrations -database "postgresql:///goal_tracker?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql:///goal_tracker?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

.PHONY: createdb dropdb mcreate migrateup migratedown sqlc test