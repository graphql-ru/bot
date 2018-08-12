FROM alpine

RUN apk update && apk add ca-certificates
RUN rm -rf /var/cache/apk/*
RUN mkdir opt

ADD bin/bot opt/bot
ENTRYPOINT ["/opt/bot"]
