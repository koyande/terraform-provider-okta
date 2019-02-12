package okta

// Not all APIs are supported by okta-sdk-golang, this is one

import (
	"fmt"

	"github.com/okta/okta-sdk-golang/okta"
	"github.com/okta/okta-sdk-golang/okta/query"
)

type AuthorizationServer struct {
	Audiences   []string               `json:"audiences,omitempty"`
	Credentials *AuthServerCredentials `json:"credentials,omitempty"`
	Description string                 `json:"descriptions,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Id          string                 `json:"id,omitempty"`
}

type AuthServerCredentials struct {
	Signing *okta.ApplicationCredentialsSigning `json:"signing,omitempty"`
}

func (m *ApiSupplement) DeleteAuthorizationServer(id string) (*okta.Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%s", id)
	req, err := m.requestExecutor.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	return m.requestExecutor.Do(req, nil)
}
func (m *ApiSupplement) ListAuthorizationServers() ([]*AuthorizationServer, *okta.Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers")
	req, err := m.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var auth []*AuthorizationServer
	resp, err := m.requestExecutor.Do(req, &auth)
	return auth, resp, err
}
func (m *ApiSupplement) CreateAuthorizationServer(body AuthorizationServer, qp *query.Params) (*AuthorizationServer, *okta.Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers")
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.requestExecutor.NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	authorizationServer := body
	resp, err := m.requestExecutor.Do(req, &authorizationServer)
	return &authorizationServer, resp, err
}

func (m *ApiSupplement) UpdateAuthorizationServer(id string, body AuthorizationServer, qp *query.Params) (*AuthorizationServer, *okta.Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%s", id)
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.requestExecutor.NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	authorizationServer := body
	resp, err := m.requestExecutor.Do(req, &authorizationServer)
	return &authorizationServer, resp, err
}

func (m *ApiSupplement) GetAuthorizationServer(id string, authorizationServerInstance AuthorizationServer) (*AuthorizationServer, *okta.Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%s", id)
	req, err := m.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	authorizationServer := authorizationServerInstance
	resp, err := m.requestExecutor.Do(req, &authorizationServer)
	return &authorizationServer, resp, err
}
func (m *ApiSupplement) ActivateAuthorizationServer(id string) (*okta.Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%s/lifecycle/activate", id)
	req, err := m.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	return m.requestExecutor.Do(req, nil)
}
func (m *ApiSupplement) DeactivateAuthorizationServer(id string) (*okta.Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%s/lifecycle/deactivate", id)
	req, err := m.requestExecutor.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	return m.requestExecutor.Do(req, nil)
}
