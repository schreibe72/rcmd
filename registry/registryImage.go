package registry

import (
	"strconv"

	"github.com/hashicorp/go-version"
)

type RegistryImage struct {
	labels       map[string]string
	configDigest string
	tag          string
}

type RegistryImages struct {
	RegistryImages []RegistryImage
	sortLabels     []string
	desc           bool
	sortCriteria   string
}

func (ris RegistryImages) Len() int {
	return len(ris.RegistryImages)
}

func (ris RegistryImages) Less(i, j int) bool {
	switch ris.sortCriteria {
	case "string":
		return ris.LessByAlpha(i, j)
	case "int":
		return ris.LessByInt(i, j)
	case "version":
		return ris.LessByVersion(i, j)
	default:
		return ris.LessByAlpha(i, j)
	}
}

func (ris RegistryImages) LessByInt(i, j int) bool {

	if ris.desc {
		for _, label := range ris.sortLabels {
			a, _ := strconv.Atoi(ris.RegistryImages[i].labels[label])
			b, _ := strconv.Atoi(ris.RegistryImages[j].labels[label])
			if a > b {
				return true
			}
			if a < b {
				return false
			}
		}
	} else {
		for _, label := range ris.sortLabels {
			a, _ := strconv.Atoi(ris.RegistryImages[i].labels[label])
			b, _ := strconv.Atoi(ris.RegistryImages[j].labels[label])
			if a < b {
				return true
			}
			if a > b {
				return false
			}
		}
	}
	return false
}

func (ris RegistryImages) LessByVersion(i, j int) bool {

	if ris.desc {
		for _, label := range ris.sortLabels {
			v1, err := version.NewVersion(ris.RegistryImages[i].labels[label])
			if err != nil {
				return true
			}
			v2, err := version.NewVersion(ris.RegistryImages[j].labels[label])
			if err != nil {
				return false
			}
			if v1.GreaterThan(v2) {
				return true
			}
			if v1.LessThan(v2) {
				return false
			}
		}
	} else {
		for _, label := range ris.sortLabels {
			v1, err := version.NewVersion(ris.RegistryImages[i].labels[label])
			if err != nil {
				return false
			}
			v2, err := version.NewVersion(ris.RegistryImages[j].labels[label])
			if err != nil {
				return true
			}
			if v1.LessThan(v2) {
				return true
			}
			if v1.GreaterThan(v2) {
				return false
			}
		}
	}
	return false
}

func (ris RegistryImages) LessByAlpha(i, j int) bool {
	if ris.desc {
		for _, label := range ris.sortLabels {
			if ris.RegistryImages[i].labels[label] > ris.RegistryImages[j].labels[label] {
				return true
			}
			if ris.RegistryImages[i].labels[label] < ris.RegistryImages[j].labels[label] {
				return false
			}
		}

	} else {
		for _, label := range ris.sortLabels {
			if ris.RegistryImages[i].labels[label] < ris.RegistryImages[j].labels[label] {
				return true
			}
			if ris.RegistryImages[i].labels[label] > ris.RegistryImages[j].labels[label] {
				return false
			}
		}
	}
	return false
}

func (ri RegistryImages) Swap(i, j int) {
	ri.RegistryImages[i], ri.RegistryImages[j] = ri.RegistryImages[j], ri.RegistryImages[i]
}
