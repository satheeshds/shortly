#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init
#RUN mkdir /go/bin/
RUN go build -o /go/bin/app/ -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app/ /app/
ENTRYPOINT /app/shortly
LABEL Name=shortly Version=0.0.1
EXPOSE 8080
