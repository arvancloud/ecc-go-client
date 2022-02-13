package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type SSHKeysApiService service

/*
SSHKeysApiService Create new SSH key.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param region Region code
 * @param body interface{}

@return interface{}
*/
func (a *SSHKeysApiService) Create(ctx context.Context, region string, body interface{}) (interface{}, *http.Response, error) {
	httpMethod := strings.ToUpper("Post")

	// create path and map variables
	path := a.client.cfg.BasePath + "/regions/{region}/ssh-keys"
	path = strings.Replace(path, "{"+"region"+"}", fmt.Sprintf("%v", region), -1)

	postBody := &body

	r, err := a.client.prepareRequest(ctx, path, httpMethod, postBody)
	if err != nil {
		return nil, nil, err
	}

	return a.client.callAPI(r)
}

/*
SSHKeysApiService Shows details of a SSH key.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param region Region code
 * @param name The name of the SSH key

@return interface{}
*/
func (a *SSHKeysApiService) Get(ctx context.Context, region string, name string) (interface{}, *http.Response, error) {
	httpMethod := strings.ToUpper("Get")

	// create path and map variables
	path := a.client.cfg.BasePath + "/regions/{region}/ssh-keys/{name}"
	path = strings.Replace(path, "{"+"region"+"}", fmt.Sprintf("%v", region), -1)
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	r, err := a.client.prepareRequest(ctx, path, httpMethod, nil)
	if err != nil {
		return nil, nil, err
	}

	return a.client.callAPI(r)
}

/*
SSHKeysApiService Return all SSH keys.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param region Region code

@return interface{}
*/
func (a *SSHKeysApiService) All(ctx context.Context, region string) (interface{}, *http.Response, error) {
	httpMethod := strings.ToUpper("Get")

	// create path and map variables
	path := a.client.cfg.BasePath + "/regions/{region}/ssh-keys"
	path = strings.Replace(path, "{"+"region"+"}", fmt.Sprintf("%v", region), -1)

	r, err := a.client.prepareRequest(ctx, path, httpMethod, nil)
	if err != nil {
		return nil, nil, err
	}

	return a.client.callAPI(r)
}

/*
SSHKeysApiService Delete a SSH key.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param region Region code
 * @param id The ID of the SSH key
 * @param forceDelete Force delete

@return interface{}
*/
func (a *SSHKeysApiService) Delete(ctx context.Context, region string, name string) (interface{}, *http.Response, error) {
	httpMethod := strings.ToUpper("Delete")

	// create path and map variables
	path := a.client.cfg.BasePath + "/regions/{region}/ssh-keys/{name}"
	path = strings.Replace(path, "{"+"region"+"}", fmt.Sprintf("%v", region), -1)
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	r, err := a.client.prepareRequest(ctx, path, httpMethod, nil)
	if err != nil {
		return nil, nil, err
	}

	return a.client.callAPI(r)
}
