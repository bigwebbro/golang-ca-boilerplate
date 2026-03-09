FROM golang:1.25.4-alpine

COPY --from=cosmtrek/air /go/bin/air /usr/local/bin/air

WORKDIR /app

ARG USER_ID
ARG GROUP_ID
USER ${USER_ID}:${GROUP_ID}

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
