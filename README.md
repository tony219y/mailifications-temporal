# mailifications-temporal
`./backend`
### 1. Start Temporal server
```
temporal server start-dev --ui-port 8080
```
### 2. Run Worker 
`Receive Email and Delay from (Starter, Frontend)`
```
go run ./cmd/worker/main.go
```
### 3. Run Starter 
`Send an example data to worker`

```
go run ./cmd/starter/main.go
```
---
### 🟢 (Optional) Play with React 


```
cd frontend
yarn dev
```
### Start API
```
cd backend
go run ./cmd/api/main.go
```

## Download Temporal CLI

https://learn.temporal.io/getting_started/go/dev_environment


