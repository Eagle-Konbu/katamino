FROM node:16 as builder

WORKDIR /myapp

COPY package*.json .

RUN npm install && \
    npm run build

FROM nginx:stable-alpine

COPY --from=builder /myapp/build /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]
