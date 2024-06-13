# syntax=docker/dockerfile:1

FROM node:18.8-alpine3.16 as front-end

WORKDIR /app
COPY frontend/package.json ./
COPY frontend/yarn.lock ./
RUN yarn install --frozen-lockfile
COPY ./frontend /app

RUN yarn build

FROM golang:1.21.0
COPY --from=front-end /app/build/ /app/build/
COPY /server/ /app/server/
WORKDIR /app/server/
RUN go mod download
RUN go build -o demyst-v1 -ldflags "-s -w"

EXPOSE $APP_PORT

CMD ["/app/server/demyst-v1"]