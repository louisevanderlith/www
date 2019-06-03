# www
Website for avosa

This is the website man.

## Run from Docker
* $ docker build -t avosa/www:dev .
* $ docker rm WWWDEV
* $ docker run -d -p 8091:8091 --network mango_net --name WWWDEV avosa/www:dev
* $ docker logs WWWDEV

## Run with docker-compose
* docker-compose up --build -d
* docker-compose logs