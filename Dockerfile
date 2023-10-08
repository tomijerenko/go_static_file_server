ARG ALPINE_VERSION=3.18
ARG GOLANG_VERSION=1.21
ARG COMPILER=golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION}
ARG CGO_ENABLED=0

###################################################
##                    COMPILE                    ##
###################################################
FROM ${COMPILER} AS build
WORKDIR /app
RUN apk update && apk upgrade --no-cache
RUN apk add --no-cache make
COPY go.mod .
RUN go mod download
RUN go mod verify
COPY  . .
RUN make build-release

###################################################
##                    SETUP                      ##
###################################################
FROM alpine:${ALPINE_VERSION}
ENV USER_ID=65535
ENV GROUP_ID=65535
ENV USER_NAME=container
ENV GROUP_NAME=container
RUN addgroup -g $GROUP_ID $GROUP_NAME && \
    adduser --shell /sbin/nologin --disabled-password \
    --no-create-home --uid $USER_ID --ingroup $GROUP_NAME $USER_NAME
RUN apk update && apk upgrade --no-cache
WORKDIR /app
RUN mkdir www
COPY --from=build /app/server .
EXPOSE 8080
USER $USER_NAME:$GROUP_NAME
ENTRYPOINT ["./server"]
CMD ["0.0.0.0", "8080", "www"]
