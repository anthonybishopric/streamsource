# Stream Source

CLI with a server that generates several streams of test messages.

```
$ streamsource -h
Usage of streamsource:
  -delay
        Add a delay into message responses
  -nil
        Sometimes return nothing
  -streams int
        Number of streams to create (default 3)
$ streamsource
2017/07/10 11:24:24 Serving messages on port 8000
2017/07/10 11:24:24 Serving messages on port 8001
2017/07/10 11:24:24 Serving messages on port 8002

-- in a seperate terminal --

$ curl localhost:8000 | jq
{
  "timestamp": 4,
  "message": "144.225.139.29: sofiagarcia341@example.org braid -- Europe/Bucharest"
}
$ curl localhost:8000 | jq
{
  "timestamp": 5,
  "message": "245.84.179.4: aubreywhite058@example.com leg -- America/Rosario"
}

```