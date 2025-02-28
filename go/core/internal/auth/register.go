package auth

import (
	ory "github.com/ory/kratos-client-go"
)

func (c *Client) Register(username string, password string, tinkId string) (Session, error) {
	flow, _, err := c.OryClient.V0alpha2Api.InitializeSelfServiceRegistrationFlowWithoutBrowser(c.ctx).Execute()
	if err != nil {
		return Session{}, err
	}

	result, _, err := c.OryClient.V0alpha2Api.SubmitSelfServiceRegistrationFlow(c.ctx).Flow(flow.Id).SubmitSelfServiceRegistrationFlowBody(
		ory.SubmitSelfServiceRegistrationFlowWithPasswordMethodBodyAsSubmitSelfServiceRegistrationFlowBody(&ory.SubmitSelfServiceRegistrationFlowWithPasswordMethodBody{
			Method:   "password",
			Password: password,
			Traits:   map[string]interface{}{"username": username, "tink_id": tinkId},
		}),
	).Execute()
	if err != nil {
		return Session{}, err
	}

	return Session{
		Session: *result.Session,
		Token:   *result.SessionToken,
	}, nil
}
