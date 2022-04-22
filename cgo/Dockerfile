FROM golang:1.18.1 AS build

WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
RUN go build -o httpd

FROM alpine:3.15
WORKDIR /app
COPY --from=build /app/httpd .
CMD ["./httpd"]
