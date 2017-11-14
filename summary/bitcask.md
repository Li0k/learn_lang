# bitcask

+ 存储系统的基本功能包括：增、删、读、改。其中读取操作有分为顺序读取和随机读取。Bitcask模型是一种日志型键值模型。所谓日志型，是指它不直接支持随机写入，而是像日志一样支持追加操作

## bitcask优点

+ 将随机写入转化为顺序写入，提高吞吐量，写操作无须查找，直接追加
+ 如果使用SSD作为存储介质，能够更好的利用新硬件的特性？

## bitcask数据结构

1. bitcask中任一时刻，只有一个文件可以写入，称为active data file，其他文件称为older data file，保留之前写入的文件。

2. bitcask文件，每行的数据格式

   ```
   crc(后几项的crc值) tstamp(时间戳) key_sz value_sz key value
   ```

3. 索引表中，记录了数据的主键与位置信息，存放于内存中。数据结构为

   ```
   file_id(文件编号) value_len value_pos tstamp 
   ```

4. hint file 的记录与数据文件的格式基本相同，唯一不同的是数据文件记录数据的值，而线索文件(hint file)则是记录数据的位置。

   ```
   tstamp key_sz value_sz value_pos key
   ```

   hint file用以bitcask重启时重建内存中的索引表，bitcask只读取hint file而不是去遍历所有数据，加快索引表的构建。

## bitcask基本功能

### insert

用户的写入直接追加到 active data file,但是无限制的追加会导致文件过大，所以到达指定大小时，新建一个active data file，将原来的转为older data file。写入记录的同时还要在索引表中添加索引记录。

### delete

不删除文件，而是追加一条相同key的记录，把value标记为删除，然后更新索引表。

### update

bitcask不支持随机写入，只能直接追加记录，并修改索引表。存在多个相同key的记录，可以通过tstamp来判断最新数据。

### read

直接根据索引表记录的file与pos找到记录在磁盘中的记录，读取key

### merge

bitcask直接追加的特性，导致会有很多冗余或者说过期的记录，例如被标记为删除或者追加修改的记录。merge可以合并这部分数据，减少占用空间。

merge操作，通过定期将所有older data file中的数据扫描一遍并生成新的data file（没有包括active data file 是因为它还在不停写入）。并且在执行merge时可以顺带生成hint file。