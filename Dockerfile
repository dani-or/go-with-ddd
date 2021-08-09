#docker image build -t new .
#docker run --network host -d new
FROM golang:alpine AS build
RUN apk add --update git
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir -p app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app1 cmd/api/main.go

FROM scratch
#esto se tieene que hacer https://medium.com/the-go-journey/x509-certificate-signed-by-unknown-authority-running-a-go-app-inside-a-docker-container-a12869337eb
##en alpine bastaría con esto RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/* pero como es from scratch
ADD resources/test.crt /etc/ssl/certs/
ENV NEQUI_CREDITS_TABLE_NAME=credit-customer-product-qa

COPY --from=build /app1 /app
EXPOSE 8080
ENTRYPOINT ["/app"]