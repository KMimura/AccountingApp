FROM nginx

COPY ./nginx/website.conf /etc/nginx/conf.d
RUN rm /etc/nginx/conf.d/default.conf
RUN mkdir -p /var/www/accounting
ADD ./website /var/www/accounting
