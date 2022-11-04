CREATE TABLE IF NOT EXISTS users (
    id             BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name           VARCHAR(100) NOT NULL,
    email          VARCHAR(2048) NOT NULL,
    oidc_provider  VARCHAR(50) NOT NULL,
    token          VARCHAR(2048) NOT NULL,
    created_at     TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at     TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
);

CREATE TABLE IF NOT EXISTS site (
    id             BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    url            VARCHAR(2048) NOT NULL,
    created_at     TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at     TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
);

CREATE TABLE IF NOT EXISTS bookmark (
    id             BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id        BIGINT(20) NOT NULL,
    site_id        BIGINT(20) NOT NULL,
    title          VARCHAR(100) NOT NULL,
    note           VARCHAR(2048) DEFAULT NULL,
    created_at     TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at     TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
);

CREATE TABLE IF NOT EXISTS tag (
    id             BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id        BIGINT(20) NOT NULL,
    site_id        BIGINT(20) NOT NULL,
    tag            VARCHAR(100) NOT NULL,
    created_at     TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at     TIMESTAMP NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
);

INSERT INTO users (name, email, oidc_provider, token) VALUES ('tooi', 'oomori.lunch@gmail.com', 'Github', 'hogehoge');

INSERT INTO site (url) VALUES ('https://www.google.com');
INSERT INTO site (url) VALUES ('https://www.yahoo.co.jp');

INSERT INTO bookmark (user_id, site_id, title, note) VALUES (1, 1, 'google', '今日の帰りに夕食の材料を買う');
INSERT INTO bookmark (user_id, site_id, title, note) VALUES (1, 2, 'yahoo', 'TOEICの勉強を１時間やる');

INSERT INTO tag (user_id, site_id, tag) VALUES (1, 1, 'web');
INSERT INTO tag (user_id, site_id, tag) VALUES (1, 2, 'web');
