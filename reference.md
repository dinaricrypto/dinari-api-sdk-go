# Reference

## Market Data

<details><summary><code>client.MarketData.GetAlloys() -> *dinari.PaginatedAlloyResponseV1</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns available `Alloys` with cursor-based pagination.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetAlloysRequest{}
client.MarketData.GetAlloys(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**symbols:** `*string` — If set, this endpoint will return Alloys that match the symbols specified

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return

</dd>
</dl>

<dl>
<dd>

**order:** `*dinari.GetAlloysRequestOrder` — Sort order

</dd>
</dl>

<dl>
<dd>

**next:** `*string` — Cursor for next page

</dd>
</dl>

<dl>
<dd>

**previous:** `*string` — Cursor for previous page

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.MarketData.GetAlloyCurrentPrice(AlloyID) -> *dinari.AlloyPriceV1</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the current price for a specified `Alloy`.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetAlloyCurrentPriceRequest{
    AlloyID: "alloy_id",
}
client.MarketData.GetAlloyCurrentPrice(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**alloyID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.MarketData.GetAlloyPriceChart(AlloyID) -> []*dinari.AlloyHistoricalPriceDataPointV1</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get historical price data for a specified `Alloy`. Each index in the array represents a single tick in a price chart.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetAlloyPriceChartRequest{
    AlloyID: "alloy_id",
    Timespan: dinari.GetAlloyPriceChartRequestTimespanDay,
}
client.MarketData.GetAlloyPriceChart(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**alloyID:** `string`

</dd>
</dl>

<dl>
<dd>

**timespan:** `*dinari.GetAlloyPriceChartRequestTimespan` — The timespan of the historical prices to query.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.MarketData.GetStocks() -> *dinari.GetStocksResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of `Stocks`.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetStocksRequest{}
client.MarketData.GetStocks(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**symbols:** `*string` — List of `Stock` symbols to query. If not provided, all `Stocks` are returned.

</dd>
</dl>

<dl>
<dd>

**page:** `*int`

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int`

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return

</dd>
</dl>

<dl>
<dd>

**order:** `*dinari.GetStocksRequestOrder` — Sort order

</dd>
</dl>

<dl>
<dd>

**next:** `*string` — Cursor for next page

</dd>
</dl>

<dl>
<dd>

**previous:** `*string` — Cursor for previous page

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.MarketData.GetStockCurrentPrice(StockID) -> *dinari.StockPrice</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get current price for a specified `Stock`.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetStockCurrentPriceRequest{
    StockID: "stock_id",
}
client.MarketData.GetStockCurrentPrice(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**stockID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.MarketData.GetStockCurrentQuotesRouteV2(StockID) -> *dinari.StockQuoteResponseOneOfSchema</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get quote for a specified `Stock`.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetStockCurrentQuotesRouteV2Request{
    StockID: "stock_id",
}
client.MarketData.GetStockCurrentQuotesRouteV2(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**stockID:** `string`

</dd>
</dl>

<dl>
<dd>

**feed:** `*dinari.MarketDataFeed`

Requested data source for the quote. Only applies when using x-api-version: 2.
Allowed values:

- `null`: (default) Selects the highest quality available free data source.
- `sip`: Consolidated quote from all U.S. exchanges (NBBO). This is a paid data source and incurs usage-based billing.

</dd>
</dl>

<dl>
<dd>

**entityID:** `*string` — Entity id of customer requesting quote. Required to use sip feed. If this is an internal query, supply your organization entity_id

</dd>
</dl>

<dl>
<dd>

**apiVersion:** `*string` — Version for this method

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.MarketData.GetStockPriceChart(StockID) -> []*dinari.StocksHistoricalPriceDataPoint</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get historical price data for a specified `Stock`. Each index in the array represents a single tick in a price chart.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetStockPriceChartRequest{
    StockID: "stock_id",
    Timespan: dinari.StockTickTimespan1Day,
}
client.MarketData.GetStockPriceChart(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**stockID:** `string`

</dd>
</dl>

<dl>
<dd>

**timespan:** `*dinari.StockTickTimespan1` — The timespan of the historical prices to query.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.MarketData.GetStockNewsByStockID(StockID) -> []*dinari.StockNewsArticleInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the most recent news articles relating to a `Stock`, including a summary of the article
and a link to the original source.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetStockNewsByStockIDRequest{
    StockID: "stock_id",
}
client.MarketData.GetStockNewsByStockID(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**stockID:** `string`

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — The number of articles to return.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.MarketData.GetMarketHours() -> *dinari.ExtendedMarketHours</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the market hours for the current trading session and next open trading session.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.MarketData.GetMarketHours(
    context.TODO(),
)
```

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

## Corporate Actions

<details><summary><code>client.CorporateActions.GetStockSplits() -> *dinari.GetStockSplitsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of stock splits for `Stocks` available for trade via Dinari.
The splits are ordered by the date they were created, with the most recent split first.

In an example 10-for-1 stock split, trading will be halted for the stock at the end of the `payable_date`, as
the split transitions from `PENDING` to `IN_PROGRESS`. This usually occurs over a weekend to minimize trading
disruptions. Each share of stock owned by a shareholder will then be converted into 10 shares, and the split
becomes `COMPLETE` as trading resumes on the `ex_date` with new split-adjusted prices.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetStockSplitsRequest{}
client.CorporateActions.GetStockSplits(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**page:** `*int`

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int`

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return

</dd>
</dl>

<dl>
<dd>

**order:** `*dinari.GetStockSplitsRequestOrder` — Sort order

</dd>
</dl>

<dl>
<dd>

**next:** `*string` — Cursor for next page

</dd>
</dl>

<dl>
<dd>

**previous:** `*string` — Cursor for previous page

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.CorporateActions.GetStockSplitsByID(StockID) -> *dinari.GetStockSplitsByIDResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of stock splits for a specific `Stock`.
The splits are ordered by the date they were created, with the most recent split first.

In an example 10-for-1 stock split, trading will be halted for the stock at the end of the `payable_date`, as
the split transitions from `PENDING` to `IN_PROGRESS`. This usually occurs over a weekend to minimize trading
disruptions. Each share of stock owned by a shareholder will then be converted into 10 shares, and the split
becomes `COMPLETE` as trading resumes on the `ex_date` with new split-adjusted prices.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetStockSplitsByIDRequest{
    StockID: "stock_id",
}
client.CorporateActions.GetStockSplitsByID(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**stockID:** `string`

</dd>
</dl>

<dl>
<dd>

**page:** `*int`

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int`

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return

</dd>
</dl>

<dl>
<dd>

**order:** `*dinari.GetStockSplitsByIDRequestOrder` — Sort order

</dd>
</dl>

<dl>
<dd>

**next:** `*string` — Cursor for next page

</dd>
</dl>

<dl>
<dd>

**previous:** `*string` — Cursor for previous page

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.CorporateActions.GetStockDividendsByStockID(StockID) -> []*dinari.StockDividendAnnouncementDetail</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of announced stock dividend details for a specified `Stock`.

Note that this data applies only to actual stocks. Yield received for holding tokenized shares may differ from this.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetStockDividendsByStockIDRequest{
    StockID: "stock_id",
}
client.CorporateActions.GetStockDividendsByStockID(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**stockID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

## Entities

<details><summary><code>client.Entities.GetEntities() -> *dinari.GetEntitiesResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of direct `Entities` your organization manages. These `Entities` represent individual customers of your organization.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetEntitiesRequest{}
client.Entities.GetEntities(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**referenceID:** `*string` — Case sensitive unique reference ID for the `Entity`.

</dd>
</dl>

<dl>
<dd>

**page:** `*int`

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int`

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return

</dd>
</dl>

<dl>
<dd>

**order:** `*dinari.GetEntitiesRequestOrder` — Sort order

</dd>
</dl>

<dl>
<dd>

**next:** `*string` — Cursor for next page

</dd>
</dl>

<dl>
<dd>

**previous:** `*string` — Cursor for previous page

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Entities.CreateEntities(request) -> *dinari.Entity1</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new `Entity` to be managed by your organization. This `Entity` represents an individual customer of your organization.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateEntityInput{
    Name: "name",
}
client.Entities.CreateEntities(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — Name of the `Entity`.

</dd>
</dl>

<dl>
<dd>

**referenceID:** `*string` — Case sensitive unique reference ID for the `Entity`. We recommend setting this to the unique ID of the `Entity` in your system.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Entities.GetCurrentEntity() -> *dinari.Entity1</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the current authenticated `Entity`, which represents your organization.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Entities.GetCurrentEntity(
    context.TODO(),
)
```

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Entities.GetEntityByID(EntityID) -> *dinari.Entity1</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific customer `Entity` of your organization by their ID.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetEntityByIDRequest{
    EntityID: "entity_id",
}
client.Entities.GetEntityByID(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**entityID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Entities.UpdateEntityByID(EntityID, request) -> *dinari.Entity1</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update a specific customer `Entity` of your organization.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.PatchEntityInput{
    EntityID: "entity_id",
}
client.Entities.UpdateEntityByID(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**entityID:** `string`

</dd>
</dl>

<dl>
<dd>

**referenceID:** `*string` — Case sensitive unique reference ID for the `Entity`. We recommend setting this to the unique ID of the `Entity` in your system.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

## Accounts

<details><summary><code>client.Accounts.GetEntityAccounts(EntityID) -> *dinari.GetEntityAccountsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of all `Accounts` that belong to a specific `Entity`. This `Entity` represents your organization itself, or an individual customer of your organization.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetEntityAccountsRequest{
    EntityID: "entity_id",
}
client.Accounts.GetEntityAccounts(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**entityID:** `string`

</dd>
</dl>

<dl>
<dd>

**page:** `*int`

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int`

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return

</dd>
</dl>

<dl>
<dd>

**order:** `*dinari.GetEntityAccountsRequestOrder` — Sort order

</dd>
</dl>

<dl>
<dd>

**next:** `*string` — Cursor for next page

</dd>
</dl>

<dl>
<dd>

**previous:** `*string` — Cursor for previous page

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Accounts.CreateEntityAccounts(EntityID, request) -> *dinari.Account1</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a new `Account` for a specific `Entity`. This `Entity` represents your organization itself, or an individual customer of your organization.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateAccountInput{
    EntityID: "entity_id",
}
client.Accounts.CreateEntityAccounts(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**entityID:** `string`

</dd>
</dl>

<dl>
<dd>

**jurisdiction:** `*dinari.Jurisdiction` — Jurisdiction of the `Account`.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Accounts.GetAccount(AccountID) -> *dinari.Account1</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific `Account` by its ID.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetAccountRequest{
    AccountID: "account_id",
}
client.Accounts.GetAccount(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Accounts.CreateAccountDeactivation(AccountID) -> *dinari.Account1</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Set the `Account` to be inactive. Inactive accounts cannot be used for trading.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateAccountDeactivationRequest{
    AccountID: "account_id",
}
client.Accounts.CreateAccountDeactivation(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Accounts.GetAccountPortfolio(AccountID) -> *dinari.Portfolio</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the portfolio of the `Account`, excluding cash equivalents such as stablecoins.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetAccountPortfolioRequest{
    AccountID: "account_id",
}
client.Accounts.GetAccountPortfolio(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**page:** `*int` — The page number.

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The number of stocks to return per page, maximum number is 200.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Accounts.GetAccountCash(AccountID) -> []*dinari.PaymentTokenBalance</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the cash balances of the `Account`, including stablecoins and other cash equivalents.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetAccountCashRequest{
    AccountID: "account_id",
}
client.Accounts.GetAccountCash(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Accounts.GetAccountInterestPayments(AccountID) -> *dinari.GetAccountInterestPaymentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get interest payments made to the `Account` from yield-bearing cash holdings.

Currently, the only yield-bearing stablecoin accepted by Dinari is [USD+](https://usd.dinari.com/).

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetAccountInterestPaymentsRequest{
    AccountID: "account_id",
    StartDate: dinari.MustParseDate(
        "2023-01-15",
    ),
    EndDate: dinari.MustParseDate(
        "2023-01-15",
    ),
}
client.Accounts.GetAccountInterestPayments(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**startDate:** `time.Time` — Start date, inclusive, in US Eastern time zone. ISO 8601 format, YYYY-MM-DD.

</dd>
</dl>

<dl>
<dd>

**endDate:** `time.Time` — End date, exclusive, in US Eastern time zone. ISO 8601 format, YYYY-MM-DD.

</dd>
</dl>

<dl>
<dd>

**page:** `*int`

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int`

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return

</dd>
</dl>

<dl>
<dd>

**order:** `*dinari.GetAccountInterestPaymentsRequestOrder` — Sort order

</dd>
</dl>

<dl>
<dd>

**next:** `*string` — Cursor for next page

</dd>
</dl>

<dl>
<dd>

**previous:** `*string` — Cursor for previous page

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Accounts.GetAccountDividendPayments(AccountID) -> *dinari.GetAccountDividendPaymentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get dividend payments made to the `Account` from dividend-bearing stock holdings.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetAccountDividendPaymentsRequest{
    AccountID: "account_id",
    StartDate: dinari.MustParseDate(
        "2023-01-15",
    ),
    EndDate: dinari.MustParseDate(
        "2023-01-15",
    ),
}
client.Accounts.GetAccountDividendPayments(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**startDate:** `time.Time` — Start date, inclusive, in US Eastern time zone. ISO 8601 format, YYYY-MM-DD.

</dd>
</dl>

<dl>
<dd>

**endDate:** `time.Time` — End date, exclusive, in US Eastern time zone. ISO 8601 format, YYYY-MM-DD.

</dd>
</dl>

<dl>
<dd>

**stockID:** `*string` — Optional ID of the `Stock` to filter by

</dd>
</dl>

<dl>
<dd>

**page:** `*int`

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int`

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return

</dd>
</dl>

<dl>
<dd>

**order:** `*dinari.GetAccountDividendPaymentsRequestOrder` — Sort order

</dd>
</dl>

<dl>
<dd>

**next:** `*string` — Cursor for next page

</dd>
</dl>

<dl>
<dd>

**previous:** `*string` — Cursor for previous page

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Accounts.CreateAccountFaucet(AccountID, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Mints 1,000 mockUSD sandbox payment tokens to the `Wallet` connected to the `Account`.

This feature is only supported in sandbox mode.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.FaucetInput{
    AccountID: "account_id",
}
client.Accounts.CreateAccountFaucet(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**chainID:** `*dinari.ChainID` — CAIP-2 chain ID of blockchain in which to mint the sandbox payment tokens. If none specified, defaults to eip155:421614. If the `Account` is linked to a Dinari-managed `Wallet`, only eip155:42161 is allowed.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Accounts.GetBrokerageAccountActivities(AccountID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of brokerage activities tied to the specified `Account`.

**⚠️ ALPHA: This endpoint is in early development and subject to breaking changes.**

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetBrokerageAccountActivitiesRequest{
    AccountID: "account_id",
}
client.Accounts.GetBrokerageAccountActivities(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The maximum number of entries to return in the response. Defaults to 100.

</dd>
</dl>

<dl>
<dd>

**pageToken:** `*string` — Pagination token. Set to the `id` field of the last Activity returned in the previous page to get the next page of results.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Accounts.GetTokenTransfers(AccountID) -> *dinari.GetTokenTransfersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get `TokenTransfer`s made from this `Account`.

A `TokenTransfer` represents a transfer of tokens through the Dinari platform from one `Account` to another.
As such, only `Account`s that are connected to Dinari-managed `Wallet`s can initiate `TokenTransfer`s.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetTokenTransfersRequest{
    AccountID: "account_id",
}
client.Accounts.GetTokenTransfers(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**page:** `*int`

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int`

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return

</dd>
</dl>

<dl>
<dd>

**order:** `*dinari.GetTokenTransfersRequestOrder` — Sort order

</dd>
</dl>

<dl>
<dd>

**next:** `*string` — Cursor for next page

</dd>
</dl>

<dl>
<dd>

**previous:** `*string` — Cursor for previous page

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Accounts.CreateTokenTransfers(AccountID, request) -> *dinari.TokenTransfer</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `TokenTransfer` from this `Account`.

A `TokenTransfer` represents a transfer of tokens through the Dinari platform from one `Account` to another.
As such, only `Account`s that are connected to Dinari-managed `Wallet`s can initiate `TokenTransfer`s.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateManagedTokenTransferInput{
    AccountID: "account_id",
    RecipientAccountID: "recipient_account_id",
    TokenAddress: "token_address",
    Quantity: 1.1,
}
client.Accounts.CreateTokenTransfers(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**recipientAccountID:** `string` — ID of the recipient account to which the tokens will be transferred.

</dd>
</dl>

<dl>
<dd>

**tokenAddress:** `string` — Address of the token to transfer.

</dd>
</dl>

<dl>
<dd>

**quantity:** `float64` — Quantity of the token to transfer.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Accounts.GetTokenTransfer(AccountID, TransferID) -> *dinari.TokenTransfer</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific `TokenTransfer` made from this `Account` by its ID.

A `TokenTransfer` represents a transfer of tokens through the Dinari platform from one `Account` to another.
As such, only `Account`s that are connected to Dinari-managed `Wallet`s can initiate `TokenTransfer`s.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetTokenTransferRequest{
    AccountID: "account_id",
    TransferID: "transfer_id",
}
client.Accounts.GetTokenTransfer(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**transferID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

## KYC

<details><summary><code>client.Kyc.GetEntityKyc(EntityID) -> *dinari.Kyc</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get most recent KYC data of the `Entity`.

If there are any completed KYC checks, data from the most recent one will be returned.
If there are no completed KYC checks, the most recent KYC check information, regardless of status, will be returned.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetEntityKycRequest{
    EntityID: "entity_id",
}
client.Kyc.GetEntityKyc(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**entityID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Kyc.CreateEntityKyc(EntityID, request) -> *dinari.Kyc</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submit KYC data directly, for partners that are provisioned to provide their own KYC data.

This feature is available for everyone in sandbox mode, and for specifically provisioned partners in production.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateEntityKycRequest{
    EntityID: "entity_id",
    Body: &dinari.CreateKycInput{
        Baseline: &dinari.CreateBaselineKycInput{
            Data: &dinari.BaselineKycCheckDataV0{
                CountryCode: "SG",
                LastName: "Doe",
                AddressCountryCode: "SG",
            },
            ProviderName: "provider_name",
        },
    },
}
client.Kyc.CreateEntityKyc(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**entityID:** `string`

</dd>
</dl>

<dl>
<dd>

**request:** `*dinari.CreateKycInput`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Kyc.GetEntityKycDocument(EntityID, KycID) -> []*dinari.KycDocument</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get uploaded documents for a KYC check

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetEntityKycDocumentRequest{
    EntityID: "entity_id",
    KycID: "kyc_id",
}
client.Kyc.GetEntityKycDocument(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**entityID:** `string`

</dd>
</dl>

<dl>
<dd>

**kycID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Kyc.CreateEntityKycDocument(EntityID, KycID, request) -> *dinari.KycDocument</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Upload KYC-related documentation for partners that are provisioned to provide their own KYC data.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateKycDocumentFileInput{
    EntityID: "entity_id",
    KycID: "kyc_id",
    DocumentType: dinari.KycDocumentTypeGovernmentID,
    File: strings.NewReader(
        "",
    ),
}
client.Kyc.CreateEntityKycDocument(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**entityID:** `string`

</dd>
</dl>

<dl>
<dd>

**kycID:** `string`

</dd>
</dl>

<dl>
<dd>

**documentType:** `*dinari.KycDocumentType` — Type of `KYCDocument` to be uploaded.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Kyc.CreateManagedEntityKycEmbed(EntityID, request) -> *dinari.ManagedKycEmbed</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a Dinari-managed KYC Check and get a URL for your end customer to interact with.

The URL points to a web-based KYC interface that can be presented to the end customer for KYC verification.
Once the customer completes this KYC flow, the KYC check will be created and available in the KYC API.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateManagedKycInput{
    EntityID: "entity_id",
}
client.Kyc.CreateManagedEntityKycEmbed(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**entityID:** `string`

</dd>
</dl>

<dl>
<dd>

**jurisdiction:** `*dinari.Jurisdiction` — Jurisdiction for the KYC check. Defaults to BASELINE.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

## Orders

<details><summary><code>client.Orders.GetEntityOrders() -> *dinari.GetEntityOrdersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of all `Orders` under the `Entity`.<br>Optionally `Orders` can be transaction hash or fulfillment transaction hash.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetEntityOrdersRequest{}
client.Orders.GetEntityOrders(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**chainID:** `*string` — CAIP-2 formatted chain ID of the blockchain the `Order` was made on.

</dd>
</dl>

<dl>
<dd>

**orderTransactionHash:** `*string` — Transaction hash of the `Order`.

</dd>
</dl>

<dl>
<dd>

**orderFulfillmentTransactionHash:** `*string` — Fulfillment transaction hash of the `Order`.

</dd>
</dl>

<dl>
<dd>

**orderRequestID:** `*string` — Order Request ID for the `Order`

</dd>
</dl>

<dl>
<dd>

**page:** `*int`

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int`

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return

</dd>
</dl>

<dl>
<dd>

**order:** `*dinari.GetEntityOrdersRequestOrder` — Sort order

</dd>
</dl>

<dl>
<dd>

**next:** `*string` — Cursor for next page

</dd>
</dl>

<dl>
<dd>

**previous:** `*string` — Cursor for previous page

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Orders.GetOrders(AccountID) -> *dinari.GetOrdersResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of all `Orders` under the `Account`.<br>Optionally `Orders` can be filtered by chain ID, transaction hash, or client order ID.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetOrdersRequest{
    AccountID: "account_id",
}
client.Orders.GetOrders(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**chainID:** `*string` — CAIP-2 formatted chain ID of the blockchain the `Order` was made on.

</dd>
</dl>

<dl>
<dd>

**orderTransactionHash:** `*string` — Transaction hash of the `Order`.

</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Customer-supplied identifier to search for `Order`s.

</dd>
</dl>

<dl>
<dd>

**page:** `*int`

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int`

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return

</dd>
</dl>

<dl>
<dd>

**order:** `*dinari.GetOrdersRequestOrder` — Sort order

</dd>
</dl>

<dl>
<dd>

**next:** `*string` — Cursor for next page

</dd>
</dl>

<dl>
<dd>

**previous:** `*string` — Cursor for previous page

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Orders.GetOrderByID(AccountID, OrderID) -> *dinari.Order</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific `Order` by its ID.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetOrderByIDRequest{
    AccountID: "account_id",
    OrderID: "order_id",
}
client.Orders.GetOrderByID(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**orderID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Orders.CreateOrderCancellation(AccountID, OrderID) -> *dinari.Order</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancel an `Order` by its ID. Note that this requires the `Order` ID, not the `OrderRequest` ID.
Once you submit a cancellation request, it cannot be undone. Be advised that orders with a
status of PENDING_FILL, PENDING_ESCROW, FILLED, REJECTED, or CANCELLED cannot be cancelled.

`Order` cancellation is not guaranteed nor is it immediate. The `Order` may still be executed
if the cancellation request is not received in time.

Check the status using the "Get Order by ID" endpoint to confirm whether the `Order` has been cancelled.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateOrderCancellationRequest{
    AccountID: "account_id",
    OrderID: "order_id",
}
client.Orders.CreateOrderCancellation(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**orderID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Orders.CreateBatchOrderCancellation(AccountID, request) -> *dinari.BatchCancelOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Cancel multiple `Orders` by their IDs in a single request. Note that this requires the `Order` IDs, not the `OrderRequest` IDs.
Once you submit a cancellation request, it cannot be undone. Be advised that orders with a
status of PENDING_FILL, PENDING_ESCROW, FILLED, REJECTED, or CANCELLED cannot be cancelled.

`Order` cancellation is not guaranteed nor is it immediate. The `Orders` may still be executed
if the cancellation request is not received in time.

The response will indicate which orders were successfully queued to cancel and which failed to queue.
Check the status using the "Get Order by ID" endpoint to confirm whether individual `Orders` have been cancelled.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.BatchCancelOrderInput{
    AccountID: "account_id",
    OrderIDs: []string{
        "order_ids",
    },
}
client.Orders.CreateBatchOrderCancellation(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**orderIDs:** `[]string` — List of `Order` IDs to cancel

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Orders.GetAccountAllOrderFulfillments(AccountID) -> *dinari.GetAccountAllOrderFulfillmentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Query `OrderFulfillments` under the `Account`.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetAccountAllOrderFulfillmentsRequest{
    AccountID: "account_id",
}
client.Orders.GetAccountAllOrderFulfillments(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**orderIDs:** `*string` — List of `Order` IDs to query `AccountOrderFulfillment` for.

</dd>
</dl>

<dl>
<dd>

**page:** `*int`

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int`

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return

</dd>
</dl>

<dl>
<dd>

**order:** `*dinari.GetAccountAllOrderFulfillmentsRequestOrder` — Sort order

</dd>
</dl>

<dl>
<dd>

**next:** `*string` — Cursor for next page

</dd>
</dl>

<dl>
<dd>

**previous:** `*string` — Cursor for previous page

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Orders.GetAccountOrderFulfillments(AccountID, OrderID) -> *dinari.GetAccountOrderFulfillmentsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get `OrderFulfillments` for a specific `Order`.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetAccountOrderFulfillmentsRequest{
    AccountID: "account_id",
    OrderID: "order_id",
}
client.Orders.GetAccountOrderFulfillments(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**orderID:** `string`

</dd>
</dl>

<dl>
<dd>

**page:** `*int`

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int`

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return

</dd>
</dl>

<dl>
<dd>

**order:** `*dinari.GetAccountOrderFulfillmentsRequestOrder` — Sort order

</dd>
</dl>

<dl>
<dd>

**next:** `*string` — Cursor for next page

</dd>
</dl>

<dl>
<dd>

**previous:** `*string` — Cursor for previous page

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Orders.GetAccountOrderFulfillmentByID(AccountID, OrderFulfillmentID) -> *dinari.OrderFulfillment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific `OrderFulfillment` by its ID.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetAccountOrderFulfillmentByIDRequest{
    AccountID: "account_id",
    OrderFulfillmentID: "order_fulfillment_id",
}
client.Orders.GetAccountOrderFulfillmentByID(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**orderFulfillmentID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

## Wallets

<details><summary><code>client.Wallets.GetWallet(AccountID) -> *dinari.Wallet</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the wallet connected to the `Account`.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetWalletRequest{
    AccountID: "account_id",
}
client.Wallets.GetWallet(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Wallets.GetAccountWalletConnectionNonce(AccountID) -> *dinari.AccountWalletConnectionMessage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a nonce and message to be signed in order to verify `Wallet` ownership.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetAccountWalletConnectionNonceRequest{
    AccountID: "account_id",
    WalletAddress: "wallet_address",
    ChainID: dinari.WalletChainIDEip1550,
}
client.Wallets.GetAccountWalletConnectionNonce(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**walletAddress:** `string` — Address of the `Wallet` to connect.

</dd>
</dl>

<dl>
<dd>

**chainID:** `*dinari.WalletChainID` — CAIP-2 formatted chain ID of the blockchain the `Wallet` is on. eip155:0 is used for EOA wallets

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Wallets.CreateAccountWalletConnection(AccountID, request) -> *dinari.Wallet</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Connect a `Wallet` to the `Account` after verifying the signature.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateAccountWalletConnectionInput{
    AccountID: "account_id",
    Signature: "signature",
    Nonce: "nonce",
    WalletAddress: "wallet_address",
    ChainID: dinari.WalletChainIDEip1550,
}
client.Wallets.CreateAccountWalletConnection(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**signature:** `string` — Signature payload from signing the connection message with the `Wallet`.

</dd>
</dl>

<dl>
<dd>

**nonce:** `string` — Nonce contained within the connection message.

</dd>
</dl>

<dl>
<dd>

**walletAddress:** `string` — Address of the `Wallet`.

</dd>
</dl>

<dl>
<dd>

**chainID:** `*dinari.WalletChainID` — CAIP-2 formatted chain ID of the blockchain the `Wallet` to link is on. eip155:0 is used for EOA wallets

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Wallets.CreateAccountWalletInternalConnection(AccountID, request) -> *dinari.Wallet</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Connect an internal `Wallet` to the `Account`.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateAccountWalletInternalConnectionInput{
    AccountID: "account_id",
    WalletAddress: "wallet_address",
    ChainID: dinari.WalletChainIDEip1550,
}
client.Wallets.CreateAccountWalletInternalConnection(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**walletAddress:** `string` — Address of the `Wallet`.

</dd>
</dl>

<dl>
<dd>

**chainID:** `*dinari.WalletChainID` — CAIP-2 formatted chain ID of the blockchain the `Wallet` to link is on. eip155:0 is used for EOA wallets

</dd>
</dl>

<dl>
<dd>

**isShared:** `*bool` — Is the linked Wallet shared or not

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

## Order Requests

<details><summary><code>client.OrderRequests.GetListOrderRequests(AccountID) -> *dinari.GetListOrderRequestsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists `OrderRequests`.<br>Optionally `OrderRequests` can be filtered by certain parameters.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetListOrderRequestsRequest{
    AccountID: "account_id",
}
client.OrderRequests.GetListOrderRequests(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**orderRequestID:** `*string` — Order Request ID for the `OrderRequest`

</dd>
</dl>

<dl>
<dd>

**orderID:** `*string` — Order ID for the `OrderRequest`

</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Customer-supplied ID to map this `OrderRequest` to an order in their own systems.

</dd>
</dl>

<dl>
<dd>

**page:** `*int`

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int`

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return

</dd>
</dl>

<dl>
<dd>

**order:** `*dinari.GetListOrderRequestsRequestOrder` — Sort order

</dd>
</dl>

<dl>
<dd>

**next:** `*string` — Cursor for next page

</dd>
</dl>

<dl>
<dd>

**previous:** `*string` — Cursor for previous page

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.OrderRequests.GetOrderRequestByID(AccountID, OrderRequestID) -> *dinari.OrderRequest</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific `OrderRequest` by its ID.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetOrderRequestByIDRequest{
    AccountID: "account_id",
    OrderRequestID: "order_request_id",
}
client.OrderRequests.GetOrderRequestByID(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**orderRequestID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.OrderRequests.CreateEip155OrderRequestPermit(AccountID, request) -> *dinari.Eip155OrderRequestPermit</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Generates a permit that can be signed and used to create an `OrderRequest` using Dinari's EVM smart contracts.

This is a convenience method to prepare the transactions needed to create an `OrderRequest` using Dinari's EVM smart contracts.
Once signed, the transactions can be sent to the EVM network to create the order.
Note that the fee quote is already included in the transactions, so no additional fee quote lookup is needed.

Fees for the `Order` can optionally be specified in the `OrderRequest` for DFN orders in USD, supporting up to 6 decimal places.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.Eip155OrderRequestPermitInput{
    AccountID: "account_id",
    ChainID: dinari.ChainIDEip1551,
    OrderTif: dinari.OrderTifDay,
    OrderSide: dinari.OrderSideBuy,
    OrderType: dinari.OrderTypeMarket,
    PaymentToken: "payment_token",
}
client.OrderRequests.CreateEip155OrderRequestPermit(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**chainID:** `*dinari.ChainID` — CAIP-2 chain ID of the blockchain where the `Order` will be placed.

</dd>
</dl>

<dl>
<dd>

**orderTif:** `*dinari.OrderTif` — Time in force. Indicates how long `Order` is valid for.

</dd>
</dl>

<dl>
<dd>

**orderSide:** `*dinari.OrderSide` — Indicates whether `Order` is a buy or sell.

</dd>
</dl>

<dl>
<dd>

**orderType:** `*dinari.OrderType` — Type of `Order`.

</dd>
</dl>

<dl>
<dd>

**paymentToken:** `string` — Address of payment token. Required for Accounts outside of US jurisdiction. Accounts with US jurisdiction must use: USDC for live, MockUSD for sandbox.

</dd>
</dl>

<dl>
<dd>

**paymentTokenQuantity:** `*float64` — Amount of payment tokens involved. Required for market buy `Orders`.

</dd>
</dl>

<dl>
<dd>

**assetTokenQuantity:** `*float64` — Amount of dShare asset tokens involved. Required for limit `Order Requests` and market sell `Order Requests`. Must be a positive number with a precision of up to 4 decimal places for limit `Order Requests` or up to 6 decimal places for market sell `Order Requests`.

</dd>
</dl>

<dl>
<dd>

**limitPrice:** `*float64` — Price per asset in the asset's native currency. USD for US equities and ETFs. Required for limit `Orders`.

</dd>
</dl>

<dl>
<dd>

**stockID:** `*string` — The ID of the `Stock` for which the `Order` is being placed.

</dd>
</dl>

<dl>
<dd>

**tokenID:** `*string` — The ID of the `Token` for which the `Order` is being placed.

</dd>
</dl>

<dl>
<dd>

**alloyID:** `*string` — The ID of the `Alloy` for which the `Order` is being placed.

</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Customer-supplied unique identifier to map this `Order` to an order in the customer's systems.

</dd>
</dl>

<dl>
<dd>

**fee:** `*float64` — Optional fee amount associated with `Order` in USD for DFN orders. Must be a positive number with a precision of up to 6 decimal places.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.OrderRequests.CreateEip155OrderRequestPermitTransaction(AccountID, request) -> *dinari.Eip155OrderRequest</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Submits a transaction for an EIP155 Order Request given the EIP155OrderRequest ID and Permit Signature.

An `EIP155OrderRequest` representing the proxied order is returned.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateEip155OrderRequestPermitTransactionRequest{
    AccountID: "account_id",
    Body: &dinari.Eip155OrderRequestPermitTransaction{
        OrderRequestID: "order_request_id",
        PermitSignature: "permit_signature",
    },
}
client.OrderRequests.CreateEip155OrderRequestPermitTransaction(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**request:** `*dinari.Eip155OrderRequestPermitTransaction`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.OrderRequests.CreateEip155PermitTransaction(AccountID, request) -> *dinari.Eip155Transaction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Prepare a transaction to be placed on EVM. The returned structure contains the necessary
data to create an `EIP155Transaction` object.

Safe wallets can send `0x` as `permit_signature`. Those orders aren't simulated before they're returned.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateEip155PermitTransactionRequest{
    AccountID: "account_id",
    Body: &dinari.Eip155OrderRequestPermitTransaction{
        OrderRequestID: "order_request_id",
        PermitSignature: "permit_signature",
    },
}
client.OrderRequests.CreateEip155PermitTransaction(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**request:** `*dinari.Eip155OrderRequestPermitTransaction`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

## Managed Orders

<details><summary><code>client.ManagedOrders.CreateMarketBuyManagedOrderRequest(AccountID, request) -> *dinari.OrderRequest</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a managed `OrderRequest` to place a market buy `Order`.

Fees for the `Order` can optionally be specified in the `OrderRequest` for DFN orders in USD, supporting up to 6 decimal places

If an `OrderRequest` with the same `client_order_id` already exists for the given account, the creation call will fail.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateMarketBuyOrderRequestInput{
    AccountID: "account_id",
    PaymentAmount: 1.1,
}
client.ManagedOrders.CreateMarketBuyManagedOrderRequest(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**paymentAmount:** `float64`

Amount of currency (USD for US equities and ETFs) to pay for the order.
Must be a positive number with a precision of up to 2 decimal places.

</dd>
</dl>

<dl>
<dd>

**paymentTokenAddress:** `*string` — Address of the payment token to be used for the payment of the order. If not provided, the default payment token (USD+) will be used.

</dd>
</dl>

<dl>
<dd>

**stockID:** `*string` — ID of `Stock`.

</dd>
</dl>

<dl>
<dd>

**alloyID:** `*string` — ID of `Alloy`.

</dd>
</dl>

<dl>
<dd>

**recipientAccountID:** `*string` — ID of `Account` to receive the `Order`.

</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Customer-supplied ID to map this order to an order in their own systems. Must be unique within the entity.

</dd>
</dl>

<dl>
<dd>

**fee:** `*float64` — Optional fee amount associated with `Order` in USD for DFN orders. Must be a positive number with a precision of up to 6 decimal places.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.ManagedOrders.CreateMarketSellManagedOrderRequest(AccountID, request) -> *dinari.OrderRequest</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a managed `OrderRequest` to place a market sell `Order`.

Fees for the `Order` can optionally be specified in the `OrderRequest` for DFN orders in USD, supporting up to 6 decimal places

If an `OrderRequest` with the same `client_order_id` already exists for the given account, the creation call will fail.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateMarketSellOrderRequestInput{
    AccountID: "account_id",
    AssetQuantity: 1.1,
}
client.ManagedOrders.CreateMarketSellManagedOrderRequest(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**assetQuantity:** `float64` — Quantity of shares to trade. Must be a positive number with a precision of up to 6 decimal places.

</dd>
</dl>

<dl>
<dd>

**stockID:** `*string` — ID of `Stock`.

</dd>
</dl>

<dl>
<dd>

**alloyID:** `*string` — ID of `Alloy`.

</dd>
</dl>

<dl>
<dd>

**recipientAccountID:** `*string` — ID of `Account` to receive the `Order`.

</dd>
</dl>

<dl>
<dd>

**paymentTokenAddress:** `*string` — Address of the payment token to be used for the sell order. If not provided, the default payment token (USD+) will be used. Should only be specified if `recipient_account_id` for a non-managed wallet account is also provided.

</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Customer-supplied ID to map this order to an order in their own systems. Must be unique within the entity.

</dd>
</dl>

<dl>
<dd>

**fee:** `*float64` — Optional fee amount associated with `Order` in USD for DFN orders. Must be a positive number with a precision of up to 6 decimal places.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.ManagedOrders.CreateLimitBuyManagedOrderRequest(AccountID, request) -> *dinari.OrderRequest</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a managed `OrderRequest` to place a limit buy `Order`.

Fees for the `Order` can optionally be specified in the `OrderRequest` for DFN orders in USD, supporting up to 6 decimal places

If an `OrderRequest` with the same `client_order_id` already exists for the given account, the creation call will fail.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateLimitOrderBuyRequestInput{
    AccountID: "account_id",
    AssetQuantity: 1.1,
    LimitPrice: 1.1,
}
client.ManagedOrders.CreateLimitBuyManagedOrderRequest(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**assetQuantity:** `float64` — Amount of dShare asset involved. Required for limit `Order Requests` and market sell `Order Requests`. Must be a positive number with a precision of up to 4 decimal places.

</dd>
</dl>

<dl>
<dd>

**limitPrice:** `float64` — Price at which to execute the order. Must be a positive number with a precision of up to 2 decimal places.

</dd>
</dl>

<dl>
<dd>

**stockID:** `*string` — ID of `Stock`.

</dd>
</dl>

<dl>
<dd>

**alloyID:** `*string` — ID of `Alloy`.

</dd>
</dl>

<dl>
<dd>

**recipientAccountID:** `*string` — ID of `Account` to receive the `Order`.

</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Customer-supplied ID to map this order to an order in their own systems. Must be unique within the entity.

</dd>
</dl>

<dl>
<dd>

**paymentTokenAddress:** `*string` — Address of the payment token to be used for the payment of the order. If not provided, the default payment token (USD+) will be used.

</dd>
</dl>

<dl>
<dd>

**fee:** `*float64` — Optional fee amount associated with `Order` in USD for DFN orders. Must be a positive number with a precision of up to 6 decimal places.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.ManagedOrders.CreateLimitSellManagedOrderRequest(AccountID, request) -> *dinari.OrderRequest</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a managed `OrderRequest` to place a limit sell `Order`.

Fees for the `Order` can optionally be specified in the `OrderRequest` for DFN orders in USD, supporting up to 6 decimal places

If an `OrderRequest` with the same `client_order_id` already exists for the given account, the creation call will fail.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateLimitOrderSellRequestInput{
    AccountID: "account_id",
    AssetQuantity: 1.1,
    LimitPrice: 1.1,
}
client.ManagedOrders.CreateLimitSellManagedOrderRequest(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**assetQuantity:** `float64` — Amount of dShare asset involved. Required for limit `Order Requests` and market sell `Order Requests`. Must be a positive number with a precision of up to 4 decimal places.

</dd>
</dl>

<dl>
<dd>

**limitPrice:** `float64` — Price at which to execute the order. Must be a positive number with a precision of up to 2 decimal places.

</dd>
</dl>

<dl>
<dd>

**stockID:** `*string` — ID of `Stock`.

</dd>
</dl>

<dl>
<dd>

**alloyID:** `*string` — ID of `Alloy`.

</dd>
</dl>

<dl>
<dd>

**recipientAccountID:** `*string` — ID of `Account` to receive the `Order`.

</dd>
</dl>

<dl>
<dd>

**paymentTokenAddress:** `*string` — Address of the payment token to be used for the sell order. If not provided, the default payment token (USD+) will be used. Should only be specified if `recipient_account_id` for a non-managed wallet account is also provided.

</dd>
</dl>

<dl>
<dd>

**clientOrderID:** `*string` — Customer-supplied ID to map this order to an order in their own systems. Must be unique within the entity.

</dd>
</dl>

<dl>
<dd>

**fee:** `*float64` — Optional fee amount associated with `Order` in USD for DFN orders. Must be a positive number with a precision of up to 6 decimal places.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

## Withdrawals

<details><summary><code>client.Withdrawals.GetAccountWithdrawalRequests(AccountID) -> *dinari.GetAccountWithdrawalRequestsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List `WithdrawalRequests` under the `Account`, sorted by most recent.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetAccountWithdrawalRequestsRequest{
    AccountID: "account_id",
}
client.Withdrawals.GetAccountWithdrawalRequests(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**page:** `*int`

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int`

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return

</dd>
</dl>

<dl>
<dd>

**order:** `*dinari.GetAccountWithdrawalRequestsRequestOrder` — Sort order

</dd>
</dl>

<dl>
<dd>

**next:** `*string` — Cursor for next page

</dd>
</dl>

<dl>
<dd>

**previous:** `*string` — Cursor for previous page

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Withdrawals.CreateAccountWithdrawalRequests(AccountID, request) -> *dinari.WithdrawalRequest1</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Request to withdraw USD+ payment tokens from a managed `Account` and send the equivalent amount of USDC to the specified recipient `Account`.

The recipient `Account` must belong to the same `Entity` as the managed `Account`.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.CreateWithdrawalRequestInput{
    AccountID: "account_id",
    RecipientAccountID: "recipient_account_id",
    PaymentTokenQuantity: 1.1,
}
client.Withdrawals.CreateAccountWithdrawalRequests(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**recipientAccountID:** `string` — ID of the `Account` that will receive payment tokens from the `Withdrawal`.

</dd>
</dl>

<dl>
<dd>

**paymentTokenQuantity:** `float64` — Amount of USD+ payment tokens to be withdrawn. Must be greater than 0 and have at most 6 decimal places.

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Withdrawals.GetAccountWithdrawalRequest(AccountID, WithdrawalRequestID) -> *dinari.WithdrawalRequest1</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific `WithdrawalRequest` by its ID.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetAccountWithdrawalRequestRequest{
    AccountID: "account_id",
    WithdrawalRequestID: "withdrawal_request_id",
}
client.Withdrawals.GetAccountWithdrawalRequest(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**withdrawalRequestID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Withdrawals.GetAccountWithdrawals(AccountID) -> *dinari.GetAccountWithdrawalsResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a list of all `Withdrawals` under the `Account`, sorted by most recent.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetAccountWithdrawalsRequest{
    AccountID: "account_id",
}
client.Withdrawals.GetAccountWithdrawals(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**withdrawalRequestID:** `*string` — ID of the `WithdrawalRequest` to find `Withdrawals` for.

</dd>
</dl>

<dl>
<dd>

**page:** `*int`

</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int`

</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Number of results to return

</dd>
</dl>

<dl>
<dd>

**order:** `*dinari.GetAccountWithdrawalsRequestOrder` — Sort order

</dd>
</dl>

<dl>
<dd>

**next:** `*string` — Cursor for next page

</dd>
</dl>

<dl>
<dd>

**previous:** `*string` — Cursor for previous page

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>

<details><summary><code>client.Withdrawals.GetAccountWithdrawal(AccountID, WithdrawalID) -> *dinari.Withdrawal1</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific `Withdrawal` by its ID.

</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &dinari.GetAccountWithdrawalRequest{
    AccountID: "account_id",
    WithdrawalID: "withdrawal_id",
}
client.Withdrawals.GetAccountWithdrawal(
    context.TODO(),
    request,
)
```

</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountID:** `string`

</dd>
</dl>

<dl>
<dd>

**withdrawalID:** `string`

</dd>
</dl>
</dd>
</dl>

</dd>
</dl>
</details>
