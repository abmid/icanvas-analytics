CREATE TABLE IF NOT EXISTS report_enrollments(
    id serial PRIMARY KEY,
    course_report_id INTEGER REFERENCES report_courses(id),
    enrollment_id INTEGER,
    user_id INTEGER,
    login_id VARCHAR(255),
    full_name VARCHAR(255),
    role_id INTEGER,
    role VARCHAR(50),
    current_score FLOAT,
    current_grade FLOAT, 
    final_score FLOAT,
    final_grade FLOAT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP    
)