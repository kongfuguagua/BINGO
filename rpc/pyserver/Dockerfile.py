FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/python:3.12
WORKDIR /app

COPY requirements.txt .
RUN pip install -r requirements.txt

COPY . .

CMD ["python", "server.py"]