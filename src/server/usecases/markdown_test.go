package usecases

import (
	"testing"
)

func TestMarkdownInteractor_DeleteIfNotExistsByPath(t *testing.T) {
	type fields struct {
		markdownRepo MarkdownRepo
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &MarkdownInteractor{
				markdownRepo: tt.fields.markdownRepo,
			}
			if err := u.DeleteIfNotExistsByPath(); (err != nil) != tt.wantErr {
				t.Errorf("MarkdownInteractor.DeleteIfNotExistsByPath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
