FROM scratch

ENV {[( .EnvPrefix )]}_LOCAL_HOST 0.0.0.0
ENV {[( .EnvPrefix )]}_LOCAL_PORT 8080
ENV {[( .EnvPrefix )]}_LOG_LEVEL 0

EXPOSE ${[( .EnvPrefix )]}_LOCAL_PORT

COPY certs /etc/ssl/
COPY bin/linux-amd64/{[( .ServiceName )]} /

CMD ["/{[( .ServiceName )]}"]
