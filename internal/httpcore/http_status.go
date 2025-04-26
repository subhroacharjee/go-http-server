package httpcore

type HttpStatus uint

// Enum-style constants for HTTP status codes
const (
	StatusContinue           HttpStatus = 100
	StatusSwitchingProtocols HttpStatus = 101
	StatusProcessing         HttpStatus = 102
	StatusEarlyHints         HttpStatus = 103

	StatusOK                   HttpStatus = 200
	StatusCreated              HttpStatus = 201
	StatusAccepted             HttpStatus = 202
	StatusNonAuthoritativeInfo HttpStatus = 203
	StatusNoContent            HttpStatus = 204
	StatusResetContent         HttpStatus = 205
	StatusPartialContent       HttpStatus = 206
	StatusMultiStatus          HttpStatus = 207
	StatusAlreadyReported      HttpStatus = 208
	StatusIMUsed               HttpStatus = 226

	StatusMultipleChoices   HttpStatus = 300
	StatusMovedPermanently  HttpStatus = 301
	StatusFound             HttpStatus = 302
	StatusSeeOther          HttpStatus = 303
	StatusNotModified       HttpStatus = 304
	StatusUseProxy          HttpStatus = 305
	StatusTemporaryRedirect HttpStatus = 307
	StatusPermanentRedirect HttpStatus = 308

	StatusBadRequest                  HttpStatus = 400
	StatusUnauthorized                HttpStatus = 401
	StatusPaymentRequired             HttpStatus = 402
	StatusForbidden                   HttpStatus = 403
	StatusNotFound                    HttpStatus = 404
	StatusMethodNotAllowed            HttpStatus = 405
	StatusNotAcceptable               HttpStatus = 406
	StatusProxyAuthRequired           HttpStatus = 407
	StatusRequestTimeout              HttpStatus = 408
	StatusConflict                    HttpStatus = 409
	StatusGone                        HttpStatus = 410
	StatusLengthRequired              HttpStatus = 411
	StatusPreconditionFailed          HttpStatus = 412
	StatusPayloadTooLarge             HttpStatus = 413
	StatusURITooLong                  HttpStatus = 414
	StatusUnsupportedMediaType        HttpStatus = 415
	StatusRangeNotSatisfiable         HttpStatus = 416
	StatusExpectationFailed           HttpStatus = 417
	StatusTeapot                      HttpStatus = 418
	StatusMisdirectedRequest          HttpStatus = 421
	StatusUnprocessableEntity         HttpStatus = 422
	StatusLocked                      HttpStatus = 423
	StatusFailedDependency            HttpStatus = 424
	StatusTooEarly                    HttpStatus = 425
	StatusUpgradeRequired             HttpStatus = 426
	StatusPreconditionRequired        HttpStatus = 428
	StatusTooManyRequests             HttpStatus = 429
	StatusRequestHeaderFieldsTooLarge HttpStatus = 431
	StatusUnavailableForLegalReasons  HttpStatus = 451

	StatusInternalServerError           HttpStatus = 500
	StatusNotImplemented                HttpStatus = 501
	StatusBadGateway                    HttpStatus = 502
	StatusServiceUnavailable            HttpStatus = 503
	StatusGatewayTimeout                HttpStatus = 504
	StatusHTTPVersionNotSupported       HttpStatus = 505
	StatusVariantAlsoNegotiates         HttpStatus = 506
	StatusInsufficientStorage           HttpStatus = 507
	StatusLoopDetected                  HttpStatus = 508
	StatusNotExtended                   HttpStatus = 510
	StatusNetworkAuthenticationRequired HttpStatus = 511
)

// Global map of HttpStatus to messages
var httpStatusMessages = map[HttpStatus]string{
	StatusContinue:           "Continue",
	StatusSwitchingProtocols: "Switching Protocols",
	StatusProcessing:         "Processing",
	StatusEarlyHints:         "Early Hints",

	StatusOK:                   "OK",
	StatusCreated:              "Created",
	StatusAccepted:             "Accepted",
	StatusNonAuthoritativeInfo: "Non-Authoritative Information",
	StatusNoContent:            "No Content",
	StatusResetContent:         "Reset Content",
	StatusPartialContent:       "Partial Content",
	StatusMultiStatus:          "Multi-Status",
	StatusAlreadyReported:      "Already Reported",
	StatusIMUsed:               "IM Used",

	StatusMultipleChoices:   "Multiple Choices",
	StatusMovedPermanently:  "Moved Permanently",
	StatusFound:             "Found",
	StatusSeeOther:          "See Other",
	StatusNotModified:       "Not Modified",
	StatusUseProxy:          "Use Proxy",
	StatusTemporaryRedirect: "Temporary Redirect",
	StatusPermanentRedirect: "Permanent Redirect",

	StatusBadRequest:                  "Bad Request",
	StatusUnauthorized:                "Unauthorized",
	StatusPaymentRequired:             "Payment Required",
	StatusForbidden:                   "Forbidden",
	StatusNotFound:                    "Not Found",
	StatusMethodNotAllowed:            "Method Not Allowed",
	StatusNotAcceptable:               "Not Acceptable",
	StatusProxyAuthRequired:           "Proxy Authentication Required",
	StatusRequestTimeout:              "Request Timeout",
	StatusConflict:                    "Conflict",
	StatusGone:                        "Gone",
	StatusLengthRequired:              "Length Required",
	StatusPreconditionFailed:          "Precondition Failed",
	StatusPayloadTooLarge:             "Payload Too Large",
	StatusURITooLong:                  "URI Too Long",
	StatusUnsupportedMediaType:        "Unsupported Media Type",
	StatusRangeNotSatisfiable:         "Range Not Satisfiable",
	StatusExpectationFailed:           "Expectation Failed",
	StatusTeapot:                      "I'm a teapot",
	StatusMisdirectedRequest:          "Misdirected Request",
	StatusUnprocessableEntity:         "Unprocessable Entity",
	StatusLocked:                      "Locked",
	StatusFailedDependency:            "Failed Dependency",
	StatusTooEarly:                    "Too Early",
	StatusUpgradeRequired:             "Upgrade Required",
	StatusPreconditionRequired:        "Precondition Required",
	StatusTooManyRequests:             "Too Many Requests",
	StatusRequestHeaderFieldsTooLarge: "Request Header Fields Too Large",
	StatusUnavailableForLegalReasons:  "Unavailable For Legal Reasons",

	StatusInternalServerError:           "Internal Server Error",
	StatusNotImplemented:                "Not Implemented",
	StatusBadGateway:                    "Bad Gateway",
	StatusServiceUnavailable:            "Service Unavailable",
	StatusGatewayTimeout:                "Gateway Timeout",
	StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
	StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
	StatusInsufficientStorage:           "Insufficient Storage",
	StatusLoopDetected:                  "Loop Detected",
	StatusNotExtended:                   "Not Extended",
	StatusNetworkAuthenticationRequired: "Network Authentication Required",
}
