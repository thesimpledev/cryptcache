package flags

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func TestParseFlagsSelective(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected map[string]interface{}
	}{
		{
			name: "init only",
			args: []string{"--init"},
			expected: map[string]interface{}{
				"InitProject": true,
			},
		},
		{
			name: "name only",
			args: []string{"-n", "testProject"},
			expected: map[string]interface{}{
				"Name": "testProject",
			},
		},
		{
			name: "New Project only",
			args: []string{"--new", "testProject"},
			expected: map[string]interface{}{
				"NewProfile": true,
			},
		},
		{
			name: "Set/Update Key only",
			args: []string{"--set", "testProject"},
			expected: map[string]interface{}{
				"Set": true,
			},
		},
		{
			name: "Get Ket Value only",
			args: []string{"--get", "testProject"},
			expected: map[string]interface{}{
				"Get": true,
			},
		},
		{
			name: "Export Project only",
			args: []string{"--export", "testProject"},
			expected: map[string]interface{}{
				"Export": true,
			},
		},
		{
			name: "Key Set only",
			args: []string{"-k", "key Test"},
			expected: map[string]interface{}{
				"Key": "key Test",
			},
		},
		{
			name: "Value Set only",
			args: []string{"-v", "value Test"},
			expected: map[string]interface{}{
				"Value": "value Test",
			},
		},
		{
			name: "File Name Set only",
			args: []string{"-f", "fileNameText.txt"},
			expected: map[string]interface{}{
				"File": "fileNameText.txt",
			},
		},
		{
			name: "Private Key Set only",
			args: []string{"-private", "privateKey"},
			expected: map[string]interface{}{
				"Private": "privateKey",
			},
		},
		{
			name: "Pubic Key Set only",
			args: []string{"-public", "publicKey"},
			expected: map[string]interface{}{
				"Public": "publicKey",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			os.Args = append([]string{os.Args[0]}, tt.args...)
			got := ParseFlags()
			gotVal := reflect.ValueOf(got)

			for field, want := range tt.expected {
				fieldVal := gotVal.FieldByName(field)
				if !fieldVal.IsValid() {
					t.Errorf("field %q not found in Flags", field)
					continue
				}
				if !reflect.DeepEqual(fieldVal.Interface(), want) {
					t.Errorf("field %q = %v, want %v", field, fieldVal.Interface(), want)
				}
			}
		})
	}
}

func TestParseFlagsDefaults(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{os.Args[0]}
	got := ParseFlags()
	want := Flags{
		InitProject: false,
		NewProfile:  false,
		Set:         false,
		Get:         false,
		Export:      false,
		Name:        "Definitely not a Mimic",
		File:        "",
		Key:         "",
		Value:       "",
		Private:     "~/.ssh/id_rsa",
		Public:      "~/.ssh/id_rsa.pub",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ParseFlags() = %+v, want %+v", got, want)
	}
}
