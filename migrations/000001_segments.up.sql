CREATE TABLE IF NOT EXISTS segments(
    biz_tag     varchar(32) not null,
    max_id      bigint       null,
    step        int          null,
    remark      varchar(200) null,
    create_time datetime       null,
    update_time datetime       null,
    constraint segments_pk
        primary key (biz_tag)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;

INSERT INTO segments(`biz_tag`, `max_id`, `step`, `remark`, `create_time`, `update_time`)
VALUES ('test', 0, 100, 'test', NOW(), NOW());