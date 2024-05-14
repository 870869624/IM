CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "account" varchar UNIQUE NOT NULL,
  "token" varchar NOT NULL,
  "account_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00',
  "phone_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00',
  "username_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00',
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "friends" (
  "id" bigserial PRIMARY KEY,
  "f_user_account" varchar NOT NULL,
  "t_user_account" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "groups" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "account" varchar UNIQUE NOT NULL,
  "owner" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "group_members" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "group_account" varchar NOT NULL,
  "user_account" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "messages" (
  "id" bigserial PRIMARY KEY,
  "content" varchar NOT NULL,
  "from_user_account" varchar NOT NULL,
  "to_user_account" varchar NOT NULL,
  "group_account" varchar NOT NULL,
  "m_type" int NOT NULL,
  "networkstatus" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "applies" (
  "id" bigserial PRIMARY KEY,
  "applicate_account" varchar NOT NULL,
  "target_account" varchar NOT NULL,
  "a_type" int NOT NULL,
  "status" int NOT NULL,
  "object" int NOT NULL
);

CREATE INDEX ON "groups" ("name");

CREATE INDEX ON "groups" ("account");

CREATE INDEX ON "group_members" ("group_account");

CREATE INDEX ON "group_members" ("user_account");

CREATE INDEX ON "messages" ("content");

COMMENT ON COLUMN "group_members"."name" IS '群内昵称应该是唯一的';

COMMENT ON COLUMN "messages"."m_type" IS '消息类型:0系统消息,1为私人消息,2为群消息';

COMMENT ON COLUMN "messages"."networkstatus" IS '0为在线,1为离线';

COMMENT ON COLUMN "applies"."a_type" IS '消息类型:1为直接添加,2为请求添加,3为同意添加,4拒绝添加';

COMMENT ON COLUMN "applies"."status" IS '0:未同意,1:已同意';

COMMENT ON COLUMN "applies"."object" IS '0.用户, 1:群组';

ALTER TABLE "friends" ADD FOREIGN KEY ("f_user_account") REFERENCES "users" ("account");

ALTER TABLE "friends" ADD FOREIGN KEY ("t_user_account") REFERENCES "users" ("account");

ALTER TABLE "groups" ADD FOREIGN KEY ("owner") REFERENCES "users" ("account");

ALTER TABLE "group_members" ADD FOREIGN KEY ("user_account") REFERENCES "users" ("account");

ALTER TABLE "applies" ADD FOREIGN KEY ("applicate_account") REFERENCES "users" ("account");
