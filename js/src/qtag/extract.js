import antlr4 from 'antlr4';
import CommentQLLexer from '../generated/CommentQLLexer.js';
import CommentQLParser from '../generated/CommentQLParser.js';
import CommentQLParserListener from '../generated/CommentQLListener.js';
import ParseTreeWalker from "antlr4/src/antlr4/tree/ParseTreeWalker.js";
import PredictionMode from "antlr4/src/antlr4/atn/PredictionMode.js";

class Listener extends CommentQLParserListener {
  comments = [];
  enterCommentBody(ctx) {
    try {
      const obj = JSON.parse(ctx.json.getText())
      this.comments.push({"prefix":ctx.prefix.getText(), "obj": obj})
    } catch(e) {
      // do nothing.
    }
  }
}

export default function parse(query) {
  const chars = new antlr4.InputStream(query);
  const lexer = new CommentQLLexer(chars);
  const tokens = new antlr4.CommonTokenStream(lexer);
  const parser = new CommentQLParser(tokens);
  parser._interp.predictionMode = PredictionMode.SLL;
  parser.buildParseTrees = true;
  const tree = parser.body();
  const walker = new ParseTreeWalker();
  const l = new Listener()
  walker.walk(l, tree);
  return l.comments.length == 0 ? null : l.comments;
}
