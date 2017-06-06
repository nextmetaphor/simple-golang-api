FROM alpine:latest
ADD simple-golang-api /tmp
ENTRYPOINT /tmp/simple-golang-api