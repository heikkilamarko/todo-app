FROM node:18-alpine AS build
WORKDIR /app
COPY . .
RUN npm ci
RUN npm run build

FROM caddy:alpine
COPY ./Caddyfile /etc/caddy/
COPY --from=build /app/build /usr/share/todo-app
