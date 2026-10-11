package echox

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/Tomdooo/spajz/internal/models"
)

func TestLookupErrorInfo(t *testing.T) {
	tests := []struct {
		name      string
		err       error
		wantInfo  errorInfo
		wantKnown bool
	}{
		{
			name: "bucket not found",
			err:  models.ErrBucketNotFound,
			wantInfo: errorInfo{
				Status:  http.StatusNotFound,
				Code:    models.CodeBucketNotFound,
				Message: "No such bucket.",
			},
			wantKnown: true,
		},
		{
			name: "wrapped bucket not found",
			err:  fmt.Errorf("loading bucket: %w", models.ErrBucketNotFound),
			wantInfo: errorInfo{
				Status:  http.StatusNotFound,
				Code:    models.CodeBucketNotFound,
				Message: "No such bucket.",
			},
			wantKnown: true,
		},
		{
			name: "file not saved",
			err:  models.ErrFileNotSaved,
			wantInfo: errorInfo{
				Status:    http.StatusInternalServerError,
				Code:      models.CodeFileNotSaved,
				Message:   "File could not be saved.",
				ShouldLog: true,
			},
			wantKnown: true,
		},
		{
			name:      "unknown error",
			err:       errors.New("unknown error"),
			wantInfo:  errorInfo{},
			wantKnown: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInfo, gotKnown := lookupErrorInfo(tt.err)

			if gotKnown != tt.wantKnown {
				t.Errorf("known = %t, want %t", gotKnown, tt.wantKnown)
			}
			if gotInfo != tt.wantInfo {
				t.Errorf("info = %+v, want %+v", gotInfo, tt.wantInfo)
			}
		})
	}
}
