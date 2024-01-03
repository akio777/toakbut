FROM postgres:15.2

COPY migrations/*.sql /docker-entrypoint-initdb.d/


ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB toakbut
ENV POSTGRES_PORT 5432

EXPOSE 5432