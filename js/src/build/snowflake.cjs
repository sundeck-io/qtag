const fs = require('fs');
const snowflake = require('snowflake-sdk');

// Error Checking
const args = process.argv
if (args.length != 5) {
  exit("Required arguments <bundle> <database> <schema>.");
}

if (!process.env.SNOWFLAKE_USERNAME || !process.env.SNOWFLAKE_ACCOUNT || !process.env.SNOWFLAKE_PASSWORD) {
  exit("This script requires that SNOWFLAKE_USERNAME, SNOWFLAKE_ACCOUNT and SNOWFLAKE_PASSWORD environment variables are set.");
}

// create a replacement behavior that doesn't use backreferences since they will cause problems with our javascript blobs
String.prototype.replaceString = function(findNoRegex, replaceNoRegex) {
  return this.split(findNoRegex).join(replaceNoRegex);
}

// Template for creating Snowflake Function
const parserFunction = `
CREATE OR REPLACE FUNCTION 
  "{{database}}"."{{schema}}"._EXTRACT(query string)
  RETURNS variant
  LANGUAGE JAVASCRIPT
AS
$$
var initialize = function() {
  {{parser}}
  parse = CommentQL;
};

// we execute the initialization code once to take advantage of Snowflake's reuse of V8 isolates.
if (typeof(initialized) === "undefined") {
  initialize();
  initialized = true;
}
return parse(QUERY);
$$;
`
const wrapperFunction = `
CREATE OR REPLACE SECURE FUNCTION 
  "{{database}}"."{{schema}}".EXTRACT(query string)
  COPY GRANTS 
  RETURNS ARRAY
  AS
$$
  case when length(query) < 100000 then "{{database}}"."{{schema}}"._extract(query)::ARRAY
  else null
  end
$$;
`

const data = fs.readFileSync(args[2], 'utf8');

const connection = snowflake.createConnection({
  account: process.env.SNOWFLAKE_ACCOUNT,
  username: process.env.SNOWFLAKE_USERNAME,
  password: process.env.SNOWFLAKE_PASSWORD
});

connection.connect(
    function(err, conn) {
      if (err) {
        exit('Unable to connect: ' + err.message);
      } else {
        console.log("Snowflake Connected.");
      }
    }
);

connection.execute({
  sqlText: replaceDBAndSchema(parserFunction).replaceString('{{parser}}', data),

  complete: function(err, stmt, rows) {
    if (err) {
      exit('Failed to create parser function: ' + err.message);
    } else {
      console.log('Created parser function.');
      connection.execute({
        sqlText: replaceDBAndSchema(wrapperFunction),
        complete: function(err, stmt, rows) {
          if (err) {
            exit('Failed to create wrapper function: ' + err.message);
          } else {
            console.log('Created wrapper function.');
          }
        }
      });
    }
  }
});

// Replace database and schema with second and third cli arguments
function replaceDBAndSchema(str) {
  return str
      .replaceString('{{database}}',args[3])
      .replaceString('{{schema}}', args[4]);
}

function exit(msg) {
  console.error(msg)
  process.exit();
}
