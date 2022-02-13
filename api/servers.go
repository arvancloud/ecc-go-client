package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type ServersApiService service

/*
ServersApiService Create new Server.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param region Region code
 * @param body interface{}

@return interface{}
*/
func (a *ServersApiService) Create(ctx context.Context, region string, body interface{}) (interface{}, *http.Response, error) {
	httpMethod := strings.ToUpper("Post")

	// create path and map variables
	path := a.client.cfg.BasePath + "/regions/{region}/servers"
	path = strings.Replace(path, "{"+"region"+"}", fmt.Sprintf("%v", region), -1)

	postBody := &body

	r, err := a.client.prepareRequest(ctx, path, httpMethod, postBody)
	if err != nil {
		return nil, nil, err
	}

	return a.client.callAPI(r)
}

/*
ServersApiService Shows details of a Server.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param region Region code
 * @param id The ID of the server

@return interface{}
*/
func (a *ServersApiService) Get(ctx context.Context, region string, id string) (interface{}, *http.Response, error) {
	httpMethod := strings.ToUpper("Get")

	path := a.client.cfg.BasePath + "/regions/{region}/servers/{id}"
	path = strings.Replace(path, "{"+"region"+"}", fmt.Sprintf("%v", region), -1)
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	r, err := a.client.prepareRequest(ctx, path, httpMethod, nil)
	if err != nil {
		return nil, nil, err
	}

	return a.client.callAPI(r)
}

/*
ServersApiService Return all Servers.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param region Region code

@return interface{}
*/
func (a *ServersApiService) All(ctx context.Context, region string) (interface{}, *http.Response, error) {
	httpMethod := strings.ToUpper("Get")

	// create path and map variables
	path := a.client.cfg.BasePath + "/regions/{region}/servers"
	path = strings.Replace(path, "{"+"region"+"}", fmt.Sprintf("%v", region), -1)

	r, err := a.client.prepareRequest(ctx, path, httpMethod, nil)
	if err != nil {
		return nil, nil, err
	}

	return a.client.callAPI(r)
}

/*
ServersApiService Delete a Server.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param region Region code
 * @param id The ID of the server
 * @param forceDelete Force delete

@return interface{}
*/
func (a *ServersApiService) Delete(ctx context.Context, region string, id string, forceDelete bool) (interface{}, *http.Response, error) {
	httpMethod := strings.ToUpper("Delete")

	// create path and map variables
	path := a.client.cfg.BasePath + "/regions/{region}/servers/{id}"
	path = strings.Replace(path, "{"+"region"+"}", fmt.Sprintf("%v", region), -1)
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	// TODO: pass forceDelete to query params
	// if forceDelete {
	// 	// queryParams.Add("forceDelete", strconv.FormatBool(forceDelete))
	// }

	r, err := a.client.prepareRequest(ctx, path, httpMethod, nil)
	if err != nil {
		return nil, nil, err
	}

	return a.client.callAPI(r)
}

/*
ServersApiService Return a region options.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param region Region code

@return interface{}
*/
func (a *ServersApiService) GetOptions(ctx context.Context, region string) (interface{}, *http.Response, error) {
	httpMethod := strings.ToUpper("Get")

	// create path and map variables
	path := a.client.cfg.BasePath + "/regions/{region}/servers/options"
	path = strings.Replace(path, "{"+"region"+"}", fmt.Sprintf("%v", region), -1)

	r, err := a.client.prepareRequest(ctx, path, httpMethod, nil)
	if err != nil {
		return nil, nil, err
	}

	return a.client.callAPI(r)
}
