import { serve } from '@hono/node-server';
import { Hono } from 'hono';
import 'dotenv/config';

import { musicRoute } from './routes/music.js';
import { healthRoute } from './routes/health.js';

const app = new Hono()

app.route('/api/v1/', healthRoute)
app.route('/api/v1/', musicRoute)

serve({
  fetch: app.fetch,
  port: 3000
}, (info) => {
  console.log(`Server is running on http://localhost:${info.port}`)
})
