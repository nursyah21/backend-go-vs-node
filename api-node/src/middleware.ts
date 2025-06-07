import { createMiddleware } from "hono/factory"
import { verifyJwt, type Jwt } from "./lib/jwt.js"

export const authMiddleware = createMiddleware<{
    Variables: {
        user: Jwt
    }
}>(async (c, next) => {
    const token = c.req.header("x-token")
    let res
    try {
        res = await verifyJwt(token!)
    } catch (error) {
        return c.json({ error: 'jwt not valid' }, 400)
    }
    c.set('user', res)

    await next()
})

