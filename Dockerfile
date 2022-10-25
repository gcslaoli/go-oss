FROM alpine
# Install dependencies
ARG TARGETARCH
RUN echo "I'm building for $TARGETARCH"
ENV TZ                      Asia/Shanghai
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && apk add tzdata  bash

###############################################################################
#                                INSTALLATION
###############################################################################
EXPOSE 8000

ENV WORKDIR                 /app
VOLUME [ "${WORKDIR}/upload/" ]
ADD resource                $WORKDIR/
COPY ./temp/linux_amd64/main $WORKDIR/main
RUN chmod +x $WORKDIR/main

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main
