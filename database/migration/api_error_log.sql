CREATE TABLE api_error_log (
    error_id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    timestamp DATETIME NOT NULL,
    error_message TEXT NOT NULL
)
