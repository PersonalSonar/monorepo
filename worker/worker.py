import asyncio
import json
import logging
from dataclasses import dataclass, asdict
from datetime import datetime
from pathlib import Path
import os
from dotenv import load_dotenv

load_dotenv()

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

@dataclass
class Job:
    id: str
    task_id: str
    job_type: str
    status: str
    created_at: str
    completed_at: str = None
    result: dict = None

class JobQueue:
    def __init__(self):
        self.queue = asyncio.Queue()
        self.jobs: dict[str, Job] = {}
        self.data_file = Path('jobs.json')
        self._load_jobs()

    def _load_jobs(self):
        if self.data_file.exists():
            with open(self.data_file) as f:
                data = json.load(f)
                for job_data in data:
                    job = Job(**job_data)
                    self.jobs[job.id] = job

    def _save_jobs(self):
        with open(self.data_file, 'w') as f:
            json.dump([asdict(job) for job in self.jobs.values()], f, indent=2)

    async def enqueue(self, task_id: str, job_type: str) -> str:
        job_id = f"job_{datetime.now().timestamp()}"
        job = Job(
            id=job_id,
            task_id=task_id,
            job_type=job_type,
            status='pending',
            created_at=datetime.now().isoformat()
        )
        self.jobs[job_id] = job
        await self.queue.put(job)
        self._save_jobs()
        logger.info(f"Enqueued job {job_id}")
        return job_id

    async def process_jobs(self):
        while True:
            job = await self.queue.get()
            logger.info(f"Processing job {job.id}")

            job.status = 'processing'
            await asyncio.sleep(1)  # Simulate work

            job.status = 'completed'
            job.completed_at = datetime.now().isoformat()
            job.result = {
                'processed_items': 42,
                'execution_time_ms': 1000
            }

            logger.info(f"Completed job {job.id}")
            self._save_jobs()
            self.queue.task_done()

    def get_job(self, job_id: str) -> Job:
        return self.jobs.get(job_id)

    def list_jobs(self) -> list[Job]:
        return list(self.jobs.values())

async def main():
    queue = JobQueue()

    logger.info("Starting job worker")

    # Start processing jobs
    processor_task = asyncio.create_task(queue.process_jobs())

    # Simulate enqueueing jobs
    for i in range(3):
        await queue.enqueue(f"task_{i}", "process_data")
        await asyncio.sleep(0.5)

    # Wait for all jobs to complete
    await queue.queue.join()

    logger.info("All jobs processed")
    logger.info(f"Final job states: {[asdict(job) for job in queue.list_jobs()]}")

if __name__ == '__main__':
    asyncio.run(main())
