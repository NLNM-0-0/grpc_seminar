FROM mysql:8.0

COPY order/deploy/db.sql /docker-entrypoint-initdb.d/

CMD ["mysqld"]
