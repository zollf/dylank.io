FROM nginx:latest

COPY ./config/prod/nginx.conf /etc/nginx/nginx.conf