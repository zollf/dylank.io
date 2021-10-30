FROM node:14.18

WORKDIR /usr/app

COPY . /usr/app

CMD yarn --production && yarn dev