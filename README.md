How to Run:  
go mod tidy  
go run main.go  


### GrapphQL Query:  
```
query {
  GetData {
    id
    title
    segments {
      id
      contents {
        id
        title
        ... on A {
          teams {
            name
          }
        }
        ... on B {
          score
        }
      }
    }
  }
}
```

### Error:
```
  "errors": [
    {
      "message": "json: cannot unmarshal object into Go struct field Segment.segments.contents of type models.Content",
      "path": [
        "GetData"
      ]
    }
  ],
  "data": {
    "GetData": null
  }
}
```
