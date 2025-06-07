import { eq } from "drizzle-orm"
import { Hono } from "hono"
import { db } from "../db/db.js"
import { users } from "../db/schema.js"
import { generateJwt } from "../lib/jwt.js"
import { hashPassword, verifyPassword } from "../lib/password.js"

export const userRoute = new Hono()

type Props = typeof users.$inferSelect

userRoute.post('/register', async (c) => {
    const { username, password: _password } = await c.req.json<Props>()

    if (!username || !_password) {
        return c.json({ error: "Missing required fields: username, password" }, 400)
    }

    const password = await hashPassword(_password)

    try {
        await db.insert(users).values({ username, password })
    } catch (error) {
        return c.json({ error: `${username} already exist` }, 400)
    }

    return c.json({ status: "success" })
})

userRoute.post('/login', async (c) => {
    const { username, password } = await c.req.json<Props>()

    if (!username || !password) {
        return c.json({ error: "Missing required fields: username, password" }, 400)
    }

    const res = await db.select().from(users).where(eq(users.username, username)).limit(1)

    if (!await verifyPassword(res[0].password, password)) {
        return c.json({ error: "username and password didnt match" }, 400)
    }

    const token = await generateJwt({ id: res[0].id, username: res[0].username })

    return c.json({ status: "success", token })
})
