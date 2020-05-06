FROM alpine:edge
RUN apk add --update --no-cache terraform go git
#COPY . /action
#WORKDIR /action
ENTRYPOINT ["/bin/sh", "-c", "ls -la; pwd; go run ." ]
