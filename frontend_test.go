package main

import (
	"testing"
)

func TestFrontEnd(t *testing.T) {
	runFrontEndTest(t, "no_object_refs", noObjectRefPrimitives)
	runFrontEndTest(t, "two_responses", twoResponsesPrimitives)
}

func runFrontEndTest(
	t *testing.T,
	testName string,
	primitivesFactory func() (*PrimitivesSystem, error)) {
	t.Logf("Runing test '%s'...", testName)
	compiler := newTestCompiler(t, testName)
	actual, err := compiler.frontEnd()
	if err != nil {
		t.Fatalf("Front end failed: %v", err)
	}
	if actual == nil {
		t.Fatalf("Front end returned nil primitives")
	}
	expected, err := primitivesFactory()
	if err != nil {
		t.Fatalf("Cannot get expected primitives system: %v", err)
	}
	assertStructsEqual(t, expected, actual)
}

func noObjectRefPrimitives() (*PrimitivesSystem, error) {
	primitives := new(PrimitivesSystem)
	// Errors
	invalidPass := &ErrorPrimitive{
		Name:        "API_ERROR_AUTH_PARAM_PASSWORD",
		Code:        1111,
		Desctiption: "Invalid password",
	}
	incorrectCode := &ErrorPrimitive{
		Name:        "API_ERROR_AUTH_PARAM_CODE",
		Code:        1110,
		Desctiption: "Incorrect code",
	}
	primitives.ErrorPrimitives = append(
		primitives.ErrorPrimitives, invalidPass, incorrectCode)
	// Methods
	authConfirmResponseBody := NewObjectPrimitive("auth_confirm_response")
	successType := NewIntegerEnumTypePrimitive().AddRecord(1, "ok")
	successParam := NewParamPrimitive("success", "1 if success", successType, false)
	authConfirmResponseBody.AddProp(successParam)
	userIdParam := NewParamPrimitive("user_id", "User ID", NewIntegerTypePrimitiveNoBounds(), false)
	authConfirmResponseBody.AddProp(userIdParam)
	authConfirmResponse := NewResponsePrimitive("response", authConfirmResponseBody)

	method := &MethodPrimitive{
		Name:        "auth.confirm",
		Description: "Completes a user's registration (begun with the [vk.com/dev/auth.signup|auth.signup] method) using an authorization code.",
		Open:        true,
	}
	method.
		AddError(invalidPass).
		AddError(incorrectCode).
		AddParam("client_id", "", NewIntegerTypePrimitiveNoBounds(), true).
		AddParam("client_secret", "", NewStringTypePrimitive(), true).
		AddParam("phone", "", NewStringTypePrimitive(), true).
		AddParam("code", "", NewStringTypePrimitive(), true).
		AddParam("password", "", NewStringTypePrimitive(), false).
		AddParam("test_mode", "", NewBooleanTypePrimitive(), false).
		AddParam("intro", "", NewIntegerTypePrimitiveMinBound(0), false).
		AddResponse(authConfirmResponse)
	primitives.MethodPrimitives = append(primitives.MethodPrimitives, method)
	// Objects
	return primitives, nil
}

