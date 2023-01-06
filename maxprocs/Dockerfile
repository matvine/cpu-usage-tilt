FROM golang:1.19
WORKDIR /maxprocs
ADD . .
RUN go install ./
ENTRYPOINT maxprocs
