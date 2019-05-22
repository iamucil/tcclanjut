FROM golang:1.12 as builder

WORKDIR /github.com/iamucil/goweb/

COPY go.mod go.sum ./

RUN go mod download

# Copy the resources
COPY . .
# Build it
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/service \
  ./main.go

# Build the runtime image
FROM alpine:3.8
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy static file into containers
COPY ./resources/assets ./resources/assets

COPY --from=builder /github.com/iamucil/goweb/bin/service ./service

EXPOSE 8080

ENTRYPOINT [ "./service" ]
