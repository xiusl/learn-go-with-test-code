package _2_reflection

import (
	"reflect"
	"testing"
)


func TestWalk(t *testing.T) {
	expected := "Jack"
	var got []string

	x := struct{
		Name string
	}{expected}

	Walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want 1", len(got))
	}

	if got[0] != expected {
		t.Errorf("got '%s' want '%s'", got[0], expected)
	}

}


func TestWalkV2(t *testing.T) {

	type Profile struct {
		City string
		Age int
	}

	type Person struct {
		Name string
		Profile Profile
	}

	testCases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			Name: "Struct with one string field",
			Input: struct {
				Name string
			}{"Jack"},
			ExpectedCalls: []string{"Jack"},
		},
		{
			Name: "Struct with two string",
			Input: struct {
				Name string
				City string
			}{"Jack", "HZ"},
			ExpectedCalls: []string{"Jack", "HZ"},
		},
		{
			Name: "Struct with int and string",
			Input: struct {
				Name string
				City string
				Age int
			}{"Jack", "HZ", 18},
			ExpectedCalls: []string{"Jack", "HZ"},
		},
		{
			Name: "Nested",
			Input: Person{
				Name: "Jack",
				Profile: Profile{
					City: "HZ",
					Age: 18,
				},
			},
			ExpectedCalls: []string{"Jack", "HZ"},
		},
		{
			Name: "Pointed",
			Input: &Person{
				Name: "Jack",
				Profile: Profile{
					City: "HZ",
					Age: 18,
				},
			},
			ExpectedCalls: []string{"Jack", "HZ"},
		},
		{
			Name: "Pointed",
			Input: []Profile{
				{
					City: "HZ",
					Age:  18,
				},
				{
					City: "SH",
					Age: 20,
				},
			},
			ExpectedCalls: []string{"HZ", "SH"},
		},
		{
			Name: "Array",
			Input: [2]Profile{
				{
					City: "HZ",
					Age:  18,
				},
				{
					City: "SH",
					Age: 20,
				},
			},
			ExpectedCalls: []string{"HZ", "SH"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			var got []string
			WalkV2(tc.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tc.ExpectedCalls) {
				t.Errorf("got %v want %v", got, tc.ExpectedCalls)
			}
		})
	}

	t.Run("Map", func(t *testing.T) {
		dict := map[string]string{
			"Foo": "aaa",
			"Baz": "bbb",
		}

		var got []string
		WalkV2(dict, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "aaa")
		assertContains(t, got, "bbb")
	})
}

func assertContains(t *testing.T, haystack []string, needle string)  {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain '%s' but it didnt", haystack, needle)
	}
}
