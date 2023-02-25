# Docker build file for fauxmailer
# https://github.com/inbucket/fauxmailer

# Build
FROM golang:1.20-alpine3.17 as builder
RUN apk add --no-cache ca-certificates git
WORKDIR /build
COPY . .
ENV CGO_ENABLED 0
RUN go build -o fauxmailer

# Run in minimal image
FROM alpine:3.17
WORKDIR /usr/bin
COPY --from=builder /build/fauxmailer .

# Run
CMD ["fauxmailer"]
