FROM golang:1.22.2-alpine3.19 AS BUILDER
WORKDIR /app
COPY ./ /app
RUN go build 

FROM golang:1.22.2-alpine3.19
WORKDIR /app
COPY --from=BUILDER /app/test-github-action /app/test-github-action
EXPOSE 8080
ENTRYPOINT [ "./test-github-action" ]
