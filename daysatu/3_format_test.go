package daysatu

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFormatJSON(t *testing.T) {
	type args struct {
		data interface{}
	}

	myData := Person{
		Name: "John",
	}

	ex, _ := json.Marshal(myData)

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			args: args{
				data: Person{
					Name: "John",
				},
			},
			want:    ex,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatJSON(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatYaml(t *testing.T) {
	type args struct {
		data interface{}
	}

	data := User{
		Name: "Jono",
		ID:   "1",
	}

	wantData := `<?xml version="1.0" encoding="UTF-8"?>
	<User>
	 <id>1</id>
	 <name>Jono</name>
	</User>,`

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Test Jono",
			args: args{
				data: data,
			},
			want: []byte(wantData),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatYaml(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatYaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormatYaml() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
