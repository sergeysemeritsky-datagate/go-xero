package integration

import (
	"context"
	"testing"
)

func Test_GetConnections(t *testing.T) {
	connections, _, err := client.Connections.GetConnections(context.TODO())
	if err != nil {
		t.Fatalf("Connections.GetConnections returned error: %v", err)
	}
	if connections == nil {
		t.Fatalf("Connections.GetConnections returned nil")
	}

	if len(connections) == 0 {
		t.Fatalf("Connections.GetConnections returned no connections")
	}

	found := false
	for _, c := range connections {
		if c.TenantId == client.TenantId {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Connections.GetConnections returned no connections of the current tenant")
	}
}
