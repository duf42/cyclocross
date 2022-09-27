FROM danger89/cmake as library

ARG LIBRARY_SOURCE_FOLDER

WORKDIR /src

COPY $LIBRARY_SOURCE_FOLDER .

RUN make

FROM golang:1.14.3-alpine AS base

# Install GCC
RUN apk add build-base libc6-compat
RUN mkdir /deps

WORKDIR /src

ENV CGO_ENABLED=1
ENV GOPATH=/
ENV GOOS=linux 
ENV GOARCH=amd64
ENV CC=gcc

COPY --from=library /src/lib /deps
COPY ./sources .
COPY ./tests .

FROM base AS build

RUN go build -o /out/server .

FROM base AS unit-test
ENV LD_LIBRARY_PATH=/deps
ENTRYPOINT ["go", "test"]
CMD ["-v", "./..."]

FROM alpine:latest

RUN mkdir /config && mkdir /deps && mkdir /install && mkdir /web
ENV LD_LIBRARY_PATH=/deps

COPY --from=build /out/server /install/server
COPY --from=build /deps /deps

COPY ./web /web
COPY VERSION /config/VERSION

CMD ["/install/server"]
