#### go_gin_api.place 
地点表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 主键 | int(11) unsigned | PRI | NO | auto_increment |  |
| 2 | name | 名称 | varchar(255) |  | NO |  |  |
| 3 | content | 详情 | text |  | YES |  |  |
| 4 | province | 省份 | varchar(255) |  | YES |  |  |
| 5 | city | 城市 | varchar(255) |  | YES |  |  |
| 6 | county | 区县 | varchar(255) |  | YES |  |  |
| 7 | img | 图片 | varchar(255) |  | YES |  |  |
| 8 | address | 地址 | varchar(255) |  | YES |  |  |
| 9 | is_deleted | 是否删除 1:是  -1:否 | tinyint(1) |  | NO |  | -1 |
| 10 | created_at | 创建时间 | timestamp |  | NO |  | CURRENT_TIMESTAMP |
| 11 | created_user | 创建人 | varchar(60) |  | NO |  |  |
| 12 | updated_at | 更新时间 | timestamp |  | NO | on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 13 | updated_user | 更新人 | varchar(60) |  | NO |  |  |
