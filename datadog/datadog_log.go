package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

type ResponseBody struct {
	Attributes Attributes
	ID         string
}

type Attributes struct {
	Attributes interface{}
	Timestamp  string
	Status     string
}

type FinalPayload struct {
	Entries []Entry
}

type Entry struct {
	InsertID    string
	JsonPayload interface{}
	Timestamp   string
	Severity    string
}

func main() {
	//var responses []ResponseBody
	body := datadogV2.LogsListRequest{
		Filter: &datadogV2.LogsQueryFilter{
			Query: datadog.PtrString("service:my-service env:prod @context:BU->Client @severity:error"),
			Indexes: []string{
				"*",
			},
			From: datadog.PtrString("2023-03-23T07:37:00.000Z"),
			To:   datadog.PtrString("2023-03-23T07:40:00.000Z"),
		},
		Sort: datadogV2.LOGSSORT_TIMESTAMP_ASCENDING.Ptr(),
		Page: &datadogV2.LogsListRequestPage{
			Limit: datadog.PtrInt32(2),
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewLogsApi(apiClient)

	resp, _ := api.ListLogsWithPagination(ctx, *datadogV2.NewListLogsOptionalParameters().WithBody(body))
	var entries []Entry
	for paginationResult := range resp {
		if paginationResult.Error != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `LogsApi.ListLogs`: %v\n", paginationResult.Error)
		}

		responseContent, _ := json.MarshalIndent(paginationResult.Item, "", "  ")
		var responseBody ResponseBody
		json.Unmarshal(responseContent, &responseBody)
		entries = append(entries, Entry{
			InsertID:    responseBody.ID,
			JsonPayload: responseBody.Attributes.Attributes,
			Timestamp:   responseBody.Attributes.Timestamp,
			Severity:    responseBody.Attributes.Status,
		})
	}

	finalPayload := FinalPayload{
		Entries: entries,
	}

	finalResponse, _ := json.Marshal(finalPayload)
	_ = ioutil.WriteFile("file.json", finalResponse, 0644)

}
