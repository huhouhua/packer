package docker_registry

//easyjson:json
type ManifestSlice []Manifest

//easyjson:json
type Manifest struct {
	MediaType     string  `json:"mediaType"`
	SchemaVersion int     `json:"schemaVersion"`
	RepoTag       RepoTag `json:"repoTag"`
	Config        Config  `json:"config"`
	Layers        []Layer `json:"layers"`
}

//easyjson:json
type RepoTag struct {
	RepoName         string `json:"repoName"`
	Tag              string `json:"tag"`
	ExternalRegistry string `json:"externalRegistry"`
	UserName         string `json:"userName"`
	PassWord         string `json:"passWord"`
}

//easyjson:json
type Layer struct {
	MediaType string `json:"mediaType"`
	Digest    string `json:"digest"`
	Size      int64  `json:"size"`
	BlobPath  string `json:"blobPath"`
}

//easyjson:json
type Config struct {
	Layer
}
