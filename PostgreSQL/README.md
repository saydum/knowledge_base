docker run --name pgsql -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres

docker exec -it pgsql bash

 psql -U postgres