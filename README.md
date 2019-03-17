# www
Website for avosa

This is the website man.

## Run from Docker
* $ docker build -t avosa/www:latest .
* $ docker rm WWWDEV
* $ docker run -d -e RUNMODE=DEV -p 8091:8091 --network mango_net --name WWWDEV avosa/www:latest
* $ docker logs WWWDEV

## Run with docker-compose
* docker-compose up --build -d
* docker-compose logs