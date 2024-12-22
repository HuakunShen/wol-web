
FROM golang:1.23 AS server_builder
WORKDIR /app
COPY ./apps/server /app
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go


FROM oven/bun:1 AS web_builder
WORKDIR /app
COPY . .
RUN bun install
RUN bun run build


FROM alpine:latest
# FROM ubuntu:latest
WORKDIR /app
COPY --from=server_builder /app/server ./server
COPY --from=web_builder /app/apps/web/build ./pb_public

EXPOSE 8090

CMD ["/app/server", "serve", "--http=0.0.0.0:8090"]
# docker build . -t huakun/wol:latest
# docker run -p 8090:8090 --rm huakun/wol:latest
