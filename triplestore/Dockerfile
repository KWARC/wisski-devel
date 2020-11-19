# This Dockerfile contains instructions to compile and run GraphDB inside a Docker container. 
# It is roughly based on https://github.com/Ontotext-AD/graphdb-docker/blob/master/free-edition/Dockerfile
# but has been modified for performance and security. 

# This image is intended to be built like:
#   docker build --build-arg graphdb_src=graphdb.zip .

# We first compile the healthchecker into an executable. 
# This will be used further down the line.  
FROM golang:1-buster as healthcheck
ADD healthcheck.go /app/
RUN go build -o /healthcheck /app/healthcheck.go

# Next make a base image to base further builds on. 
# We don't use alpine here, as that uses significantly slower musl instead of glibc. 
FROM adoptopenjdk/openjdk11:debian-slim as base

# Create a user called graphdb
RUN useradd -ms /bin/bash graphdb

# make a base images, to add the sources to. 
FROM base as sources

# install unzip
RUN apt-get update && apt-get install -y unzip

# add the source file (by default graphdb.zip) to the image
ARG src=graphdb.zip
ADD ${src} /graphdb.zip

# unpack it into a temporary directory
RUN unzip "$src" -d "/unpack/"

# Move it into /opt/graphdb, and chown it to graphdb
RUN mv "/unpack"/* /opt/graphdb
RUN chown -R graphdb:graphdb /opt/graphdb

# finally make an image that will run
FROM base as final

# copy over the sources
COPY --from=sources /opt/graphdb /opt/graphdb

# set environment variables for graphdb_home and path
ENV GRAPHDB_HOME=/opt/graphdb
ENV PATH=$GRAPHDB_HOME/bin:$PATH

# expose a port
EXPOSE 7200

# setup a healthcheck, that uses the compiled executable from above.
COPY --from=healthcheck /healthcheck /healthcheck
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD /healthcheck

# Create directories for data, work and logs and chown it to the graphdb user
RUN mkdir -p /opt/graphdb/data && chown graphdb:graphdb /opt/graphdb/data && \
    mkdir -p /opt/graphdb/work && chown graphdb:graphdb /opt/graphdb/work && \
    mkdir -p /opt/graphdb/logs && chown graphdb:graphdb /opt/graphdb/logs

# make volumes for them
VOLUME /opt/graphdb/data
VOLUME /opt/graphdb/work
VOLUME /opt/graphdb/logs

# setup command and entrypoint
USER graphdb:graphdb
CMD ["/opt/graphdb/bin/graphdb", "-Dgraphdb.home=/opt/graphdb"]
