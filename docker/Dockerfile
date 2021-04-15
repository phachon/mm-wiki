FROM golang:alpine AS build

LABEL maintainer="chaiyd<chaiyd.cn@gmail.com>"

# aliyun mirros && time zone
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update

WORKDIR /app

RUN apk add git \
    && git clone https://github.com/phachon/mm-wiki.git \
    && cd mm-wiki \
    && sh build.sh


FROM alpine:3.13

LABEL maintainer="chaiyd<chaiyd.cn@gmail.com>"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk add tzdata dumb-init \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

ENV HTTP_PORT=8081 \
    DB_HOST=127.0.0.1 \
    DB_PORT=3306 \
    DB_NAME=mm_wiki \
    DB_USER=mm_wiki \
    DB_PASS=123456 \
    COOKIE=2160000

COPY --from=build /app/mm-wiki/release /app/mm-wiki

WORKDIR /app/mm-wiki
VOLUME /app/mm-wiki/data

ADD start.sh /app/mm-wiki/
RUN chmod +x start.sh
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["sh", "-c", "./start.sh"]