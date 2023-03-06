FROM alpine:3.17
RUN apk add --no-cache composer php81-tokenizer php81-pecl-swoole go
COPY php.ini /etc/php81/conf.d/php.ini
WORKDIR /opt
EXPOSE 9501
