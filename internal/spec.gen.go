// Package internal provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/pseudo-su/oapi-ui-codegen DO NOT EDIT.
package internal

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"strings"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xVy27jRhD8lUEnhwSgRXk3yIG3ZGMgAvJYIHtb+NBLNqnZzMs9TVmCwH8PZoa0vZIc",
	"J0AC+CQ+pqurqrvEI7TeBu/ISYTmCLHdksV8ecPsOV0E9oFYNOXHre8o/faeLQo0oJ28fQMVyCFQuaWB",
	"GKYKLMWIQz49v4zC2g0wTRUw3Y2aqYPmY8F8PH/7AOY/faZWEtZ7knMuujtl8v13F5k4tJdoVCA4vExP",
	"dzBDPMOskBGy+eJrph4a+Kp+tLaefa2TjOkBBJnxAFPqp13vU7HRLbmYyRbS8OvmQy7RYtLtH/c4DMQq",
	"tRXPybcdcdTeQQPXq/VqnU77QA6Dhgbe5kcVBJRtpleHmfBQLE2GomjvNh00YHSUrChVMFoS4gjNxyN0",
	"FFvWQUqnn/29sugOKstW4hWTjOwUivKOlGhL6huLe3W9Xn8LSSA0cDcSHxYzUzOrBZ6a3aOJVM1r+I/W",
	"bLpN9TF4F8tSvFmvy546IZcVYghGt1lj/Tkm+scnHV4YVyzz+VL9DyrgQJ3KA1S+V6FYtiXssl9H2F85",
	"2uf2p6VGuz+TYbIllc5krATyKOOpA6e7Wdj0OBr5z3SWrF8QOjraB2qFOkXzmQriaC3yARr4RUdRaMyi",
	"X3BIywL59naqIPh4YclaJhSa1+xkeNfnlv02GvNgDrwi+e+yDoVJ/rn6qSpZq4+BZNNNz2Yubv39e5If",
	"D5vupdh92JLSXVqWtD2BZE4ea9rRkrIU9ceQ5eZfhEx4pL/bsP85UZeMvllsXhonXah2aHR6djdSlNc0",
	"+Y3rveo9K1QxUKt73T6zBKmMeLeMc2QDDdSDN+iGq/RGt3QlZINBoXp3DadM3o3M5ET95C1qB1O1gGxF",
	"Qmzq+tjlN9O/Af1dtsQLZAU7ZI2fTBl3gSsxnN0GDHpFe7TB0Kr19gzvxgmxasco3qoCkP/ikId54E9Z",
	"N3UORvp8rWL5nq20Tzyn2+mvAAAA//+J82hQkwgAAA==",
}

// GetOpenAPISpec returns the Swagger specification corresponding to the generated code
// in this file.
func GetOpenAPISpec() (*openapi3.Swagger, error) {
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

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
