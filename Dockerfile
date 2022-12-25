FROM golang:1.19.0-alpine3.16 as builder

ARG VERSION="devel"

WORKDIR /app
COPY . /app/
ENV VERSION=$VERSION
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV CGO_ENABLED=0

## UPX package is used for compressing binaries file to reduce build size
# RUN apk update && apk --no-cache add upx
RUN mkdir -p bin \
    && echo "Build version $VERSION" \
    && export SRC_APP=$(pwd)/cmd \
    && go build -ldflags="-s -w -X main.version=$VERSION" -o ./bin/app $SRC_APP 

RUN upx --force ./bin/app

FROM alpine:3.16.2

COPY --from=builder /app/bin/app /usr/bin/app 

ENTRYPOINT [ "/usr/bin/app"]