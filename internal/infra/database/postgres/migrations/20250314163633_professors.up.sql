CREATE TABLE professors (
    id VARCHAR(250) PRIMARY KEY,
    name VARCHAR(250) NOT NULL,
    faculty VARCHAR(200) NOT NULL,
    major VARCHAR(200) NOT NULL,
    profile_img_link VARCHAR(250),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);