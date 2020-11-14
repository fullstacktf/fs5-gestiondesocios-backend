FROM mariadb:10.5.8-focal

ADD scripts/ /docker-entrypoint-initdb.d

ENV MYSQL_ROOT_PASSWORD test123
ENV MYSQL_DATABASE testDB
ENV MYSQL_USER toto
ENV MYSQL_PASSWORD test123

#RUN apt-get update

EXPOSE 3306

CMD ["mysqld"]
