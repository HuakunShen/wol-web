# Build Backend
FROM golang:1.17.6-alpine3.15 AS server-builder
WORKDIR /go/src
COPY backend .
RUN apk add build-base && go build -o server .

FROM golang:1.17.6-alpine3.15 AS wol-cli-builder
WORKDIR /root
RUN apk update && apk add git && git clone https://github.com/HuakunShen/wol.git
WORKDIR /root/wol
RUN go build .

# Add compiled frontend and backend for deployment
FROM alpine:3.13.4
WORKDIR /wol-server

# Environment Variables that can be set
ENV PORT 9090
ENV JWT_SECRET secret
ENV JWT_VALID_TIME 14400
ENV NUM_USER_ALLOWED 1
# Copy server executable
COPY --from=server-builder /go/src/server /wol-server/server
# Copy environment variable file
COPY --from=server-builder /go/src/.env /wol-server/.env
# Copy wol cli
COPY --from=wol-cli-builder /root/wol/wol /usr/local/bin
# Copy frontend, make sure the production build of frontend is in ./frontend/dist/
COPY .//dist/pwa/ /frontend/dist/
EXPOSE 9090
CMD ["./server"]

