# QTag: Turbocharge Your SQL Comments

QTag is a standard and a set of tools for extracting comment-based metadata from your SQL queries. QTag is built to work
with any kind of SQL but was initially used primarily with Snowflake.

The standard and all tools are released under the Apache2.0 OSS License.

Find out more at [qtag.dev](https://qtag.dev)

## Quickstart (via Snowflake marketplace)

1. Install the freely-available [Decktools Marketplace App](https://app.snowflake.com/marketplace/listing/GZTYZT5BVK/sundeck-decktools).
2. Enrich your query history with QTags using the provided scalar function.
   ```
   SELECT 
     DECKTOOLS..QTAG(query_text) qtags 
   FROM 
     SNOWFLAKE.ACCOUNT_USAGE.QUERY_HISTORY
3. Do more complex analyses using the provided table function.
   ```
   SELECT
     TAGS.* 
   FROM
     SNOWFLAKE.ACCOUNT_USAGE.QUERY_HISTORY,
     LATERAL DECKTOOLS..QTAG_TABLE(QUERY_TEXT) TAGS 
   ```

### Find out more at [sundeck.io/oss/qtag](https://sundeck.io/oss/qtag)

----
## In-depth Details

### QTag via Command Line

QTag also provides a `qtag` cli for common platforms that will extract QTags from a provided SQL query. This can be used by
piping in a SQL query and getting back the collection of relevant QTags. The QTag CLI tool has two optional parameters:

```bash
Usage of qtag:
      --all             include all unknown comments and attributes (default false)
      --allattributes   include all unknown attributes for known comments (default false)
      --usage           whether or not to show usage
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

The QTag CLI tool can be installed in your environment using the following steps:
1. Install Golang
2. Checkout QTag repo
3. cd into `go` directory
4. run go generate `go generate ./...`
5. run go install `go install ./cmd/qtag/...`

### QTag via Snowflake Function

QTag includes support to deploy the QTag parser as a self-contained Javascript UDF within Snowflake. There are two ways
you can add this functionality to your Snowflake cluster: install the [marketplace app](https://app.snowflake.com/marketplace/listing/GZTYZT5BVK/sundeck-decktools) 
or manually build and install the function.

By default, the QTag function installed in Snowflake will only extract QTags for queries of Hex, dbt, Sigma and Mode. 

#### Snowflake Marketplace App (a.k.a. the easy way)

Install the [Decktools Snowflake App](https://app.snowflake.com/marketplace/listing/GZTYZT5BVK/sundeck-decktools) listed 
in the Snowflake Marketplace. This is the easiest way to access the QTag functionality and 
other future capabilities as they are added to decktools. This is free to use.

#### Manually Build and Deploy

You can checkout the github repo here and generate/deploy the functions to your Snowflake instance directly. In this
scenario, you'll need to manually keep your functions up to date as new comment definitions are created. To do this,
follow the following steps:

1. Checkout this git repository `git clone git@github.com:sundeck-io/qtag.git`
1. Install recent node/npm on your machine
1. Install all npm tools (`cd js` && `npm install`)
1. Export SNOWFLAKE_USERNAME, SNOWFLAKE_ACCOUNT and SNOWFLAKE_PASSWORD to your environment
1. Run `yarn snow` to deploy the function to your snowflake warehouse. It will be uploaded into the `DECKTOOLS.PUBLIC`
   schema as `QTAG()` (including several variants).

## Javascript Package

In the `js` directory, you can interact with the entrypoint in `entry.js`. 

You can also interact with the exports within `qtag/extract.js`. This allows you to use two different parsers: a antlr-based parser (similar to the Go library) and a hand-written parser. (The hand-written parser is the default when deploying to Snowflake.) 

## Go Library

In the `go` directory, we make a QTag extraction library available. To use this library, follow these instruction:

1. Add the QTag package to your go project `go get github.com/sundeck-io/qtag/go@latest`
1. Import the QTag package `import github.com/sundeck-io/qtag/go/pkg/qtag`
1. Create a new QTag extarctor `extractor := qtag.NewExtractor()`
1. Invoke the extractor with a set of
   options `var qtags []QTag = extractor.extract(<myquery>, <unidentified>, <unexpected>)`

## Well-known Comment Types

QTag automatically identifies a number of comment patterns including those from dbt, Hex, Sigma and Mode. One 
goal of this project to collect up all the known patterns so we can make it easy for people to use QTag. Comment types
are declared within the project at incorporated into the code at build time. You can find the well defined comments here.

If there is a tool that produces well-known structured comments that you think should be included in QTag, please propose
a PR to add that to the project. When adding a new declared comment type, please also add [test cases](https://github.com/sundeck-io/qtag/tree/main/tests).

## Support & Community

QTag was created by the folks at [Sundeck](https://sundeck.io). It's released under the Apache 2.0 license and free for
all to use. Contributions are welcome.

