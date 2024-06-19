FROM golang:1.22.4-bullseye as build
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
RUN git clone https://github.com/rvanderp3/lmsensors-exporter.git
COPY . .
RUN go build cmd/exporter.go

FROM ubuntu:24.04
RUN apt update
RUN apt install -y lm-sensors
COPY --from=build /usr/src/app/exporter .
CMD exporter