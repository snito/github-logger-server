FROM golang

RUN go get github.com/oxequa/realize
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/jinzhu/gorm