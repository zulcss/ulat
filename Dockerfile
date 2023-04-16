FROM debian:bookworm

RUN apt-get update -y && \
    apt-get install -y build-essential \
		       ca-certificates \
		       git \
		       golang \
		       mmdebstrap \
		       ostree \
		       systemd-standalone-tmpfiles \
		       vim 
ADD hooks/ostree /usr/share/mmdebstrap/hooks/ostree
RUN mkdir -p /workspace /artifacts
RUN export PATH="$PATH:/worspace/bin"
WORKDIR /artifacts
