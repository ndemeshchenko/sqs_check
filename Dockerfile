FROM golang:1.9

RUN mkdir -p /sqs_check
WORKDIR /sqs_check

COPY . /sqs_check
RUN go get github.com/aws/aws-sdk-go/aws
RUN go get github.com/aws/aws-sdk-go/aws/credentials
RUN go get github.com/aws/aws-sdk-go/aws/session
RUN go get github.com/aws/aws-sdk-go/service/sqs

RUN go build main.go

CMD ./main
