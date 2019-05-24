FROM node:11.14-stretch-slim

WORKDIR /usr/src/app
RUN npm i --registry=https://registry.npm.taobao.org
COPY . .
CMD ["node","main.js"]