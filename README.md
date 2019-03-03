# www
Website for avosa

This is the website man.

## Run from Docker
* $ GOOS=linux GOARCH=amd64 go build
* $ gulp
* $ docker build -t avosa/www:latest .
* $ docker rm WWWDEV
* $ docker run -d -e RUNMODE=DEV -p 8091:8091 --network mango_net --name WWWDEV avosa/www:latest
* $ docker logs WWWDEV