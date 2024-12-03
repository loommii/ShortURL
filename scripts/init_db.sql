CREATE TABLE short_urls (
    id INT AUTO_INCREMENT PRIMARY KEY,
    short_url VARCHAR(10) NOT NULL,  -- 短链接，长度可以根据需要调整
    long_url TEXT NOT NULL,          -- 长链接
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 创建时间
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- 更新时间
    UNIQUE KEY(short_url)           -- 确保短链接唯一
);
