CREATE TABLE reviews (
    id VARCHAR(250) NOT NULL PRIMARY KEY,
    user_id VARCHAR(250) NOT NULL,
    prof_id VARCHAR(250) NOT NULL,
    comment TEXT,
    difficulty_rating DOUBLE PRECISION CHECK (difficulty_rating BETWEEN 0 AND 5),
    friendly_rating DOUBLE PRECISION CHECK (friendly_rating BETWEEN 0 AND 5),
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (prof_id) REFERENCES professors(id) ON DELETE CASCADE,
    CONSTRAINT unique_user_prof_id UNIQUE (user_id, prof_id)
);
