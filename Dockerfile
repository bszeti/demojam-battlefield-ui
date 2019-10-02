FROM registry.access.redhat.com/ubi7/ubi

ADD ./battlefield-ui /
ADD ./static /static

EXPOSE 8080

ENTRYPOINT ["/battlefield-ui"]