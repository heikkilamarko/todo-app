FROM node:16-alpine AS build
WORKDIR /app
COPY . .
RUN npm ci -w todo-app
RUN npm run build -w todo-app

FROM nginx:alpine
ENV NGINX_ENVSUBST_TEMPLATE_DIR=/etc/todo-app
ENV NGINX_ENVSUBST_OUTPUT_DIR=/usr/share/nginx/html
COPY packages/todo-app/nginx.conf /etc/nginx/conf.d/default.conf
COPY packages/todo-app/nginx/ /etc/todo-app/
RUN rm -rf /usr/share/nginx/html/*
COPY --from=build /app/packages/todo-app/dist /usr/share/nginx/html

# RUN cat /etc/nginx/nginx.conf
# RUN cat /etc/nginx/conf.d/default.conf