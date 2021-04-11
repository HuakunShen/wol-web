# wol-web

A web app hosted locally for wakeonlan, turn on computers in private network.

Frontend: Vuejs

Backend: golang + fiber

# Docker

Modify the image within `docker-compose.yml`

## x64

### golang image:

- golang:1.16.3-alpine3.13

### postgres image:

- postgres:13.2

## raspberry pi (arm)

### golang image:

- arm32v7/golang:1.16.3-alpine3.13

### postgres image:

- arm32v7/postgres:13.2

### Other Arm Libraries

> replace arm32v7 with any of the following based on your arm device

- arm32v5
- arm32v6
- arm32v7
- arm64v8
