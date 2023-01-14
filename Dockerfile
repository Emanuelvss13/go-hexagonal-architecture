FROM golang:1.16


WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN go get -u github.com/spf13/cobra@latest
RUN go install github.com/spf13/cobra-cli@latest
RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN go get github.com/stretchr/testify/require

RUN apt-get update && apt-get install sqlite3 -y

CMD ["tail", "-f", "/dev/null"]