FROM golang:1.16-buster as base
  WORKDIR /app
  ENV PORT=80 \
    APP_NAME=trellenge-go

FROM base as moduler
  COPY . .
  RUN go mod vendor

FROM moduler as builder
  RUN go build -o build/${APP_NAME} cmd/main.go

FROM moduler as development
  RUN GO111MODULE=on go get github.com/cortesi/modd/cmd/modd

  CMD ["modd"]

FROM base as production
  COPY --from=builder /app/build/${APP_NAME} bin/${APP_NAME}

  CMD [ "./bin/trellenge-go" ]
