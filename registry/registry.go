package registry

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/docker/docker/api/types"
	"github.com/schreibe72/docker-registry-client/registry"
)

type privatRegistry struct {
	registry.Registry
}

func New(url, username, password string, verbose bool) (*privatRegistry, error) {
	transport := http.DefaultTransport
	logger := registry.Log
	if !verbose {
		logger = registry.Quiet
	}
	r, err := registry.NewFromTransport(url, username, password, transport, logger)
	if err != nil {
		return nil, err
	}

	var pr privatRegistry
	pr.Registry = *r
	return &pr, nil
}

func (r *privatRegistry) Labels(repo, tag string) (map[string]string, error) {
	m, err := r.ManifestV2(repo, tag)
	if err != nil {
		return map[string]string{}, err
	}
	reader, err := r.DownloadLayer(repo, m.Config.Digest)
	if err != nil {
		return map[string]string{}, err
	}
	if reader != nil {
		defer reader.Close()
	}
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return map[string]string{}, err
	}
	var image types.ImageInspect
	err = json.Unmarshal(b, &image)
	if err != nil {
		return map[string]string{}, err
	}
	return image.Config.Labels, nil
}

func (r *privatRegistry) SortedTagsByLabel(repo string, sortLabels []string, desc bool, sortCriteria string) ([]string, error) {

	var ris RegistryImages
	ris.desc = desc
	ris.sortCriteria = sortCriteria
	ris.sortLabels = sortLabels

	tags, err := r.Tags(repo)
	if err != nil {
		return []string{}, err
	}
	for _, tag := range tags {
		lables, err := r.Labels(repo, tag)
		if err != nil {
			return []string{}, err
		}
		ris.RegistryImages = append(ris.RegistryImages, RegistryImage{labels: lables, tag: tag})
	}
	sort.Sort(ris)
	var result []string
	for _, ri := range ris.RegistryImages {
		result = append(result, ri.tag)
	}
	return result, nil
}

func (r *privatRegistry) DeleteTag(repo, tag string) error {
	digest, err := r.ManifestDigestV2(repo, tag)
	if err != nil {
		return err
	}
	err = r.DeleteManifest(repo, digest)
	return err
}
