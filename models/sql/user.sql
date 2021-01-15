-- 为什么不使用自增id作为用户id
-- 1.别人可以通过注册用户来获取到数据库的用户量
-- 2.分库分表的时候用户id可能会重复
-- 使用分布式ID生成器
create table user (
    id bigint(20) not null auto_increment,
    user_id bigint(20) not null ,
    username varchar(64) not null,
    password varchar(64) not null,
    email varchar (64),
    gender tinyint(4) not null default 0,
    class_id tinyint(4),
    phone varchar(11),
    create_time timestamp null default current_timestamp ,
    update_time timestamp null  default current_timestamp on update current_timestamp ,
    primary key(id),
    unique key idx_username(username),
    unique key idx_user_id(user_id),
    unique key idx_class_id(class_id)
) engine=innodb default charset=utf8mb4 collate=utf8mb4_general_ci;