From centos
ADD  ./gin-web /
COPY ./config.yml $GOPATH/src/gin-web/config/
RUN chmod +x  /gin-web
ENTRYPOINT ["./gin-web"]