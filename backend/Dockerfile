# Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o mm-wiki .

# Runtime stage
FROM alpine:3.18

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/mm-wiki .
COPY --from=builder /app/conf ./conf

EXPOSE 8088

ENTRYPOINT ["./mm-wiki"]
CMD ["-conf", "./conf/prod/app.yaml"]
