import { jwtVerify, SignJWT, type JWTPayload } from "jose"
import type { users } from "../db/schema.js"

type Props = Omit<typeof users.$inferSelect, 'password'>

export interface Jwt extends JWTPayload {
    id: number;
    username: string;
}

const secret = new TextEncoder().encode('secret-key-123')

export const generateJwt = async ({ id, username }: Props) => await new SignJWT({
    id,
    username,
    iat: Math.floor(Date.now() / 1000),
    exp: Math.floor(Date.now() / 1000) + 60 * 60
})
    .setProtectedHeader({ alg: 'HS256' })
    .sign(secret)

export const verifyJwt = async (token: string) =>
    (await jwtVerify(token, secret)).payload as Jwt