FROM gcr.io/tetratelabs/tetrate-base:v0.1
ADD build/sink /usr/local/bin/sink
ENTRYPOINT ["/usr/local/bin/sink"]
