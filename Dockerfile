FROM docker.io/library/golang:1.23 AS server_builder
WORKDIR /app
COPY ./apps/server /app
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go


FROM docker.io/oven/bun:1 AS web_builder
WORKDIR /app
COPY . .
RUN bun install
RUN bun run build


FROM docker.io/library/alpine:latest
# FROM ubuntu:latest
WORKDIR /app
COPY --from=server_builder /app/server ./server
COPY --from=web_builder /app/apps/web/build ./pb_public

# Add environment variable with default value
ENV PORT=8090
EXPOSE $PORT

CMD ["/bin/sh", "-c", "/app/server serve --http=0.0.0.0:${PORT}"]
# docker build . -t huakunshen/wol:latest
# docker run -p 8090:8090 --rm huakunshen/wol:latest
