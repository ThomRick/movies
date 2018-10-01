FROM node:carbon-alpine as dependencies
WORKDIR /bin/movies/webapp
COPY package.json package.json
COPY package-lock.json package-lock.json
RUN npm install

FROM node:carbon-alpine as builder
WORKDIR /bin/movies/webapp
COPY --from=dependencies /bin/movies/webapp .
COPY public public
COPY src src
RUN npm run -s build

FROM node:carbon-alpine
RUN npm install -g serve
WORKDIR /bin/movies
COPY --from=builder /bin/movies/webapp/build .
CMD serve -l $PORT -s .
