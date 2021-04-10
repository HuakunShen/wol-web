FROM golang:latest
WORKDIR /wol_backend
ADD . /wol_backend
RUN go build -o server .
EXPOSE 9090
ENTRYPOINT ["./server"]
