FROM golang:1.24.0-alpine AS builder

WORKDIR /usr/local/src

RUN apk add --no-cache build-base sqlite-dev

COPY ["app/go.mod", "app/go.sum", "./"]
RUN go mod download

COPY app ./
RUN go build -o bin/app ./cmd/main.go

FROM alpine AS runner

WORKDIR /app
COPY app ./

RUN apk add --no-cache sqlite

COPY --from=builder /usr/local/src/bin/app ./app

CMD ["./app"]