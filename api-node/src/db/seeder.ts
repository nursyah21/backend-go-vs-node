import { db } from "./db.js";
import { musics, users } from "./schema.js";

console.log('seeding database...')


await db.insert(users).values([
    {username: "admin", password: "password"}
]);

const formatNumber = (num: number) => num.toString().padStart(2,'0')

await db.insert(musics).values(
    Array.from({length:10}, (_, idx)=>(
        {artist: `artist-${formatNumber(idx)}`, title: `title-${formatNumber(idx)}`, link: `link-${formatNumber(idx)}`}
    ))
);
