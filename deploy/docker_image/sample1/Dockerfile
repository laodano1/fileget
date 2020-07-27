FROM alpine:latest

RUN mkdir -p /home/game/bin

COPY gwapi gwapi.json launch.sh version.txt /home/game/bin/

LABEL VERSION="__VER__" \
      BRANCH="origin/master" \
      COMMIT="__COMIT__" \
      COMPILE="go version go1.13 linux/amd64" \
      DATE="__DAT__"


# add environment variables.
ENV ISDOCKER 1

WORKDIR /home/game

CMD ["sh", "/home/game/bin/launch.sh"]

EXPOSE 80



