FROM node:14.18

WORKDIR /usr/app

COPY /frontend /usr/app

RUN yarn
RUN yarn build

CMD yarn start