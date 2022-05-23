PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE posts(
id integer primary key autoincrement,
title text not null,
content text not null,
published_at datetime not null,
created_at datetime not null,
updated_at datetime not null
);
DELETE FROM sqlite_sequence;
COMMIT;
