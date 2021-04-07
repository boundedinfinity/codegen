package model

import (
	"path"

	"github.com/boundedinfinity/optional"
)

type DataModelType string

const (
	DataModelType_Null    DataModelType = "null"
	DataModelType_Boolean DataModelType = "boolean"
	DataModelType_Object  DataModelType = "object"
	DataModelType_Array   DataModelType = "array"
	DataModelType_Number  DataModelType = "number"
	DataModelType_String  DataModelType = "string"
)

type OpenApiV310 struct {
	Openapi    optional.StringOptional        `json:"openapi,omitempty" yaml:"openapi,omitempty"`
	Info       *OpenApiV310Info               `json:"info,omitempty" yaml:"info,omitempty"`
	Servers    []OpenApiV310Server            `json:"servers,omitempty" yaml:"servers,omitempty"`
	Components *OpenApiV310Components         `json:"components,omitempty" yaml:"components,omitempty"`
	Paths      map[string]OpenApiV310PathItem `json:"paths,omitempty" yaml:"paths,omitempty"`
	X_Bi_Go    *OpenApiV310Extention          `json:"x-bi-go,omitempty" yaml:"x-bi-go,omitempty" xml:"x-bi-go,omitempty"`
}

type OpenApiV310Info struct {
	Title          *string             `json:"title,omitempty" yaml:"title,omitempty"`
	Summary        *string             `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description    *string             `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService *string             `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Contact        *OpenApiV310Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
	License        *OpenApiV310License `json:"license,omitempty" yaml:"license,omitempty"`
	Version        *string             `json:"version,omitempty" yaml:"version,omitempty"`
}

type OpenApiV310Contact struct {
	Name  *string `json:"name,omitempty" yaml:"name,omitempty"`
	Url   *string `json:"url,omitempty" yaml:"url,omitempty"`
	Email *string `json:"email,omitempty" yaml:"email,omitempty"`
}

type OpenApiV310License struct {
	Name       *string `json:"name,omitempty" yaml:"name,omitempty"`
	Identifier *string `json:"identifier,omitempty" yaml:"identifier,omitempty"`
	Url        *string `json:"url,omitempty" yaml:"url,omitempty"`
}

type OpenApiV310Server struct {
	Url         *string                              `json:"url,omitempty" yaml:"url,omitempty"`
	Description *string                              `json:"description,omitempty" yaml:"description,omitempty"`
	Variables   map[string]OpenApiV310ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
}

type OpenApiV310ServerVariable struct {
	Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`
	Description *string  `json:"description,omitempty" yaml:"description,omitempty"`
	Default     *string  `json:"default,omitempty" yaml:"default,omitempty"`
}

type OpenApiV310PathItem struct {
	Ref         string                 `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Description string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Summary     string                 `json:"summary,omitempty" yaml:"summary,omitempty"`
	Get         *OpenApiV310Operation  `json:"get,omitempty" yaml:"get,omitempty"`
	Put         *OpenApiV310Operation  `json:"put,omitempty" yaml:"put,omitempty"`
	Post        *OpenApiV310Operation  `json:"post,omitempty" yaml:"post,omitempty"`
	Delete      *OpenApiV310Operation  `json:"delete,omitempty" yaml:"delete,omitempty"`
	Options     *OpenApiV310Operation  `json:"options,omitempty" yaml:"options,omitempty"`
	Head        *OpenApiV310Operation  `json:"head,omitempty" yaml:"head,omitempty"`
	Patch       *OpenApiV310Operation  `json:"patch,omitempty" yaml:"patch,omitempty"`
	Trace       *OpenApiV310Operation  `json:"trace,omitempty" yaml:"trace,omitempty"`
	Servers     []OpenApiV310Server    `json:"servers,omitempty" yaml:"servers,omitempty"`
	Parameters  []OpenApiV310Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

type OpenApiV310Operation struct {
	Tags         []string                          `json:"tags,omitempty" yaml:"tags,omitempty"`
	Description  *string                           `json:"description,omitempty" yaml:"description,omitempty"`
	Summary      *string                           `json:"summary,omitempty" yaml:"summary,omitempty"`
	ExternalDocs *OpenApiV310ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	OperationId  *string                           `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Parameters   []OpenApiV310Parameter            `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  *OpenApiV310RequestBody           `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses    map[string]OpenApiV310Response    `json:"responses,omitempty" yaml:"responses,omitempty"`
	Callbacks    *OpenApiV310Callback              `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	Deprecated   *bool                             `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security     *OpenApiV310SecurityRequirement   `json:"security,omitempty" yaml:"security,omitempty"`
	Servers      []OpenApiV310Server               `json:"servers,omitempty" yaml:"servers,omitempty"`
}

type OpenApiV310Parameter struct {
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	Summary     *string `json:"summary,omitempty" yaml:"summary,omitempty"`
	OperationId *string `json:"operationId,omitempty" yaml:"operationId,omitempty"`
}

