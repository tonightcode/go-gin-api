#### go_gin_api.swiper 
轮播图

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 主键 | int(11) unsigned | PRI | NO | auto_increment |  |
| 2 | title | 标题 | varchar(255) |  | NO |  |  |
| 3 | url | 地址 | varchar(255) |  | NO |  |  |
| 4 | ref_type | 类型 1事件 2地点 3人物 4购物 5自定义 | tinyint(1) |  | NO |  | 1 |
| 5 | ref_id | 引用的ID | int(11) |  | NO |  | 0 |
| 6 | is_deleted | 是否删除 1:是  -1:否 | tinyint(1) |  | NO |  | -1 |
| 7 | created_at | 创建时间 | timestamp |  | NO |  | CURRENT_TIMESTAMP |
| 8 | created_user | 创建人 | varchar(60) |  | NO |  |  |
| 9 | updated_at | 更新时间 | timestamp |  | NO | on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 10 | updated_user | 更新人 | varchar(60) |  | NO |  |  |
