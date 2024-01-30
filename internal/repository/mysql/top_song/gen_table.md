#### go_gin_api.top_song 
排行榜-歌曲

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | topid | 排行榜ID | int(11) unsigned |  | NO |  |  |
| 2 | songid | 声音ID | int(11) unsigned |  | NO |  |  |
| 3 | is_deleted | 是否删除 1:是 0 隐藏 -1:否 | tinyint(1) |  | NO |  | -1 |
| 4 | created_at | 创建时间 | timestamp |  | NO |  | CURRENT_TIMESTAMP |
| 5 | created_user | 创建人 | varchar(60) |  | NO |  |  |
| 6 | updated_at | 更新时间 | timestamp |  | NO | on update CURRENT_TIMESTAMP | CURRENT_TIMESTAMP |
| 7 | updated_user | 更新人 | varchar(60) |  | NO |  |  |
