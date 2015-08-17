FROM scratch
MAINTAINER Johannes Scheuermann<johannes.scheuermann@inovex.de>
ADD arg-example arg-example
ENTRYPOINT ["/arg-example"]
