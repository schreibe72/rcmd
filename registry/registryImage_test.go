package registry

import (
	"sort"
	"testing"
)

func testEq(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
func TestSortAlphaASC(t *testing.T) {
	var ris RegistryImages
	ris.sortInt = false
	ris.desc = false
	ris.sortLabels = []string{"VERSION", "OS"}
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{"VERSION": "0.0.1", "OS": "14.4.2"}, tag: "c"})
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{"VERSION": "0.0.1", "OS": "14.4.1"}, tag: "b"})
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{"VERSION": "0.2.1", "OS": "14.4.0"}, tag: "d"})
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{}, tag: "a"})
	sort.Sort(ris)
	result := []string{}
	for _, i := range ris.RegistryImages {
		result = append(result, i.tag)
	}
	expected := []string{"a", "b", "c", "d"}
	if !testEq(result, expected) {
		t.Errorf("Sort ASC: %v != %v", result, expected)
	}
}

func TestSortAlphaDESC(t *testing.T) {
	var ris RegistryImages
	ris.sortInt = false
	ris.desc = true
	ris.sortLabels = []string{"VERSION", "OS"}
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{"VERSION": "0.0.1", "OS": "14.4.2"}, tag: "c"})
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{"VERSION": "0.0.1", "OS": "14.4.1"}, tag: "b"})
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{"VERSION": "0.2.1", "OS": "14.4.0"}, tag: "d"})
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{}, tag: "a"})
	sort.Sort(ris)
	result := []string{}
	for _, i := range ris.RegistryImages {
		result = append(result, i.tag)
	}
	expected := []string{"d", "c", "b", "a"}
	if !testEq(result, expected) {
		t.Errorf("Sort DESC: %v != %v", result, expected)
	}
}

func TestSortIntASC(t *testing.T) {
	var ris RegistryImages
	ris.sortInt = true
	ris.desc = false
	ris.sortLabels = []string{"VERSION", "OS"}
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{"VERSION": "1", "OS": "3"}, tag: "c"})
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{"VERSION": "1", "OS": "1"}, tag: "b"})
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{"VERSION": "2", "OS": "1"}, tag: "d"})
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{"VERSION": "10", "OS": "34"}, tag: "e"})
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{}, tag: "a"})
	sort.Sort(ris)
	result := []string{}
	for _, i := range ris.RegistryImages {
		result = append(result, i.tag)
	}
	expected := []string{"a", "b", "c", "d", "e"}
	if !testEq(result, expected) {
		t.Errorf("INT Sort ASC: %v != %v", result, expected)
	}
}

func TestSortIntDESC(t *testing.T) {
	var ris RegistryImages
	ris.sortInt = true
	ris.desc = true
	ris.sortLabels = []string{"VERSION", "OS"}
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{"VERSION": "1", "OS": "3"}, tag: "c"})
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{"VERSION": "1", "OS": "1"}, tag: "b"})
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{"VERSION": "2", "OS": "1"}, tag: "d"})
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{"VERSION": "10", "OS": "34"}, tag: "e"})
	ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: map[string]string{}, tag: "a"})
	sort.Sort(ris)
	result := []string{}
	for _, i := range ris.RegistryImages {
		result = append(result, i.tag)
	}
	expected := []string{"e", "d", "c", "b", "a"}
	if !testEq(result, expected) {
		t.Errorf("INT Sort DESC: %v != %v", result, expected)
	}
}
