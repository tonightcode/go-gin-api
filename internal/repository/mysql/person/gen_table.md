#### go_gin_api.person 
人物表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 主键 | int(11) unsigned | PRI | NO | auto_increment |  |
| 2 | username | 名字 | varchar(255) |  | NO |  |  |
| 3 | intro | 详情 | text |  | YES |  |  |
| 4 | icon | 头像 | varchar(255) |  | YES |  |  |
| 5 | is_deleted | 是否删除 1:是  -1:否 | tinyint(1) |  | NO |  | -1 |
| 6 | created_at | 创建时间 | timestamp |  | NO |  | CURRENT_TIMESTAMP |
| 7 | created_user | 创建人 | varchar(60) |  | NO |  |  |
| 8 | updated_at | 更新时间 | timestamp |  | NO | on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 9 | updated_user | 更新人 | varchar(60) |  | NO |  |  |
