# syntax=docker/dockerfile:1
FROM golang:1.22.1-alpine AS base
WORKDIR /src
COPY go.mod go.sum .
RUN go mod download
COPY . .

FROM base AS build-server
RUN go build -o /bin/server ./cmd/film-library

FROM scratch AS server
COPY --from=build-server /bin/server /bin/
ENTRYPOINT [ "/bin/server" ]
