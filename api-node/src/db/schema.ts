import { pgTable, serial, text } from "drizzle-orm/pg-core";

export const musics = pgTable("musics", {
  id: serial('id').primaryKey(),
  title: text().notNull(),
  artist: text().notNull(),
  link: text().notNull(),
});
