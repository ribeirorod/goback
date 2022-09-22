FROM postgres:13

ENV POSTGRES_USER=postgres
ENV POSTGRES_DB=postgres


COPY init_db.sh /docker-entrypoint-initdb.d/
COPY postgres.sql /docker-entrypoint-initdb.d/


