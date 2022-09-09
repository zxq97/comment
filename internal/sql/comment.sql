CREATE TABLE comment_subject
(
    `id`          BIGINT    NOT NULL AUTO_INCREMENT,
    `obj_id`      BIGINT    NOT NULL DEFAULT 0,
    `obj_type`    TINYINT   NOT NULL DEFAULT 0,
    `uid`         BIGINT    NOT NULL DEFAULT 0 COMMENT '被评论作品用户id',
    `count`       INT       NOT NULL DEFAULT 0,
    `root_count`  INT       NOT NULL DEFAULT 0,
    `state`       TINYINT   NOT NULL DEFAULT 0 COMMENT '0: 放出 1:隐藏',
    `attre`       INT       NOT NULL DEFAULT 0 COMMENT 'bit 0:人工置顶 1:up置顶',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE unid_objid_objtype (`obj_id`, `obj_type`)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE comment_index
(
    `id`          BIGINT    NOT NULL DEFAULT 0,
    `obj_id`      BIGINT    NOT NULL DEFAULT 0,
    `obj_type`    TINYINT   NOT NULL DEFAULT 0,
    `uid`         BIGINT    NOT NULL DEFAULT 0,
    `root`        BIGINT    NOT NULL DEFAULT 0,
    `parent`      BIGINT    NOT NULL DEFAULT 0 COMMENT '回复id',
    `floor`       INT       NOT NULL DEFAULT 0 COMMENT '楼层',
    `count`       INT       NOT NULL DEFAULT 0 COMMENT '该评论子评论总数',
    `like`        INT       NOT NULL DEFAULT 0,
    `hate`        INT       NOT NULL DEFAULT 0,
    `state`       TINYINT   NOT NULL DEFAULT 0 COMMENT '0: 放出 1:隐藏',
    `attre`       INT       NOT NULL DEFAULT 0 COMMENT 'bit 0:人工置顶 1:up置顶',
    `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY           idx_objid_objtype_floor (`obj_id`, `obj_type`, `floor`)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

ALTER TABLE comment_index ADD INDEX idx_objid_objtype_root_parent (`obj_id`, `obj_type`, `root`, `parent`);

CREATE TABLE comment_content
(
    `comment_id`  BIGINT        NOT NULL DEFAULT 0,
    `uid`         BIGINT        NOT NULL DEFAULT 0,
    `ip`          INT           NOT NULL DEFAULT 0,
    `message`     VARCHAR(1024) NOT NULL DEFAULT '',
    `create_time` TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`comment_id`)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
