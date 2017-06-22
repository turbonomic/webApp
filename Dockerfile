FROM php:7.0-apache
COPY www/* /var/www/html/
COPY ./conf/status.conf /etc/apache2/mods-available/
COPY ./conf/charset.conf /etc/apache2/conf-available/

#change default port from 80 to 8080
EXPOSE 8080
COPY ./conf/ports.conf /etc/apache2/
COPY ./conf/000-default.conf /etc/apache2/sites-available/

#RUN mkdir /var/run/apache2/
RUN chmod 777 /var/run/apache2/

#RUN mkdir /var/log/apache2
RUN chmod 777 /var/log/apache2/

#RUN mkdir /var/lock/apache2
RUN chmod 777 /var/lock/apache2

