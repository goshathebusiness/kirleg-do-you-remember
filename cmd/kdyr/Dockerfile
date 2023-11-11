FROM golang:1.21.4 as build
WORKDIR /code
COPY .. /code
RUN go build -buildvcs=false -o ./bin ./cmd/kdyr

#############
FROM alpine:3.15
WORKDIR /bin
COPY --from=build /code/bin/ ./

ENTRYPOINT ["./kdyr"]