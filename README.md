<div align="center">
    <picture>
        <source srcset="https://user-images.githubusercontent.com/6267663/221572723-e77f55a3-5d19-4a13-94f8-e7b0b340d71e.svg" media="(prefers-color-scheme: dark)">
        <img src="https://user-images.githubusercontent.com/6267663/221572726-6982541c-d1cf-4d9f-9bbf-cd774a2713e6.svg">
    </picture>
   <h1>Go SDK</h1>
   <p><strong>Open Source Ledger for money-moving platforms</strong></p>
   <p>Build and track custom fit money flows on a scalable financial infrastructure.</p>
   <a href="https://docs.formance.com/api/stack/v1.0#section/Introduction"><img src="https://img.shields.io/static/v1?label=Docs&message=Docs&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/formance-go-sdk/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/formance-go-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://join.slack.com/t/formance-community/shared_invite/zt-1of48xmgy-Jc6RH8gzcWf5D0qD2HBPQA"><img src="https://img.shields.io/static/v1?label=Slack&message=Join&color=7289da&style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/formance-go-sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := formancegosdk.New(
		formancegosdk.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.GetServerInfo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.ServerInfo != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Formance SDK](docs/sdks/formance/README.md)

* [GetServerInfo](docs/sdks/formance/README.md#getserverinfo) - Get server info
* [PaymentsgetServerInfo](docs/sdks/formance/README.md#paymentsgetserverinfo) - Get server info
* [SearchgetServerInfo](docs/sdks/formance/README.md#searchgetserverinfo) - Get server info

### [Clients](docs/sdks/clients/README.md)

* [AddScopeToClient](docs/sdks/clients/README.md#addscopetoclient) - Add scope to client
* [CreateClient](docs/sdks/clients/README.md#createclient) - Create client
* [CreateSecret](docs/sdks/clients/README.md#createsecret) - Add a secret to a client
* [DeleteClient](docs/sdks/clients/README.md#deleteclient) - Delete client
* [DeleteScopeFromClient](docs/sdks/clients/README.md#deletescopefromclient) - Delete scope from client
* [DeleteSecret](docs/sdks/clients/README.md#deletesecret) - Delete a secret from a client
* [ListClients](docs/sdks/clients/README.md#listclients) - List clients
* [ReadClient](docs/sdks/clients/README.md#readclient) - Read client
* [UpdateClient](docs/sdks/clients/README.md#updateclient) - Update client

### [Scopes](docs/sdks/scopes/README.md)

* [AddTransientScope](docs/sdks/scopes/README.md#addtransientscope) - Add a transient scope to a scope
* [CreateScope](docs/sdks/scopes/README.md#createscope) - Create scope
* [DeleteScope](docs/sdks/scopes/README.md#deletescope) - Delete scope
* [DeleteTransientScope](docs/sdks/scopes/README.md#deletetransientscope) - Delete a transient scope from a scope
* [ListScopes](docs/sdks/scopes/README.md#listscopes) - List scopes
* [ReadScope](docs/sdks/scopes/README.md#readscope) - Read scope
* [UpdateScope](docs/sdks/scopes/README.md#updatescope) - Update scope

### [Users](docs/sdks/users/README.md)

* [ListUsers](docs/sdks/users/README.md#listusers) - List users
* [ReadUser](docs/sdks/users/README.md#readuser) - Read user

### [Server](docs/sdks/server/README.md)

* [GetInfo](docs/sdks/server/README.md#getinfo) - Show server information

### [Ledger](docs/sdks/ledger/README.md)

* [GetLedgerInfo](docs/sdks/ledger/README.md#getledgerinfo) - Get information about a ledger

### [Accounts](docs/sdks/accounts/README.md)

* [AddMetadataToAccount](docs/sdks/accounts/README.md#addmetadatatoaccount) - Add metadata to an account
* [CountAccounts](docs/sdks/accounts/README.md#countaccounts) - Count the accounts from a ledger
* [GetAccount](docs/sdks/accounts/README.md#getaccount) - Get account by its address
* [ListAccounts](docs/sdks/accounts/README.md#listaccounts) - List accounts from a ledger

### [Balances](docs/sdks/balances/README.md)

* [GetBalances](docs/sdks/balances/README.md#getbalances) - Get the balances from a ledger's account
* [GetBalancesAggregated](docs/sdks/balances/README.md#getbalancesaggregated) - Get the aggregated balances from selected accounts

### [Logs](docs/sdks/logs/README.md)

* [ListLogs](docs/sdks/logs/README.md#listlogs) - List the logs from a ledger

### [Mapping](docs/sdks/mapping/README.md)

* [GetMapping](docs/sdks/mapping/README.md#getmapping) - Get the mapping of a ledger
* [UpdateMapping](docs/sdks/mapping/README.md#updatemapping) - Update the mapping of a ledger

### [Script](docs/sdks/script/README.md)

* [~~RunScript~~](docs/sdks/script/README.md#runscript) - Execute a Numscript :warning: **Deprecated**

### [Stats](docs/sdks/stats/README.md)

* [ReadStats](docs/sdks/stats/README.md#readstats) - Get statistics from a ledger

### [Transactions](docs/sdks/transactions/README.md)

* [CreateTransactions](docs/sdks/transactions/README.md#createtransactions) - Create a new batch of transactions to a ledger
* [AddMetadataOnTransaction](docs/sdks/transactions/README.md#addmetadataontransaction) - Set the metadata of a transaction by its ID
* [CountTransactions](docs/sdks/transactions/README.md#counttransactions) - Count the transactions from a ledger
* [CreateTransaction](docs/sdks/transactions/README.md#createtransaction) - Create a new transaction to a ledger
* [GetTransaction](docs/sdks/transactions/README.md#gettransaction) - Get transaction from a ledger by its ID
* [ListTransactions](docs/sdks/transactions/README.md#listtransactions) - List transactions from a ledger
* [RevertTransaction](docs/sdks/transactions/README.md#reverttransaction) - Revert a ledger transaction by its ID

### [Orchestration](docs/sdks/orchestration/README.md)

* [CreateWorkflow](docs/sdks/orchestration/README.md#createworkflow) - Create workflow
* [GetFlow](docs/sdks/orchestration/README.md#getflow) - Get a flow by id
* [GetWorkflowOccurrence](docs/sdks/orchestration/README.md#getworkflowoccurrence) - Get a workflow occurrence by id
* [ListFlows](docs/sdks/orchestration/README.md#listflows) - List registered flows
* [ListRuns](docs/sdks/orchestration/README.md#listruns) - List occurrences of a workflow
* [OrchestrationgetServerInfo](docs/sdks/orchestration/README.md#orchestrationgetserverinfo) - Get server info
* [RunWorkflow](docs/sdks/orchestration/README.md#runworkflow) - Run workflow

### [Payments](docs/sdks/payments/README.md)

* [ConnectorsStripeTransfer](docs/sdks/payments/README.md#connectorsstripetransfer) - Transfer funds between Stripe accounts
* [GetConnectorTask](docs/sdks/payments/README.md#getconnectortask) - Read a specific task of the connector
* [GetPayment](docs/sdks/payments/README.md#getpayment) - Get a payment
* [InstallConnector](docs/sdks/payments/README.md#installconnector) - Install a connector
* [ListAllConnectors](docs/sdks/payments/README.md#listallconnectors) - List all installed connectors
* [ListConfigsAvailableConnectors](docs/sdks/payments/README.md#listconfigsavailableconnectors) - List the configs of each available connector
* [ListConnectorTasks](docs/sdks/payments/README.md#listconnectortasks) - List tasks from a connector
* [ListPayments](docs/sdks/payments/README.md#listpayments) - List payments
* [PaymentslistAccounts](docs/sdks/payments/README.md#paymentslistaccounts) - List accounts
* [ReadConnectorConfig](docs/sdks/payments/README.md#readconnectorconfig) - Read the config of a connector
* [ResetConnector](docs/sdks/payments/README.md#resetconnector) - Reset a connector
* [UninstallConnector](docs/sdks/payments/README.md#uninstallconnector) - Uninstall a connector

### [Search](docs/sdks/search/README.md)

* [Search](docs/sdks/search/README.md#search) - Search

### [Wallets](docs/sdks/wallets/README.md)

* [ConfirmHold](docs/sdks/wallets/README.md#confirmhold) - Confirm a hold
* [CreateBalance](docs/sdks/wallets/README.md#createbalance) - Create a balance
* [CreateWallet](docs/sdks/wallets/README.md#createwallet) - Create a new wallet
* [CreditWallet](docs/sdks/wallets/README.md#creditwallet) - Credit a wallet
* [DebitWallet](docs/sdks/wallets/README.md#debitwallet) - Debit a wallet
* [GetBalance](docs/sdks/wallets/README.md#getbalance) - Get detailed balance
* [GetHold](docs/sdks/wallets/README.md#gethold) - Get a hold
* [GetHolds](docs/sdks/wallets/README.md#getholds) - Get all holds for a wallet
* [GetTransactions](docs/sdks/wallets/README.md#gettransactions)
* [GetWallet](docs/sdks/wallets/README.md#getwallet) - Get a wallet
* [ListBalances](docs/sdks/wallets/README.md#listbalances) - List balances of a wallet
* [ListWallets](docs/sdks/wallets/README.md#listwallets) - List all wallets
* [UpdateWallet](docs/sdks/wallets/README.md#updatewallet) - Update a wallet
* [VoidHold](docs/sdks/wallets/README.md#voidhold) - Cancel a hold
* [WalletsgetServerInfo](docs/sdks/wallets/README.md#walletsgetserverinfo) - Get server info

### [Webhooks](docs/sdks/webhooks/README.md)

* [ActivateConfig](docs/sdks/webhooks/README.md#activateconfig) - Activate one config
* [ChangeConfigSecret](docs/sdks/webhooks/README.md#changeconfigsecret) - Change the signing secret of a config
* [DeactivateConfig](docs/sdks/webhooks/README.md#deactivateconfig) - Deactivate one config
* [DeleteConfig](docs/sdks/webhooks/README.md#deleteconfig) - Delete one config
* [GetManyConfigs](docs/sdks/webhooks/README.md#getmanyconfigs) - Get many configs
* [InsertConfig](docs/sdks/webhooks/README.md#insertconfig) - Insert a new config
* [TestConfig](docs/sdks/webhooks/README.md#testconfig) - Test one config
<!-- End Available Resources and Operations [operations] -->





<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := formancegosdk.New(
		formancegosdk.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.GetServerInfo(ctx)
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `http://localhost` | None |
| 1 | `https://{organization}.sandbox.formance.cloud` | `organization` |

#### Example

```go
package main

import (
	"context"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := formancegosdk.New(
		formancegosdk.WithServerIndex(1),
		formancegosdk.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.GetServerInfo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.ServerInfo != nil {
		// handle response
	}
}

```

#### Variables

Some of the server options above contain variables. If you want to set the values of those variables, the following options are provided for doing so:
 * `WithOrganization string`

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := formancegosdk.New(
		formancegosdk.WithServerURL("http://localhost"),
		formancegosdk.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.GetServerInfo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.ServerInfo != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name            | Type            | Scheme          |
| --------------- | --------------- | --------------- |
| `Authorization` | oauth2          | OAuth2 token    |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	formancegosdk "github.com/speakeasy-sdks/formance-go-sdk"
	"log"
)

func main() {
	s := formancegosdk.New(
		formancegosdk.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.GetServerInfo(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.ServerInfo != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
