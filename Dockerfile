FROM alpine:latest

COPY www .
COPY conf conf
COPY views views
COPY static static

CMD [ "./www"]