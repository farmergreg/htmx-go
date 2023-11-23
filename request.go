package htmx

import (
	"net/http"
)

// IsHTMXRequest returns true if the given request
// was made by HTMX.
//
// This can be used to add special logic for HTMX requests.
//
// Checks if 'HX-Request' is 'true'.
func IsHTMXRequest(r *http.Request) bool {
	return r.Header.Get(HeaderRequest) == "true"
}

// IsBoosted returns true if the given request
// was made via an element using 'hx-boost'.
//
// This can be used to add special logic for boosted requests.
//
// Checks if 'HX-Boosted' is 'true'.
func IsBoosted(r *http.Request) bool {
	return r.Header.Get(HeaderBoosted) == "true"
}

// IsHistoryRestoreRequest returns true if the given request
// is for history restoration after a miss in the local history cache.
//
// Checks if 'HX-History-Restore-Request' is 'true'.
func IsHistoryRestoreRequest(r *http.Request) bool {
	return r.Header.Get(HeaderHistoryRestoreRequest) == "true"
}

// GetCurrentURL returns the current URL that HTMX made this request from.
//
// Returns false if header 'HX-Current-URL' does not exist.
func GetCurrentURL(r *http.Request) (string, bool) {
	if _, ok := r.Header[http.CanonicalHeaderKey(HeaderCurrentURL)]; !ok {
		return "", false
	}
	return r.Header.Get(HeaderCurrentURL), true
}

// GetPrompt returns the user response to an hx-prompt from a given request.
//
// Returns false if header 'HX-Prompt' does not exist.
func GetPrompt(r *http.Request) (string, bool) {
	if _, ok := r.Header[http.CanonicalHeaderKey(HeaderPrompt)]; !ok {
		return "", false
	}
	return r.Header.Get(HeaderPrompt), true
}

// GetTarget returns the ID of the target element if it exists from a given request.
//
// Returns false if header 'HX-Target' does not exist.
func GetTarget(r *http.Request) (string, bool) {
	if _, ok := r.Header[http.CanonicalHeaderKey(HeaderTarget)]; !ok {
		return "", false
	}
	return r.Header.Get(HeaderTarget), true
}

// GetTriggerName returns the 'name' of the triggered element if it exists from a given request.
//
// Returns false if header 'HX-Trigger-Name' does not exist.
func GetTriggerName(r *http.Request) (string, bool) {
	if _, ok := r.Header[http.CanonicalHeaderKey(HeaderTriggerName)]; !ok {
		return "", false
	}
	return r.Header.Get(HeaderTriggerName), true
}

// GetTrigger returns the ID of the triggered element if it exists from a given request.
//
// Returns false if header 'HX-Trigger' does not exist.
func GetTrigger(r *http.Request) (string, bool) {
	if _, ok := r.Header[http.CanonicalHeaderKey(HeaderTrigger)]; !ok {
		return "", false
	}
	return r.Header.Get(HeaderTrigger), true
}
