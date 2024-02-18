FROM golang:1.18.2-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app


FROM debian:bullseye-slim as depoy

RUN apt-get update

COPY --from=deploy-builder /app/app .
CMD ["./app"]

FROM golang:1.18.2 as dev
WORKDIR /app