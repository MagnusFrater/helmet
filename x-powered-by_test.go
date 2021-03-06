package helmet

import (
	"testing"
)

func TestXPoweredBy_NewXPoweredBy(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		hide        bool
		replacement string
	}{
		{name: "Empty", hide: false, replacement: ""},
		{name: "Hide", hide: true, replacement: ""},
		{name: "Replacement", hide: false, replacement: "PHP 4.2.0"},
		{name: "Empty, Replacement", hide: true, replacement: "PHP 4.2.0"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			xpb := NewXPoweredBy(tc.hide, tc.replacement)

			if xpb.Hide != tc.hide {
				t.Errorf("Incorrect Hide\tExpected: %t\tActual: %t\n", tc.hide, xpb.Hide)
			}

			if xpb.Replacement != tc.replacement {
				t.Errorf("Incorrect Replacement\tExpected: %s\tActual: %s\n", tc.replacement, xpb.Replacement)
			}
		})
	}
}

func TestXPoweredBy_EmptyXPoweredBy(t *testing.T) {
	t.Parallel()

	xpb := EmptyXPoweredBy()

	if xpb.Hide != false {
		t.Errorf("Hide should be false\tActual: %t\n", xpb.Hide)
	}

	if xpb.Replacement != "" {
		t.Errorf("Replacement should be empty\tActual: %s\n", xpb.Replacement)
	}
}

func TestXPoweredBy_Empty(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		xPoweredBy    *XPoweredBy
		expectedEmpty bool
	}{
		{name: "Empty", xPoweredBy: EmptyXPoweredBy(), expectedEmpty: true},
		{name: "Hide", xPoweredBy: NewXPoweredBy(true, ""), expectedEmpty: false},
		{name: "Replacement", xPoweredBy: NewXPoweredBy(false, "PHP 4.2.0"), expectedEmpty: false},
		{name: "Hide, Replacement", xPoweredBy: NewXPoweredBy(true, "PHP 4.2.0"), expectedEmpty: false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.xPoweredBy.Empty()
			if exists != tc.expectedEmpty {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedEmpty, exists)
			}
		})
	}
}
