import axios from "axios";
import { db } from "./db.js";
import { musics, users } from "./schema.js";

console.log('seeding database...')

try {
    try {
        await axios.post('http://localhost:5000/api/v1/register', {
            username: "admin", password: "password"
        })
    } catch {
        throw Error("admin already exists")
    }

    const formatNumber = (num: number) => num.toString().padStart(2, '0')
    const userid = (await db.select().from(users).limit(1))[0].id

    await db.insert(musics).values(
        Array.from({ length: 10 }, (_, idx) => (
            { artist: `artist-${formatNumber(idx)}`, title: `title-${formatNumber(idx)}`, link: `link-${formatNumber(idx)}`, userid }
        ))
    );
} catch (error) {
    console.log(error)
}
