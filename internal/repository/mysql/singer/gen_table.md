#### go_gin_api.singer 
歌手

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | singerid | 主键 | int(11) unsigned | PRI | NO | auto_increment |  |
| 2 | name | 名称 | varchar(255) |  | NO |  |  |
| 3 | head | 头像 | varchar(255) |  | NO |  |  |
| 4 | is_deleted | 是否删除 1:是 0 隐藏 -1:否 | tinyint(1) |  | NO |  | -1 |
| 5 | created_at | 创建时间 | timestamp |  | NO |  | CURRENT_TIMESTAMP |
| 6 | created_user | 创建人 | varchar(60) |  | NO |  |  |
| 7 | updated_at | 更新时间 | timestamp |  | NO | on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 8 | updated_user | 更新人 | varchar(60) |  | NO |  |  |
