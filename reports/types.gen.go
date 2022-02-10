// Package reports provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package reports

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// CreateReportResponse defines model for CreateReportResponse.
type CreateReportResponse struct {

	// The identifier for the report. This identifier is unique only in combination with a seller ID.
	ReportId string `json:"reportId"`
}

// CreateReportScheduleResponse defines model for CreateReportScheduleResponse.
type CreateReportScheduleResponse struct {

	// The identifier for the report schedule. This identifier is unique only in combination with a seller ID.
	ReportScheduleId string `json:"reportScheduleId"`
}

// CreateReportScheduleSpecification defines model for CreateReportScheduleSpecification.
type CreateReportScheduleSpecification struct {

	// A list of marketplace identifiers for the report schedule.
	MarketplaceIds []string `json:"marketplaceIds"`

	// The date and time when the schedule will create its next report, in ISO 8601 date time format.
	NextReportCreationTime *time.Time `json:"nextReportCreationTime,omitempty"`

	// One of a set of predefined ISO 8601 periods that specifies how often a report should be created.
	Period string `json:"period"`

	// Additional information passed to reports. This varies by report type.
	ReportOptions *ReportOptions `json:"reportOptions,omitempty"`

	// The report type.
	ReportType string `json:"reportType"`
}

// CreateReportSpecification defines model for CreateReportSpecification.
type CreateReportSpecification struct {

	// The end of a date and time range, in ISO 8601 date time format, used for selecting the data to report. The default is now. The value must be prior to or equal to the current date and time. Not all report types make use of this.
	DataEndTime *time.Time `json:"dataEndTime,omitempty"`

	// The start of a date and time range, in ISO 8601 date time format, used for selecting the data to report. The default is now. The value must be prior to or equal to the current date and time. Not all report types make use of this.
	DataStartTime *time.Time `json:"dataStartTime,omitempty"`

	// A list of marketplace identifiers. The report document's contents will contain data for all of the specified marketplaces, unless the report type indicates otherwise.
	MarketplaceIds []string `json:"marketplaceIds"`

	// Additional information passed to reports. This varies by report type.
	ReportOptions *ReportOptions `json:"reportOptions,omitempty"`

	// The report type.
	ReportType string `json:"reportType"`
}

// Error defines model for Error.
type Error struct {

	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`

	// Additional details that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`

	// A message that describes the error condition in a human-readable form.
	Message string `json:"message"`
}

// ErrorList defines model for ErrorList.
type ErrorList struct {
	Errors []Error `json:"errors"`
}

// GetReportsResponse defines model for GetReportsResponse.
type GetReportsResponse struct {

	// Returned when the number of results exceeds pageSize. To get the next page of results, call getReports with this token as the only parameter.
	NextToken *string `json:"nextToken,omitempty"`

	// A list of reports.
	Reports ReportList `json:"reports"`
}

// Report defines model for Report.
type Report struct {

	// The date and time when the report was created.
	CreatedTime time.Time `json:"createdTime"`

	// The end of a date and time range used for selecting the data to report.
	DataEndTime *time.Time `json:"dataEndTime,omitempty"`

	// The start of a date and time range used for selecting the data to report.
	DataStartTime *time.Time `json:"dataStartTime,omitempty"`

	// A list of marketplace identifiers for the report.
	MarketplaceIds *[]string `json:"marketplaceIds,omitempty"`

	// The date and time when the report processing completed, in ISO 8601 date time format.
	ProcessingEndTime *time.Time `json:"processingEndTime,omitempty"`

	// The date and time when the report processing started, in ISO 8601 date time format.
	ProcessingStartTime *time.Time `json:"processingStartTime,omitempty"`

	// The processing status of the report.
	ProcessingStatus string `json:"processingStatus"`

	// The identifier for the report document. Pass this into the getReportDocument operation to get the information you will need to retrieve the report document's contents.
	ReportDocumentId *string `json:"reportDocumentId,omitempty"`

	// The identifier for the report. This identifier is unique only in combination with a seller ID.
	ReportId string `json:"reportId"`

	// The identifier of the report schedule that created this report (if any). This identifier is unique only in combination with a seller ID.
	ReportScheduleId *string `json:"reportScheduleId,omitempty"`

	// The report type.
	ReportType string `json:"reportType"`
}

