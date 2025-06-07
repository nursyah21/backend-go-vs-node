import { db } from "./db.js";
import { musics, users } from "./schema.js";

console.log('dropping database...')

await db.delete(users)
await db.delete(musics)