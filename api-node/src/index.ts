import { serve } from '@hono/node-server'
import { Hono } from 'hono'

const app = new Hono()

const data  = {
  'title': 'ayafuya',
  'artist': 'rokudenashi',
  'link': 'https://yt.nurs.my.id/ayafuya'
}

app.get('/api/v1/data', (c) => {
  return c.json(data)
})

serve({
  fetch: app.fetch,
  port: 3000
}, (info) => {
  console.log(`Server is running on http://localhost:${info.port}`)
})
