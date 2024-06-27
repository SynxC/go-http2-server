FROM golang:1.22.2 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /goapp
FROM alpine:latest
RUN apk update && apk add --no-cache tshark
COPY --from=builder /goapp /goapp
EXPOSE 8080
CMD ["/goapp"]