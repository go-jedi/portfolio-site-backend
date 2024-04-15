FROM golang:1.22.2-alpine AS builder

WORKDIR /github.com/go-jedi/portfolio/app/
COPY . /github.com/go-jedi/portfolio/app/

RUN go mod download
RUN go build -o .bin/rest_server cmd/rest_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/go-jedi/portfolio/app/.bin/rest_server .
COPY .env /root/

CMD ["./rest_server"]