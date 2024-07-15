CREATE TABLE IF NOT EXISTS articles (
  `id`         VARCHAR(36)           NOT NULL,
  `type`       ENUM('BLOG', 'PAGE')  NOT NULL,
  `title`      VARCHAR(500)          NOT NULL,
  `createdAt`  DATETIME              NOT NULL,
  `updatedAt`  DATETIME              NOT NULL,
  `summary`    VARCHAR(1000)         NOT NULL,
  `content`    TEXT                  NOT NULL,

  PRIMARY KEY (`id`),
  FULLTEXT(`title`, `summary`)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS tags (
  `key`   VARCHAR(30)  NOT NULL,
  `name`  VARCHAR(30)  NOT NULL,

  PRIMARY KEY (`key`)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS article_tag_links (
  `article_id`  VARCHAR(36)  NOT NULL,
  `tag_key`     VARCHAR(30)  NOT NULL,

  PRIMARY KEY (`article_id`, `tag_key`),
  FOREIGN KEY (`article_id`) REFERENCES articles(`id`),
  FOREIGN KEY (`tag_key`) REFERENCES tags(`key`)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS file_entries (
  `id`             VARCHAR(36)   NOT NULL,
  `title`          VARCHAR(500)  NOT NULL,
  `mime_type`      VARCHAR(30)   NOT NULL,
  `size_in_bytes`  INT UNSIGNED  NOT NULL,

  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS pages (
  `key`         VARCHAR(30)  NOT NULL,
  `menu_name`   VARCHAR(30)  NOT NULL,
  `article_id`  VARCHAR(36)  NOT NULL,

  PRIMARY KEY (`key`),
  FOREIGN KEY (`article_id`) REFERENCES articles(`id`)
) ENGINE=InnoDB;
