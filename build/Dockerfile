FROM alpine:3.6

ARG USER_ID=1000
ARG GROUP_ID=1000

RUN set -ex \
  && apk add --no-cache --update \
    bash \
    tzdata \
    ca-certificates \
  && apk add --update --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/edge/community shadow


RUN set -ex \
  && groupadd -g "${GROUP_ID}" docker \
  && useradd -r -M -N -u "${USER_ID}" -g docker -s /bin/sh -d /home/docker docker \
  && mkdir -p /home/docker \
  && chown -R docker:docker /home/docker


RUN set -ex \
  && apk del --purge git \
  && rm -rf "${GOPATH}" /usr/local/go

COPY --chown=docker:docker ./build/files/docker-entrypoint.sh /
COPY --chown=docker:docker ./bin /
COPY --chown=docker:docker ./src/checker/testdata/url_list.txt /

USER docker

ENTRYPOINT ["/docker-entrypoint.sh"]

CMD ["/service"]
