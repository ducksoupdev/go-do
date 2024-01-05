package functions

import "testing"

func TestErrorHandler(t *testing.T) {
	type args struct {
		title      string
		statusCode int
		message    string
	}
	tests := []struct {
		name string
		args args
		want *Response
	}{
		{
			name: "Test error handler",
			args: args{
				title:      "title",
				statusCode: 400,
				message:    "message",
			},
			want: &Response{
				StatusCode: 400,
				Headers:    Headers,
				Body:       "{\"title\":\"title\",\"status\":400,\"detail\":\"message\"}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorHandler(tt.args.title, tt.args.statusCode, tt.args.message); got.Body != tt.want.Body {
				t.Errorf("ErrorHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusCodeTitle(t *testing.T) {
	type args struct {
		statusCode int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test status code title",
			args: args{
				statusCode: 200,
			},
			want: "OK",
		},
		{
			name: "Test status code title",
			args: args{
				statusCode: 201,
			},
			want: "Created",
		},
		{
			name: "Test status code title",
			args: args{
				statusCode: 204,
			},
			want: "No Content",
		},
		{
			name: "Test status code title",
			args: args{
				statusCode: 400,
			},
			want: "Bad Request",
		},
		{
			name: "Test status code title",
			args: args{
				statusCode: 401,
			},
			want: "Unauthorized",
		},
		{
			name: "Test status code title",
			args: args{
				statusCode: 403,
			},
			want: "Forbidden",
		},
		{
			name: "Test status code title",
			args: args{
				statusCode: 404,
			},
			want: "Not Found",
		},
		{
			name: "Test status code title",
			args: args{
				statusCode: 405,
			},
			want: "Method Not Allowed",
		},
		{
			name: "Test status code title",
			args: args{
				statusCode: 406,
			},
			want: "Not Acceptable",
		},
		{
			name: "Test status code title",
			args: args{
				statusCode: 409,
			},
			want: "Conflict",
		},
		{
			name: "Test status code title",
			args: args{
				statusCode: 415,
			},
			want: "Unsupported Media Type",
		},
		{
			name: "Test status code title",
			args: args{
				statusCode: 500,
			},
			want: "Internal Server Error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StatusCodeTitle(tt.args.statusCode); got != tt.want {
				t.Errorf("StatusCodeTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		s   []string
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test contains",
			args: args{
				s:   []string{"a", "b", "c"},
				str: "a",
			},
			want: true,
		},
		{
			name: "Test contains",
			args: args{
				s:   []string{"a", "b", "c"},
				str: "d",
			},
			want: false,
		},
		{
			name: "Test contains",
			args: args{
				s:   []string{},
				str: "a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Contains(tt.args.s, tt.args.str) != tt.want {
				t.Errorf("Contains() = %v, want %v", Contains(tt.args.s, tt.args.str), tt.want)
			}
		})
	}
}

func TestToStringArray(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test to string array",
			args: args{
				i: []interface{}{"a", "b", "c"},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "Test to string array",
			args: args{
				i: []string{"a", "b", "c"},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "Test to string array",
			args: args{
				i: []interface{}{},
			},
			want: []string{},
		},
		{
			name: "Test to string array",
			args: args{
				i: []string{},
			},
			want: []string{},
		},
		{
			name: "Test to string array",
			args: args{
				i: nil,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToStringArray(tt.args.i); !Equal(got, tt.want) {
				t.Errorf("ToStringArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Equal(got []string, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
