FROM node:16-alpine as builder

WORKDIR /build
COPY . .

RUN yarn install --frozen-lockfile && \
    NODE_ENV=production NODE_OPTIONS="--max_old_space_size=8192" yarn build


FROM node:16-alpine

RUN apk upgrade --update-cache --available && \
    apk add tzdata

ENV TZ=Asia/Tokyo
ENV NODE_ENV=production
WORKDIR /app

COPY --from=builder /build/package.json /build/yarn.lock ./
COPY --from=builder /build/.next ./.next
COPY --from=builder /build/node_modules ./node_modules

CMD yarn start