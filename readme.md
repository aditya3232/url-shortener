# API Spec

## 1 UrlShortener

### 1.1 Create

Request :
- Method : POST
- URL : `{{local}}:3636/short`
- Body (form-data) :
    - long_url : string, required
- Response :

```json 
{
    "meta": {
        "message": "Successfully created new data.",
        "code": 201
    },
    "data": {
        "long_url": "https://www.youtube.com/watch?v=XepkKrzPD5g",
        "short_url": "http://localhost:3636/short/Ccl8ZO"
    }
}
```
### 1.2 Get All

Request :
- Method : GET
- URL : `{{local}}:3636/short`
- Params (filter)
    - limit:10
    - page:1
    - sort:id
    - order:desc
    - id:
    - long_url:
    - short_url:
    - click:
    - created_at:
    - updated_at:
- Response :

```json 
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "pagination": {
        "page": 1,
        "limit": 10,
        "total": 3,
        "total_filtered": 3
    },
    "data": [
        {
            "id": 3,
            "long_url": "https://www.youtube.com/watch?v=MTF4hFkTASA",
            "short_url": "http://localhost:3636/short/4YOvo7",
            "click": 1,
            "created_at": "2024-01-09T09:54:13+07:00",
            "updated_at": "2024-01-09T09:54:56+07:00"
        },
        {
            "id": 2,
            "long_url": "https://www.youtube.com/watch?v=KbNL9ZyB49c&list=RDnxXusN7Yyb8&index=8",
            "short_url": "http://localhost:3636/short/DNWYQ2",
            "click": 0,
            "created_at": "2024-01-08T15:40:52+07:00",
            "updated_at": "2024-01-08T15:40:52+07:00"
        },
        {
            "id": 1,
            "long_url": "https://www.youtube.com/watch?v=KbNL9ZyB49c&list=RDnxXusN7Yyb8&index=8",
            "short_url": "http://localhost:3636/short/zGNDHn",
            "click": 2,
            "created_at": "2024-01-08T14:54:25+07:00",
            "updated_at": "2024-01-08T15:40:13+07:00"
        }
    ]
}
```

