# QTag: Supercharge your Query Metadata

QTag is a standard and a set of tools for extracting comment based metadata from your SQL queries

The standard and all tools are released under the Apache2.0 OSS License

Getting started:

The easiest way to start with QTag is to install the function into Snowflake and use it against your query history.
Steps:

1. Install Snowflake Marketplace App
2. Run a SQL query using QTAG. For example:
   ```
   SELECT DECKTOOLS..QTAG(query_text) qtags FROM SNOWFLAKE.ACCOUNT_USAGE.QUERY_HISTORY
   ```

## Command Line Tools

QTag provides the `qtag` cli for common platforms that will extract qtags from a provided SQL query. This can be used by
piping in a SQL query and getting back the collection of relevant QTags. The QTag CLI tool has two optional parameters:

```bash
qtag --usage
Usage of qtag:
      --unexpected     whether to include unexpected attributes of known of comment patterns
      --unidentified   whether or not to include unidentified comment patterns
      --usage          whether or not to show usage
```

Example Usage:

```bash
# echo 'select 1 -- {"app":"dbt"}' | qtag
[
  {
    "Source": "dbt",
    "Type": "APPLICATION",
    "Key": "app",
    "Value": "dbt"
  }
]
```

## QTag Snowflake Function

QTag includes support to deploy the qtag parser as a self-contained Javascript UDF within Snowflake. There are two ways
you can add this functionality to your Snowflake cluster: install the marketplace app or manually build and install the
function.

By default, the QTag function installed in Snowflake will only extract QTags for queries of 

### Snowflake Marketplace App (a.k.a. the easy way)

Install the Decktools Snowflake App listed in the Snowflake Marketplace. This is the easiest way to access the QTag
functionality and other future capabilities as they are added to decktools. This is free to use.

### Manually Build and Deploy

You can checkout the github repo here and generate/deploy the functions to your Snowflake instance directly. In this
scenario, you'll need to manually keep your functions up to date as new comment definitions are created. To do this,
follow the following steps:

1. Checkout this git repository `git clone git@github.com:sundeck-io/qtag.git`
1. Install recent node/npm on your machine
1. Install all npm tools (`cd js` && `npm install`)
1. Export SNOWFLAKE_USERNAME, SNOWFLAKE_ACCOUNT and SNOWFLAKE_PASSWORD to your environment
1. Run `yarn snow` to deploy the function to your snowflake warehouse. It will be uploaded into the `DECKTOOLS.PUBLIC`
   schema as `QTAG()`

## Javascript Package

In the `js` directory, you export tools for working with QTags using Javascript.

## Go Library

In the `go` directory, we make a QTag extraction library available. To use this library, follow these instruction:

1. Add the qtag package to your go project `go get github.com/sundeck-io/qtag/go@latest`
1. Import the qtag package `import github.com/sundeck-io/qtag/go/pkg/qtag`
1. Create a new qtag extarctor `extractor := qtag.NewExtractor()`
1. Invoke the extractor with a set of
   options `var qtags []QTag = extractor.extract(<myquery>, <unidentified>, <unexpected>)`

## Adding additional pre-declared tools

QTag includes a number of well-known tools within the repository. If there is a tool that produces standard comments,
please propose a PR to add that to the project

## Support

QTag was created by the folks at [Sundeck](https://sundeck.io).

Find out more at [qtag.dev](https://qtag.dev)