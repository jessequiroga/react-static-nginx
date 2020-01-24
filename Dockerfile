ARG NGINX_IMAGE

# Build Entry point execuable
FROM golang:1.13.6-alpine3.11 as builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o env-replacer cmd/entrypoint/main.go
RUN chmod +x env-replacer

# Build Image
FROM $NGINX_IMAGE

RUN rm -rf /usr/share/nginx/html/*

COPY nginx-default.conf /etc/nginx/conf.d/default.conf 
COPY entrypoint.sh /usr/share/entrypoint.sh
COPY __deregister-service-worker.html /usr/share/nginx/html/
COPY --from=builder /app/env-replacer /usr/local/bin/

CMD ["nginx", "-g", "daemon off;"]
ENTRYPOINT ["/usr/share/entrypoint.sh"]