FROM node:12 as node-builder

WORKDIR /build

COPY ./ui /build

RUN npm install . && npm run build

FROM golang:1.14 AS golang-builder

WORKDIR /build

COPY . /build

RUN go build -o agility --ldflags "-linkmode 'external' -extldflags '-static'" .

FROM alpine:3

WORKDIR /bin

COPY --from=node-builder /build/dist /usr/share/agility/assets
COPY --from=golang-builder /build/agility /bin/agility

EXPOSE 80

ENTRYPOINT [ "/bin/agility" ]
