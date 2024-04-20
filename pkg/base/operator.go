package base

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/gofrs/uuid"
	pipelinePB "github.com/instill-ai/protogen-go/vdp/pipeline/v1beta"
)

type IOperator interface {
	IComponent

	LoadOperatorDefinition(definitionJSON []byte, tasksJSON []byte, additionalJSONBytes map[string][]byte) error
	GetOperatorDefinition(sysVars map[string]any, component *pipelinePB.OperatorComponent) (*pipelinePB.OperatorDefinition, error)
	CreateExecution(sysVars map[string]any, task string) (*ExecutionWrapper, error)
}

type BaseOperator struct {
	Logger       *zap.Logger
	UsageHandler UsageHandler

	taskInputSchemas  map[string]string
	taskOutputSchemas map[string]string

	definition *pipelinePB.OperatorDefinition
}

type IOperatorExecution interface {
	IExecution

	GetOperator() IOperator
}

type BaseOperatorExecution struct {
	Operator        IOperator
	SystemVariables map[string]any
	Task            string
}

func (o *BaseOperator) GetID() string {
	return o.definition.Id
}
func (o *BaseOperator) GetUID() uuid.UUID {
	return uuid.FromStringOrNil(o.definition.Uid)
}
func (o *BaseOperator) GetLogger() *zap.Logger {
	return o.Logger
}
func (o *BaseOperator) GetUsageHandler() UsageHandler {
	return o.UsageHandler
}
func (o *BaseOperator) GetTaskInputSchemas() map[string]string {
	return o.taskInputSchemas
}
func (o *BaseOperator) GetTaskOutputSchemas() map[string]string {
	return o.taskOutputSchemas
}

func (o *BaseOperator) GetOperatorDefinition(sysVars map[string]any, component *pipelinePB.OperatorComponent) (*pipelinePB.OperatorDefinition, error) {
	return o.definition, nil
}

// LoadOperatorDefinition loads the operator definitions from json files
func (o *BaseOperator) LoadOperatorDefinition(definitionJSONBytes []byte, tasksJSONBytes []byte, additionalJSONBytes map[string][]byte) error {
	var err error
	var definitionJSON any

	err = json.Unmarshal(definitionJSONBytes, &definitionJSON)
	if err != nil {
		return err
	}
	renderedTasksJSON, nil := renderTaskJSON(tasksJSONBytes, additionalJSONBytes)
	if err != nil {
		return nil
	}

	availableTasks := []string{}
	for _, availableTask := range definitionJSON.(map[string]interface{})["available_tasks"].([]interface{}) {
		availableTasks = append(availableTasks, availableTask.(string))
	}

	tasks, taskStructs, err := loadTasks(availableTasks, renderedTasksJSON)
	if err != nil {
		return err
	}

	o.taskInputSchemas = map[string]string{}
	o.taskOutputSchemas = map[string]string{}
	for k := range taskStructs {
		var s []byte
		s, err = protojson.Marshal(taskStructs[k].Fields["input"].GetStructValue())
		if err != nil {
			return err
		}
		o.taskInputSchemas[k] = string(s)

		s, err = protojson.Marshal(taskStructs[k].Fields["output"].GetStructValue())
		if err != nil {
			return err
		}
		o.taskOutputSchemas[k] = string(s)
	}

	o.definition = &pipelinePB.OperatorDefinition{}
	err = protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(definitionJSONBytes, o.definition)
	if err != nil {
		return err
	}

	o.definition.Name = fmt.Sprintf("operator-definitions/%s", o.definition.Id)
	o.definition.Tasks = tasks
	o.definition.Spec.ComponentSpecification, err = generateComponentSpec(o.definition.Title, tasks, taskStructs)
	if err != nil {
		return err
	}
	o.definition.Spec.DataSpecifications, err = generateDataSpecs(taskStructs)
	if err != nil {
		return err
	}

	return nil
}

func (e *BaseOperatorExecution) GetTask() string {
	return e.Task
}
func (e *BaseOperatorExecution) GetOperator() IOperator {
	return e.Operator
}
func (e *BaseOperatorExecution) GetSystemVariables() map[string]any {
	return e.SystemVariables
}
func (e *BaseOperatorExecution) GetLogger() *zap.Logger {
	return e.Operator.GetLogger()
}
func (e *BaseOperatorExecution) GetUsageHandler() UsageHandler {
	return e.Operator.GetUsageHandler()
}
func (e *BaseOperatorExecution) GetTaskInputSchema() string {
	return e.Operator.GetTaskInputSchemas()[e.Task]
}
func (e *BaseOperatorExecution) GetTaskOutputSchema() string {
	return e.Operator.GetTaskOutputSchemas()[e.Task]
}
