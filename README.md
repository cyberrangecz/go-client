# crczp-go-client
A [CRCZP]() client library written in Go.

## Supported API calls:
- Login to Keycloak
- Sandbox Definition - Get, Create, Delete
- Sandbox Pool - Get, Create, Delete, Cleanup
- Sandbox Allocation Unit - Get, CreateAllocation, CreateAllocationAwait, CancelAllocation, CreateCleanup, CreateCleanupAwait, GetAllocationOutput
- Training Definition - Get, Create, Delete
- Training Definition Adaptive - Get, Create, Delete

## Usage
```go
import "github.com/cyberrangecz/go-client"
```

Create a client with username and password:
```go
client, err := czcrp.NewClient("https://your.crczp.ex", "CRCZP-Client", "username", "password")
if err != nil {
    log.Fatalf("Failed to create CRCZP client: %v", err)
}
```

Use the client to create a sandbox definition:
```go
sandboxDefinition, err := client.CreateSandboxDefinition(context.Background(), 
	"https://github.com/cyberrangecz/library-demo-training.git", "master")

if err != nil {
    log.Fatalf("Failed to create sandbox definition: %v", err)
}
```
