import os
import time
import requests
from datetime import datetime, timedelta

LOG_DIR = "./data/logs"

def log(message):
    timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    log_entry = f"[{timestamp}] {message}\n"
    with open(f"{LOG_DIR}/sitemap-generator.log", "a+") as f:
        # 如果日志大于10000行，则先删除最旧的100行日志
        lines = f.readlines()
        if len(lines) > 10000:
            f.seek(0)
            f.writelines(lines[100:])
            f.truncate()
        f.write(log_entry)

class Article:
    def __init__(self, url, lastmod, changefreq='weekly', priority='0.8'):
        self.url = url
        self.lastmod = lastmod
        self.changefreq = changefreq
        self.priority = priority

def fetch_articles_from_api(api_url):
    """
    从后台API获取文章列表。
    
    :param api_url: 后台接口的URL
    :return: 文章列表的JSON数据
    """
    try:
        response = requests.get(api_url)
        response.raise_for_status()  # 检查请求是否成功
        json_data = response.json()
        return [Article(
            url=f"https://www.hitori.cn/article/{article['ID']}",
            lastmod=article['ModifiedAt']
        ) for article in json_data]
    except requests.exceptions.RequestException as e:
        log(f"获取文章列表失败: {e}")
        return []
    
def add_url_to_sitemap(sitemap_xml: str, url: str, lastmod: str, changefreq: str = 'weekly', priority: str = '0.8') -> str:
    """
    向Sitemap XML字符串中添加一个<url>条目。
    
    :param sitemap_xml: 当前的Sitemap XML字符串
    :param url: 页面的URL
    :param lastmod: 最后修改时间
    :param changefreq: 更新频率
    :param priority: 优先级
    :return: 更新后的Sitemap XML字符串
    """
    if not url:
        log(f"警告: 缺少URL字段，跳过。")
        return sitemap_xml

    # 添加<url>条目
    sitemap_xml += f'''  <url>
    <loc>{url}</loc>
'''
    if lastmod:
        sitemap_xml += f'''    <lastmod>{lastmod}</lastmod>
'''
    sitemap_xml += f'''    <changefreq>{changefreq}</changefreq>
    <priority>{priority}</priority>
  </url>
'''
    
    return sitemap_xml

def generate_sitemap(list_article: list[Article], output_file="sitemap.xml"):
    """
    根据文章列表生成Sitemap XML文件。
    
    :param articles: 文章列表，每个文章应包含 'url' 和 'lastmod' 字段
    :param output_file: 输出的Sitemap文件名
    """
    # Sitemap XML的头部
    sitemap_xml = '''<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
'''

    # 添加首页<url>条目
    sitemap_xml = add_url_to_sitemap(sitemap_xml, "https://www.hitori.cn/home", "", "daily", "1.0")

    # 遍历每个文章，添加<url>条目
    for article in list_article:
        url = article.url
        lastmod = article.lastmod
        changefreq = article.changefreq
        priority = article.priority

        # 格式化日期为YYYY-MM-DD
        datetime_lastmod = datetime.strptime(lastmod, "%Y-%m-%dT%H:%M:%S.%f%z")
        lastmod = datetime_lastmod.strftime("%Y-%m-%d") if lastmod else ""

        # 添加<url>条目
        sitemap_xml = add_url_to_sitemap(sitemap_xml, url, lastmod, changefreq, priority)

    # Sitemap XML的尾部
    sitemap_xml += '''</urlset>'''
    
    # 写入文件
    with open(output_file, 'w', encoding='utf-8') as f:
        f.write(sitemap_xml)

    log(f"Sitemap已生成: {output_file}")

# 每个整点执行一次任务
def schedule_hourly(task):
    while True:
        now = datetime.now()
        # 计算下一个整点
        target = now.replace(minute=0, second=0, microsecond=0) + timedelta(hours=1)
        delay = (target - now).total_seconds()
        log(f"Next task scheduled at {target.strftime('%Y-%m-%d %H:%M:%S')}")
        time.sleep(delay)
        task()

def run():
    # 配置API信息
    API_URL = "https://www.hitori.cn/api/article/all"  # 替换为你的实际API地址
    
    # 获取文章列表
    list_article = fetch_articles_from_api(API_URL)
    if not list_article:
        log("未能获取文章列表，无法生成Sitemap。")
        return
    
    # log(f"获取到 {len(list_article)} 篇文章。")
    # for idx, article in enumerate(list_article):
    #     log(f"{idx}. URL: {article.url}, 最后修改: {article.lastmod}")
    
    # 生成Sitemap
    generate_sitemap(list_article, output_file="./data/sitemap.xml")

if __name__ == "__main__":
    os.makedirs(LOG_DIR, exist_ok=True)

    run()
    
    schedule_hourly(run)