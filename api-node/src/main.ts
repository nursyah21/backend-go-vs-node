import { serve } from '@hono/node-server';
import { Hono } from 'hono';

import { musicRoute } from './routes/music.js';
import { healthRoute } from './routes/health.js';
import { userRoute } from './routes/user.js';

const app = new Hono()

app.route('/api/v1/', healthRoute)
app.route('/api/v1/', musicRoute)
app.route('/api/v1/', userRoute)

serve({
  fetch: app.fetch,
  port: 3000
}, (info) => {
  console.log(`Server is running on http://localhost:${info.port}`)
})
