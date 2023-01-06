FROM golang:1.19
WORKDIR /nomaxprocs
ADD . .
RUN go install ./
ENTRYPOINT nomaxprocs
