FROM golang:1.23-alpine AS builder

RUN mkdir -p /app
COPY . /app/
WORKDIR /app

# build
RUN go build -o srv cmd/test/*

FROM alpine

ENV GOOSE_VERSION v3.7.0
ENV USER docker
ENV UID 12345
ENV GID 23456

RUN apk --no-cache add ca-certificates curl

# install goose
ADD https://github.com/pressly/goose/releases/download/$GOOSE_VERSION/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

# copy migrations
COPY db/migrations /app/migrations

# copy executable
COPY --from=builder /app/srv /app/srv

# copy entrypoint
COPY docker/docker-entrypoint.sh /app/docker-entrypoint.sh
RUN chmod +x /app/docker-entrypoint.sh

# use non-root user
RUN addgroup --gid $GID -S $USER
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --ingroup "$USER" \
    --no-create-home \
    --uid "$UID" \
    "$USER"

RUN chown -R ${UID}:${UID} /app/

USER $USER:$USER

ENTRYPOINT ["/app/docker-entrypoint.sh"]