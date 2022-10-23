BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "officers" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"email"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "faculties" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "prefixes" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "educationals" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "teachers" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"email"	text,
	"officer_id"	integer,
	"faculty_id"	integer,
	"prefix_id"	integer,
	"educational_id"	integer,
	CONSTRAINT "fk_faculties_teacher" FOREIGN KEY("faculty_id") REFERENCES "faculties"("id"),
	CONSTRAINT "fk_officers_teacher" FOREIGN KEY("officer_id") REFERENCES "officers"("id"),
	CONSTRAINT "fk_educationals_teacher" FOREIGN KEY("educational_id") REFERENCES "educationals"("id"),
	PRIMARY KEY("id"),
	CONSTRAINT "fk_prefixes_teacher" FOREIGN KEY("prefix_id") REFERENCES "prefixes"("id")
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_officers_email" ON "officers" (
	"email"
);
CREATE INDEX IF NOT EXISTS "idx_officers_deleted_at" ON "officers" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_faculties_deleted_at" ON "faculties" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_prefixes_deleted_at" ON "prefixes" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_educationals_deleted_at" ON "educationals" (
	"deleted_at"
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_teachers_email" ON "teachers" (
	"email"
);
CREATE INDEX IF NOT EXISTS "idx_teachers_deleted_at" ON "teachers" (
	"deleted_at"
);
COMMIT;
