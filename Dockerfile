FROM alpine/git

ENV TZ=Asia/Shanghai

WORKDIR /app

RUN git clone https://github.com/phachon/mm-wiki.git


FROM golang:1.14.1-alpine

COPY --from=0 /app/mm-wiki /app/mm-wiki

WORKDIR /app/mm-wiki

# 如果国内网络不好，可添加以下环境
# RUN go env -w GO111MODULE=on
# RUN go env -w GOPROXY=https://goproxy.cn,direct
# RUN export GO111MODULE=on
# RUN export GOPROXY=https://goproxy.cn

RUN mkdir /opt/mm-wiki && ls /app/mm-wiki
RUN go build -o /opt/mm-wiki/mm-wiki ./ \
    && cp -r ./conf/ /opt/mm-wiki \
    && cp -r ./install/ /opt/mm-wiki\
    && cp ./scripts/run.sh /opt/mm-wiki\
    && cp -r ./static/ /opt/mm-wiki\
    && cp -r ./views/ /opt/mm-wiki\
    && cp -r ./logs/ /opt/mm-wiki\
    && cp -r ./docs/ /opt/mm-wiki
CMD ["/opt/mm-wiki/mm-wiki", "--conf", "/opt/mm-wiki/conf/mm-wiki.conf"]