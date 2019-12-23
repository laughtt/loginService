
#build stage
FROM golang:latest
WORKDIR /app
RUN chmod 700 /app
COPY . .
ENV GO111MODULE=on
RUN go mod download
#RUN protoc --go_out=plugins=grpc:. api/proto/v1/proto-service.proto
#CMD [ "protoc" , "--go_out=plugins=grpc:." ,"api/proto/v1/proto-service.proto"]
EXPOSE 80
RUN go build pkg/cmd/main.go
CMD [ "./main" ]