package serialize

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type friend struct {
	Friend string `json:"friend"`
}

type data struct {
	Name      string   `json:"name"`
	Abilities []string `json:"abilities"`
	Combos    friend   `json:"combos"`
}

func TestEncode(t *testing.T) {
	t.Run("valid struct encoding", func(t *testing.T) {
		jeff := &data{
			Name:      "jeff",
			Abilities: []string{"wallride", "spit", "cute"},
			Combos:    friend{"groot"},
		}

		want := `{"name":"jeff","abilities":["wallride","spit","cute"],"combos":{"friend":"groot"}}`
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		err := EncodeJSON(response, request, 200, jeff)

		if err != nil {
			t.Errorf("got unwanted err %v", err)
		}

		if response.Code != 200 {
			t.Errorf("got %d status code but want %d", response.Code, 200)
		}

		if strings.TrimSuffix(response.Body.String(), "\n") != want {
			t.Errorf("got %s but want %s", response.Body.String(), want)
		}
	})
}

func TestDecode(t *testing.T) {

	t.Run("valid decoding", func(t *testing.T) {
		want := &data{
			Name:      "jeff",
			Abilities: []string{"wallride", "spit", "cute"},
			Combos:    friend{"groot"},
		}

		blob := `{"name":"jeff","abilities":["wallride","spit","cute"],"combos":{"friend":"groot"}}`

		request := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader([]byte(blob)))

		got, err := DecodeJSON[data](request)

		if err != nil {
			t.Errorf("got unwanted err %v", err)
		}

		if !reflect.DeepEqual(&got, want) {
			t.Errorf("got %q but want %q", got, want)
		}

	})

	t.Run("invalid decode", func(t *testing.T) {

		blob := `aggba`

		request := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader([]byte(blob)))

		got, err := DecodeJSON[data](request)

		if err == nil {
			t.Errorf("did not get err on decoding, instead got %v", got)
		}
	})
}

// Tests that an ecoding then decoding operation produces the exact same reference
func TestSymmetry(t *testing.T) {
	cases := []struct {
		Name      string
		Json      string
		Reference interface{}
	}{
		{
			Name: "A",
			Json: `{"name":"jeff","abilities":["wallride","spit","cute"],"combos":{"friend":"groot"}}`,
			Reference: data{
				Name:      "jeff",
				Abilities: []string{"wallride", "spit", "cute"},
				Combos:    friend{"groot"},
			},
		},
	}

	for _, test := range cases {

		t.Run(test.Name, func(t *testing.T) {

			// First encode our reference into json
			response := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, "/", nil)

			err := EncodeJSON(response, request, 200, test.Reference)

			if err != nil {
				t.Errorf("unwanted error on encoding %v", err)
			}

			// Now decode our json back into the reference
			request = httptest.NewRequest(http.MethodGet, "/", strings.NewReader(response.Body.String()))
			got, err := DecodeJSON[data](request)

			if err != nil {
				t.Errorf("unwanted error on decoding %v", err)
			}

			if !reflect.DeepEqual(got, test.Reference) {
				t.Errorf("got %q but want %q", got, test.Reference)
			}
		})

	}
}
