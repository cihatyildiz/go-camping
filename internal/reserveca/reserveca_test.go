package reserveca

import (
    "testing"
)

func TestClient_CheckCampsite(t *testing.T) {
    client := NewClient()

    req := CampsiteRequest{
        FacilityId: 490,
        StartDate:  "2024-07-05",
        EndDate:    "2024-07-12",
    }

    resp, err := client.CheckCampsite(req)
    if err != nil {
        t.Fatalf("Client.CheckCampsite failed: %s", err)
    }

    if resp.Facility.FacilityId != req.FacilityId {
        t.Errorf("Expected FacilityId %d, got %d", req.FacilityId, resp.Facility.FacilityId)
    }

    // Add more assertions as needed
}
