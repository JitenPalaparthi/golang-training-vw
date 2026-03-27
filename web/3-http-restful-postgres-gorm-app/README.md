### go get 

```bash
#gin
go get github.com/gin-gonic/gin

# GORM
go get -u gorm.io/gorm

# Postgres
gorm.io/driver/postgres
```

### Docker commands

```bash
# create docker network
docker network create demo-network
docker pull postgres:latest

# to create a postgres container

docker run -d --name pg -p 5432:5432 --network demo-network -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=user-profiles-db postgres:latest

# to create admin panel container

docker pull adminer

# make sure two containers are running
docker ps

```
