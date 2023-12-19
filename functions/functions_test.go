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
