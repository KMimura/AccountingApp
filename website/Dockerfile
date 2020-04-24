FROM nginx:alpine

COPY ./nginx/website.conf /etc/nginx/conf.d
RUN rm /etc/nginx/conf.d/default.conf
WORKDIR mkdir -p /var/www/accounting
COPY ./website /var/www/accounting
