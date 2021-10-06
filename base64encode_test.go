// package base64url provides base64url encoding/decoding support
package base64url

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "simlpe string",
			args: args{
				"U2VuZCByZWluZm9yY2VtZW50cw",
			},
			want:    []byte("Send reinforcements"),
			wantErr: false,
		},
		{
			name: "longer string",
			args: args{"Tm93IGlzIHRoZSB0aW1lIGZvciBhbGwgZ29vZCBjb2RlcnN0byBsZWFybiBSdWJ5"},
			want: []byte("Now is the time for all good coders" +
				"to learn Ruby"),
			wantErr: false,
		},
		{
			name:    "zero",
			args:    args{"AA"},
			want:    []byte{00},
			wantErr: false,
		}, {
			name:    "zero zero",
			args:    args{"AAA"},
			want:    []byte{00, 00},
			wantErr: false,
		},
		{
			name:    "zero zero zero",
			args:    args{"AAAA"},
			want:    []byte{00, 00, 00},
			wantErr: false,
		},
		{
			name:    "LF",
			args:    args{"Cg"},
			want:    []byte("\n"),
			wantErr: false,
		},
		{
			name:    "CR LF",
			args:    args{"DQo"},
			want:    []byte{13, 10},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple string",
			args: args{[]byte("Send reinforcements")},
			want: "U2VuZCByZWluZm9yY2VtZW50cw",
		},
		{
			name: "longer string",
			args: args{[]byte("Now is the time for all good coders" +
				"to learn Ruby")},
			want: "Tm93IGlzIHRoZSB0aW1lIGZvciBhbGwgZ29vZCBjb2RlcnN0byBsZWFybiBSdWJ5",
		},
		{
			name: "multiline string",
			args: args{[]byte("This is line one" +
				"This is line two" +
				"This is line three" +
				"And so on..." +
				"")},
			want: "VGhpcyBpcyBsaW5lIG9uZVRoaXMgaXMgbGluZSB0d29UaGlzIGlzIGxpbmUgdGhyZWVBbmQgc28gb24uLi4",
		}, {
			name: "zero",
			args: args{
				data: []byte{00},
			},
			want: "AA",
		},
		{
			name: "zero zero",
			args: args{
				data: []byte{00, 00},
			},
			want: "AAA",
		},
		{
			name: "zero zero zero",
			args: args{
				data: []byte{00, 00, 00},
			},
			want: "AAAA",
		},
		{
			name: "LF",
			args: args{
				data: []byte{10},
			},
			want: "Cg",
		},
		{
			name: "CR",
			args: args{
				data: []byte{13},
			},
			want: "DQ",
		},
		{
			name: "CR LF",
			args: args{
				data: []byte{13, 10},
			},
			want: "DQo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args.data); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
