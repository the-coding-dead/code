FROM ubuntu:latest
RUN apt-get update \
 && apt-get install -y --no-install-recommends \
    gcc \
    make \
    binutils \
    libc6-dev \
    gdb \
 && apt-get clean -y \
 && rm -rf /var/lib/apt/lists/*
