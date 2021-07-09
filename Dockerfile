FROM ubuntu:16.04

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && apt-get install -y \
    autoconf automake build-essential curl git libsnappy-dev libtool pkg-config

RUN git clone https://github.com/openvenues/libpostal

EXPOSE 8080

WORKDIR /libpostal
RUN ./bootstrap.sh && mkdir -p /opt/libpostal_data && ./configure --datadir=/opt/libpostal_data && make && make install && ldconfig

WORKDIR /go_package
ADD https://golang.org/dl/go1.16.5.linux-amd64.tar.gz .
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

WORKDIR /corrector
COPY . .
RUN go build cmd/corrector/main.go
RUN cp main binary/

WORKDIR binary/

CMD ["./main"]
