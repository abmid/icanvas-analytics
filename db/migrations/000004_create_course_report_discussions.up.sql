CREATE TABLE IF NOT EXISTS report_discussions(
    id serial PRIMARY KEY,
    course_report_id INTEGER REFERENCES report_courses(id),
    discussion_id INTEGER, -- this is from canvas discussion ID
    title VARCHAR(250),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
)