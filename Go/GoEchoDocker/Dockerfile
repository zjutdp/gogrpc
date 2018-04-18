FROM alpine:3.6
RUN apk add bash
ENV WORK_DIR=/docker/bin/
WORKDIR $WORK_DIR
COPY ./go_docker /docker/bin/
ENTRYPOINT /docker/bin/go_docker
EXPOSE 8080