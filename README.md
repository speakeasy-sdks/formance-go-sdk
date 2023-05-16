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

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/formance-go-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/formance-go-sdk"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
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
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations

### [Formance SDK](docs/formance/README.md)

* [GetServerInfo](docs/formance/README.md#getserverinfo) - Get server info
* [PaymentsgetServerInfo](docs/formance/README.md#paymentsgetserverinfo) - Get server info
* [SearchgetServerInfo](docs/formance/README.md#searchgetserverinfo) - Get server info

### [Accounts](docs/accounts/README.md)

* [AddMetadataToAccount](docs/accounts/README.md#addmetadatatoaccount) - Add metadata to an account
* [CountAccounts](docs/accounts/README.md#countaccounts) - Count the accounts from a ledger
* [GetAccount](docs/accounts/README.md#getaccount) - Get account by its address
* [ListAccounts](docs/accounts/README.md#listaccounts) - List accounts from a ledger

### [Balances](docs/balances/README.md)

* [GetBalances](docs/balances/README.md#getbalances) - Get the balances from a ledger's account
* [GetBalancesAggregated](docs/balances/README.md#getbalancesaggregated) - Get the aggregated balances from selected accounts

### [Clients](docs/clients/README.md)

* [AddScopeToClient](docs/clients/README.md#addscopetoclient) - Add scope to client
* [CreateClient](docs/clients/README.md#createclient) - Create client
* [CreateSecret](docs/clients/README.md#createsecret) - Add a secret to a client
* [DeleteClient](docs/clients/README.md#deleteclient) - Delete client
* [DeleteScopeFromClient](docs/clients/README.md#deletescopefromclient) - Delete scope from client
* [DeleteSecret](docs/clients/README.md#deletesecret) - Delete a secret from a client
* [ListClients](docs/clients/README.md#listclients) - List clients
* [ReadClient](docs/clients/README.md#readclient) - Read client
* [UpdateClient](docs/clients/README.md#updateclient) - Update client

### [Ledger](docs/ledger/README.md)

* [GetLedgerInfo](docs/ledger/README.md#getledgerinfo) - Get information about a ledger

### [Logs](docs/logs/README.md)

* [ListLogs](docs/logs/README.md#listlogs) - List the logs from a ledger

### [Mapping](docs/mapping/README.md)

* [GetMapping](docs/mapping/README.md#getmapping) - Get the mapping of a ledger
* [UpdateMapping](docs/mapping/README.md#updatemapping) - Update the mapping of a ledger

### [Orchestration](docs/orchestration/README.md)

* [CreateWorkflow](docs/orchestration/README.md#createworkflow) - Create workflow
* [GetFlow](docs/orchestration/README.md#getflow) - Get a flow by id
* [GetWorkflowOccurrence](docs/orchestration/README.md#getworkflowoccurrence) - Get a workflow occurrence by id
* [ListFlows](docs/orchestration/README.md#listflows) - List registered flows
* [ListRuns](docs/orchestration/README.md#listruns) - List occurrences of a workflow
* [OrchestrationgetServerInfo](docs/orchestration/README.md#orchestrationgetserverinfo) - Get server info
* [RunWorkflow](docs/orchestration/README.md#runworkflow) - Run workflow

### [Payments](docs/payments/README.md)

* [ConnectorsStripeTransfer](docs/payments/README.md#connectorsstripetransfer) - Transfer funds between Stripe accounts
* [GetConnectorTask](docs/payments/README.md#getconnectortask) - Read a specific task of the connector
* [GetPayment](docs/payments/README.md#getpayment) - Get a payment
* [InstallConnector](docs/payments/README.md#installconnector) - Install a connector
* [ListAllConnectors](docs/payments/README.md#listallconnectors) - List all installed connectors
* [ListConfigsAvailableConnectors](docs/payments/README.md#listconfigsavailableconnectors) - List the configs of each available connector
* [ListConnectorTasks](docs/payments/README.md#listconnectortasks) - List tasks from a connector
* [ListPayments](docs/payments/README.md#listpayments) - List payments
* [PaymentslistAccounts](docs/payments/README.md#paymentslistaccounts) - List accounts
* [ReadConnectorConfig](docs/payments/README.md#readconnectorconfig) - Read the config of a connector
* [ResetConnector](docs/payments/README.md#resetconnector) - Reset a connector
* [UninstallConnector](docs/payments/README.md#uninstallconnector) - Uninstall a connector

### [Scopes](docs/scopes/README.md)

* [AddTransientScope](docs/scopes/README.md#addtransientscope) - Add a transient scope to a scope
* [CreateScope](docs/scopes/README.md#createscope) - Create scope
* [DeleteScope](docs/scopes/README.md#deletescope) - Delete scope
* [DeleteTransientScope](docs/scopes/README.md#deletetransientscope) - Delete a transient scope from a scope
* [ListScopes](docs/scopes/README.md#listscopes) - List scopes
* [ReadScope](docs/scopes/README.md#readscope) - Read scope
* [UpdateScope](docs/scopes/README.md#updatescope) - Update scope

### [Script](docs/script/README.md)

* [~~RunScript~~](docs/script/README.md#runscript) - Execute a Numscript :warning: **Deprecated**

### [Search](docs/search/README.md)

* [Search](docs/search/README.md#search) - Search

### [Server](docs/server/README.md)

* [GetInfo](docs/server/README.md#getinfo) - Show server information

### [Stats](docs/stats/README.md)

* [ReadStats](docs/stats/README.md#readstats) - Get statistics from a ledger

### [Transactions](docs/transactions/README.md)

* [CreateTransactions](docs/transactions/README.md#createtransactions) - Create a new batch of transactions to a ledger
* [AddMetadataOnTransaction](docs/transactions/README.md#addmetadataontransaction) - Set the metadata of a transaction by its ID
* [CountTransactions](docs/transactions/README.md#counttransactions) - Count the transactions from a ledger
* [CreateTransaction](docs/transactions/README.md#createtransaction) - Create a new transaction to a ledger
* [GetTransaction](docs/transactions/README.md#gettransaction) - Get transaction from a ledger by its ID
* [ListTransactions](docs/transactions/README.md#listtransactions) - List transactions from a ledger
* [RevertTransaction](docs/transactions/README.md#reverttransaction) - Revert a ledger transaction by its ID

### [Users](docs/users/README.md)

* [ListUsers](docs/users/README.md#listusers) - List users
* [ReadUser](docs/users/README.md#readuser) - Read user

### [Wallets](docs/wallets/README.md)

* [ConfirmHold](docs/wallets/README.md#confirmhold) - Confirm a hold
* [CreateBalance](docs/wallets/README.md#createbalance) - Create a balance
* [CreateWallet](docs/wallets/README.md#createwallet) - Create a new wallet
* [CreditWallet](docs/wallets/README.md#creditwallet) - Credit a wallet
* [DebitWallet](docs/wallets/README.md#debitwallet) - Debit a wallet
* [GetBalance](docs/wallets/README.md#getbalance) - Get detailed balance
* [GetHold](docs/wallets/README.md#gethold) - Get a hold
* [GetHolds](docs/wallets/README.md#getholds) - Get all holds for a wallet
* [GetTransactions](docs/wallets/README.md#gettransactions)
* [GetWallet](docs/wallets/README.md#getwallet) - Get a wallet
* [ListBalances](docs/wallets/README.md#listbalances) - List balances of a wallet
* [ListWallets](docs/wallets/README.md#listwallets) - List all wallets
* [UpdateWallet](docs/wallets/README.md#updatewallet) - Update a wallet
* [VoidHold](docs/wallets/README.md#voidhold) - Cancel a hold
* [WalletsgetServerInfo](docs/wallets/README.md#walletsgetserverinfo) - Get server info

### [Webhooks](docs/webhooks/README.md)

* [ActivateConfig](docs/webhooks/README.md#activateconfig) - Activate one config
* [ChangeConfigSecret](docs/webhooks/README.md#changeconfigsecret) - Change the signing secret of a config
* [DeactivateConfig](docs/webhooks/README.md#deactivateconfig) - Deactivate one config
* [DeleteConfig](docs/webhooks/README.md#deleteconfig) - Delete one config
* [GetManyConfigs](docs/webhooks/README.md#getmanyconfigs) - Get many configs
* [InsertConfig](docs/webhooks/README.md#insertconfig) - Insert a new config
* [TestConfig](docs/webhooks/README.md#testconfig) - Test one config
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
