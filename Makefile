dbname = swift_db
dbuser = admin
dbpassword = secretpassword
dbhost = localhost
dbport = 5432

migrateup:
	migrate -path db/migration -database "postgresql://${dbuser}:${dbpassword}@${dbhost}:${dbport}/${dbname}?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://${dbuser}:${dbpassword}@${dbhost}:${dbport}/${dbname}?sslmode=disable" -verbose down

sqlc:
	sqlc generate

mock: 
	mockgen -package mockdb -destination db/mock/store.go example.com/m/v2/db/sqlc Store

run:
	go run main.go

.PHONY: migrateup migratedown sqlc mock run