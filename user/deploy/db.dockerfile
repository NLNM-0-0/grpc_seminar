FROM mysql:8.0.41-bookworm

COPY user/deploy/db.sql /docker-entrypoint-initdb.d/

CMD ["mysqld"]