// ReportDocument defines model for ReportDocument.
type ReportDocument struct {

	// If present, the report document contents have been compressed with the provided algorithm.
	CompressionAlgorithm *string `json:"compressionAlgorithm,omitempty"`

	// The identifier for the report document. This identifier is unique only in combination with a seller ID.
	ReportDocumentId string `json:"reportDocumentId"`

	// A presigned URL for the report document. This URL expires after 5 minutes.
	Url string `json:"url"`
}

// ReportList defines model for ReportList.
type ReportList []Report

// ReportOptions defines model for ReportOptions.
type ReportOptions struct {
	AdditionalProperties map[string]string `json:"-"`
}

// ReportSchedule defines model for ReportSchedule.
type ReportSchedule struct {

	// A list of marketplace identifiers. The report document's contents will contain data for all of the specified marketplaces, unless the report type indicates otherwise.
	MarketplaceIds *[]string `json:"marketplaceIds,omitempty"`

	// The date and time when the schedule will create its next report, in ISO 8601 date time format.
	NextReportCreationTime *time.Time `json:"nextReportCreationTime,omitempty"`

	// An ISO 8601 period value that indicates how often a report should be created.
	Period string `json:"period"`

	// Additional information passed to reports. This varies by report type.
	ReportOptions *ReportOptions `json:"reportOptions,omitempty"`

	// The identifier for the report schedule. This identifier is unique only in combination with a seller ID.
	ReportScheduleId string `json:"reportScheduleId"`

	// The report type.
	ReportType string `json:"reportType"`
}

// ReportScheduleList defines model for ReportScheduleList.
type ReportScheduleList struct {
	ReportSchedules []ReportSchedule `json:"reportSchedules"`
}

// GetReportsParams defines parameters for GetReports.
type GetReportsParams struct {

	// A list of report types used to filter reports. When reportTypes is provided, the other filter parameters (processingStatuses, marketplaceIds, createdSince, createdUntil) and pageSize may also be provided. Either reportTypes or nextToken is required.
	ReportTypes *[]string `json:"reportTypes,omitempty"`

	// A list of processing statuses used to filter reports.
	ProcessingStatuses *[]string `json:"processingStatuses,omitempty"`

	// A list of marketplace identifiers used to filter reports. The reports returned will match at least one of the marketplaces that you specify.
	MarketplaceIds *[]string `json:"marketplaceIds,omitempty"`

	// The maximum number of reports to return in a single call.
	PageSize *int `json:"pageSize,omitempty"`

	// The earliest report creation date and time for reports to include in the response, in ISO 8601 date time format. The default is 90 days ago. Reports are retained for a maximum of 90 days.
	CreatedSince *time.Time `json:"createdSince,omitempty"`

	// The latest report creation date and time for reports to include in the response, in ISO 8601 date time format. The default is now.
	CreatedUntil *time.Time `json:"createdUntil,omitempty"`

	// A string token returned in the response to your previous request. nextToken is returned when the number of results exceeds the specified pageSize value. To get the next page of results, call the getReports operation and include this token as the only parameter. Specifying nextToken with any other parameters will cause the request to fail.
	NextToken *string `json:"nextToken,omitempty"`
}

// CreateReportJSONBody defines parameters for CreateReport.
type CreateReportJSONBody CreateReportSpecification

// GetReportSchedulesParams defines parameters for GetReportSchedules.
type GetReportSchedulesParams struct {

	// A list of report types used to filter report schedules.
	ReportTypes []string `json:"reportTypes"`
}

// CreateReportScheduleJSONBody defines parameters for CreateReportSchedule.
type CreateReportScheduleJSONBody CreateReportScheduleSpecification

// CreateReportRequestBody defines body for CreateReport for application/json ContentType.
type CreateReportJSONRequestBody CreateReportJSONBody

// CreateReportScheduleRequestBody defines body for CreateReportSchedule for application/json ContentType.
type CreateReportScheduleJSONRequestBody CreateReportScheduleJSONBody

// Getter for additional properties for ReportOptions. Returns the specified
// element and whether it was found
func (a ReportOptions) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ReportOptions
func (a *ReportOptions) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ReportOptions to handle AdditionalProperties
func (a *ReportOptions) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ReportOptions to handle AdditionalProperties
func (a ReportOptions) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}
