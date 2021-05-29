FROM node:16-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM nginx:1.21.0-alpine
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY nginx/ /etc/todo-app/
RUN rm -rf /usr/share/nginx/html/*
COPY --from=build /app/dist /usr/share/nginx/html

# RUN cat /etc/nginx/nginx.conf
# RUN cat /etc/nginx/conf.d/default.conf