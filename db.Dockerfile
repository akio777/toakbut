FROM postgres:15.2

RUN apt-get update && apt-get install -y tzdata
ENV TZ=Asia/Bangkok
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY migrations/*.sql /docker-entrypoint-initdb.d/


ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB toakbut
ENV POSTGRES_PORT 5432
ENV POSTGRES_HOST localhost

EXPOSE 5432