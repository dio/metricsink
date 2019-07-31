FROM alpine:latest
ADD build/sink /usr/local/bin/sink
ENTRYPOINT ["/usr/local/bin/sink"]
