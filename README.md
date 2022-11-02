# rest-api
migrate -path ./schema -database 'postgres://postgers:qwerty@localhost:5436/postgres?sslmode=disable' up

docker exec -it 7e2a67f159c1 /bin/bash
psql -U postgres