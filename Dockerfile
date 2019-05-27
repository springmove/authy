FROM alpine:3.6

RUN apk update && apk add curl
RUN mkdir -p /etc/authy/conf /etc/authy/log

COPY ./entrypoint.sh /
COPY ./conf/config.yml /etc/authy/config.yml
#COPY ./etc/authy/api.json /etc/authy
COPY ./build/authy /usr/bin

ENTRYPOINT ["/entrypoint.sh"]
CMD authy --config /etc/authy/conf/config.yml
