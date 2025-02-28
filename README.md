# Home Excercise 2025

## Run locally (developement)
1. Run on your command line
```
cp sample.env app.env
```

2. Check DB_SOURCE value at app.env

**For local developement:**

**Correct path**
✅ DB_SOURCE=postgres://admin:secretpassword@localhost:5432/swift_db?sslmode=disable

**Wrong path**❌ DB_SOURCE=postgres://admin:secretpassword@postgres:5432/swift_db?sslmode=disable


3. Run postgres service from docker-compose.yaml

This will automatically create swift_db database

4. Run application

**For developement**
```
make run
```

**For deployment**
```
go build -o main main.go
```
```
./main
```

### Optional steps

!Make sure to set:
**dbhost = localhost** in Makefile

- Upgrading db schema
```
make migrateup
```

- Downgrading db schema version
```
make migratedown
```

- Changing sql queries

Edit/Add new .sql files at /db/query folder. Then

```
make sqlc
```


## Run containerised application (release)
1. Create app.env
```
cp sample.env app.env
```

2. Check DB_SOURCE value at app.env

**For containerised version:**

**Correct path**
✅ DB_SOURCE=postgres://admin:secretpassword@postgres:5432/swift_db?sslmode=disable

**Wrong path** 
❌ DB_SOURCE=postgres://admin:secretpassword@localhost:5432/swift_db?sslmode=disable

3. Run docker-compose
```
docker-compose up -d
```
2 containers should be created:
- my_postgres
- swift_service

swift_service automatically launches script for adding swift_codes.xlsx data to my_postgres

To dismiss this part remove this code from main.go file :
```
	err = ExtractXLSXDataToDB(conn)
	if err != nil {
		log.Fatalf("unable to parse .xlsx data: %v", err)
	}
```

4. Application is ready

Now you can send requests to **swift_service** application via **localhost:8080**

## Troubleshooting

Sometimes docker-compose.yaml could not be executed properly because of permissions set to local folders (access denied). It's because of this lines:

```
volumes:
    - pgdata:/var/lib/postgresql/data
    - ./db/migration/000001_migration_name.up.sql:/docker-entrypoint-initdb.d/init.sql
```

To solve this problem:
- add local permissions (to ensure application has access to db folder)
```
chmod 777 db
```
- use sudo
```
sudo docker-compose up -d
```

