FROM nginx:1.13.11-alpine

RUN rm -rf /usr/share/nginx/html/*

COPY nginx-default.conf /etc/nginx/conf.d/default.conf 
COPY entrypoint.sh /usr/share/entrypoint.sh

CMD ["nginx", "-g", "daemon off;"]
ENTRYPOINT ["/usr/share/entrypoint.sh"]