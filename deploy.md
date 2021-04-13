# Deployment Instructions

## Raspberry Pi

### Option 1: docker-compose

```bash
docker-compose up
# then ctrl c to stop and run in detach mode
docker-compose up -d
```

Some of the following options work for raspberry pi but not guaranteed

## amd64 or macos

### Option 1: docker-compose

```bash
docker-compose up -d
```

### Option 2: Makefile

```bash
make deploy-test
# make sure it works, then stop, and run next line
make deploy
```

### Option 3: Docker image

```bash
docker volume create wol-web-db

docker run --name wol-db \
-e POSTGRES_PASSWORD=wakeonlan \
-e POSTGRES_USER=wol \
-e POSTGRES_DB=wol \
-e PGDATA=/var/lib/postgresql/data/pgdata \
-v wol-web-db:/var/lib/postgresql/data \
--restart unless-stopped \
--network host \
-d \
postgres:13.2-alpine

docker run --name wol-server \
--restart=unless-stopped \
--network host \
-e POSTGRES_PASSWORD=wakeonlan \
-e JWT_SECRET=wol-jwt \
-d \
huakunshen/wol:latest
```

### Option 4: Manual Setup

1. Install nodejs and golang
2. Start a postgres database
   ```bash
   docker volume create wol-web-db
    docker run --name wol-db \
    -e POSTGRES_PASSWORD=wakeonlan \
    -e POSTGRES_USER=wol \
    -e POSTGRES_DB=wol \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v wol-web-db:/var/lib/postgresql/data \
    --restart unless-stopped \
    --network host \
    -d \
    postgres:13.2-alpine
   ```
3. Build frontend
   ```bash
   npm i -g @vue/cli
   cd frontend
   npm install
   npm run build
   ```
4. Build backend server
   ```bash
   cd backend
   go build -o server .
   ./server
   ```
