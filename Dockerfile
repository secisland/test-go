FROM centos:latest
LABEL maintainer="body"


ENV DEMO_PATH /root/demo
ENV PATH ${DEMO_PATH}:${PATH}

RUN mkdir -p ${DEMO_PATH}
COPY demo ${DEMO_PATH}
WORKDIR ${DEMO_PATH}

CMD ["demo"]