package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-participation-creates-rmq-kube/DPFM_API_Input_Reader"
	//dpfm_api_processing_formatter "data-platform-api-participation-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-participation-creates-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) (*Header, error) {
	data := subfuncSDC.Message.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToHeader(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	subfuncSDC.Message.Header = &sub_func_complementer.Header{
		Participation:				*input.Header.Participation,
		ParticipationDate:			input.Header.ParticipationDate,
		ParticipationTime:			input.Header.ParticipationTime,
		Participator:				input.Header.Participator,
		ParticipationObjectType:	input.Header.ParticipationObjectType,
		ParticipationObject:		input.Header.ParticipationObject,
		Attendance:					input.Header.Attendance,
		CreationDate:				input.Header.CreationDate,
		CreationTime:				input.Header.CreationTime,
		IsCancelled:				input.Header.IsCancelled,
	}

	return subfuncSDC
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
