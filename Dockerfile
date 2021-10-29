#build stage
FROM golang:alpine AS builder
# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --no-cache git

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./app ./cmd/api/main.go
WORKDIR /app

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /build/app /app/restapi
COPY --from=builder /build/config /config
LABEL Name=gorray Version=0.0.1
EXPOSE 3000
ENTRYPOINT ["/app/restapi"]