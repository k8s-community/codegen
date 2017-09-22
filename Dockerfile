FROM scratch

ENV CODEGEN_LOCAL_HOST 0.0.0.0
ENV CODEGEN_LOCAL_PORT 8080
ENV CODEGEN_LOG_LEVEL 0

EXPOSE $CODEGEN_LOCAL_PORT

COPY certs /etc/ssl/
COPY bin/linux-amd64/codegen /

COPY templates /templates
COPY _code-templates /code-templates
COPY static /static

CMD ["/codegen"]