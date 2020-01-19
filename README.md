# QueueWriter




## Benchmarking Results
Total requests: 1000
Concurrent Requests: 10
Payload Size: 1Kb

ab -p post.txt -T application/json -n 1000 -c 10 http://localhost:8000/message
This is ApacheBench, Version 2.3 <$Revision: 1826891 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8000

Document Path:          /message
Document Length:        0 bytes

Concurrency Level:      10
Time taken for tests:   4.450 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      81000 bytes
Total body sent:        1045000
HTML transferred:       0 bytes
Requests per second:    224.69 [#/sec] (mean)
Time per request:       44.505 [ms] (mean)
Time per request:       4.450 [ms] (mean, across all concurrent requests)
Transfer rate:          17.77 [Kbytes/sec] received
                        229.30 kb/s sent
                        247.08 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       1
Processing:    15   43  54.3     27     465
Waiting:       15   42  54.3     27     465
Total:         15   43  54.3     27     465

Percentage of the requests served within a certain time (ms)
  50%     27
  66%     33
  75%     39
  80%     44
  90%     63
  95%     94
  98%    281
  99%    312
 100%    465 (longest request)
