# https://www.linkedin.com/pulse/go-dockerized-grpc-server-example-tiago-melo?trk=pulse-article_more-articles_related-content-card

FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server  .

EXPOSE 50051

CMD [ "/app/server" ]