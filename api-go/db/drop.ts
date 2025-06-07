import { sql } from "drizzle-orm";
import { db } from "./db.js";

console.log('dropping table...')

await db.execute(sql`DROP TABLE IF EXISTS musics, users `)