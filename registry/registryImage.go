package registry

import "strconv"

type RegistryImage struct {
	labels       map[string]string
	configDigest string
	tag          string
}

type RegistryImages struct {
	RegistryImages []RegistryImage
	sortLabels     []string
	desc           bool
	sortInt        bool
}

func (ris RegistryImages) Len() int {
	return len(ris.RegistryImages)
}

func (ris RegistryImages) Less(i, j int) bool {
	if ris.sortInt {
		return ris.LessByInt(i, j)
	}
	return ris.LessByAlpha(i, j)

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
