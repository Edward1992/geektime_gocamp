### 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

``` shell 
uname -a
# Darwin XX 19.6.0 Darwin Kernel Version 19.6.0: Tue Nov 10 00:10:30 PST 2020; root:xnu-6153.141.10~1/RELEASE_X86_64 x86_64
```
``` shell
redis-server -v
# Redis server v=5.0.6 sha=00000000:0 malloc=libc bits=64 build=ff6bda879ec1
```

---
``` shell
redis-benchmark -t get,set -d 10
```
```
====== SET ======
  100000 requests completed in 1.47 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

95.99% <= 1 milliseconds
99.43% <= 2 milliseconds
99.72% <= 3 milliseconds
99.79% <= 4 milliseconds
99.79% <= 5 milliseconds
99.81% <= 6 milliseconds
99.83% <= 9 milliseconds
99.86% <= 10 milliseconds
99.93% <= 11 milliseconds
99.95% <= 14 milliseconds
99.95% <= 15 milliseconds
99.99% <= 16 milliseconds
100.00% <= 16 milliseconds
67842.61 requests per second

====== GET ======
  100000 requests completed in 1.38 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1

96.86% <= 1 milliseconds
99.58% <= 2 milliseconds
99.85% <= 3 milliseconds
99.97% <= 4 milliseconds
99.99% <= 5 milliseconds
100.00% <= 5 milliseconds
72202.16 requests per second
```
---
``` shell
redis-benchmark -t get,set -d 20
```
```
====== SET ======
  100000 requests completed in 1.39 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

97.62% <= 1 milliseconds
99.65% <= 2 milliseconds
99.82% <= 3 milliseconds
99.87% <= 4 milliseconds
99.90% <= 5 milliseconds
99.91% <= 6 milliseconds
99.93% <= 8 milliseconds
99.93% <= 9 milliseconds
99.95% <= 17 milliseconds
99.96% <= 18 milliseconds
100.00% <= 18 milliseconds
71684.59 requests per second

====== GET ======
  100000 requests completed in 1.33 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1

98.32% <= 1 milliseconds
99.94% <= 2 milliseconds
100.00% <= 2 milliseconds
75471.70 requests per second
```
---
``` shell
redis-benchmark -t get,set -d 50
```
```
====== SET ======
  100000 requests completed in 1.43 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

96.90% <= 1 milliseconds
99.78% <= 2 milliseconds
99.87% <= 3 milliseconds
99.99% <= 4 milliseconds
100.00% <= 4 milliseconds
69930.07 requests per second

====== GET ======
  100000 requests completed in 1.43 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1

96.52% <= 1 milliseconds
99.51% <= 2 milliseconds
99.90% <= 3 milliseconds
99.97% <= 4 milliseconds
100.00% <= 5 milliseconds
100.00% <= 5 milliseconds
70028.02 requests per second
```
---
``` shell
redis-benchmark -t get,set -d 100
```
```
====== SET ======
  100000 requests completed in 1.45 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

96.41% <= 1 milliseconds
99.72% <= 2 milliseconds
99.90% <= 3 milliseconds
99.93% <= 4 milliseconds
99.98% <= 5 milliseconds
99.99% <= 6 milliseconds
100.00% <= 6 milliseconds
69013.11 requests per second

====== GET ======
  100000 requests completed in 1.39 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1

97.15% <= 1 milliseconds
99.78% <= 2 milliseconds
99.92% <= 3 milliseconds
99.95% <= 10 milliseconds
99.99% <= 11 milliseconds
100.00% <= 11 milliseconds
71787.51 requests per second
```
---
``` shell
redis-benchmark -t get,set -d 200
```
```
====== SET ======
  100000 requests completed in 1.53 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1

93.98% <= 1 milliseconds
99.06% <= 2 milliseconds
99.85% <= 3 milliseconds
99.95% <= 4 milliseconds
100.00% <= 4 milliseconds
65530.80 requests per second

====== GET ======
  100000 requests completed in 1.38 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1

97.01% <= 1 milliseconds
99.70% <= 2 milliseconds
99.93% <= 3 milliseconds
99.95% <= 10 milliseconds
99.97% <= 11 milliseconds
100.00% <= 12 milliseconds
100.00% <= 12 milliseconds
72358.90 requests per second
```
---
``` shell
redis-benchmark -t get,set -d 1024
```
```
====== SET ======
  100000 requests completed in 1.54 seconds
  50 parallel clients
  1024 bytes payload
  keep alive: 1

94.51% <= 1 milliseconds
99.34% <= 2 milliseconds
99.77% <= 3 milliseconds
99.86% <= 4 milliseconds
99.92% <= 5 milliseconds
99.95% <= 7 milliseconds
99.95% <= 8 milliseconds
99.96% <= 9 milliseconds
99.98% <= 10 milliseconds
100.00% <= 11 milliseconds
64724.92 requests per second

====== GET ======
  100000 requests completed in 1.38 seconds
  50 parallel clients
  1024 bytes payload
  keep alive: 1

97.12% <= 1 milliseconds
99.70% <= 2 milliseconds
99.96% <= 3 milliseconds
99.96% <= 4 milliseconds
99.97% <= 5 milliseconds
100.00% <= 5 milliseconds
72202.16 requests per second
```
---
``` shell
redis-benchmark -t get,set -d 5120
```
```
====== SET ======
  100000 requests completed in 1.66 seconds
  50 parallel clients
  5120 bytes payload
  keep alive: 1

91.64% <= 1 milliseconds
99.10% <= 2 milliseconds
99.76% <= 3 milliseconds
99.86% <= 4 milliseconds
99.91% <= 5 milliseconds
99.93% <= 6 milliseconds
99.94% <= 7 milliseconds
99.95% <= 12 milliseconds
99.96% <= 13 milliseconds
99.96% <= 19 milliseconds
100.00% <= 19 milliseconds
60132.29 requests per second

====== GET ======
  100000 requests completed in 1.55 seconds
  50 parallel clients
  5120 bytes payload
  keep alive: 1

95.82% <= 1 milliseconds
99.26% <= 2 milliseconds
99.76% <= 3 milliseconds
99.81% <= 4 milliseconds
99.89% <= 5 milliseconds
99.92% <= 6 milliseconds
99.94% <= 7 milliseconds
99.94% <= 10 milliseconds
99.95% <= 41 milliseconds
99.98% <= 42 milliseconds
100.00% <= 42 milliseconds
64474.53 requests per second
```
---

