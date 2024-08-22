#docker file to build this application
FROM golang:1.22.2-alpine AS builder

WORKDIR /src

# Restore dependencies
COPY . .

RUN go mod tidy

# Build executable
RUN go build -o /src/api-go ./cmd/webserver/main.go

FROM scratch
WORKDIR /src
COPY --from=builder /src/api-go ./
EXPOSE 1337
CMD ["/src/api-go"]
