CREATE TABLE IF NOT EXISTS report_assigments(
    id serial PRIMARY KEY,
    course_report_id INTEGER REFERENCES report_courses(id),
    assigment_id INTEGER, -- this is assigment_id from canvas
    name VARCHAR(250),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)