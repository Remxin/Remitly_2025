dbname = swift_db
dbuser = admin
dbpassword = secretpassword
dbhost = localhost
dbport = 5432

migrateup:
	migrate -path db/migration -database "postgresql://${dbuser}:${dbpassword}@${dbhost}:${dbport}/${dbname}?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://${dbuser}:${dbpassword}@${dbhost}:${dbport}/${dbname}?sslmode=disable" -verbose up

sqlc:
	sqlc generate

.PHONY: migrateup migratedown sqlc