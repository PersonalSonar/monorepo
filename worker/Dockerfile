FROM python:3.11-slim

WORKDIR /app/worker

COPY worker/requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt

COPY worker/worker.py ./

CMD ["python", "worker.py"]
