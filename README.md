# wol-web

> A web app hosted locally for wakeonlan
>
> This is a rewrite with sveltekit + PocketBase to replace the old version I wrote a few years ago
>
> In the new rewrite, instead of writing an entire rest API server and manage database with gorm, I use PocketBase as backend and database. Its golang extension feature allows me to add the wakeonlan feature easily. The most complicated part, Auth, is also fully handled by PocketBase, saving lots of time. PocketBase also has a built-in database management UI, making user creation/management much easier.

![](https://i.imgur.com/93GMAf8.png)

## Deployment

Deployment with docker is super simple.

Docker image [`huakunshen/wol`](https://hub.docker.com/repository/docker/huakunshen/wol/) is available on docker hub.

Both `linux/amd64` and `linux/arm64` are supported.

> [!IMPORTANT]
> 1. The container must be in the same network as the target hosts
> 2. The container must be started with `--network=host`
> 3. Mac doesn't support `--network=host` with docker. On Mac you have to run server with go directly. It's recommended to use linux.

### Step 1: Start Server

Here is a full command to start the server with a superuser initialized

```bash
docker run -p 8090:8090 --rm \
  --network=host \
  -e SUPERUSER_EMAIL=<root@example.com> \
  -e SUPERUSER_PASSWORD=<your password> \
  -v ./pb_data:/app/pb_data \
  huakunshen/wol:latest
```

- In this example I used `--rm` to clean up the container after it's closed for demo purpose
- In production, you should replace `--rm` with `-d` to run it in detach mode

The 2 environment variables and volume are optional but recommended.

- The volume is for data persistence, so you don't lose your data after container is destroyed.
  - Instead of using a local directory, it's better to create a volume and bind to it.
- The 2 environment variables are used to create an initial superuser in database
  - This project uses pocketbase as its backend and database, there is no user register feature as we shouldn't allow random person to register and send magic packets in your network. The only way to create user is log into pocketbase admin console with a superuser account and manually create user in the `users` collection/table.
  - When both `SUPERUSER_EMAIL` and `SUPERUSER_PASSWORD` are set, the server will create this superuser the first time it starts and you could login directly.
  - If you didn't set initial superuser credentials, you could also create a superuser
    - A long URL should be printed to console, open it in browser. The token in the URL allows you to create a superuser
    ```
    (!) Launch the URL below in the browser if it hasn't been open already to create your first superuser account:
    http://0.0.0.0:8090/_/#/pbinstal/eyJhbGciOiJIUzI1NiIs...
    ```
    - If you ran `docker run -d` in detach mode, run `docker logs <container name>` to find the long URL for superuser creation.

### Step 2: Create a Regular User

The superuser we discussed previously is like a database admin, you need to create a regular user to login to the website.

1. Go to [`http://localhost:8090/_/`](http://localhost:8090/_/) (or the url of your environment), login with superuser credentials
2. Go to `users` collection, create a user, remember your emial and password

### Step 3: Login to WOL Web

You can go to `http://localhost:8090/auth` and login with your regular user credentials.

Create a host then your can wake up your computer from browser.

## Develop

**Prerequisite:**
1. Bun
2. Golang

bun workspace is used to manage this monorepo, `dev` script will start frontend and backend together.

```bash
bun install
bun run dev # start both sveltekit dev server and golang pocketbase server
```

### Server

The golang server is in `apps/server`. It's a golang pocketbase extension with some custom routes.

```bash
air # start development
go run main.go serve # start the server without hot reload
```

#### Migrations

When table is modified, run `go run . migrate collections` to generate a migration `.go` file that will be auto loaded.


### Frontend

The frontend is in `apps/web`, written with sveltekit + `@sveltejs/adapter-static`.

In development, use `http://localhost:5173`. 

Running `bun run build` will generate a `apps/web/build` directory, 
and will be automatically copied to `apps/server/pb_public` ready to be served directly by the golang server as static assets.

In production the website is accesssible at `http://localhost:8090/`.

### Docker

`make buildx` automatically builds docker image for `linux/amd64` and `linux/arm64` and push to dockerhub.
