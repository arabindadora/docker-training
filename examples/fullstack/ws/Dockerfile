FROM nginx

COPY src/default.conf /etc/nginx/conf.d/

EXPOSE 80
ENTRYPOINT ["nginx", "-g", "daemon off;"]
