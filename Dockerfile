FROM golang:1.18 as build
MAINTAINER lizhang

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app
ADD . ./
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o app main.go


FROM alpine as prod
COPY --from=build /app/application.toml /config/application.toml
COPY --from=build /app/templates /templates
COPY --from=build /app/static /static
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

COPY --from=build /app/app /
ENTRYPOINT ./app