package shikimori_test

import (
	"context"
	"testing"

	"github.com/SevereCloud/shikimori"
)

func TestPeople(t *testing.T) {
	t.Parallel()

	resp, err := shiki.People(context.Background(), 3, nil)

	NoError(t, err)
	NotEmpty(t, resp)
}

func TestGrouppedRole_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	type fields struct {
		Name  string
		Count int
	}

	tests := []struct {
		fields  fields
		data    []byte
		wantErr bool
	}{
		{
			fields: fields{
				Name:  "name",
				Count: 1,
			},
			data:    []byte(`["name",1]`),
			wantErr: false,
		},
		{
			data:    []byte(`{}`),
			wantErr: true,
		},
		{
			data:    []byte(`[1,"name"]`),
			wantErr: true,
		},
		{
			data:    []byte(`["name","name"]`),
			wantErr: true,
		},
		{
			data:    []byte(`["name"]`),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt //nolint:varnamelen

		t.Run(string(tt.data), func(t *testing.T) {
			t.Parallel()

			r := &shikimori.RoleStat{
				Name:  tt.fields.Name,
				Count: tt.fields.Count,
			}
			if err := r.UnmarshalJSON(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("GrouppedRole.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
