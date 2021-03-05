#build
FROM golang:1.15.6-buster as builder
COPY . /unico
RUN ls
WORKDIR /unico
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o unico ./cmd/unico/main.go

#execute
FROM alpine:latest
WORKDIR /root/
COPY --from=builder unico .
EXPOSE 3000
CMD ["./unico"]
