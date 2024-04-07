FROM alpine:3.19.1
RUN --mount=type=secret,id=secret-file,target=/run/secrets/secret.txt cat /run/secrets/secret.txt
