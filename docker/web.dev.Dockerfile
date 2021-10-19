FROM node:14

WORKDIR /usr/app

COPY . /usr/app

CMD yarn && yarn dev