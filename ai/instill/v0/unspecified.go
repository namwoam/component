package instill

import (
	"fmt"

	modelPB "github.com/instill-ai/protogen-go/model/model/v1alpha"
	"google.golang.org/protobuf/types/known/structpb"
)

func (e *execution) executeUnspecified(grpcClient modelPB.ModelPublicServiceClient, nsID string, modelID string, version string, inputs []*structpb.Struct) ([]*structpb.Struct, error) {
	if len(inputs) <= 0 {
		return nil, fmt.Errorf("invalid input: %v for model: %s/%s", inputs, nsID, modelID)
	}
	//TODO: figure out what to do here?
	/*
		modelInput := &modelPB.TaskInput_Unspecified{
			Unspecified: &modelPB.UnspecifiedInput{
				RawInputs: []*structpb.Struct{},
			},
		}
	*/
	return inputs, nil
}
