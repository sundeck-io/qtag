import {NewExtractor} from "./extract.js";
import hand from './handparser.js'

export default function extract(query, arbitraryComments, arbitraryAttributes) {
  let e = NewExtractor(hand);
  return e.extract(query, arbitraryComments, arbitraryAttributes);
}
