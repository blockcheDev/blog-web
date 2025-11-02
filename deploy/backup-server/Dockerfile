FROM python:3

RUN apt-get update && apt-get install -y \
    default-mysql-client \
    gzip \
    && rm -rf /var/lib/apt/lists/* \
    pip config set global.index-url http://mirrors.cloud.tencent.com/pypi/simple \
    && pip install -U cos-python-sdk-v5

WORKDIR /app
COPY deploy/backup-server/backup-mysql.py ./

ENTRYPOINT python backup-mysql.py