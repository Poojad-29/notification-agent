CREATE TABLE notifications (
    id VARCHAR(50) PRIMARY KEY,
    recipient VARCHAR(255),
    channel VARCHAR(50),
    message TEXT,
    status VARCHAR(50),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
