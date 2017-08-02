FROM golang:1.8
WORKDIR /go
RUN go get github.com/ivanthescientist/tournament_service
CMD /go/bin/tournament_service