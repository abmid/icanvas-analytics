CREATE TABLE IF NOT EXISTS report_course_results(
    id serial PRIMARY KEY,
    report_course_id integer references report_courses(id),
    assigment_count integer,
    discussion_count integer,
    student_count integer,
    finish_grading_count integer,
    final_score float,
    created_at timestamp,
    updated_at timestamp
);