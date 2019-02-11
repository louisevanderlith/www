# www
Website for avosa

This is the website man.

## Run from Docker
*$ go build
*$ docker build -t avosa/www:dev .
*$ docker rm wwwDEV
*$ docker run -d --network host --name wwwDEV avosa/www:dev
*$ docker logs wwwDEV