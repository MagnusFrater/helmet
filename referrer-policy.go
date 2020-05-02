package helmet

import (
	"net/http"
	"strings"
)

// HeaderReferrerPolicy is the Referrer-Policy HTTP security header.
const HeaderReferrerPolicy = "Referrer-Policy"

// X-Frame-Options options.
const (
	DirectiveNoReferrer                  ReferrerPolicyDirective = "no-referrer"
	DirectiveNoReferrerWhenDowngrade     ReferrerPolicyDirective = "no-referrer-when-downgrade"
	DirectiveOrigin                      ReferrerPolicyDirective = "origin"
	DirectiveOriginWhenCrossOrigin       ReferrerPolicyDirective = "origin-when-cross-origin"
	DirectiveSmaeOrigin                  ReferrerPolicyDirective = "same-origin"
	DirectiveStrictOrigin                ReferrerPolicyDirective = "strict-origin"
	DirectiveStrictOriginWhenCrossOrigin ReferrerPolicyDirective = "strict-origin-when-cross-origin"
	DirectiveUnsafeURL                   ReferrerPolicyDirective = "unsafe-url"
)

type (
	// ReferrerPolicyDirective represents a Referrer-Policy directive.
	ReferrerPolicyDirective string

	// ReferrerPolicy represents the Referrer-Policy HTTP security header.
	ReferrerPolicy struct {
		// Make note that if there is more than 1 directive, the desired directive should be specified last.
		// Every other directive is a fallback, prioritized in the order from right-to-left.
		policies []ReferrerPolicyDirective

		cache string
	}
)

// NewReferrerPolicy creates a new Referrer-Policy.
func NewReferrerPolicy(directives ...ReferrerPolicyDirective) *ReferrerPolicy {
	return &ReferrerPolicy{directives, ""}
}

// EmptyReferrerPolicy creates a blank slate Referrer-Policy.
func EmptyReferrerPolicy() *ReferrerPolicy {
	return NewReferrerPolicy()
}

func (rp *ReferrerPolicy) String() string {
	if rp.cache != "" {
		return rp.cache
	}

	directivesAsStrings := []string{}
	for _, directive := range rp.policies {
		directivesAsStrings = append(directivesAsStrings, string(directive))
	}

	rp.cache = strings.Join(directivesAsStrings, ", ")
	return rp.cache
}

// Exists returns whether the Referrer-Policy has been set.
func (rp *ReferrerPolicy) Exists() bool {
	return len(rp.policies) > 0
}

// Header adds the Referrer-Policy HTTP header to the given http.ResponseWriter.
func (rp *ReferrerPolicy) Header(w http.ResponseWriter) {
	if rp.Exists() {
		w.Header().Set(HeaderReferrerPolicy, rp.String())
	}
}