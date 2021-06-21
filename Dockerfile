FROM golang:alpine AS build
LABEL Author="Daniela Osorio R <danosori@bancolombia.com.co>"

ENV NEQUI_CREDITS_TABLE_NAME="credit-customer-product-qa"

WORKDIR /go/src/myapp
COPY . .

EXPOSE 9443
RUN go build -o /go/bin/myapp cmd/api/main.go

FROM scratch
COPY --from=build /go/bin/myapp /go/bin/myapp
ENTRYPOINT ["/go/bin/myapp"]