FROM golang:1.25.4-alpine AS builder

WORKDIR /app

COPY go.work go.work.sum ./
COPY cmd/http/api/go.mod ./cmd/http/api/
COPY internal/user/adapter/go.mod ./internal/user/adapter/
COPY internal/user/application/go.mod ./internal/user/application/
COPY internal/user/domain/go.mod ./internal/user/domain/
COPY pkg/go.mod ./pkg/

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/http/api

FROM alpine:3.19

WORKDIR /app

ARG USER=appuser
ARG GROUP=appgroup

COPY --from=builder /app/api .

EXPOSE 8080

RUN addgroup ${GROUP} && \
    adduser -D -G ${GROUP} ${USER} && \
    chown ${USER}:${GROUP} /app/api && \
    chmod +x /app/api

USER ${USER}

CMD ["./api"]
