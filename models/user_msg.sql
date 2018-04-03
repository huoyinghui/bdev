create table `user_message`
    -- --------------------------------------------------
    --  Table Structure for `bdev/models.Message`
    -- --------------------------------------------------
   CREATE TABLE IF NOT EXISTS "user_message" (
        "id" serial NOT NULL PRIMARY KEY,
        "from_user_id" bigint NOT NULL DEFAULT 0 ,
        "to_user_id" bigint NOT NULL DEFAULT 0 ,
        "created_at" timestamp with time zone,
        "update_at" timestamp with time zone,
        "title" text,
        "message" text,
        "is_delete" bool NOT NULL DEFAULT FALSE ,
        "status" text
    );
    CREATE INDEX "user_message_id_from_user_id" ON "user_message" ("id", "from_user_id");
    CREATE INDEX "user_message_id_to_user_id" ON "user_message" ("id", "to_user_id");
    CREATE INDEX "user_message_id_created_at" ON "user_message" ("id", "created_at");


-- Insert
INSERT INTO "user_message" ("from_user_id", "to_user_id", "created_at", "update_at", "title", "text", "is_delete", "status")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING "id"
-- `2`,
-- `1001`,
-- `2018-04-02 17:15:47.641417322 +0800 CST`,
-- `2018-04-02 17:15:47.641424145 +0800 CST`,
-- `Welcome to join APKPURE developer`,
--  `1.you xxx 2.xxxxx 3.xxxx`, `false`, `UNSEEN`

-- Select by toUserId    分页
SELECT T0."id", T0."from_user_id", T0."to_user_id", T0."created_at", T0."update_at", T0."title", T0."text", T0."is_delete", T0."status"
FROM "user_message" T0
WHERE T0."to_user_id" = $1
AND T0."status" = $2
AND T0."is_delete" = $3
LIMIT 1000
-- `1001`, `UNSEEN`|`SEEN`, `false`

-- Select by fromUserId  分页
SELECT T0."id", T0."from_user_id", T0."to_user_id", T0."created_at", T0."update_at", T0."title", T0."text", T0."is_delete", T0."status"
FROM "user_message" T0
WHERE T0."from_user_id" = $1
AND T0."is_delete" = $2
LIMIT 1000
-- `2`, `false`



-- Update by mid
UPDATE "user_message" SET "from_user_id" = $1, "to_user_id" = $2, "created_at" = $3, "update_at" = $4, "title" = $5, "text" = $6, "is_delete" = $7, "status" = $8
WHERE "id" = $9
 -- `109`,
 -- `0`,
 -- `2018-04-02 17:14:57.964735 +0800 CST`,
 -- `2018-04-02 17:22:35.944305943 +0800 CST`,
 -- `Welcome to join APKPURE developer`,
 -- `1.you xxx 2.xxxxx 3.xxxx`,
 -- `false`,
 -- `SEEN`,
 -- `10`

-- Delete by mid
UPDATE "user_message" SET "is_delete" = $1
WHERE "id" = $2
-- `true`,
-- `10`


﻿SELECT id, from_user_id, to_user_id, created_at, update_at, title, is_delete, status, message
FROM public.user_message
LIMIT 1
OFFSET 10;

SELECT T0."id", T0."from_user_id", T0."to_user_id", T0."created_at", T0."update_at", T0."title", T0."message", T0."is_delete", T0."status" FROM "user_message" T0
WHERE T0."to_user_id" = 0
AND T0."from_user_id" = 116
AND T0."status" = $3
AND T0."is_delete" = $4
ORDER BY T0."id" ASC
LIMIT 5 OFFSET 20
-- `0`, `116`, `UNSEEN`, `false`