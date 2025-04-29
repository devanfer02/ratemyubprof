CREATE TABLE review_reactions (
    user_id VARCHAR(250) NOT NULL,
    review_id VARCHAR(250) NOT NULL,
    reaction_type INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (review_id) REFERENCES reviews(id) ON DELETE CASCADE,
    CONSTRAINT unique_user_review_id UNIQUE (user_id, review_id)
);