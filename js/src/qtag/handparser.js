
export default function parseHand(query) {
  const q = Array.from(query);

  // all comments need at least 6 characters for us to care about them. Having the outside loop only scan at that level
  // means we never have to check boundary conditions during lookahead
  // --a{}\n (6)
  // //a{}\n (6)
  // /*a{}*/ (7)
  // --{"a":1}\n (10)
  const end = q.length - 6;
  const comments = [];
  outside: for (let i = 0; i < end; i++) {
    const char = q[i];

    { // single character checks.
      let quoteCharTemp = "\\";
      switch (char) {
          // these are sing
        case "\"":
          quoteCharTemp = "\"";
          //fallthrough
        case "`":
          //fallthrough
        case "'":
          // consume next character
          i++;
          // we need at least 6 characters plus single close to care.
          const end2 = end - 1;
          const quoteChar = quoteCharTemp;
          for (; i < end2; i++) {
            let char2 = q[i];
            switch (char2) {
              case quoteChar:
                if (q[i + 1] === char) {
                  // this is a quoted character, move forward one so we don't accidentally close the item.
                  i++;
                  continue;
                }
                // fallthrough
              case char:
                // finished this block
                continue outside;
            }
          }
          // we never found the end of this block, no more things could be found.
          continue;
        default:
          // we didn't find single char match, fall through to double character match.
          break;
      }
    }

    const twoChar = char + q[i + 1]; //safe given early end calculated above.

    let prefixEnd;
    let jsonStart;
    switch (twoChar) {

        // Single line comments.
      case "--":
        //fallthrough
      case "//": {
        // end at line endings.
        // consume the starting token
        i += 2;
        const prefixStart = i;
        const end2 = q.length
        for (; i < end2; i++) {
          const char2 = q[i];
          switch (char2) {
            case '{':
              prefixEnd = i;
              jsonStart = i;
              i++;
              endOfComment: for (; i < end2; i++) {
                const char3 = q[i];
                switch (char3) {
                  case '\n':
                    // fallthrough
                  case '\r':
                    break endOfComment;
                }
              }
              addToComment(comments, query.substring(prefixStart, prefixEnd), query.substring(prefixEnd, i));
              continue outside;
            case '\n':
            case '\r':
              // no json, not a jsoncomment, move on.
              continue outside;
          }
        }
        break;
      }

        // Dollar Dollar String
      case "$$": {
        i+=2;
        const end2 = q.length - 1
        for (; i < end2; i++) {
          if (q[i] === '$' && q[i+1] === '$') {
            break;
          }
        }
        i++;
        // end of dollar dollar string.
        continue;
      }


      case "/*": {
        // end at close tag.
        const endChar1 = "*";
        const endChar2 = "/";
        i += 2;
        const prefixStart = i;
        const end2 = q.length - 1 // need at least two characters to close this tag.
        for (; i < end2; i++) {
          const char2 = q[i];
          switch (char2) {

              // this could be the end of the sequence
            case endChar1: {
              if (q[i + 1] === endChar2) {
                i++;
                // this is a valid block of the desired type but not interesting
                continue outside;
              }
              break;
            }

              // this is the end of a json comment prefix
            case '{': {
              prefixEnd = i;
              jsonStart = i;
              i++;
              for (; i < end2; i++) {
                if (q[i] === endChar1 && q[i + 1] === endChar2) {
                  // this is a valid block of the desired type but not interesting
                  addToComment(comments, query.substring(prefixStart, prefixEnd), query.substring(prefixEnd, i));
                  i++;
                  continue outside;
                }
              }
              // we ran out of characters, this is an invalid comment string, ignore.
              continue outside;
            }
          }
        }
      }

        // we didn't find a double character match, let's restart at the next position
        // note that even though we used lookahead for the two character tokens above, we can't move ahead two on the outer loop
        // since the second character we saw in our two char token could actually match a single char token above.
    }

  }
  return comments;
}

function addToComment(comments, prefix, jsonText) {
  try {
    const obj = JSON.parse(jsonText)
    comments.push({"prefix": prefix.trim(), "metadata": obj, "body": jsonText,})
  } catch (e) {
    // do nothing.
  }
}
