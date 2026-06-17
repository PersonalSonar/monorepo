# Worker Service

Python async job processor service.

## Setup

```bash
cd worker
python3 -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
pip install -r requirements.txt
```

## Running

```bash
python worker.py
```

## Features

- Async job queue processing
- Job persistence to JSON
- Status tracking (pending → processing → completed)
- Job result storage

## Job Structure

```python
{
    "id": "job_1234567890",
    "task_id": "task_0",
    "job_type": "process_data",
    "status": "completed",
    "created_at": "2024-01-01T12:00:00",
    "completed_at": "2024-01-01T12:00:01",
    "result": {
        "processed_items": 42,
        "execution_time_ms": 1000
    }
}
```

## Configuration

Create a `.env` file:

```
LOG_LEVEL=INFO
WORKER_THREADS=4
```
