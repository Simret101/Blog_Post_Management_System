FROM golang:1.23-alpine AS builder

WORKDIR /workspace
COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o uuid main.go

# Use a minimal base image for the runtime
FROM alpine:latest
RUN apk update && apk add --no-cache bash curl jq
WORKDIR /
COPY --from=builder /workspace/uuid .

CMD ["./uuid"]
