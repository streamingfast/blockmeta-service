FROM golang:1.22-alpine as build
WORKDIR /app

COPY ./server/go.mod ./server/go.sum ./
RUN go mod download

COPY ./server/ ./

RUN go build -o /app/blockmeta ./cmd/blockmeta/*

####

FROM alpine:edge

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=build /app/blockmeta /app/blockmeta

