How to Run:  
go mod tidy  
go run main.go  


GrapphQL Query:  
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
