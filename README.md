# www
Website for avosa

This is the website man.

## Run from Docker
* $ GOOS=linux GOARCH=amd64 go build
* $ gulp
* $ docker build -t avosa/www:dev .
* $ docker rm wwwDEV
* $ docker run -d -p 8091:8091 --network mango_net --name wwwDEV avosa/www:dev
* $ docker logs wwwDEV