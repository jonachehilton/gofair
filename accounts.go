package gofair

// Accounts API Operations
const (
	getAccountFunds = "getAccountFunds/"
)

// Account object
type Account struct {
	Client *Client
}

func (a *Account) GetAccountFunds() (AccountFundsResponse, error) {
	// create url
	url := createURL(Endpoints.Account, getAccountFunds)

	// build request
	params := struct {
		Wallet string
	}{
		Wallet: "UK",
	}

	var response AccountFundsResponse

	// make request
	err := a.Client.request(url, params, &response)
	if err != nil {
		return response, err
	}
	return response, err
}
