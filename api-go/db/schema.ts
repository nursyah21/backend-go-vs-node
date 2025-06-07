import { integer, pgTable, serial, text } from "drizzle-orm/pg-core";

export const musics = pgTable("musics", {
  id: serial('id').primaryKey(),
  userid: integer().notNull().references(()=>users.id),
  title: text().notNull(),
  artist: text().notNull(),
  link: text().notNull(),
});

export const users = pgTable("users", {
  id: serial('id').primaryKey(),
  username: text().notNull().unique(),
  password: text().notNull(),
});