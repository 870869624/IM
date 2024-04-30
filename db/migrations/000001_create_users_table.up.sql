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

CREATE TABLE "friend" (
  "id" bigserial PRIMARY KEY,
  "f_user_id" varchar NOT NULL,
  "t_user_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "groups" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "account" varchar UNIQUE NOT NULL,
  "owner" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "group_member" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "group_id" varchar NOT NULL,
  "user_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "messages" (
  "id" bigserial PRIMARY KEY,
  "content" varchar,
  "from_user_id" varchar NOT NULL,
  "to_user_id" varchar,
  "group_id" varchar,
  "m_type" int NOT NULL,
  "networktatus" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "apply" (
  "id" bigserial PRIMARY KEY,
  "applicate_id" varchar NOT NULL,
  "target_id" varchar NOT NULL,
  "a_type" int NOT NULL
);

CREATE INDEX ON "groups" ("name");

CREATE INDEX ON "groups" ("account");

CREATE INDEX ON "group_member" ("group_id");

CREATE INDEX ON "group_member" ("user_id");

CREATE INDEX ON "messages" ("content");

COMMENT ON COLUMN "messages"."m_type" IS '消息类型,0为群消息,1为私人消息';

COMMENT ON COLUMN "messages"."networktatus" IS '0为在线,1为离线';

COMMENT ON COLUMN "apply"."a_type" IS '消息类型,0为群申请,1为私人申请';

ALTER TABLE "friend" ADD FOREIGN KEY ("f_user_id") REFERENCES "users" ("account");

ALTER TABLE "friend" ADD FOREIGN KEY ("t_user_id") REFERENCES "users" ("account");

ALTER TABLE "groups" ADD FOREIGN KEY ("owner") REFERENCES "users" ("account");

ALTER TABLE "group_member" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("account");

ALTER TABLE "group_member" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("account");

ALTER TABLE "messages" ADD FOREIGN KEY ("from_user_id") REFERENCES "users" ("account");

ALTER TABLE "messages" ADD FOREIGN KEY ("to_user_id") REFERENCES "users" ("account");

ALTER TABLE "messages" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("account");

ALTER TABLE "apply" ADD FOREIGN KEY ("applicate_id") REFERENCES "users" ("account");

ALTER TABLE "apply" ADD FOREIGN KEY ("target_id") REFERENCES "users" ("account");
