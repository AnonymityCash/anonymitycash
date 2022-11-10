# Build Anonymitycash in a stock Go builder container
FROM golang:1.9-alpine as builder

RUN apk add --no-cache make git

ADD . /go/src/github.com/anonymitycash/anonymitycash
RUN cd /go/src/github.com/anonymitycash/anonymitycash && make anonymitycashd && make anonymitycashcli

# Pull Anonymitycash into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/anonymitycash/anonymitycash/cmd/anonymitycashd/anonymitycashd /usr/local/bin/
COPY --from=builder /go/src/github.com/anonymitycash/anonymitycash/cmd/anonymitycashcli/anonymitycashcli /usr/local/bin/

EXPOSE 1999 46655 46654 1313
