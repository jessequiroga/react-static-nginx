FROM nginx:1.15.10-alpine

RUN rm -rf /usr/share/nginx/html/*

COPY nginx-default.conf /etc/nginx/conf.d/default.conf 
COPY entrypoint.sh /usr/share/entrypoint.sh
COPY __deregister-service-worker.html /usr/share/nginx/html/

CMD ["nginx", "-g", "daemon off;"]
ENTRYPOINT ["/usr/share/entrypoint.sh"]