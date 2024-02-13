FROM golang: 18.2.2-bullseye as deploy-builder

WORKDIR /app

COPY go.mod

ENTRYPOINT ["top", "-b"]