FROM alpine
MAINTAINER Julian Kleinhans <julian.kleinhans@aoe.com>

RUN apk --no-cache add ca-certificates && update-ca-certificates
RUN mkdir -p /usr/bin
ADD k8s-ingress-linklist /usr/bin/
RUN chmod +x /usr/bin/k8s-ingress-linklist
ENV PATH $PATH:/usr/bin

WORKDIR /app
COPY src/internal/frontend /app/internal/frontend
COPY src/web /app/web

EXPOSE 8080

CMD ["/usr/bin/k8s-ingress-linklist"]