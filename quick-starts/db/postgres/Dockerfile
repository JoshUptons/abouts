FROM postgres:latest

ENV POSTGRES_PASSWORD password

WORKDIR /database/

COPY sql/*.sql /docker-entrypoint-initdb.d/

EXPOSE 5432:5432
