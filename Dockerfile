FROM scratch
ADD alpine/alpine-minirootfs-3.10.3-x86_64.tar.gz /
CMD ["/bin/sh"]
