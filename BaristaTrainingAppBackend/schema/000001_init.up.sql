CREATE TABLE IF NOT EXISTS "lesson_status" (
	"id" int NOT NULL UNIQUE,
	"name" varchar NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "lection_information" (
	"lection_id" int NOT NULL,
	"information_id" int NOT NULL,
	PRIMARY KEY ("lection_id", "information_id")
);

CREATE TABLE IF NOT EXISTS "employee" (
	"id" int NOT NULL UNIQUE,
	"name" varchar NOT NULL,
	"lastname" varchar NOT NULL,
	"email" varchar NOT NULL,
	"password_hash" int NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "role" (
	"id" int NOT NULL UNIQUE,
	"name" varchar NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "course_status" (
	"id" int NOT NULL UNIQUE,
	"name" varchar NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "course" (
	"id" int NOT NULL UNIQUE,
	"name" varchar NOT NULL,
	"description" text NOT NULL,
	"passing_score" int NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "test" (
	"id" int NOT NULL UNIQUE,
	"theme" varchar NOT NULL,
	"description" text NOT NULL,
	"role_id" int NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "lesson" (
	"id" int NOT NULL UNIQUE,
	"name" varchar NOT NULL,
	"description" text NOT NULL,
	"passing_score" int NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "question" (
	"id" int NOT NULL UNIQUE,
	"theme" varchar NOT NULL,
	"text" varchar NOT NULL,
	"right_answer_id" int NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "employee_lesson" (
	"employee_id" int NOT NULL,
	"lesson_id" int NOT NULL,
	"lesson_status" int NOT NULL,
	PRIMARY KEY ("employee_id", "lesson_id")
);

CREATE TABLE IF NOT EXISTS "lection" (
	"id" int NOT NULL UNIQUE,
	"theme" varchar NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "information" (
	"id" int NOT NULL UNIQUE,
	"theme" varchar NOT NULL,
	"text" varchar NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "test_question" (
	"test_id" int NOT NULL,
	"question" int NOT NULL,
	PRIMARY KEY ("test_id", "question")
);

CREATE TABLE IF NOT EXISTS "employee_test" (
	"employee_id" int NOT NULL,
	"test_id" int NOT NULL,
	"score" int NOT NULL,
	PRIMARY KEY ("employee_id", "test_id")
);

CREATE TABLE IF NOT EXISTS "lesson_test" (
	"lesson_id" int NOT NULL,
	"test_id" int NOT NULL,
	PRIMARY KEY ("lesson_id", "test_id")
);

CREATE TABLE IF NOT EXISTS "employee_course" (
	"employee_id" int NOT NULL,
	"course_id" int NOT NULL,
	"course_status" int NOT NULL,
	PRIMARY KEY ("employee_id", "course_id")
);

CREATE TABLE IF NOT EXISTS "lesson_lection" (
	"lesson_id" int NOT NULL,
	"lection_id" int NOT NULL,
	PRIMARY KEY ("lesson_id", "lection_id")
);

CREATE TABLE IF NOT EXISTS "course_lesson" (
	"course_id" int NOT NULL,
	"lesson_id" int NOT NULL,
	PRIMARY KEY ("course_id", "lesson_id")
);

CREATE TABLE IF NOT EXISTS "employee_role" (
	"employee_id" int NOT NULL,
	"role_id" int NOT NULL,
	PRIMARY KEY ("employee_id", "role_id")
);

CREATE TABLE IF NOT EXISTS "answer" (
	"id" int NOT NULL UNIQUE,
	"text" varchar NOT NULL,
	PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "question_answer" (
	"question_id" int NOT NULL,
	"answer_id" int NOT NULL,
	PRIMARY KEY ("question_id", "answer_id")
);


ALTER TABLE "lection_information" ADD CONSTRAINT "lection_information_fk0" FOREIGN KEY ("lection_id") REFERENCES "lection"("id");

ALTER TABLE "lection_information" ADD CONSTRAINT "lection_information_fk1" FOREIGN KEY ("information_id") REFERENCES "information"("id");




ALTER TABLE "test" ADD CONSTRAINT "test_fk3" FOREIGN KEY ("role_id") REFERENCES "role"("id");

ALTER TABLE "question" ADD CONSTRAINT "question_fk3" FOREIGN KEY ("right_answer_id") REFERENCES "answer"("id");
ALTER TABLE "employee_lesson" ADD CONSTRAINT "employee_lesson_fk0" FOREIGN KEY ("employee_id") REFERENCES "employee"("id");

ALTER TABLE "employee_lesson" ADD CONSTRAINT "employee_lesson_fk1" FOREIGN KEY ("lesson_id") REFERENCES "lesson"("id");

ALTER TABLE "employee_lesson" ADD CONSTRAINT "employee_lesson_fk2" FOREIGN KEY ("lesson_status") REFERENCES "lesson_status"("id");


ALTER TABLE "test_question" ADD CONSTRAINT "test_question_fk0" FOREIGN KEY ("test_id") REFERENCES "test"("id");

ALTER TABLE "test_question" ADD CONSTRAINT "test_question_fk1" FOREIGN KEY ("question") REFERENCES "question"("id");
ALTER TABLE "employee_test" ADD CONSTRAINT "employee_test_fk0" FOREIGN KEY ("employee_id") REFERENCES "employee"("id");

ALTER TABLE "employee_test" ADD CONSTRAINT "employee_test_fk1" FOREIGN KEY ("test_id") REFERENCES "test"("id");
ALTER TABLE "lesson_test" ADD CONSTRAINT "lesson_test_fk0" FOREIGN KEY ("lesson_id") REFERENCES "lesson"("id");

ALTER TABLE "lesson_test" ADD CONSTRAINT "lesson_test_fk1" FOREIGN KEY ("test_id") REFERENCES "test"("id");
ALTER TABLE "employee_course" ADD CONSTRAINT "employee_course_fk0" FOREIGN KEY ("employee_id") REFERENCES "employee"("id");

ALTER TABLE "employee_course" ADD CONSTRAINT "employee_course_fk1" FOREIGN KEY ("course_id") REFERENCES "course"("id");

ALTER TABLE "employee_course" ADD CONSTRAINT "employee_course_fk2" FOREIGN KEY ("course_status") REFERENCES "course_status"("id");
ALTER TABLE "lesson_lection" ADD CONSTRAINT "lesson_lection_fk0" FOREIGN KEY ("lesson_id") REFERENCES "lesson"("id");

ALTER TABLE "lesson_lection" ADD CONSTRAINT "lesson_lection_fk1" FOREIGN KEY ("lection_id") REFERENCES "lection"("id");
ALTER TABLE "course_lesson" ADD CONSTRAINT "course_lesson_fk0" FOREIGN KEY ("course_id") REFERENCES "course"("id");

ALTER TABLE "course_lesson" ADD CONSTRAINT "course_lesson_fk1" FOREIGN KEY ("lesson_id") REFERENCES "lesson"("id");
ALTER TABLE "employee_role" ADD CONSTRAINT "employee_role_fk0" FOREIGN KEY ("employee_id") REFERENCES "employee"("id");

ALTER TABLE "employee_role" ADD CONSTRAINT "employee_role_fk1" FOREIGN KEY ("role_id") REFERENCES "role"("id");

ALTER TABLE "question_answer" ADD CONSTRAINT "question_answer_fk0" FOREIGN KEY ("question_id") REFERENCES "question"("id");

ALTER TABLE "question_answer" ADD CONSTRAINT "question_answer_fk1" FOREIGN KEY ("answer_id") REFERENCES "answer"("id");