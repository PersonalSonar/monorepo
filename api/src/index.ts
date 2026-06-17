import express, { Request, Response } from 'express';
import cors from 'cors';

const app = express();
const PORT = process.env.PORT || 3000;

app.use(cors());
app.use(express.json());

interface Task {
  id: string;
  title: string;
  completed: boolean;
  createdAt: Date;
}

const tasks: Map<string, Task> = new Map();

app.get('/health', (_req: Request, res: Response) => {
  res.json({ status: 'ok', service: 'api' });
});

app.get('/tasks', (_req: Request, res: Response) => {
  res.json(Array.from(tasks.values()));
});

app.post('/tasks', (req: Request, res: Response) => {
  const { title } = req.body;
  if (!title) {
    return res.status(400).json({ error: 'Title is required' });
  }

  const id = Date.now().toString();
  const task: Task = {
    id,
    title,
    completed: false,
    createdAt: new Date(),
  };

  tasks.set(id, task);
  res.status(201).json(task);
});

app.get('/tasks/:id', (req: Request, res: Response) => {
  const task = tasks.get(req.params.id);
  if (!task) {
    return res.status(404).json({ error: 'Task not found' });
  }
  res.json(task);
});

app.patch('/tasks/:id', (req: Request, res: Response) => {
  const task = tasks.get(req.params.id);
  if (!task) {
    return res.status(404).json({ error: 'Task not found' });
  }

  if (req.body.completed !== undefined) {
    task.completed = req.body.completed;
  }
  if (req.body.title !== undefined) {
    task.title = req.body.title;
  }

  res.json(task);
});

app.delete('/tasks/:id', (req: Request, res: Response) => {
  const deleted = tasks.delete(req.params.id);
  if (!deleted) {
    return res.status(404).json({ error: 'Task not found' });
  }
  res.status(204).send();
});

app.listen(PORT, () => {
  console.log(`API server running on port ${PORT}`);
});
