const fs = require('fs');
const snowflake = require('snowflake-sdk');
const {mapErrorCodeToSqlState} = require("snowflake-sdk/lib/errors.js");

// Error Checking
const args = process.argv
if (args.length !== 5) {
  exit("Required arguments <bundle> <database> <schema>.");
}

if (!process.env.SNOWFLAKE_USERNAME || !process.env.SNOWFLAKE_ACCOUNT || !process.env.SNOWFLAKE_PASSWORD) {
  exit("This script requires that SNOWFLAKE_USERNAME, SNOWFLAKE_ACCOUNT and SNOWFLAKE_PASSWORD environment variables are set.");
}

// create a replacement behavior that doesn't use backreferences since they will cause problems with our javascript blobs
String.prototype.replaceString = function(findNoRegex, replaceNoRegex) {
  return this.split(findNoRegex).join(replaceNoRegex);
}

const templates = [
  {
    name: "_QTAG",
    query: `
      CREATE OR REPLACE FUNCTION 
        "{{database}}"."{{schema}}"._QTAG(query string, allcomments boolean, allattributes boolean)
        RETURNS variant
        LANGUAGE JAVASCRIPT
      AS
      $$
      var initialize = function() {
        {{parser}}
        parse = qtag;
      };
      
      // we execute the initialization code once to take advantage of Snowflake's reuse of V8 isolates.
      if (typeof(initialized) === "undefined") {
        initialize();
        initialized = true;
      }
      return parse(QUERY, ALLCOMMENTS, ALLATTRIBUTES);
      $$;
      `
  },
  {
    name: "QTAG(complex)",
    query: `
      CREATE OR REPLACE SECURE FUNCTION 
        "{{database}}"."{{schema}}".QTAG(query string, allcomments boolean, allattributes boolean)
        COPY GRANTS 
        RETURNS ARRAY
        AS
      $$
        case when length(query) > 100000 then null 
        else "{{database}}"."{{schema}}"._QTAG(query, allcomments, allattributes)::ARRAY
        end
      $$;`
  } ,
  {
    name: "QTAG(complex)",
    query: `
      CREATE OR REPLACE SECURE FUNCTION 
        "{{database}}"."{{schema}}".QTAG(query string, allcomments boolean, allattributes boolean)
        COPY GRANTS 
        RETURNS ARRAY
        AS
      $$
        case when length(query) > 100000 then null 
        else "{{database}}"."{{schema}}"._QTAG(query, allcomments, allattributes)::ARRAY
        end
      $$;`
  },
  {
    name: "QTAG(simple)",
    query: `
      CREATE OR REPLACE SECURE FUNCTION 
        "{{database}}"."{{schema}}".QTAG(query string)
        COPY GRANTS 
        RETURNS ARRAY
        AS
      $$
        "{{database}}"."{{schema}}".QTAG(query, false, false)
      $$;`
  },
  {
    name: "QTAG_TABLE(complex)",
    query: `
      CREATE OR REPLACE SECURE FUNCTION 
        "{{database}}"."{{schema}}".QTAG_TABLE(query string, allcomments boolean, allattributes boolean)
        COPY GRANTS 
        RETURNS TABLE(key text, source text, type text, value text)
        LANGUAGE JAVASCRIPT
        AS
      $$
      {
        processRow: function f(row, rowWriter, context) {
          if (row.QUERY.length > 100000) {
            return; 
          }  
          {{parser}}
          let items = qtag(row.QUERY, row.ALLCOMMENTS, row.ALLATTRIBUTES);
          for (const i of items) {
            rowWriter.writeRow(i);
          }
        }
      }
      $$;
      `
  },
  {
    name: "QTAG_TABLE(simple)",
    query: `
      CREATE OR REPLACE SECURE FUNCTION 
        "{{database}}"."{{schema}}".QTAG_TABLE(query string)
        COPY GRANTS 
        RETURNS TABLE(key text, source text, type text, value text)
        LANGUAGE JAVASCRIPT
        AS
      $$
      {
        processRow: function f(row, rowWriter, context) {
          if (row.QUERY.length > 100000) {
            return; 
          }  
          {{parser}}
          let items = qtag(row.QUERY, false, false);
          for (const i of items) {
            rowWriter.writeRow(i);
          }
        }
      }
      $$;`
  }

]


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

for(const template of templates) {
  const sql = replaceDBAndSchema(template.query).replaceString('{{parser}}', data.replaceString("$$", "\\$\\$"));
  let running = false;
  let statement = connection.execute({
    sqlText: sql,
    complete: function(err, stmt, rows) {
      if (err) {
        exit(template.name + ': Failed to create: ' + err.message);
      } else {
        console.log(template.name + ':Created.');
      }
      running = true;
    }
  });
  while(running){}

}

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

