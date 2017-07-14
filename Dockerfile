FROM alpine:latest
ADD simple-golang-api /tmp
CMD ["/tmp/simple-golang-api"]