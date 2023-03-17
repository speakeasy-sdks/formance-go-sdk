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

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/formance-go-sdk"
    "github.com/speakeasy-sdks/formance-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/formance-go-sdk/pkg/models/operations"
)

func main() {
    s := formance.New(
        WithSecurity(        shared.Security{
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
## SDK Available Operations

### Formance SDK

* `GetServerInfo` - Get server info
* `PaymentsgetServerInfo` - Get server info
* `SearchgetServerInfo` - Get server info

### Accounts

* `AddMetadataToAccount` - Add metadata to an account
* `CountAccounts` - Count the accounts from a ledger
* `GetAccount` - Get account by its address
* `ListAccounts` - List accounts from a ledger

### Balances

* `GetBalances` - Get the balances from a ledger's account
* `GetBalancesAggregated` - Get the aggregated balances from selected accounts

### Clients

* `AddScopeToClient` - Add scope to client
* `CreateClient` - Create client
* `CreateSecret` - Add a secret to a client
* `DeleteClient` - Delete client
* `DeleteScopeFromClient` - Delete scope from client
* `DeleteSecret` - Delete a secret from a client
* `ListClients` - List clients
* `ReadClient` - Read client
* `UpdateClient` - Update client

### Ledger

* `GetLedgerInfo` - Get information about a ledger

### Logs

* `ListLogs` - List the logs from a ledger

### Mapping

* `GetMapping` - Get the mapping of a ledger
* `UpdateMapping` - Update the mapping of a ledger

### Orchestration

* `CreateWorkflow` - Create workflow
* `GetFlow` - Get a flow by id
* `GetWorkflowOccurrence` - Get a workflow occurrence by id
* `ListFlows` - List registered flows
* `ListRuns` - List occurrences of a workflow
* `OrchestrationgetServerInfo` - Get server info
* `RunWorkflow` - Run workflow

### Payments

* `ConnectorsStripeTransfer` - Transfer funds between Stripe accounts
* `GetConnectorTask` - Read a specific task of the connector
* `GetPayment` - Get a payment
* `InstallConnector` - Install a connector
* `ListAllConnectors` - List all installed connectors
* `ListConfigsAvailableConnectors` - List the configs of each available connector
* `ListConnectorTasks` - List tasks from a connector
* `ListPayments` - List payments
* `PaymentslistAccounts` - List accounts
* `ReadConnectorConfig` - Read the config of a connector
* `ResetConnector` - Reset a connector
* `UninstallConnector` - Uninstall a connector

### Scopes

* `AddTransientScope` - Add a transient scope to a scope
* `CreateScope` - Create scope
* `DeleteScope` - Delete scope
* `DeleteTransientScope` - Delete a transient scope from a scope
* `ListScopes` - List scopes
* `ReadScope` - Read scope
* `UpdateScope` - Update scope

### Script

* `RunScript` - Execute a Numscript

### Search

* `Search` - Search

### Server

* `GetInfo` - Show server information

### Stats

* `ReadStats` - Get statistics from a ledger

### Transactions

* `CreateTransactions` - Create a new batch of transactions to a ledger
* `AddMetadataOnTransaction` - Set the metadata of a transaction by its ID
* `CountTransactions` - Count the transactions from a ledger
* `CreateTransaction` - Create a new transaction to a ledger
* `GetTransaction` - Get transaction from a ledger by its ID
* `ListTransactions` - List transactions from a ledger
* `RevertTransaction` - Revert a ledger transaction by its ID

### Users

* `ListUsers` - List users
* `ReadUser` - Read user

### Wallets

* `ConfirmHold` - Confirm a hold
* `CreateBalance` - Create a balance
* `CreateWallet` - Create a new wallet
* `CreditWallet` - Credit a wallet
* `DebitWallet` - Debit a wallet
* `GetBalance` - Get detailed balance
* `GetHold` - Get a hold
* `GetHolds` - Get all holds for a wallet
* `GetTransactions`
* `GetWallet` - Get a wallet
* `ListBalances` - List balances of a wallet
* `ListWallets` - List all wallets
* `UpdateWallet` - Update a wallet
* `VoidHold` - Cancel a hold
* `WalletsgetServerInfo` - Get server info

### Webhooks

* `ActivateConfig` - Activate one config
* `ChangeConfigSecret` - Change the signing secret of a config
* `DeactivateConfig` - Deactivate one config
* `DeleteConfig` - Delete one config
* `GetManyConfigs` - Get many configs
* `InsertConfig` - Insert a new config
* `TestConfig` - Test one config
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
