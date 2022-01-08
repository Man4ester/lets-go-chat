// Package Api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package Api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ActiveUsersResponse defines model for ActiveUsersResponse.
type ActiveUsersResponse struct {
	Count int `json:"count"`
}

// CreateUserRequest defines model for CreateUserRequest.
type CreateUserRequest struct {
	Password string `json:"password"`
	UserName string `json:"userName"`
}

// CreateUserResponse defines model for CreateUserResponse.
type CreateUserResponse struct {
	Id       *string `json:"id,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

// LoginUserRequest defines model for LoginUserRequest.
type LoginUserRequest struct {
	// The password for login in clear text
	Password string `json:"password"`

	// The user name for login
	UserName string `json:"userName"`
}

// LoginUserResonse defines model for LoginUserResonse.
type LoginUserResonse struct {
	// A url for websoket API with a one-time token for starting chat
	Url string `json:"url"`
}

// WsRTMStartParams defines parameters for WsRTMStart.
type WsRTMStartParams struct {
	// One time token for a loged user
	Token string `json:"token"`
}

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody CreateUserRequest

// LoginUserJSONBody defines parameters for LoginUser.
type LoginUserJSONBody LoginUserRequest

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody CreateUserJSONBody

// LoginUserJSONRequestBody defines body for LoginUser for application/json ContentType.
type LoginUserJSONRequestBody LoginUserJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// WsRTMStart request
	WsRTMStart(ctx context.Context, params *WsRTMStartParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateUser request with any body
	CreateUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateUser(ctx context.Context, body CreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetActiveUsers request
	GetActiveUsers(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// LoginUser request with any body
	LoginUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	LoginUser(ctx context.Context, body LoginUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) WsRTMStart(ctx context.Context, params *WsRTMStartParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewWsRTMStartRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateUserRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateUser(ctx context.Context, body CreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateUserRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetActiveUsers(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetActiveUsersRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) LoginUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewLoginUserRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) LoginUser(ctx context.Context, body LoginUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewLoginUserRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewWsRTMStartRequest generates requests for WsRTMStart
func NewWsRTMStartRequest(server string, params *WsRTMStartParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/chat/ws.rtm.start")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "token", runtime.ParamLocationQuery, params.Token); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateUserRequest calls the generic CreateUser builder with application/json body
func NewCreateUserRequest(server string, body CreateUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateUserRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateUserRequestWithBody generates requests for CreateUser with any type of body
func NewCreateUserRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetActiveUsersRequest generates requests for GetActiveUsers
func NewGetActiveUsersRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user/active")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewLoginUserRequest calls the generic LoginUser builder with application/json body
func NewLoginUserRequest(server string, body LoginUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewLoginUserRequestWithBody(server, "application/json", bodyReader)
}

// NewLoginUserRequestWithBody generates requests for LoginUser with any type of body
func NewLoginUserRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/user/login")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// WsRTMStart request
	WsRTMStartWithResponse(ctx context.Context, params *WsRTMStartParams, reqEditors ...RequestEditorFn) (*WsRTMStartResponse, error)

	// CreateUser request with any body
	CreateUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateUserResponse, error)

	CreateUserWithResponse(ctx context.Context, body CreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateUserResponse, error)

	// GetActiveUsers request
	GetActiveUsersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetActiveUsersResponse, error)

	// LoginUser request with any body
	LoginUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*LoginUserResponse, error)

	LoginUserWithResponse(ctx context.Context, body LoginUserJSONRequestBody, reqEditors ...RequestEditorFn) (*LoginUserResponse, error)
}

type WsRTMStartResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r WsRTMStartResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r WsRTMStartResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateUserResponseV1 struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CreateUserResponseV1
}

// Status returns HTTPResponse.Status
func (r CreateUserResponseV1) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateUserResponseV1) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetActiveUsersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ActiveUsersResponse
}

// Status returns HTTPResponse.Status
func (r GetActiveUsersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetActiveUsersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type LoginUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *LoginUserResonse
}

// Status returns HTTPResponse.Status
func (r LoginUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r LoginUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// WsRTMStartWithResponse request returning *WsRTMStartResponse
func (c *ClientWithResponses) WsRTMStartWithResponse(ctx context.Context, params *WsRTMStartParams, reqEditors ...RequestEditorFn) (*WsRTMStartResponse, error) {
	rsp, err := c.WsRTMStart(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseWsRTMStartResponse(rsp)
}

// CreateUserWithBodyWithResponse request with arbitrary body returning *CreateUserResponse
func (c *ClientWithResponses) CreateUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateUserResponseV1, error) {
	rsp, err := c.CreateUserWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateUserResponse(rsp)
}

func (c *ClientWithResponses) CreateUserWithResponse(ctx context.Context, body CreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateUserResponseV1, error) {
	rsp, err := c.CreateUser(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateUserResponse(rsp)
}

// GetActiveUsersWithResponse request returning *GetActiveUsersResponse
func (c *ClientWithResponses) GetActiveUsersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetActiveUsersResponse, error) {
	rsp, err := c.GetActiveUsers(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetActiveUsersResponse(rsp)
}

// LoginUserWithBodyWithResponse request with arbitrary body returning *LoginUserResponse
func (c *ClientWithResponses) LoginUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*LoginUserResponse, error) {
	rsp, err := c.LoginUserWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseLoginUserResponse(rsp)
}

func (c *ClientWithResponses) LoginUserWithResponse(ctx context.Context, body LoginUserJSONRequestBody, reqEditors ...RequestEditorFn) (*LoginUserResponse, error) {
	rsp, err := c.LoginUser(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseLoginUserResponse(rsp)
}

// ParseWsRTMStartResponse parses an HTTP response from a WsRTMStartWithResponse call
func ParseWsRTMStartResponse(rsp *http.Response) (*WsRTMStartResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &WsRTMStartResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseCreateUserResponse parses an HTTP response from a CreateUserWithResponse call
func ParseCreateUserResponse(rsp *http.Response) (*CreateUserResponseV1, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateUserResponseV1{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CreateUserResponseV1
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetActiveUsersResponse parses an HTTP response from a GetActiveUsersWithResponse call
func ParseGetActiveUsersResponse(rsp *http.Response) (*GetActiveUsersResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetActiveUsersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ActiveUsersResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseLoginUserResponse parses an HTTP response from a LoginUserWithResponse call
func ParseLoginUserResponse(rsp *http.Response) (*LoginUserResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &LoginUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest LoginUserResonse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Endpoint to start real time chat
	// (GET /chat/ws.rtm.start)
	WsRTMStart(ctx echo.Context, params WsRTMStartParams) error
	// Register (create) user
	// (POST /user)
	CreateUser(ctx echo.Context) error
	// Number of active users in a chat
	// (GET /user/active)
	GetActiveUsers(ctx echo.Context) error
	// Logs user into the system
	// (POST /user/login)
	LoginUser(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// WsRTMStart converts echo context to params.
func (w *ServerInterfaceWrapper) WsRTMStart(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params WsRTMStartParams
	// ------------- Required query parameter "token" -------------

	err = runtime.BindQueryParameter("form", true, true, "token", ctx.QueryParams(), &params.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter token: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.WsRTMStart(ctx, params)
	return err
}

// CreateUser converts echo context to params.
func (w *ServerInterfaceWrapper) CreateUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateUser(ctx)
	return err
}

// GetActiveUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetActiveUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetActiveUsers(ctx)
	return err
}

// LoginUser converts echo context to params.
func (w *ServerInterfaceWrapper) LoginUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.LoginUser(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/chat/ws.rtm.start", wrapper.WsRTMStart)
	router.POST(baseURL+"/user", wrapper.CreateUser)
	router.GET(baseURL+"/user/active", wrapper.GetActiveUsers)
	router.POST(baseURL+"/user/login", wrapper.LoginUser)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xW32/bNhD+VwhuDxtgSU7aDYOAPaRFVmTIuiFNsAJdHmjpLDGRSOXuZMUI/L8PR8l2",
	"ainLhrR7smgej999H+/Hg8583XgHjkmnD5qyEmoTPk8ytiu4IkC6AGq8I5C/G/QNIFsIRplvHcsHrxvQ",
	"qbaOoQDUm81MI9y1FiHX6afB7nq2tfOLG8hYb2b6LYLhcM0F3LVAPL6kMUSdx1y+a+vOwRVc6vSnnTdi",
	"tK4Qby0Bvjc1HJi+Hpke4Ntd8cjFc2if4sQGoEuPtWGd6ra14vYfoY7BjW4+94V1/5qmHChD27D1Tqf6",
	"sgS13VVLj6oSZ8o6lVVgUDHc83MQxx5lVzlTw96l/lI8P4qWpllusRrDOlEtVgFOBwvyt8Dq5I8z1Vku",
	"lVHeQcS2BsX+FlwwIzbI1hUqK41QAPembirB0lGaJEvjsnUke7H1SUd/tfP58Y/h+M9bb1FYPhu54B1H",
	"KlbWLf04lF9bYmUUWcET4CkCXNkM9ExXNoOBFRf00SeNyUpQx/FcqBVudMncpEnSdV1swm7ssUiGo5Sc",
	"n709ff/hNDqO53HJdSW0s+UQ/C8St3rnK7OnZgVIPbSjeB7Pxdw34ExjdapfxfP4SM90Y7gM8iRyKuko",
	"Rq7jwLL8W0D4ER2NxHmW61T/SReXv30IJuIATQ0MSDr9dMjJ7w7UgYBGHh7k4TFq4VKn+q4FXOvZlpqt",
	"Pns1GFuYDbVuKv2uxbhP7xDM0fxoLNBVU6DJBUv/2DJ5bQ169pkPZL6ez8enztzKVDbvIxCrH6atGNCZ",
	"Sn0AXAGqU0TfF1Vq69rgWqf61OWNt44FQCBYIZiq52eQjE1BofrK8lqOJ4EnSSZPE1Lsy9vAFxC/8fm6",
	"L/WOoS/2pmkqm4VjyQ0J4odHbH6LsNSp/ibZt5Zk6CvJuNqHFPg8+t6o11QNmXIo3+ZAouOexi8Ocijy",
	"EygDuqyH+qTcb0yuBh5nCuqG1yGqUDU9Kpu/6AlcQGGJAdV3PY7vt2mwFT4sr2f6Psp8DgW4aAATLXy+",
	"joYEkW+9ex2JCY3/yXR9B/xoNNBfUYepCWRCCGqzDIiWbaV2SGcKgVt0pFxbL+QZLVUfV6CIXkT7+0mX",
	"0k7NYeYNehwkYNI3yyfTcNf9vlIWjmaJCVavhuedg8utqeh/zcBR+//PslfW3UplvPEy5fSilGDy0Fge",
	"9Mfo9L6xCBSdLLmviJ97zw2DKHp1+VZ1Jbih5UB/Sj/uHrtJT86EkWByGPgYXcj+ua0tj+/LTFWRagBV",
	"6VtUpqp8B7larBUPo9b0ndbxq+P9ffv5e/NcB9rWoWQ3k70kJ859QX3Bto59AE1rYqhH1Sggo+Cp7/H9",
	"tJKsjrQ03sF41Pq3EpMyC9/ylpKhhIXV5nrzdwAAAP//DnAkQNAMAAA=",
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

