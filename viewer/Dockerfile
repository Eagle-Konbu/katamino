FROM node:16 as builder

WORKDIR /myapp

COPY . .

RUN npm install && \
    npm run build

FROM nginx:stable-alpine

COPY --from=builder /myapp/build /var/www
COPY ./nginx/default.conf /etc/nginx/conf.d/

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
