# wol-web

[Deployment](#deployment)

- [wol-web](#wol-web)
  - [Use Cases](#use-cases)
  - [UI](#ui)
  - [To Develop](#to-develop)
    - [Frontend](#frontend)
    - [Database](#database)
    - [Backend](#backend)
      - [Run backend with docker compose](#run-backend-with-docker-compose)
    - [API Documentation](#api-documentation)
  - [Docker Environment](#docker-environment)
      - [golang image:](#golang-image)
      - [nodejs image:](#nodejs-image)
  - [Build Frontend](#build-frontend)
    - [Build with Makefile (recommended)](#build-with-makefile-recommended)
    - [Manual Installation and Build](#manual-installation-and-build)
    - [Build with docker compose (easier)](#build-with-docker-compose-easier)
  - [Backend](#backend-1)
  - [Deployment](#deployment)
    - [Docker Image](#docker-image)
    - [Supported Platform](#supported-platform)
    - [docker run](#docker-run)
    - [Dockerfile](#dockerfile)
  - [Environment Variables](#environment-variables)
    - [Sample Command](#sample-command)
      - [Option 1](#option-1)
      - [Option 2](#option-2)

## Use Cases

A web app hosted locally for **wakeonlan**, turn on computers in private network.

Use VPN to go into your network and wake up your computers with a simple click in your browser.

**Frontend:** Vue.js + TypeScript

**Backend:** golang + fiber

## UI

![image](https://user-images.githubusercontent.com/33727687/186250678-93f4c626-e8d7-4eee-ab63-d636069636cc.png)


## To Develop

### Frontend

```bash
npm i -g @vue/cli
cd frontend
npm install
npm run serve
```

### Database

Database has switched from postgresql to sqlite. So there is no need to set up a database, the server will handle sqlite.

### Backend

```bash
cd backend
mkdir -p data
go get -u github.com/cosmtrek/air
# add go/bin to path
export PATH=$HOME/go/bin:$PATH    # on linux, similar on mac
air                                     # start live reload
```

#### Run backend with docker compose

```bash
docker compose -f docker-compose-helpers.yml run dev-backend
```

or `make dev-backend`

See [backend](./backend/README.md) and [Environment Variables](#environment-variables) for more configuration options.

### API Documentation

https://documenter.getpostman.com/view/UVRAJ7MZ?version=latest

## Docker Environment

The app can be hosted with docker which requires the machine to have `docker` and `docker compose` installed.

If you want to run it without docker (with native golang), see later sections: [Deploy Without Docker](#deploy-without-docker)

Modify the image tags within `docker compose.yml` and `docker docker-compose-helpers.yml` depend on what machine you are running.

#### golang image:

- golang:1.17.6-alpine3.15

#### nodejs image:

- node:16

## Build Frontend

To deploy the app, you don't need to build frontend, just download the release from github.

Or using this command `make download-frontend`.

If you need to build it, read the following instructions.

The frontend is written in vuejs and needs to be built manually to generate a `dist` folder which contains the `index.html` and other resources.

### Build with Makefile (recommended)

```bash
make build-frontend     # exactly the same as the docker compose method, just a simplified wrapper
```

### Manual Installation and Build

If you have nodejs 15+, npm on your machine, you can cd into **frontend** and run

```bash
sudo npm install -g @vue/cli        # if you don't have vue on your machine.
npm install
npm run build
```

### Build with docker compose (easier)

If you don't have the dependencies installed, you can use **docker compose** to build the frontend production build.

```bash
docker compose -f docker-compose-helpers.yml run build-frontend
```

## Backend

For more information and configuration related to backend, check [backend README](./backend/README.md)

You can configure

- port of the server
- number of users allowed to sign up
- JWT Secret and Login Time (JWT_VALID_TIME)

Run `make build-backend` to build the backend binary.

If you have golang installed, you can also run `go build . -o server` within `backend` folder.

## Deployment

### Docker Image

[`huakunshen/wol:latest`](https://hub.docker.com/repository/docker/huakunshen/wol)

### Supported Platform

- linux/arm64/v8
- linux/arm/v6
- inux/arm/v7
- linux/amd64

The docker image contains everything you need to run the app, including a wakeonlan cli called `wol`. You can run a container with network=host to use the `wol` cli tool.

### docker run

```bash
docker volume create wol
docker run -d --network=host --restart=unless-stopped --name wol-web -v wol:/wol-server/data huakunshen/wol:latest
```

or just run `make deploy` (alias of the `docker run` command above).

run `make deploy-test` to run without detach mode.

You may add customized environment variables following the [instruction](#environment-variables).

### Dockerfile

There 2 versions of Dockerfile used to build docker image.

- [Dockerfile](./Dockerfile)
  - build both vue frontend and golang server into the image
  - the `Dockerfile` assumes that the frontend is already compiled (in `frontend/dist`)
    - run `make build-frontend` or `make download-frontend` to generate a production build
  - then `docker build -t huakunshen/wol:latest .` to build the image
  - `make buildx` will generate multi-platform image

## Environment Variables

Environment Variables can be added/overwritten by:

- adding `environment:` to `docker compose` service or
- adding `-e env_name=env_value` to `docker run`

The following variables are the default environment variables.

```
PORT=9090
JWT_SECRET=secret
JWT_VALID_TIME=14400    # in minute

NUM_USER_ALLOWED=1
```

`NUM_USER_ALLOWED` env variable can be used to specify the number of user allowed to sign up. Default is 1 if you are the only user.

During development, both Database and Server environment variables can be modifed in `backend/.env`

check [backend](./backend/README.md) too.

### Sample Command

#### Option 1

1. edit `backend/.env`
2. cd into this directory (wol-web)

```bash
docker run -d \
  --network=host --name wol-web \
  -v ${PWD}/wol-web-data:/wol-server/data \
  --env-file backend/.env \
  huakunshen/wol:latest
```

#### Option 2

```bash
docker run -d \
  --network=host --name wol-web \
  -v ${PWD}/wol-web-data:/wol-server/data \
  -e PORT=9090 \
  -e JWT_SECRET=wol-secret \
  -e JWT_VALID_TIME=20000 \
  -e NUM_USER_ALLOWED=1 \
  huakunshen/wol:latest
```