type OpenApiV310Reference struct {
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	Summary     *string `json:"summary,omitempty" yaml:"summary,omitempty"`
	Ref         *string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

type OpenApiV310RequestBody struct {
	Description *string                         `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]OpenApiV310MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Required    *bool                           `json:"required,omitempty" yaml:"required,omitempty"`
}

type OpenApiV310MediaType struct {
	Schema   *JsonSchema_Draft07            `json:"schema,omitempty" yaml:"schema,omitempty"`
	Encoding map[string]OpenApiV310Encoding `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

type OpenApiV310Encoding struct {
	ContentType   *string                      `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers       map[string]OpenApiV310Header `json:"headers,omitempty" yaml:"headers,omitempty"`
	Style         *string                      `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       *bool                        `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved *bool                        `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
}

type OpenApiV310Header struct {
	Name        *string `json:"name,omitempty" yaml:"name,omitempty"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
}

type OpenApiV310SecurityRequirement struct {
}

type OpenApiV310Callback struct {
}

type OpenApiV310ExternalDocumentation struct {
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	Url         *string `json:"url,omitempty" yaml:"url,omitempty"`
}

type OpenApiV310Response struct {
	Description *string                         `json:"description,omitempty" yaml:"description,omitempty"`
	Headers     map[string]OpenApiV310Header    `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     map[string]OpenApiV310MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]OpenApiV310Link      `json:"links,omitempty" yaml:"links,omitempty"`
}

type OpenApiV310Link struct {
	OperationRef *string `json:"operationRef,omitempty" yaml:"operationRef,omitempty"`
	OperationId  *string `json:"operationId,omitempty" yaml:"operationId,omitempty"`
}

type OpenApiV310Components struct {
	Schemas         map[string]JsonSchema_Draft07        `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	X_Bi_Go_Schemas *OpenApiV310Extention                `json:"x-bi-go-schemas,omitempty" yaml:"x-bi-go-schemas,omitempty" xml:"x-bi-go-schemas,omitempty"`
	Responses       map[string]OpenApiV310Response       `json:"responses,omitempty" yaml:"responses,omitempty"`
	Parameters      map[string]OpenApiV310Parameter      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Examples        map[string]OpenApiV310Example        `json:"examples,omitempty" yaml:"examples,omitempty"`
	RequestBodies   map[string]OpenApiV310RequestBody    `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Headers         map[string]OpenApiV310Header         `json:"headers,omitempty" yaml:"headers,omitempty"`
	SecuritySchemas map[string]OpenApiV310SecurityScheme `json:"securitySchemas,omitempty" yaml:"securitySchemas,omitempty"`
	Links           map[string]OpenApiV310Link           `json:"links,omitempty" yaml:"links,omitempty"`
	Callbacks       map[string]OpenApiV310Callback       `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	PathItems       map[string]OpenApiV310PathItem       `json:"pathItems,omitempty" yaml:"pathItems,omitempty"`
}

func (t OpenApiV310Components) LookupSchema(p string) (string, JsonSchema_Draft07, bool, error) {
	n := path.Base(p)

	if t.Schemas == nil {
		return "", JsonSchema_Draft07{}, false, nil
	}

	if v, ok := t.Schemas[n]; ok {
		return n, v, true, nil
	}

	return "", JsonSchema_Draft07{}, false, nil
}

func (t OpenApiV310Components) LookupRequestBody(p string) (string, OpenApiV310RequestBody, bool, error) {
	n := path.Base(p)

	if t.RequestBodies == nil {
		return "", OpenApiV310RequestBody{}, false, nil
	}

	if v, ok := t.RequestBodies[n]; ok {
		return n, v, true, nil
	}

	return "", OpenApiV310RequestBody{}, false, nil
}

func (t OpenApiV310Components) LookupRequestResponse(p string) (string, OpenApiV310Response, bool, error) {
	n := path.Base(p)

	if t.Responses == nil {
		return "", OpenApiV310Response{}, false, nil
	}

	if v, ok := t.Responses[n]; ok {
		return n, v, true, nil
	}

	return "", OpenApiV310Response{}, false, nil
}

func (t OpenApiV310Components) LookupPathItem(p string) (string, OpenApiV310PathItem, bool, error) {
	n := path.Base(p)

	if t.PathItems == nil {
		return "", OpenApiV310PathItem{}, false, nil
	}

	if v, ok := t.PathItems[n]; ok {
		return n, v, true, nil
	}

	return "", OpenApiV310PathItem{}, false, nil
}

type OpenApiV310Example struct {
	Summary       *string `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description   *string `json:"description,omitempty" yaml:"description,omitempty"`
	Value         *string `json:"value,omitempty" yaml:"value,omitempty"`
	ExternalValue *string `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`
}

type OpenApiV310SecuritySchemeType string

const (
	DataModelType_ApiKey        OpenApiV310SecuritySchemeType = "apiKey"
	DataModelType_Http          OpenApiV310SecuritySchemeType = "http"
	DataModelType_MutualTLS     OpenApiV310SecuritySchemeType = "mutualTLS"
	DataModelType_Oauth2        OpenApiV310SecuritySchemeType = "oauth2"
	DataModelType_OpenIdConnect OpenApiV310SecuritySchemeType = "openIdConnect"
)

type OpenApiV310SecurityScheme struct {
	Type             *OpenApiV310SecuritySchemeType `json:"type,omitempty" yaml:"type,omitempty"`
	Description      *string                        `json:"description,omitempty" yaml:"description,omitempty"`
	Name             *string                        `json:"name,omitempty" yaml:"name,omitempty"`
	In               *string                        `json:"in,omitempty" yaml:"in,omitempty"`
	Scheme           *string                        `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	BearerFormat     *string                        `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
	Flows            *OpenApiV310OauthFlows         `json:"flows,omitempty" yaml:"flows,omitempty"`
	OpenIdConnectUrl *string                        `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"`
}

type OpenApiV310OauthFlows struct {
	Implicit          *OpenApiV310OauthFlow `json:"implicit,omitempty" yaml:"implicit,omitempty"`
	Password          *OpenApiV310OauthFlow `json:"password,omitempty" yaml:"password,omitempty"`
	ClientCredentials *OpenApiV310OauthFlow `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
	AuthorizationCode *OpenApiV310OauthFlow `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
}

type OpenApiV310OauthFlow struct {
	AuthorizationUrl *string           `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`
	TokenUrl         *string           `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
	RefreshUrl       *string           `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes,omitempty" yaml:"scopes,omitempty"`
}
