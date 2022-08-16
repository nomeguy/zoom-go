package zoomAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
   API Documentation
   https://marketplace.zoom.us/docs/api-reference/zoom-api/methods/#operation/webinarRegistrantStatus
*/

type UpdateMeetingRegistrantStatusRequest struct {
	Action      string       `json:"action"`
	Registrants []Registrant `json:"registrants"`
}

type Registrant struct {
	Email string `json:"email"`
	Id    string `json:"id"`
}

func (client Client) UpdateMeetingRegistrantStatus(meetingId int, status string, registrants []Registrant) (err error) {

	updateMeetingRegistrantStatusRequest := UpdateMeetingRegistrantStatusRequest{
		Action:      status,
		Registrants: registrants,
	}

	var reqBody []byte
	reqBody, err = json.Marshal(updateMeetingRegistrantStatusRequest)
	if err != nil {
		return
	}

	endpoint := fmt.Sprintf("/%s/%d/registrants/status", client.getType(), meetingId)
	httpMethod := http.MethodPut

	_, err = client.executeRequestWithBody(endpoint, httpMethod, reqBody)
	if err != nil {
		return
	}

	return

}
