// Package spec provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package spec

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"RgwzAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}