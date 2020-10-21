FROM latonaio/pylib-lite:latest

# Definition of a Device & Service
ENV SERVICE=sample-microservice

ADD . .
RUN python3 setup.py install

CMD ["/bin/sh", "docker-entrypoint.sh"]
