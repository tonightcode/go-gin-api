#### go_gin_api.relations 
事件-人物-地点关系表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | int(10) unsigned | PRI | NO | auto_increment |  |
| 2 | event_id |  | int(10) unsigned |  | YES |  |  |
| 3 | place_id |  | int(11) |  | YES |  |  |
| 4 | person_id |  | int(11) |  | YES |  |  |
| 5 | createtime |  | datetime |  | YES |  |  |
| 6 | updatetime |  | datetime |  | YES |  |  |
