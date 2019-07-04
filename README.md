# GoProxyPoll

### Introduction

GoProxyPoll is a lightweight proxy poll aiming to help you get free proxy.

### Start

*   Make sure your Go version >= 1.10

*   Make sure you've installed mysql and version >= 8.0

*   Download

```$xslt
go get -u github.com/authetic-x/GoProxyPoll
```

*   Start up

```$xslt
cd GoProxyPoll
go run main.go 
```

### Usage

*   Get an Ip

```$xslt
GET http://localhost:8000/get
```

response

```$xslt
{   "Ip": "21.32.189.24",
    "Protocol": "HTTPS"
}
```

*   Get the count

```$xslt
GET http://localhost:8000/count
```

## Data Source

*   http://www.feiyiproxy.com/?page_id=1457
