FROM golang:alpine as build
ADD go.mod simple-golang-api.go /simple-golang-api/
RUN cd /simple-golang-api && go mod tidy && go build -o /simple-golang-api/simple-golang-api
FROM alpine:latest
COPY --from=build /simple-golang-api/simple-golang-api /simple-golang-api
CMD ["/simple-golang-api"]