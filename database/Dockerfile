FROM postgres:latest
ADD ./init/01-createtable.sql /docker-entrypoint-initdb.d/
ADD ./init/02-setupdb.sh /docker-entrypoint-initdb.d/
