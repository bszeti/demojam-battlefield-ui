FROM registry.access.redhat.com/ubi7/ubi

ADD ./battlefield-ui /

ENTRYPOINT ["/battlefield-ui"]