# QueueWriter

This is a simple service that exposes a REST endpoint to POST a payload, which is then written to an SQS queue provided by the user.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

1. Install Go - https://golang.org/doc/install

2. Setup AWS credentials in your environment -

```
// Linux or OS X
$ export AWS_ACCESS_KEY_ID=YOUR_AKID
$ export AWS_SECRET_ACCESS_KEY=YOUR_SECRET_KEY

// Windows
> set AWS_ACCESS_KEY_ID=YOUR_AKID
> set AWS_SECRET_ACCESS_KEY=YOUR_SECRET_KEY
```

3. Create an SQS queue - https://docs.aws.amazon.com/cli/latest/reference/sqs/create-queue.html

## Running the service

1. Clone the git repo and navigate to the parent directory

```
$ git clone https://github.com/akhani18/QueueWriter.git
...
$ cd QueueWriter
```

2. Start the Service

```
$ go run main.go -n <SQS-Queue-Name> -r <aws-region>
```

3. Send a payload to the Queue

```
curl -d '{"key1":"value1", "key2":"value2"}' -H "Content-Type: application/json" -X POST http://localhost:8000/message
```

4. or, send a payload with a file

```
curl -d "@data.json" -H "Content-Type: application/json" -X POST http://localhost:8000/message
```

## Benchmarking Results

Total requests: 1000

Concurrent Requests: 10

Payload Size: 1Kb

```
$ ab -p post.txt -T application/json -n 1000 -c 10 http://localhost:8000/message
```

```
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

```
