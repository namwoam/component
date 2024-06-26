package text

import (
	"context"
	"testing"

	"google.golang.org/protobuf/types/known/structpb"
)

// TestSplitByToken tests the split by token task
func TestSplitByToken(t *testing.T) {
	input := &structpb.Struct{
		Fields: map[string]*structpb.Value{
			"text":  {Kind: &structpb.Value_StringValue{StringValue: "Hello world. This is a test."}},
			"model": {Kind: &structpb.Value_StringValue{StringValue: "gpt-3.5-turbo"}},
		},
	}
	inputs := []*structpb.Struct{
		input,
	}

	e := &execution{}
	e.Task = "TASK_SPLIT_BY_TOKEN"

	if _, err := e.Execute(context.Background(), inputs); err != nil {
		t.Fatalf("splitByToken returned an error: %v", err)
	}
}