func twoResponsesPrimitives() (*PrimitivesSystem, error) {
	primitives := new(PrimitivesSystem)
	// Errors

	// Objects
	photoSizeEnum := NewStringEnumTypePrimitive()
	photoSizeEnum.
		AddRecord("s", "").
		AddRecord("m", "").
		AddRecord("x", "").
		AddRecord("o", "").
		AddRecord("p", "").
		AddRecord("q", "").
		AddRecord("r", "").
		AddRecord("y", "").
		AddRecord("z", "").
		AddRecord("w", "")

	photosPhotoSizes := NewObjectPrimitive("photos_photo_sizes")
	photosPhotoSizes.
		AddProp(NewParamPrimitive("src", "URL of the image", NewStringTypePrimitive(), true)).
		AddProp(NewParamPrimitive("width", "Width in px", NewIntegerTypePrimitiveMinBound(0), true)).
		AddProp(NewParamPrimitive("height", "Height in px", NewIntegerTypePrimitiveMinBound(0), true)).
		AddProp(NewParamPrimitive("type", "Size type", photoSizeEnum, true))

	photosPhoto := NewObjectPrimitive("photos_photo")
	photosPhoto.
		AddProp(NewParamPrimitive("id", "Photo ID", NewIntegerTypePrimitiveMinBound(0), true)).
		AddProp(NewParamPrimitive("album_id", "Album ID", NewIntegerTypePrimitiveNoBounds(), true)).
		AddProp(NewParamPrimitive("owner_id", "Photo owner's ID", NewIntegerTypePrimitiveNoBounds(), true)).
		AddProp(NewParamPrimitive("user_id", "ID of the user who have uploaded the photo", NewIntegerTypePrimitiveMinBound(1), false)).
		AddProp(NewParamPrimitive("sizes", "", NewArrayTypePrimitiveNoLimits(NewObjectTypePrimitive(photosPhotoSizes)), false)).
		AddProp(NewParamPrimitive("photo_75", "URL of image with 75 px width", NewStringTypePrimitive(), false)).
		AddProp(NewParamPrimitive("photo_130", "URL of image with 130 px width", NewStringTypePrimitive(), false)).
		AddProp(NewParamPrimitive("photo_604", "URL of image with 604 px width", NewStringTypePrimitive(), false)).
		AddProp(NewParamPrimitive("photo_807", "URL of image with 807 px width", NewStringTypePrimitive(), false)).
		AddProp(NewParamPrimitive("photo_1280", "URL of image with 1280 px width", NewStringTypePrimitive(), false)).
		AddProp(NewParamPrimitive("photo_2560", "URL of image with 2560 px width", NewStringTypePrimitive(), false)).
		AddProp(NewParamPrimitive("post_id", "Post ID", NewIntegerTypePrimitiveMinBound(1), false)).
		AddProp(NewParamPrimitive("width", "Original photo width", NewIntegerTypePrimitiveMinBound(0), false)).
		AddProp(NewParamPrimitive("height", "Original photo height", NewIntegerTypePrimitiveMinBound(0), false)).
		AddProp(NewParamPrimitive("text", "Photo caption", NewStringTypePrimitive(), false)).
		AddProp(NewParamPrimitive("date", "Date when uploaded", NewStringTypePrimitive(), true)).
		AddProp(NewParamPrimitive("lat", "Latitude", NewNumberTypePrimitive(), false)).
		AddProp(NewParamPrimitive("long", "Longitude", NewNumberTypePrimitive(), false)).
		AddProp(NewParamPrimitive("access_key", "Access key for the photo", NewStringTypePrimitive(), false))

	baseBoolInt := NewIntegerEnumTypePrimitive()
	baseBoolInt.
		AddRecord(0, "no").
		AddRecord(1, "yes")

	baseLikes := NewObjectPrimitive("base_likes")
	baseLikes.
		AddProp(NewParamPrimitive("user_likes", "Information whether current user likes the photo", baseBoolInt, false)).
		AddProp(NewParamPrimitive("count", "Likes number", NewIntegerTypePrimitiveNoBounds(), false))

	baseObjectCount := NewObjectPrimitive("base_object_count")
	baseObjectCount.AddProp(NewParamPrimitive("count", "Items count", NewIntegerTypePrimitiveNoBounds(), false))

	photosPhotoFull := NewObjectPrimitive("photos_photo_full")
	photosPhoto.
		AddProp(NewParamPrimitive("id", "Photo ID", NewIntegerTypePrimitiveMinBound(0), true)).
		AddProp(NewParamPrimitive("album_id", "Album ID", NewIntegerTypePrimitiveNoBounds(), true)).
		AddProp(NewParamPrimitive("owner_id", "Photo owner's ID", NewIntegerTypePrimitiveNoBounds(), true)).
		AddProp(NewParamPrimitive("user_id", "ID of the user who have uploaded the photo", NewIntegerTypePrimitiveMinBound(1), false)).
		AddProp(NewParamPrimitive("sizes", "", NewArrayTypePrimitiveNoLimits(NewObjectTypePrimitive(photosPhotoSizes)), false)).
		AddProp(NewParamPrimitive("photo_75", "URL of image with 75 px width", NewStringTypePrimitive(), false)).
		AddProp(NewParamPrimitive("photo_130", "URL of image with 130 px width", NewStringTypePrimitive(), false)).
		AddProp(NewParamPrimitive("photo_604", "URL of image with 604 px width", NewStringTypePrimitive(), false)).
		AddProp(NewParamPrimitive("photo_807", "URL of image with 807 px width", NewStringTypePrimitive(), false)).
		AddProp(NewParamPrimitive("photo_1280", "URL of image with 1280 px width", NewStringTypePrimitive(), false)).
		AddProp(NewParamPrimitive("photo_2560", "URL of image with 2560 px width", NewStringTypePrimitive(), false)).
		AddProp(NewParamPrimitive("post_id", "Post ID", NewIntegerTypePrimitiveMinBound(1), false)).
		AddProp(NewParamPrimitive("width", "Original photo width", NewIntegerTypePrimitiveMinBound(0), false)).
		AddProp(NewParamPrimitive("height", "Original photo height", NewIntegerTypePrimitiveMinBound(0), false)).
		AddProp(NewParamPrimitive("text", "Photo caption", NewStringTypePrimitive(), false)).
		AddProp(NewParamPrimitive("date", "Date when uploaded", NewStringTypePrimitive(), true)).
		AddProp(NewParamPrimitive("lat", "Latitude", NewNumberTypePrimitive(), false)).
		AddProp(NewParamPrimitive("long", "Longitude", NewNumberTypePrimitive(), false)).
		AddProp(NewParamPrimitive("access_key", "Access key for the photo", NewStringTypePrimitive(), false)).
		AddProp(NewParamPrimitive("likes", "", NewObjectTypePrimitive(baseLikes), false)).
		AddProp(NewParamPrimitive("reposts", "", NewObjectTypePrimitive(baseObjectCount), false)).
		AddProp(NewParamPrimitive("comments", "", NewObjectTypePrimitive(baseObjectCount), false)).
		AddProp(NewParamPrimitive("can_comment", "Information whether current user can comment the photo", baseBoolInt, false)).
		AddProp(NewParamPrimitive("tags", "", NewObjectTypePrimitive(baseObjectCount), false))

	// Responses
	photosGetByIdResponseBody := NewObjectPrimitive("photos_getById_response")
	photosGetByIdResponseBody.AddProp(NewParamPrimitive("response", "", NewArrayTypePrimitiveNoLimits(NewObjectTypePrimitive(photosPhoto)), false))
	photosGetByIdResponse := NewResponsePrimitive("response", photosGetByIdResponseBody)

	photosGetByIdExtendedResponseBody := NewObjectPrimitive("photos_getById_extended_response")
	photosGetByIdExtendedResponseBody.AddProp(NewParamPrimitive("response", "", NewArrayTypePrimitiveNoLimits(NewObjectTypePrimitive(photosPhotoFull)), false))
	photosGetByIdExtendedResponse := NewResponsePrimitive("extendedResponse", photosGetByIdExtendedResponseBody)

	// Methods
	method := &MethodPrimitive{
		Name:        "auth.confirm",
		Description: "Completes a user's registration (begun with the [vk.com/dev/auth.signup|auth.signup] method) using an authorization code.",
		Open:        true,
	}
	method.
		AddParam(
			"photos",
			"IDs separated with a comma, that are IDs of users who posted photos and IDs of photos themselves with an underscore character between such IDs. To get information about a photo in the group album, you shall specify group ID instead of user ID. Example:; \"1_129207899,6492_135055734, ; -20629724_271945303\"",
			NewArrayTypePrimitiveNoLimits(NewStringTypePrimitive()),
			true).
		AddParam("extended", "'1' — to return additional fields; '0' — (default)", NewBooleanTypePrimitive(), false).
		AddParam("photo_sizes", "'1' — to return photo sizes in a", NewBooleanTypePrimitive(), false).
		AddResponse(photosGetByIdResponse).
		AddResponse(photosGetByIdExtendedResponse)
	primitives.MethodPrimitives = append(primitives.MethodPrimitives, method)
	return primitives, nil
}
