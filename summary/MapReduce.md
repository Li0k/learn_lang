# MapReduce

## 梗概

- mr架构能在大量普通配置的计算机上实现并行化。这个系统在运行时只关心如何分割输入数据，集群调度，集群错误处理，集群中的通信。
- 灵感！！！来自函数式语言的map,reduce操作！！！在数据上应用Map操作得出一个中间key/value pair集合，然后在所以具有相同key值的value值上应用Reduce操作，合并中间数据得到想要的结果。

##  运行过程

1. 用户调用mr，将输入文件分成M个数据片段，大小一般（16m-64m），然后创建大量的程序副本
2. 有master分配任务，M个map，R个Reduce任务被分配给空闲的worker
3. 被分配了 map 任务的 worker 程序读取相关的输入数据片段，从输入的数据片段中解析出 key/value pair，然后把 key/value pair 传递给用户自定义的 MapF  函数，由 Map 函数生成并输出的中间 key/valuepair，并缓存在内存中。
4. 通过分区函数，分成R个分片，周期性写入磁盘，再将存储位置发给master。master再将位置通知Reduce worker
5. Reduce worker接受到master的数据后，通过rpc从map worker所在主机粗盘上读取数据。读取完所有数据后，对key进行排序，使得具有相同key的值聚合在一起
6. Reducer worker将排序好的中间数据（key,list(value)）传给相应的ReduceF函数。处理后输出到所属的分区。
7. Map,Reduce任务完成后，master返回给用户
8. mr一般输出R个文件，可以作为下一次mr程序的输入。

## Master数据结构

+ Master持有一些数据结构，存储每个Map，Reduce任务的状态，以及Worker机器的标识。
+ 对于每个状态为completed的map task，master会保存该map task对应的R个region的位置信息和大小信息。

## 容错

### worker故障

+ master 周期性的 ping 每个 worker。如果在一个约定的时间范围内没有收到 worker 返回的信息，master 将把这个 worker 标记为失效。
+ 如果某个worker崩溃了，那么在该worker上完成的所有map task都会回退，并调度到其他空间的worker上重新执行map。同样的，worker失效时正在执行的任务也将被重新调度。
+ ！针对运行好的map的任务，因为输出保存于本机磁盘，因此worker失效后数据也将无法访问，必须重新调度执行。但是reduce任务的输出存储在全局文件系统上（如gfs），因此无需重新调度执行。

### master故障

+ 让master定期的将数据结构写入磁盘，即checkpoint
+ 实际上只有一个master假定为故障率很低，若时效，可以根据需要重新执行mr



## 存储位置

+ 带宽资源匮乏，尽量把输入数据存在由gfs管理的机器的本地磁盘来节省带宽，gfs把每个文件按64mb一个block分割，一般存放多份拷贝（3）
+ MapReduce 的 master 在调度 Map 任务时会考虑输入文件的位置信息，尽量将一个 Map 任务调度在包含相关输入数据拷贝的机器上执行，若失败，则尝试在附近机器尝试执行。

## 任务粒度

+ 我们把 Map 拆分成了 M 个片段、把 Reduce 拆分成 R 个片段执行。理想情况下，M 和 R 应当比集群中 worker 的机器数量要多得多。在每台 worker 机器都执行大量的不同任务能够提高集群的动态的负载均衡能力，并且能够加快故障恢复的速度：失效机器上执行的大量 Map 任务都可以分布到所有其他的 worker机器上去执行。
+ 但实际上，master需要进行（M+R）次调度，并在内存中保存O(M*R)个中间文件，限制了任务的数量。
+ R 值通常是由用户指定的，因为每个 Reduce 任务最终都会生成一个独立的输出文件。实际使用时，倾向于选择合适的 M 值，以使得每一个独立任务都是处理大约 16M 到 64M 的输入数据（这样，上面描写的输入数据本地存储优化策略才最有效）。R设置为机器数目M的小倍数。（eg:我们通常会用这样的比例来执行 MapReduce：M=200000，R=5000，使用 2000 台 worker 机器。）

## 备用任务

+ 影响整个mr系统总执行时间，通常是有某个“落后者”，所以设计了一个通用机制来杜绝。当mr任务执行要完成时，调用backup-task来执行剩下的，处于in-progress的任务。无论是原始的task还是backup-task执行完成，都把task标记为完成。

## 技巧

+ 在某些情况下，Map 函数产生的中间 key 值的重复数据会占很大的比重。（eg：每个 Map 任务将产生成千上万个这样的记录）。所以允许用户指定一个可选的 combiner 函数，combiner 函数首先在本地将这些记录进行一次合并，然后将合并的结果再通过网络发送出去。一般情况下，Combiner 和 Reduce 函数是一样的。Combiner 函数和 Reduce 函数之间唯一的区别是 MapReduce 库怎样控制函数的输出。Reduce 函数的输出被保存在最终的输出文件里，而 Combiner 函数的输出被写到中间文件里，然后被发送Reduce 任务。









