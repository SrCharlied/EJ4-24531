FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/data ./data

ENV PORT=24531
EXPOSE 24531

CMD [ "./server" ]