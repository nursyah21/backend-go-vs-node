import { bcrypt, bcryptVerify } from 'hash-wasm'

const salt = new Uint8Array(16)
const costFactor = 10

export const hashPassword = async (password: string) =>
    bcrypt({ password, salt, costFactor })

export const verifyPassword = async (hash: string, password: string) =>
    bcryptVerify({ password, hash })
