FROM mysql:8.0

COPY product/deploy/db.sql /docker-entrypoint-initdb.d/

CMD ["mysqld"]
