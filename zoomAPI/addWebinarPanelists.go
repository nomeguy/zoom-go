package zoomAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Panelist struct {
	Email               string `json:"email"`
	Name                string `json:"name"`
	VirtualBackgroundId string `json:"virtual_background_id"`
	NameTagId           string `json:"name_tag_id"`
	NameTagName         string `json:"name_tag_name"`
	NameTagPronouns     string `json:"name_tag_pronouns"`
	NameTagDescription  string `json:"name_tag_description"`
}

type AddWebinarPanelistsRequest struct {
	Panelists []Panelist `json:"panelists"`
}

// AddWebinarPanelists /*
/*
   API Documentation
   https://marketplace.zoom.us/docs/api-reference/zoom-api/methods/#operation/webinarPanelistCreate
*/
func (client Client) AddWebinarPanelists(webinarId int, panelists []Panelist) (err error) {
	addWebinarPanelistsRequest := AddWebinarPanelistsRequest{
		Panelists: panelists,
	}

	endpoint := fmt.Sprintf("/webinars/%d/panelists", webinarId)
	httpMethod := http.MethodPost

	var reqBody []byte
	reqBody, err = json.Marshal(addWebinarPanelistsRequest)
	if err != nil {
		return
	}

	_, err = client.executeRequestWithBody(endpoint, httpMethod, reqBody)
	if err != nil {
		return
	}

	return
}
