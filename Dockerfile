FROM golang:1.15-alpine3.12 AS build-env
RUN apk add build-base linux-headers
ARG VERSION
ENV BUILD_HOME=/BUILD
RUN mkdir -p ${BUILD_HOME}
WORKDIR ${BUILD_HOME}

ADD . ${BUILD_HOME}
RUN go mod download
RUN go build -o /bin/circuit-breaker \
    -ldflags "-X main.Version=$VERSION" \
    ./cmd

FROM alpine:3.12
COPY --from=build-env /bin/circuit-breaker /bin/circuit-breaker
ENTRYPOINT  ["/bin/circuit-breaker", "--config.path", "/circuit-breaker.yml" ]
CMD [ "combined" ]