-- create user modulith with superuser password 'modulith'
-- create database b2k_abac
-- psql -U modulith -P modulith
-- ALTER USER postgres WITH PASSWORD 'postgres_123'
-- database language : psql

-- DROP TABLE if exists comments, posts CASCADE;
-- DROP TABLE if exists role_M2M_permission ;
-- DROP TABLE if exists role_M2M_user ;
-- DROP TABLE if exists users, roles, permissions CASCADE;
-- DROP TABLE IF EXISTS schema_migrations;

CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" varchar(50),
  "full_name" varchar(50),
  "email" varchar(50),
  "hashed_password" varchar(128),
  "password_changed_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "is_email_verified" boolean NOT NULL default false,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "created_by" BIGINT REFERENCES users(id),
  "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "updated_by" BIGINT REFERENCES users(id),
  "version" BIGINT default 1,
  "latest" boolean NOT NULL default true
);

CREATE TABLE "roles" (
  "id" BIGSERIAL PRIMARY KEY,
  "title" varchar(50) ,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "created_by" BIGINT REFERENCES users(id),
  "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "updated_by" BIGINT REFERENCES users(id),
  "version" BIGINT default 1,
  "latest" boolean NOT NULL default true,
  CONSTRAINT uq_roles_title UNIQUE (title)
);

CREATE TABLE "role_m2m_user" (
  "id" BIGSERIAL PRIMARY KEY,
  "role_id" BIGINT,
  "user_id" BIGINT,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "created_by" BIGINT REFERENCES users(id),
  "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "updated_by" BIGINT REFERENCES users(id),
  "version" BIGINT default 1,
  "latest" boolean NOT NULL default true,
  CONSTRAINT fk_role_id_m2m_user FOREIGN KEY (role_id) REFERENCES roles(id),
  CONSTRAINT fk_role_m2m_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE "permissions" (
  "id" BIGSERIAL PRIMARY KEY,
  "title" varchar(50),
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "created_by" BIGINT REFERENCES users(id),
  "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "updated_by" BIGINT REFERENCES users(id),
  "version" BIGINT default 1,
  "latest" boolean NOT NULL default true,
  CONSTRAINT uq_permissions_title UNIQUE(title)
);

CREATE TABLE "role_m2m_permission" (
  "id" BIGSERIAL PRIMARY KEY,
  "role_id" BIGINT REFERENCES roles(id),
  "permission_id" BIGINT REFERENCES permissions(id),
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "created_by" BIGINT REFERENCES users(id),
  "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "updated_by" BIGINT REFERENCES users(id),
  "version" BIGINT default 1,
  "latest" boolean NOT NULL default true
);

CREATE TABLE "posts" (
  "id" BIGSERIAL PRIMARY KEY,
  "title" varchar,
  "body" text,
  "user_id" BIGINT REFERENCES users(id),
  "status" varchar,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "created_by" BIGINT REFERENCES users(id),
  "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "updated_by" BIGINT REFERENCES users(id),
  "version" BIGINT default 1,
  "latest" boolean NOT NULL default true
);


CREATE TABLE "comments" (
  "id" BIGSERIAL PRIMARY KEY,
  "post_id" BIGINT REFERENCES posts(id),
  "user_id" BIGINT REFERENCES users(id),
  "body" text,
  "status" varchar,
  "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "created_by" BIGINT REFERENCES users(id),
  "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  "updated_by" BIGINT REFERENCES users(id),
  "version" BIGINT default 1,
  "latest" boolean NOT NULL default true
);
COMMENT ON COLUMN posts.body IS 'Content of the post';
COMMENT ON COLUMN comments.body IS 'Content of the comment';
CREATE INDEX if not exists idx_role_m2m_user_latest on role_m2m_user(latest desc);
CREATE INDEX if not exists idx_roles_latest on roles(latest desc);
CREATE INDEX if not exists idx_posts_latest on posts(latest desc);
CREATE INDEX if not exists idx_role_m2m_permission_latest on role_m2m_permission(latest desc);
CREATE INDEX if not exists idx_permissions_latest on permissions(latest desc);
CREATE INDEX if not exists idx_roles_latest on roles(latest desc);
CREATE INDEX if not exists idx_users_latest on users(latest desc);

