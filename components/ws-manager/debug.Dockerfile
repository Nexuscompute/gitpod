FROM alpine:3.17

# Ensure latest packages are present, like security updates.
RUN  apk upgrade --no-cache \
  && apk add --no-cache ca-certificates

COPY ws-manager /app/ws-manager
ENTRYPOINT [ "/app/ws-manager" ]
CMD [ "-v", "help" ]