FROM node:16

WORKDIR /usr/app

COPY . /usr/app

CMD yarn && yarn dev