FROM golang:latest
WORKDIR /wol_backend
ADD ./backend /wol_backend
ADD ./frontend/dist /frontend/dist
RUN go build -o server .
EXPOSE 9090
ENTRYPOINT ["./server"]
