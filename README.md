This repository contains various [Joe query files](https://github.com/evilsocket/joe) for the [PwnGRID](http://github.com/evilsocket/pwngrid) database running on [api.pwnagotchi.ai](https://pwnagotchi.ai/api/grid/).

## Joe Queries

This document has been automatically generated with [joe v1.0.0a](https://github.com/evilsocket/joe).

### GET|POST /api/v1/auth  

Authenticate to the API with username and password in order to get a JWT token.

|Parameter|Default|
|--|--|
| `user` | _none_ |
| `pass` | _none_ |

#### Request

    curl --data "user=admin&pass=admin" http://127.0.0.1:8080/api/v1/auth

#### Response

```json
{
  "token": "..."
}
```

### GET /api/v1/queries/

Get a list of the queries that are currently loaded in the system.

#### Request

    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/queries

#### Response

```json
[
  {
    "created_at": "2019-11-09T10:57:01.784331255+01:00",
    "updated_at": "2019-11-09T10:57:01.784494235+01:00",
    "name": "top_players",
    "cache": {
      "type": 1,
      "keys": [
        "limit"
      ],
      "ttl": 3600
    },
    "description": "Top players by access points.",
    "defaults": {
      "limit": 25
    },
    "query": "SELECT  u.updated_at as active_at, u.name, u.fingerprint, u.country, COUNT(a.id) AS networks FROM units u INNER JOIN access_points a ON u.id = a.unit_id GROUP BY u.id ORDER BY networks DESC LIMIT {limit}",
    "views": {
      "bars": "top_players_bars.go"
    },
    "access": [
      "admin"
    ]
  },
  ...
}
```





### GET /api/v1/query/units_per_day/view  

Show information about the units_per_day query.

#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/units_per_day/view


#### Response

```json
{
  "created_at": "2019-11-09T14:52:22.549358143+01:00",
  "updated_at": "2019-11-09T14:52:22.550049163+01:00",
  "name": "units_per_day",
  "cache": {
    "type": 1,
    "keys": [
      "days"
    ],
    "ttl": 3600
  },
  "description": "Registrations per day.",
  "defaults": {
    "days": 30
  },
  "query": "SELECT DATE_FORMAT(created_at, '%Y-%m-%d') AS day, COUNT(id) AS registered FROM units GROUP BY day ORDER BY day DESC LIMIT {days};",
  "views": {
    "chart": "units_per_day_timeseries.go"
  },
  "access": [
    "admin"
  ]
}
```

### GET|POST /api/v1/query/units_per_day(.json|.csv)

Registrations per day.


|Parameter|Default|
|--|--|
| `days` | 30 |



#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/units_per_day.json?days=30


#### Response

```json
{
  "created_at": "2019-11-09T14:52:22.549358143+01:00",
  "updated_at": "2019-11-09T14:52:22.550049163+01:00",
  "name": "units_per_day",
  "cache": {
    "type": 1,
    "keys": [
      "days"
    ],
    "ttl": 3600
  },
  "description": "Registrations per day.",
  "defaults": {
    "days": 30
  },
  "query": "SELECT DATE_FORMAT(created_at, '%Y-%m-%d') AS day, COUNT(id) AS registered FROM units GROUP BY day ORDER BY day DESC LIMIT {days};",
  "views": {
    "chart": "units_per_day_timeseries.go"
  },
  "access": [
    "admin"
  ]
}
```

### GET|POST /api/v1/query/units_per_day/explain

Return results for an EXPLAIN operation on the units_per_day main query.


|Parameter|Default|
|--|--|
| `days` | 30 |



#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/units_per_day/explain?days=30


#### Response

```json
{
  "cached_at": null,
  "exec_time": 2763263,
  "num_records": 1,
  "records": [
    {
      "Extra": "Zero limit",
      "filtered": null,
      "id": 1,
      "key": null,
      "key_len": null,
      "partitions": null,
      "possible_keys": null,
      "ref": null,
      "rows": null,
      "select_type": "SIMPLE",
      "table": null,
      "type": null
    }
  ]
}
```
 






### GET /api/v1/query/units_per_day/chart.(png|svg)

Return a PNG or SVG representation of a chart chart for the units_per_day query.


|Parameter|Default|
|--|--|
| `days` | 30 |



#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/units_per_day/chart.png?days=30






### GET /api/v1/query/aps_per_day/view  

Show information about the aps_per_day query.

#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/aps_per_day/view


#### Response

```json
{
  "created_at": "2019-11-09T16:28:16.506359608+01:00",
  "updated_at": "2019-11-09T16:28:16.506491743+01:00",
  "name": "aps_per_day",
  "cache": {
    "type": 1,
    "keys": [
      "days"
    ],
    "ttl": 3600
  },
  "description": "Reported access points per day.",
  "defaults": {
    "days": 30
  },
  "query": "SELECT DATE_FORMAT(created_at, '%Y-%m-%d') AS day, COUNT(id) AS reported FROM access_points GROUP BY day ORDER BY day DESC LIMIT {days};",
  "views": {
    "chart": "aps_per_day_timeseries.go"
  },
  "access": [
    "admin"
  ]
}
```

### GET|POST /api/v1/query/aps_per_day(.json|.csv)

Reported access points per day.


|Parameter|Default|
|--|--|
| `days` | 30 |



#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/aps_per_day.json?days=30


#### Response

```json
{
  "created_at": "2019-11-09T16:28:16.506359608+01:00",
  "updated_at": "2019-11-09T16:28:16.506491743+01:00",
  "name": "aps_per_day",
  "cache": {
    "type": 1,
    "keys": [
      "days"
    ],
    "ttl": 3600
  },
  "description": "Reported access points per day.",
  "defaults": {
    "days": 30
  },
  "query": "SELECT DATE_FORMAT(created_at, '%Y-%m-%d') AS day, COUNT(id) AS reported FROM access_points GROUP BY day ORDER BY day DESC LIMIT {days};",
  "views": {
    "chart": "aps_per_day_timeseries.go"
  },
  "access": [
    "admin"
  ]
}
```

### GET|POST /api/v1/query/aps_per_day/explain

Return results for an EXPLAIN operation on the aps_per_day main query.


|Parameter|Default|
|--|--|
| `days` | 30 |



#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/aps_per_day/explain?days=30


#### Response

```json
{
  "cached_at": null,
  "exec_time": 2763263,
  "num_records": 1,
  "records": [
    {
      "Extra": "Zero limit",
      "filtered": null,
      "id": 1,
      "key": null,
      "key_len": null,
      "partitions": null,
      "possible_keys": null,
      "ref": null,
      "rows": null,
      "select_type": "SIMPLE",
      "table": null,
      "type": null
    }
  ]
}
```
 






### GET /api/v1/query/aps_per_day/chart.(png|svg)

Return a PNG or SVG representation of a chart chart for the aps_per_day query.


|Parameter|Default|
|--|--|
| `days` | 30 |



#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/aps_per_day/chart.png?days=30






### GET /api/v1/query/top_players/view  

Show information about the top_players query.

#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/top_players/view


#### Response

```json
{
  "created_at": "2019-11-09T10:57:01.784331255+01:00",
  "updated_at": "2019-11-09T10:57:01.784494235+01:00",
  "name": "top_players",
  "cache": {
    "type": 1,
    "keys": [
      "limit"
    ],
    "ttl": 3600
  },
  "description": "Top players by access points.",
  "defaults": {
    "limit": 25
  },
  "query": "SELECT  u.updated_at as active_at, u.name, u.fingerprint, u.country, COUNT(a.id) AS networks FROM units u INNER JOIN access_points a ON u.id = a.unit_id GROUP BY u.id ORDER BY networks DESC LIMIT {limit}",
  "views": {
    "bars": "top_players_bars.go"
  },
  "access": [
    "admin"
  ]
}
```

### GET|POST /api/v1/query/top_players(.json|.csv)

Top players by access points.


|Parameter|Default|
|--|--|
| `limit` | 25 |



#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/top_players.json?limit=25


#### Response

```json
{
  "created_at": "2019-11-09T10:57:01.784331255+01:00",
  "updated_at": "2019-11-09T10:57:01.784494235+01:00",
  "name": "top_players",
  "cache": {
    "type": 1,
    "keys": [
      "limit"
    ],
    "ttl": 3600
  },
  "description": "Top players by access points.",
  "defaults": {
    "limit": 25
  },
  "query": "SELECT  u.updated_at as active_at, u.name, u.fingerprint, u.country, COUNT(a.id) AS networks FROM units u INNER JOIN access_points a ON u.id = a.unit_id GROUP BY u.id ORDER BY networks DESC LIMIT {limit}",
  "views": {
    "bars": "top_players_bars.go"
  },
  "access": [
    "admin"
  ]
}
```

### GET|POST /api/v1/query/top_players/explain

Return results for an EXPLAIN operation on the top_players main query.


|Parameter|Default|
|--|--|
| `limit` | 25 |



#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/top_players/explain?limit=25


#### Response

```json
{
  "cached_at": null,
  "exec_time": 2763263,
  "num_records": 1,
  "records": [
    {
      "Extra": "Zero limit",
      "filtered": null,
      "id": 1,
      "key": null,
      "key_len": null,
      "partitions": null,
      "possible_keys": null,
      "ref": null,
      "rows": null,
      "select_type": "SIMPLE",
      "table": null,
      "type": null
    }
  ]
}
```
 






### GET /api/v1/query/top_players/bars.(png|svg)

Return a PNG or SVG representation of a bars chart for the top_players query.


|Parameter|Default|
|--|--|
| `limit` | 25 |



#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/top_players/bars.png?limit=25






### GET /api/v1/query/last_active/view  

Show information about the last_active query.

#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/last_active/view


#### Response

```json
{
  "created_at": "2019-11-09T11:08:43.390943648+01:00",
  "updated_at": "2019-11-09T11:08:43.391060126+01:00",
  "name": "last_active",
  "cache": {
    "type": 1,
    "keys": [
      "hours",
      "limit"
    ],
    "ttl": 1800
  },
  "description": "Units seen in the last hour.",
  "defaults": {
    "hours": 1,
    "limit": 25
  },
  "query": "SELECT updated_at, name, fingerprint, country FROM units WHERE updated_at \u003e= NOW() - INTERVAL {hours} HOUR ORDER BY updated_at DESC LIMIT {limit}",
  "views": null,
  "access": [
    "admin"
  ]
}
```

### GET|POST /api/v1/query/last_active(.json|.csv)

Units seen in the last hour.


|Parameter|Default|
|--|--|
| `hours` | 1 |
| `limit` | 25 |



#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/last_active.json?hours=1&limit=25


#### Response

```json
{
  "created_at": "2019-11-09T11:08:43.390943648+01:00",
  "updated_at": "2019-11-09T11:08:43.391060126+01:00",
  "name": "last_active",
  "cache": {
    "type": 1,
    "keys": [
      "hours",
      "limit"
    ],
    "ttl": 1800
  },
  "description": "Units seen in the last hour.",
  "defaults": {
    "hours": 1,
    "limit": 25
  },
  "query": "SELECT updated_at, name, fingerprint, country FROM units WHERE updated_at \u003e= NOW() - INTERVAL {hours} HOUR ORDER BY updated_at DESC LIMIT {limit}",
  "views": null,
  "access": [
    "admin"
  ]
}
```

### GET|POST /api/v1/query/last_active/explain

Return results for an EXPLAIN operation on the last_active main query.


|Parameter|Default|
|--|--|
| `hours` | 1 |
| `limit` | 25 |



#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/last_active/explain?limit=25&hours=1


#### Response

```json
{
  "cached_at": null,
  "exec_time": 2763263,
  "num_records": 1,
  "records": [
    {
      "Extra": "Zero limit",
      "filtered": null,
      "id": 1,
      "key": null,
      "key_len": null,
      "partitions": null,
      "possible_keys": null,
      "ref": null,
      "rows": null,
      "select_type": "SIMPLE",
      "table": null,
      "type": null
    }
  ]
}
```
 









### GET /api/v1/query/num_messages/view  

Show information about the num_messages query.

#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/num_messages/view


#### Response

```json
{
  "created_at": "2019-11-07T19:11:18.871797845+01:00",
  "updated_at": "2019-11-07T19:11:18.872002647+01:00",
  "name": "num_messages",
  "cache": {
    "type": 2,
    "keys": null,
    "ttl": 60
  },
  "description": "Number of pwngrid messages.",
  "defaults": null,
  "query": "SELECT count(id) AS num_messages FROM messages",
  "views": null,
  "access": [
    "admin"
  ]
}
```

### GET|POST /api/v1/query/num_messages(.json|.csv)

Number of pwngrid messages.



#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/num_messages.json


#### Response

```json
{
  "created_at": "2019-11-07T19:11:18.871797845+01:00",
  "updated_at": "2019-11-07T19:11:18.872002647+01:00",
  "name": "num_messages",
  "cache": {
    "type": 2,
    "keys": null,
    "ttl": 60
  },
  "description": "Number of pwngrid messages.",
  "defaults": null,
  "query": "SELECT count(id) AS num_messages FROM messages",
  "views": null,
  "access": [
    "admin"
  ]
}
```

### GET|POST /api/v1/query/num_messages/explain

Return results for an EXPLAIN operation on the num_messages main query.



#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/num_messages/explain


#### Response

```json
{
  "cached_at": null,
  "exec_time": 2763263,
  "num_records": 1,
  "records": [
    {
      "Extra": "Zero limit",
      "filtered": null,
      "id": 1,
      "key": null,
      "key_len": null,
      "partitions": null,
      "possible_keys": null,
      "ref": null,
      "rows": null,
      "select_type": "SIMPLE",
      "table": null,
      "type": null
    }
  ]
}
```
 









### GET /api/v1/query/messages_per_day/view  

Show information about the messages_per_day query.

#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/messages_per_day/view


#### Response

```json
{
  "created_at": "2019-11-09T11:29:56.269663688+01:00",
  "updated_at": "2019-11-09T11:29:56.269804774+01:00",
  "name": "messages_per_day",
  "cache": {
    "type": 1,
    "keys": [
      "days"
    ],
    "ttl": 3600
  },
  "description": "PwnMAIL messages per day.",
  "defaults": {
    "days": 30
  },
  "query": "SELECT DATE_FORMAT(created_at, '%Y-%m-%d') AS day, COUNT(id) AS sent FROM messages GROUP BY day ORDER BY day DESC LIMIT {days};",
  "views": {
    "chart": "messages_per_day_timeseries.go"
  },
  "access": [
    "admin"
  ]
}
```

### GET|POST /api/v1/query/messages_per_day(.json|.csv)

PwnMAIL messages per day.


|Parameter|Default|
|--|--|
| `days` | 30 |



#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/messages_per_day.json?days=30


#### Response

```json
{
  "created_at": "2019-11-09T11:29:56.269663688+01:00",
  "updated_at": "2019-11-09T11:29:56.269804774+01:00",
  "name": "messages_per_day",
  "cache": {
    "type": 1,
    "keys": [
      "days"
    ],
    "ttl": 3600
  },
  "description": "PwnMAIL messages per day.",
  "defaults": {
    "days": 30
  },
  "query": "SELECT DATE_FORMAT(created_at, '%Y-%m-%d') AS day, COUNT(id) AS sent FROM messages GROUP BY day ORDER BY day DESC LIMIT {days};",
  "views": {
    "chart": "messages_per_day_timeseries.go"
  },
  "access": [
    "admin"
  ]
}
```

### GET|POST /api/v1/query/messages_per_day/explain

Return results for an EXPLAIN operation on the messages_per_day main query.


|Parameter|Default|
|--|--|
| `days` | 30 |



#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/messages_per_day/explain?days=30


#### Response

```json
{
  "cached_at": null,
  "exec_time": 2763263,
  "num_records": 1,
  "records": [
    {
      "Extra": "Zero limit",
      "filtered": null,
      "id": 1,
      "key": null,
      "key_len": null,
      "partitions": null,
      "possible_keys": null,
      "ref": null,
      "rows": null,
      "select_type": "SIMPLE",
      "table": null,
      "type": null
    }
  ]
}
```
 






### GET /api/v1/query/messages_per_day/chart.(png|svg)

Return a PNG or SVG representation of a chart chart for the messages_per_day query.


|Parameter|Default|
|--|--|
| `days` | 30 |



#### Request


    curl -H "Authorization: Bearer ..token.." http://127.0.0.1:8080/api/v1/query/messages_per_day/chart.png?days=30






