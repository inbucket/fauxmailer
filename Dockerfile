# Docker build file for fauxmailer
# https://github.com/jhillyerd/fauxmailer

# Build
FROM golang:1.9-alpine as builder
RUN apk add --no-cache ca-certificates git
WORKDIR /go/src/app
COPY . .
RUN go-wrapper download
RUN go-wrapper install

# Run in minimal image
FROM alpine
RUN apk add --no-cache ca-certificates
WORKDIR /usr/bin
COPY --from=builder /go/bin/fauxmailer .

# Run
CMD ["fauxmailer"]
