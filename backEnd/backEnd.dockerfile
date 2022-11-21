FROM golang:1.19.3

ADD . /backEnd

WORKDIR /backEnd
#RUN go mod init backEnd
RUN go mod tidy
RUN go build -o main .
#RUN CHMOD +x /backEnd

