FROM mysql:9


# 定义被容器自动执行的目录
ENV AUTO_RUN_DIR /docker-entrypoint-initdb.d


# 把执行初始化的脚本放到 /docker-entrypoint-initdb.d 目录下
COPY ./deploy/mysql/blog.sql ${AUTO_RUN_DIR}/