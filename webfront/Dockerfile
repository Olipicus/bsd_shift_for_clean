FROM node:slim

WORKDIR /src
ADD . .

RUN npm install -g gulp-cli
RUN npm install
RUN gulp serve:production
