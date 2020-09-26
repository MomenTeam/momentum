FROM golang:latest AS build-env
COPY . /src
RUN cd /src && CGO_ENABLED=0 go build -o main

FROM alpine:latest
WORKDIR /app
COPY --from=build-env /src/main .
COPY --from=build-env /src/configs configs
CMD [ "./main" ]
