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
RUN mkdir www

###################################################
##                   SCRATCH                     ##
###################################################
FROM scratch
WORKDIR /app
COPY --from=build /app/server .
COPY --from=build /app/www .
USER 1000
EXPOSE 8080
ENTRYPOINT ["./server"]
CMD ["0.0.0.0", "8080", "www"]
