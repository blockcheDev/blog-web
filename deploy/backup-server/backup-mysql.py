from datetime import datetime
import os
import subprocess
import time
from qcloud_cos import CosConfig
from qcloud_cos import CosS3Client
from qcloud_cos.cos_exception import CosClientError, CosServiceError

# 从环境变量获取配置
MYSQL_HOST = os.getenv("MYSQL_HOST", "172.17.0.1")
MYSQL_PORT = os.getenv("MYSQL_PORT", "3306")
MYSQL_USER = os.getenv("MYSQL_USER", "root")
MYSQL_ROOT_PASSWORD = os.getenv("MYSQL_ROOT_PASSWORD")
MYSQL_DATABASE = os.getenv("MYSQL_DATABASE", "blog")
BACKUP_DIR = "/backups/mysql"
LOG_DIR = "/logs"

# 1. 设置用户属性, 包括 secret_id, secret_key, region 等。Appid 已在 CosConfig 中移除，请在参数 Bucket 中带上 Appid。Bucket 由 BucketName-Appid 组成
COS_SECRET_ID = os.getenv("COS_SECRET_ID")     # 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
COS_SECRET_KEY = os.getenv("COS_SECRET_KEY")   # 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
REGION = 'ap-guangzhou'      # 替换为用户的 region，已创建桶归属的 region 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket
                           # COS 支持的所有 region 列表参见 https://cloud.tencent.com/document/product/436/6224
config = CosConfig(Region=REGION, SecretId=COS_SECRET_ID, SecretKey=COS_SECRET_KEY)
client = CosS3Client(config)

def get_host_machine_name():
    with open('/etc/host_machine_name', 'r') as f:
        name = f.read().strip()
    return name if name else "unknown_host"

def log(message):
    timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    log_entry = f"[{timestamp}] {message}\n"
    with open(f"{LOG_DIR}/backup-mysql.log", "a") as f:
        f.write(log_entry)

def upload_to_cos(file_name, file_path):
    try:
        client.upload_file(
            Bucket='blog-mysql-backups-1340568599',
            Key=file_name,
            LocalFilePath=file_path,
            EnableMD5=False,
            progress_callback=None
        )
        log(f"Upload to cos succeeded: {file_name}")
    except CosClientError as e:
        log(f"Upload CosClientError: {e.stderr}")
    except CosServiceError as e:
        log(f"Upload CosServiceError: {e.stderr}")

def backup_mysql():
    while True:
        try:
            # 检查MySQL服务是否可用
            subprocess.run(["mysqladmin", "-h", MYSQL_HOST, "-P", MYSQL_PORT, "-u", MYSQL_USER, "-p" + MYSQL_ROOT_PASSWORD, "ping"], check=True)
            break
        except subprocess.CalledProcessError:
            log("MySQL service is not available yet. Retrying...")
            time.sleep(5)
    try:
        os.makedirs(BACKUP_DIR, exist_ok=True)
        os.makedirs(LOG_DIR, exist_ok=True)
        
        timestamp = datetime.now().strftime("%Y%m%d%H")

        file_name = f"{MYSQL_DATABASE}_{get_host_machine_name()}_{timestamp}.sql.gz"
        file_path = f"{BACKUP_DIR}/{file_name}"
        
        # 使用环境变量避免密码暴露在命令行
        cmd = f"mysqldump --host={MYSQL_HOST} --port={MYSQL_PORT} -u{MYSQL_USER} -p{MYSQL_ROOT_PASSWORD} --databases {MYSQL_DATABASE} | gzip > {file_path}"
        
        subprocess.run(cmd, shell=True, check=True) # check=True会抛出异常
        log(f"Backup succeeded: {file_path}")

        # 上传到COS
        upload_to_cos(file_name, file_path)
    except subprocess.CalledProcessError as e:
        log(f"Backup failed: {e.stderr}")

def schedule_daily(task, hour=4):
    # 每天固定时间执行任务
    while True:
        now = datetime.now()
        target = now.replace(hour=hour, minute=0, second=0)
        if now >= target:
            target = target.replace(day=target.day + 1)  # 如果今天已过，改为明天
        delay = (target - now).total_seconds()

        log(f"Next backup scheduled at {target.strftime('%Y-%m-%d %H:%M:%S')}")
        time.sleep(delay)
        
        task()



if __name__ == "__main__":
    # 先执行一次备份，测试用
    backup_mysql()
    
    # 每天4点执行备份
    schedule_daily(backup_mysql, 4)