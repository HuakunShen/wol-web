# Build Frontend
FROM node:latest AS frontend-builder
WORKDIR /wol/frontend
COPY frontend .
RUN npm install -g @vue/cli && \
    rm -rf dist node_modules && \
    npm ci && \
    npm run build

# Build Backend
FROM golang:1.16.3-alpine3.13 AS server-builder
WORKDIR /go/src
COPY backend .
RUN go build -o server .

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

COPY --from=server-builder /go/src/ /wol-server/
COPY --from=frontend-builder /wol/frontend/dist/ /frontend/dist/
COPY ./backend/docker_postgres_init.sql /docker-entrypoint-initdb.d/docker_postgres_init.sql
EXPOSE 9090
CMD ["./server"]

