# Docker

## 1. Postgres and SQLx

### choose the good version -> copy the whole file into the command line to export env variables

- for local development [local.standalone.env](./local.standalone.env)
- for docker compose local development [local.compose.env](./local.compose.env)
- for docker compose dev environment (cloud) [dev.compose.env](./dev.compose.env) -> this version may not be used here, as it will be used by the pipeline for the docker swarm build

## 2. depending on the setup, run these commands

### local development setup

```bash
# if first run
docker run -d --name inv-db-standalone \
-e POSTGRES_USER=some-postgres \
-e POSTGRES_DB=some-postgres \
-e POSTGRES_PASSWORD=mysecretpassword \
-p 5432:5432 postgres:16.3
sudo chmod +x ./scripts/connect_db.sh
./scripts/connect_db.sh
sql_script=$(cat "./scripts/db_script.sql")
psql $DATABASE_URL -c "$sql_script"
```

```bash
# if not the first run
docker start inv-db-standalone
cd app
cargo run
```

```bash
# to stop the database
docker stop inv-db-standalone
```

### local compose setup or dev compose setup -> no need to build for dev setup -> but you need to login to the github container registry

```bash
# use swarm (recommended)
docker swarm init
docker secret create regcred ~/.docker/config.json
docker stack deploy -c docker-compose.yml inv --with-registry-auth -d --prune
docker stack ps inv
```

```bash
# to delete the node in swarm
docker swarm leave --force
```

```bash
# compose method (alternative)
# clean up the images and volumes you don't need -> warning: this cleans ALL, clean what you need
docker stop $(docker ps -a -q)
docker system prune
docker rm $(docker ps -a -q)
docker rmi $(docker images -q)
docker volume rm $(docker volume ls -q)
# make sure to build the inv-frontend:local before! in the inv-frontend directory -> assuming that INV-Frontend is in ../INV-Frontend. to be extra careful, we remove cache from the build to avoid problems, simple! change the options as you need
cd ../INV-Frontend
docker build -t inv-frontend:local . --force-rm --no-cache
cd ../INV-Backend
docker build -t inv-backend:local . --force-rm --no-cache
docker compose up --build -d
```

```bash
# to stop docker compose
docker compose stop
```

```bash
# to run stopped docker compose
docker compose start
```

```bash
# to cleanup docker compose
docker compose down
```

#### check this link to connect with the github container registry

https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry#authenticating-with-a-personal-access-token-classic

## 3. access the database in the docker compose with the .env variables

```bash
docker-compose exec inv-db sh -c 'psql -U $POSTGRES_USER $POSTGRES_DB'
```

### useful commands for the psql terminal (you can also copy paste normal sql commands)

```bash
\dt
\q
```
