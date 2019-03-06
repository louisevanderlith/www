FROM alpine:latest

COPY www .
COPY conf conf
COPY views views
COPY dist dist

##Download the latest templates
RUN mkdir -p /views/_shared
RUN apk add --update curl && rm -rf /var/cache/apk/*
RUN apk --no-cache add jq
RUN for k in $(curl -XGET 172.18.0.1:8093/v1/asset/html | jq -r ".Data | .[]"); do curl -o views/_shared/$k 172.18.0.1:8093/v1/asset/html/$k; done

EXPOSE 8091

ENTRYPOINT [ "./www" ]