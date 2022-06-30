FROM node:16

WORKDIR /usr/app

COPY /frontend .

RUN yarn
RUN yarn build

CMD yarn start