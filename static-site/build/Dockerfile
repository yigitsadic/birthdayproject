FROM node:19.7-alpine3.17 as builder

WORKDIR /app
COPY package.json package.json
COPY package-lock.json package-lock.json 

RUN npm i 

COPY . .

RUN npm run build

FROM nginx:stable

COPY --from=builder /app/dist /usr/share/nginx/html

