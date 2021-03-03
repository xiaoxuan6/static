# Myisan和Innodb索引的区别？

MyISAM 索引的实现，和 InnoDB 索引的实现是一样使用 B+Tree 

    差别：在于 MyISAM 索引文件和数据文件是分离的，索引文件仅保存数据记录的地址。