FROM alpine:3.12.0

COPY www .
COPY build/*.dart.js dist/js/
COPY views views

RUN mkdir -p /views/_shared

EXPOSE 8091

ENTRYPOINT [ "./www" ]
