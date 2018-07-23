package cloudformation

import (
	"encoding/json"

	"reflect"

	"github.com/mitchellh/mapstructure"
)

// UntypedAWSServerlessFunction_Properties is a helper struct that can hold either a S3Event, SNSEvent, KinesisEvent, DynamoDBEvent, ApiEvent, ScheduleEvent, CloudWatchEventEvent, IoTRuleEvent, or AlexaSkillEvent value
type UntypedAWSServerlessFunction_Properties struct {
	S3Event              *UntypedAWSServerlessFunction_S3Event
	SNSEvent             *UntypedAWSServerlessFunction_SNSEvent
	KinesisEvent         *UntypedAWSServerlessFunction_KinesisEvent
	DynamoDBEvent        *UntypedAWSServerlessFunction_DynamoDBEvent
	ApiEvent             *UntypedAWSServerlessFunction_ApiEvent
	ScheduleEvent        *UntypedAWSServerlessFunction_ScheduleEvent
	CloudWatchEventEvent *UntypedAWSServerlessFunction_CloudWatchEventEvent
	IoTRuleEvent         *UntypedAWSServerlessFunction_IoTRuleEvent
	AlexaSkillEvent      *UntypedAWSServerlessFunction_AlexaSkillEvent
}

func (r UntypedAWSServerlessFunction_Properties) value() interface{} {

	if r.S3Event != nil && !reflect.DeepEqual(r.S3Event, &UntypedAWSServerlessFunction_S3Event{}) {
		return r.S3Event
	}

	if r.SNSEvent != nil && !reflect.DeepEqual(r.SNSEvent, &UntypedAWSServerlessFunction_SNSEvent{}) {
		return r.SNSEvent
	}

	if r.KinesisEvent != nil && !reflect.DeepEqual(r.KinesisEvent, &UntypedAWSServerlessFunction_KinesisEvent{}) {
		return r.KinesisEvent
	}

	if r.DynamoDBEvent != nil && !reflect.DeepEqual(r.DynamoDBEvent, &UntypedAWSServerlessFunction_DynamoDBEvent{}) {
		return r.DynamoDBEvent
	}

	if r.ApiEvent != nil && !reflect.DeepEqual(r.ApiEvent, &UntypedAWSServerlessFunction_ApiEvent{}) {
		return r.ApiEvent
	}

	if r.ScheduleEvent != nil && !reflect.DeepEqual(r.ScheduleEvent, &UntypedAWSServerlessFunction_ScheduleEvent{}) {
		return r.ScheduleEvent
	}

	if r.CloudWatchEventEvent != nil && !reflect.DeepEqual(r.CloudWatchEventEvent, &UntypedAWSServerlessFunction_CloudWatchEventEvent{}) {
		return r.CloudWatchEventEvent
	}

	if r.IoTRuleEvent != nil && !reflect.DeepEqual(r.IoTRuleEvent, &UntypedAWSServerlessFunction_IoTRuleEvent{}) {
		return r.IoTRuleEvent
	}

	if r.AlexaSkillEvent != nil && !reflect.DeepEqual(r.AlexaSkillEvent, &UntypedAWSServerlessFunction_AlexaSkillEvent{}) {
		return r.AlexaSkillEvent
	}

	if r.S3Event != nil {
		return r.S3Event
	}

	if r.SNSEvent != nil {
		return r.SNSEvent
	}

	if r.KinesisEvent != nil {
		return r.KinesisEvent
	}

	if r.DynamoDBEvent != nil {
		return r.DynamoDBEvent
	}

	if r.ApiEvent != nil {
		return r.ApiEvent
	}

	if r.ScheduleEvent != nil {
		return r.ScheduleEvent
	}

	if r.CloudWatchEventEvent != nil {
		return r.CloudWatchEventEvent
	}

	if r.IoTRuleEvent != nil {
		return r.IoTRuleEvent
	}

	if r.AlexaSkillEvent != nil {
		return r.AlexaSkillEvent
	}

	return nil

}

func (r *UntypedAWSServerlessFunction_Properties) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value())
}

// Hook into the marshaller
func (r *UntypedAWSServerlessFunction_Properties) UnmarshalJSON(b []byte) error {

	// Unmarshal into interface{} to check it's type
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case map[string]interface{}:

		mapstructure.Decode(val, &r.S3Event)

		mapstructure.Decode(val, &r.SNSEvent)

		mapstructure.Decode(val, &r.KinesisEvent)

		mapstructure.Decode(val, &r.DynamoDBEvent)

		mapstructure.Decode(val, &r.ApiEvent)

		mapstructure.Decode(val, &r.ScheduleEvent)

		mapstructure.Decode(val, &r.CloudWatchEventEvent)

		mapstructure.Decode(val, &r.IoTRuleEvent)

		mapstructure.Decode(val, &r.AlexaSkillEvent)

	case []interface{}:

	}

	return nil
}
