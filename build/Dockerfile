FROM alpine:3.11.3

WORKDIR /app
RUN apk update \
 && apk upgrade \
 && apk add --no-cache musl-dev \
 && rm -rf /var/cache/apk/*