package zoomAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
   API Documentation
   https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetingregistrants
*/
func (client Client) ListMeetingRegistrants(meetingId int, status string) (apiResponse ListMeetingRegistrantsResponse, err error) {

	endpoint := fmt.Sprintf("/%s/%d/registrants?status=%s", client.getType(), meetingId, status)
	httpMethod := http.MethodGet

	var b []byte
	b, err = client.executeRequest(endpoint, httpMethod)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &apiResponse)
	if err != nil {
		return
	}

	return
}
