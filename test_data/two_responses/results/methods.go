package vkapiclient

import (
	"encoding/json"
)

// Returns information about photos by their IDs.
func PhotosGetById(
	// IDs separated with a comma, that are IDs of users who posted photos and IDs of photos themselves with an underscore character between such IDs. To get information about a photo in the group album, you shall specify group ID instead of user ID. Example:; \"1_129207899,6492_135055734, ; -20629724_271945303\"
	photos []string,
	// '1' — to return additional fields; '0' — (default)
	extended bool,
	// '1' — to return photo sizes in a.
	photoSizes bool) (response PhotosGetByIdResponse, extendedResponse PhotosGetByIdExtendedResponse, err error) {
	request := NewRequest("photos.getById")
	request.AddStrArrParam("photos", photos)
	request.AddBoolParam("client_secret", extended)
	request.AddBoolParam("photo_sizes", photoSizes)
	byteResponse := request.Send()
	errResponse := json.Unmarshal(byteResponse, &response)
	errExtendedResponse := json.Unmarshal(byteResponse, &extendedResponse)
	if errResponse != nil {
		err = errResponse
	} else if errExtendedResponse != nil {
		err = errExtendedResponse
	} else {
		err = nil
	}
	return
}
