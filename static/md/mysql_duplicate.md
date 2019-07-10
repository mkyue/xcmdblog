**首先，任何情况下慎用INSERT ... ON DUPLICATE KEY UPDATE**

#####1. 多个唯一性索引配合DUPLICATE KEY UPDATE引起更新不安预期进行

建立唯一性索引可以确定数据表不存在重复的记录行,但是在特殊情况下可能出现问题。如下表
```sql
CREATE TABLE `test` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`a` int(11) unsigned NOT NULL,
`b` int(11) unsigned NOT NULL,
`c` int(11) unsigned NOT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
建立两个唯一索引
```sql
ALTER TABLE `test` ADD UNIQUE INDEX `idx_unique_a` (`a`) USING BTREE ;
ALTER TABLE `test` ADD UNIQUE INDEX `idx_unique_b` (`b`) USING BTREE ;
```
写入两条记录
```sql
INSERT INTO `ycx`.`test` (`id`, `a`, `b`, `c`) VALUES ('1', '1', '1', '1');
INSERT INTO `ycx`.`test` (`id`, `a`, `b`, `c`) VALUES ('2', '2', '2', '2');
```
执行insert on duplicate key
```sql
INSERT INTO test (a, b, c) VALUES(1, 2, 3) ON DUPLICATE KEY UPDATE a = 1111
```
执行结束后结果如下

| id        | a      |  b      | c       |
| --------  | -----: | :----:  | :----:  |
| 1         | 1111   |   1     | 1       |
| 2         |   2    |   2     |   2     |

`在多个唯一索引情况下,使用DUPLICATE KEY UPDATE只会更新一条记录,因此在多个唯一索引的表中应慎用该语句`

#####2. INSERT ... ON DUPLICATE KEY UPDATE引起主键不连续
mysql默认配置:innodb_autoinc_lock_mode = 1,MySQL 5.1开始引入该配置,并且线上环境推荐使用配置1,用于生成自动增量值的锁定模式。，允许值分别为0,1或2。默认设置为1（连续）。
- 设置为0： 在该模式下， 所有的insert语句都会获得表级AUTO-INC锁，用于插入存在AUTO_INCREMENT字段的表,锁在执行实际写入后才会释放,确保自增列连续。但是意味着高并发下大量写入等待问题
- 设置为1(默认)： 这是默认的锁定模式, 在此模式下,对于已知写入行数的情况下(非批量写入),INSERT语句通过互斥锁控制获取需要的自增值后解除锁定,不会等待实际执行INSERT，除非另一个事务持有AUTO-INC锁，否则不使用表级AUTO-INC锁。速度更快,同时保证连续性
- 设置为2： 此语句下,所有INSERT语句不使用AUTO-INC锁,并且多个语句可以同时执行。这是最快和最具有扩展性的锁定模式,但是并不安全

`当执行INSERT ... ON DUPLICATE KEY UPDATE语句时,在默认配置1的情况下，先通过互斥锁获取自增值,然后开始实际写入。当唯一索引出发update时,并没有实际写入.导致主键不连续,大量占用主键范围`


 
