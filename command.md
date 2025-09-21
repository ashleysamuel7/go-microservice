

docker run -d \
  --name my-postgres \
  -e POSTGRES_USER=root \
  -e POSTGRES_PASSWORD=root \
  -e POSTGRES_DB=mydb \
  -v pgdata:/var/lib/postgresql/data \
  -p 5432:5432 \
  postgres

  CREATE ROLE root WITH LOGIN PASSWORD 'root';