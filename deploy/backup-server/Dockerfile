FROM python:3

RUN apt-get update && apt-get install -y \
    default-mysql-client \
    gzip \
    && rm -rf /var/lib/apt/lists/* \
    && pip install -U cos-python-sdk-v5

WORKDIR /app
COPY deploy/backup-server/backup-mysql.py ./

ENTRYPOINT python backup-mysql.py