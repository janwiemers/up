FROM golang:1.15.2-alpine3.12 as builder

WORKDIR /build
COPY . .

RUN apk add upx make gcc libc-dev

ENV GO111MODULE=auto
ENV CGO_ENABLED: 0
ENV GOOS: linux
RUN make build.server
RUN upx -9 server

FROM golang:1.15.2-alpine3.12

WORKDIR /app

RUN addgroup -g 1000 -S app && \
  adduser -u 1000 -S app -G app && \
  date -u > BUILD_TIME

COPY --from=Builder --chown=app:app /build/server /app/server
COPY --from=Builder --chown=app:app /build/env /app/env
COPY --from=Builder --chown=app:app /build/config.example.yaml /app/config.example.yaml
RUN chown -R app:app /app
USER app
CMD ["./up"]