### 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

示例程序
``` go
for i := 0; i < KEY_AMOUNT; i++ {
		key := faker.StringWithSize(12)
		value := faker.StringWithSize(5120) // 10 20 50 100 200 1024 5120
		redis_client.Set(key, value, 0)
	}
```
---
1. 写入 10 bytes key,前后内存占用 used_memory:1051424 --> used_memory:5576528, 平均每个字节占用 90.5 bytes
2. 写入 20 bytes key,前后内存占用 used_memory:1052432 --> used_memory:6376720, 平均每个字节占用 106.5 bytes
3. 写入 50 bytes key,前后内存占用 used_memory:1052624 --> used_memory:7976912, 平均每个字节占用 138.5 bytes
4. 写入 100 bytes key,前后内存占用 used_memory:1052816 --> used_memory:10377104, 平均每个字节占用 186.5 bytes
5. 写入 200 bytes key,前后内存占用 used_memory:1053008 --> used_memory:15177296, 平均每个字节占用 282.5 bytes
6. 写入 1024 bytes key,前后内存占用 used_memory:1053200 --> used_memory:81577488, 平均每个字节占用 1610.5 bytes
7. 写入 5120 bytes key,前后内存占用 used_memory:1053776 --> used_memory:286378064, 平均每个字节占用 5076.5 bytes
