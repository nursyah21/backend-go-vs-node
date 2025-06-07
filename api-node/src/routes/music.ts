import { and, eq } from "drizzle-orm"
import { Hono } from "hono"
import { db } from "../db/db.js"
import { musics } from "../db/schema.js"
import { authMiddleware } from "../middleware.js"

type Props = typeof musics.$inferSelect

export const musicRoute = new Hono()

musicRoute.get('/music', async (c) => {
    const res = await db.select().from(musics)

    return c.json(res)
})

musicRoute.post('/music', authMiddleware, async (c) => {
    const user = c.get('user')

    const { title, artist, link } = await c.req.json<Props>()

    if (!title || !artist || !link) {
        return c.json({ error: "Missing required fields: title, artist, or link" }, 400)
    }

    try {
        await db.insert(musics).values({ userid: user.id, title, artist, link })
    } catch (error) {
        return c.json({ status: "user not found" }, 400)
    }

    return c.json({ status: "success" })
})

musicRoute.put('/music/:id', authMiddleware, async (c) => {
    const user = c.get('user')

    const id = parseInt(c.req.param('id'))

    if (isNaN(id)) {
        return c.json({ error: "Invalid ID format" }, 400)
    }

    const { title, artist, link } = await c.req.json<Props>()

    if (!title || !artist || !link) {
        return c.json({ error: "Missing required fields: title, artist, or link" }, 400)
    }

    try {
        await db.update(musics).set({ title, artist, link })
            .where(and(eq(musics.id, id), eq(musics.userid, user.id)))
    } catch { }

    return c.json({ status: "success" })
})

musicRoute.delete('/music/:id', authMiddleware, async (c) => {
    const user = c.get('user')

    const id = parseInt(c.req.param('id'))

    if (isNaN(id)) {
        return c.json({ error: "Invalid ID format" }, 400)
    }

    try {
        await db.delete(musics)
            .where(and(eq(musics.id, id), eq(musics.userid, user.id)))
    } catch { }

    return c.json({ status: "success" })
})
