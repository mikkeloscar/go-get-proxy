FROM alpine:latest
MAINTAINER Mikkel Oscar Lyderik Larsen <m@moscar.net>

# add scm-source
ADD scm-source.json /

# add binary
ADD build/linux/go-get-proxy /

ENTRYPOINT ["/go-get-proxy"]
