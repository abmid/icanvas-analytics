CREATE TABLE IF NOT EXISTS report_courses(
   id serial PRIMARY KEY,
   course_id integer NOT NULL,
   course_name VARCHAR(250) NOT NULL,
   account_id integer NOT NULL,
   created_at timestamp, 
   updated_at timestamp,
   deleted_at timestamp
);