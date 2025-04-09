CREATE TABLE reviews (
    id VARCHAR(250) PRIMARY KEY,
    user_id VARCHAR(250) NOT NULL,
    prof_id VARCHAR(250) NOT NULL,
    comment TEXT,
    difficulty_rating DOUBLE PRECISION CHECK (difficulty_rating BETWEEN 0 AND 5),
    friendly_rating DOUBLE PRECISION CHECK (friendly_rating BETWEEN 0 AND 5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (prof_id) REFERENCES professors(id) ON DELETE CASCADE
);
