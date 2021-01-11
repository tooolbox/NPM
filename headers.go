package fiserv

// HeaderData is the body of data vital to the auth signature
// which are sent as request headers.
type HeaderData struct {
	ApiKey           ApiKeyParam
	ContentType      ContentTypeParam
	ClientRequestId  ClientRequestIdParam
	Timestamp        TimestampParam
	MessageSignature MessageSignatureParam
}
