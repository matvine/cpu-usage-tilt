FROM golang:1.19
WORKDIR /app
ADD . .
RUN go install ./
ENTRYPOINT cpu-usage-tilt
