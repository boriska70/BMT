FROM alpine:3.3

COPY .dist/bmt /usr/bin/bmt

CMD ["/usr/bin/bmt"]