FROM golang:1.13.8-alpine3.11
RUN mkdir /hfcp-rack-backend
ADD . /hfcp-rack-backend
WORKDIR /hfcp-rack-backend
RUN go build -o server .
EXPOSE 80
CMD ["/hfcp-rack-backend/server"]