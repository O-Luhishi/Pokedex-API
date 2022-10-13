# Build environment
# -----------------
FROM golang:1.17-alpine as build-env
WORKDIR /app

RUN apk update && apk add --no-cache gcc musl-dev git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/pokedex_api ./cmd/api  \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate \
    && chmod +x ./build/init.sh

# Deployment environment
# ----------------------
FROM alpine
RUN apk update && apk add --no-cache bash

COPY --from=build-env /app/bin/pokedex_api /app/
COPY --from=build-env /app/bin/migrate /app/
COPY --from=build-env /app/migrations /app/migrations
COPY --from=build-env /app/build/init.sh /app/

EXPOSE 8080
CMD "./app/init.sh"