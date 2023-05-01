FROM node:19.7-alpine3.17

WORKDIR /app
COPY package.json package.json
COPY package-lock.json package-lock.json 

RUN npm i 

COPY . .

RUN npm run build

CMD npm run start

