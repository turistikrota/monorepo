package list

import (
	"testing"

	"github.com/turistikrota/api/pkg/ptr"
)

func TestNewPagiResponse(t *testing.T) {
	req := &PagiRequest{
		Limit: ptr.Int(10),
		Page:  ptr.Int(1),
	}

	list := []string{"item1", "item2", "item3"}
	total := int64(3)
	filteredTotal := int64(3)

	pagiResponse := NewPagiResponse(req, list, total, filteredTotal)

	if pagiResponse == nil {
		t.Error("Expected non-nil PagiResponse, got nil")
	}

	if pagiResponse.List == nil {
		t.Error("Expected non-nil List, got nil")
	}

	if len(pagiResponse.List) != len(list) {
		t.Errorf("Expected List length %d, got %d", len(list), len(pagiResponse.List))
	}

	if pagiResponse.Total != total {
		t.Errorf("Expected Total %d, got %d", total, pagiResponse.Total)
	}

	if pagiResponse.Limit != *req.Limit {
		t.Errorf("Expected Limit %d, got %d", *req.Limit, pagiResponse.Limit)
	}

	if pagiResponse.TotalPage != req.TotalPage(filteredTotal) {
		t.Errorf("Expected TotalPage %d, got %d", req.TotalPage(filteredTotal), pagiResponse.TotalPage)
	}

	if pagiResponse.FilteredTotal != filteredTotal {
		t.Errorf("Expected FilteredTotal %d, got %d", filteredTotal, pagiResponse.FilteredTotal)
	}

	if pagiResponse.Page != *req.Page {
		t.Errorf("Expected Page %d, got %d", *req.Page, pagiResponse.Page)
	}
}
