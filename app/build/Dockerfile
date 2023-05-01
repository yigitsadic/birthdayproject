FROM node:19.7-alpine3.17 as builder

WORKDIR /app
COPY package.json package.json
COPY package-lock.json package-lock.json 

RUN npm i 

COPY . .

ARG API_URL
RUN echo "VITE_API_URL=${API_URL}" > .env

RUN npm run build

FROM nginx:stable

RUN rm /etc/nginx/conf.d/default.conf
COPY --from=builder /app/dist /usr/share/nginx/html
COPY build/nginx.conf /etc/nginx/nginx.conf

