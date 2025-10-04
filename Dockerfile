FROM golang:1.21-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o /go-service ./cmd/server

FROM alpine:3.18
COPY --from=build /go-service /go-service
EXPOSE 8080
ENTRYPOINT ["/go-service"]
