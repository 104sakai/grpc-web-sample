FROM mariadb:10.9-rc
COPY ./init.sql /docker-entrypoint-initdb.d/
VOLUME [ "/var/lib/mysql" ]