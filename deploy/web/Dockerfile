FROM node:23 AS node
WORKDIR /usr/local/front
COPY ./front/ ./
# COPY ./front/.env.production ./
RUN npm install -g pnpm \ 
    && pnpm install
RUN pnpm build



FROM nginx:1.27
WORKDIR /usr/local/dist
COPY --from=node /usr/local/front/dist/ ./
COPY ./deploy/web/nginx.conf /etc/nginx/nginx.conf
COPY ./deploy/web/hitori.cn.* /usr/local/nginx/conf/cert/