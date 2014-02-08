# API 

## Containers

### Running a single command

```
netcat 127.0.0.1 5000 -o /dev/stdout
```
```JSON
{
  "id":3,
  "method":"LXC.Attach",
  "params":[
    {
      "name":"ubuntu",
      "command":[
        "uptime"
      ]
    }
  ]
}
```

**Response**
```JSON
{
  "id":3,
  "result":{
    "StatusCode":0,
    "Stdout":" 15:25:15 up 17:36,  0 users,  load average: 0.00, 0.01, 0.05\n",
    "Stderr":""
  },
  "error":null
}
```
