import { Hono } from "hono"
import { db } from "../db/db.js"
import { musics } from "../db/schema.js"
import { eq } from "drizzle-orm"

export const musicRoute = new Hono()

musicRoute.get('/music', async (c) => {
    const res = await db.select().from(musics)

    return c.json(res)
})

musicRoute.post('/music', async (c) => {
    const { title, artist, link } = await c.req.json()

    if (!title || !artist || !link) {
        return c.json({ error: "Missing required fields: title, artist, or link" }, 400)
    }

    await db.insert(musics).values({ title, artist, link })

    return c.json({ status: "success" })
})

musicRoute.delete('/music/:id', async (c) => {
    const id = parseInt(c.req.param('id'))

    if (isNaN(id)) {
        return c.json({ error: "Invalid ID format" }, 400)
    }
    
    
    const existing = await db.select().from(musics).where(eq(musics.id, id)).limit(1)

    if (!existing.length) {
        return c.json({ error: 'id not found' }, 404)
    }

    await db.delete(musics).where(eq(musics.id, id))
    return c.json({ status: "success" })
})
