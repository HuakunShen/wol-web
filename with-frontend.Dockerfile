# Build Backend
FROM golang:1.17.6-alpine3.15 AS server-builder
WORKDIR /go/src
COPY backend .
RUN go build -o server .

FROM golang:1.16.3-alpine3.13 AS wol-cli-builder
WORKDIR /root
RUN apk update && apk add git && git clone https://github.com/HuakunShen/wol.git
WORKDIR /root/wol
RUN go build .

# Add compiled frontend and backend to postgres db for deployment
FROM alpine:3.13.4
WORKDIR /wol-server

# Environment Variables that can be set
ENV POSTGRES_PASSWORD wakeonlan
ENV POSTGRES_USER wol
ENV POSTGRES_DB wol
ENV PORT 9090
ENV JWT_SECRET secret
ENV JWT_VALID_TIME 14400
ENV POSTGRES_HOST localhost
ENV POSTGRES_PORT 5432
ENV POSTGRES_TIMESZONE America/Toronto
ENV NUM_USER_ALLOWED 1

COPY --from=server-builder /go/src/server /wol-server/server
COPY --from=server-builder /go/src/.env /wol-server/.env
COPY --from=wol-cli-builder /root/wol/wol /usr/local/bin
COPY ./frontend/dist/ /frontend/dist/
COPY ./backend/docker_postgres_init.sql /docker-entrypoint-initdb.d/docker_postgres_init.sql
EXPOSE 9090
CMD ["./server"]

