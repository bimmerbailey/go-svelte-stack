FROM node:20.12-alpine AS base

WORKDIR /app

COPY package*.json .
RUN yarn install

COPY . .

FROM base AS development
EXPOSE 3000

ENTRYPOINT yarn dev

FROM base AS build
RUN yarn build
