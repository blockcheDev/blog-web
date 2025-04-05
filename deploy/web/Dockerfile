FROM node:23 AS node
WORKDIR /usr/local/front
COPY ./front/ ./
RUN npm config set registry https://registry.npmmirror.com \
    && npm install -g pnpm \ 
    && pnpm install
RUN pnpm build



FROM nginx:1.27
WORKDIR /usr/local/dist
COPY --from=node /usr/local/front/dist/ ./