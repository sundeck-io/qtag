import dbt from '../../declarations/dbt.yaml';
import hex from '../../declarations/hex.yaml';
import mode from '../../declarations/mode.yaml';
import sigma from '../../declarations/sigma.yaml';
import sundeck from '../../declarations/sundeck.yaml';
import metabase from '../../declarations/metabase.yaml';


export function NewExtractor(parser) {
  let decls = [dbt, hex, mode, sigma, sundeck, metabase];

  let e = new Extractor()
  e.parser = parser;
  e.declarations = decls;
  return e;
}

export default function extract(parser, query, arbitraryComments, arbitraryAttributes) {
  let e = NewExtractor(parser);
  return e.extract(query, arbitraryComments, arbitraryAttributes);
}

class Extractor {
  parser = null;
  declarations = [];

  extract(query, arbitraryComments, arbitraryAttributes) {
    let allTags = [];
    const comments = this.parser(query);
    for (const comment of comments) {
      let foundDeclaration = false;

      let isSelfDescribing = comment.prefix != null &&
          comment.prefix.length >= 5 && comment.prefix.toUpperCase().startsWith("QTAG");

      if (!isSelfDescribing) {
        // TODO: Avoid the linear traversal here. This is just v1 for simplicity.

        for (const declaration of this.declarations) {
          let tags = extractInner(declaration, comment, arbitraryComments || arbitraryAttributes)

          if (tags != null) {
            foundDeclaration = true
            allTags = allTags.concat(tags)
            break
          }
        }

        if (foundDeclaration || !arbitraryComments) {
          continue;
        }
      }

      // is selfdescribing or arbitrary
      let tags = [];
      let source = "other";
      if (isSelfDescribing) {
        let splitSource = splitOnce(comment.prefix)
        if (splitSource != null) {
          source = splitSource[1]
        }
      } else {
        if (comment.prefix != null) {
          source = comment.prefix
        }
      }

      for (const k in comment.metadata) {
        let keyName = k;
        let valueType = "UNKNOWN";

        if (isSelfDescribing) {
          let splitKey = splitOnce(k);
          if (splitKey != null) {
            keyName = splitKey[0];
            valueType = asType(splitKey[1]);
          }
        }

        tags.push({
          "SOURCE": source,
          "TYPE": valueType,
          "KEY": keyName,
          "VALUE": comment.metadata[k],
        })
      }

      allTags = allTags.concat(tags)
    }

    return allTags
  }
}

function asType(s) {
  switch (s.toUpperCase()) {
    case "TRACE":
    case "USER":
    case "DIMENSION":
    case "MEASURE":
    case "APPLICATION":
    case "UNDECLARED":
      return s.toUpperCase();
  }
  return "UNKNOWN"
}

function splitOnce(value) {
  let i = value.indexOf(':');
  if (i === -1) {
    return null;
  }
  return [value.slice(0, i), value.slice(i + 1)];
}

function findAttribute(declaration, name) {
  // init map on first call
  if (declaration.fm === undefined) {
    declaration.fm = new Map();
    for (const field of declaration.fields) {
      declaration.fm.set(field.name, field);
    }
  }
  return declaration.fm.get(name);
}

// extract the attributes from the comment if they match the expected identifier
function extractInner(d, comment, includeUnknownAttributes) {

  if (d.identifier?.prefix !== undefined) {
    if (d.identifier.prefix.toUpperCase() !== comment.prefix.toUpperCase()) {
      return null;
    }
  }

  if (d.identifier?.bodymatch !== undefined) {
    if (comment.body.indexOf(d.identifier.bodymatch) === -1) {
      return null;
    }
  }

  if (d.identifier?.fields !== undefined) {
    for (const field in d.identifier.fields) {
      if (d.identifier.fields[field] !== comment.metadata[field]) {
        return null;
      }
    }
  }

  // we matched by identifier.
  let tags = [];
  for (const key in comment.metadata) {
    let found = findAttribute(d, key);
    if (found != null) {
      tags.push({
        "SOURCE": d.name,
        "TYPE": found.type,
        "KEY": key,
        "VALUE": comment.metadata[key]
       });
    } else if (includeUnknownAttributes) {
      tags.push({
        "SOURCE": d.name,
        "TYPE": "UNDECLARED",
        "KEY": key,
        "VALUE": comment.metadata[key]
      });
    }
  }

  if (tags.length === 0) {
    return null;
  }

  return tags;
}
