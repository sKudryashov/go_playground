FROM golang:1.14

RUN mkdir ./graphsrv
COPY . ./graphsrv

ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on

WORKDIR /go/graphsrv
#race detector enabled
RUN cd ./cmd/servd && go build -race -o graphsrv grpc.go
#race detector disabled
#RUN cd ./cmd/servd && go build -o graphsrv grpc.go
RUN chmod 0766 ./cmd/servd/graphsrv

CMD [ "./cmd/servd/graphsrv", "-port=50051" ]