import { Hono } from "hono"

export const healthRoute = new Hono()

healthRoute.get('/health', async (c) => {
    
    return c.json({status: 'ok'})
})