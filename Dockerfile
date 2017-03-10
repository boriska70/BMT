FROM alpine:3.3

COPY .dist/bmt /usr/bin/bmt
COPY ./queries.yml /usr/

CMD ["/usr/bin/bmt"]