# arg-example
A simple tool to demonstrate arguments written in golang

You can use the Dockerfile to build a container running the arg-example.

## Build
´´´Bash
go build
sudo docker build -t johscheuer/arg-example .
´´´

## Run
´´´Bash
$ sudo docker run -ti  johscheuer/arg-example -name Bert -street "Sesame Street"
I'm Bert from the Sesame Street
´´´
