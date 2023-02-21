import antlr4 from "antlr4";
import QTagLexer from "../generated/QTagLexer.js";
import QTagParser from "../generated/QTagParser.js";
import PredictionMode from "antlr4/src/antlr4/atn/PredictionMode.js";
import ParseTreeWalker from "antlr4/src/antlr4/tree/ParseTreeWalker.js";
import QTagListener from "../generated/QTagListener.js";

export default function parseAntlr(query) {
  const chars = new antlr4.InputStream(query);
  const lexer = new QTagLexer(chars);
  const tokens = new antlr4.CommonTokenStream(lexer);
  const parser = new QTagParser(tokens);
  parser._interp.predictionMode = PredictionMode.SLL;
  parser.buildParseTrees = true;
  const tree = parser.body();
  const walker = new ParseTreeWalker();
  const listener = new Listener();
  walker.walk(listener, tree);
  return listener.comments;
}

class Listener extends QTagListener {
  comments = [];

  enterCommentBody(ctx) {
    try {
      let prefix = ctx.prefix.getText()
      let jsonText = ctx.json.getText()
      const obj = JSON.parse(jsonText)
      this.comments.push({"prefix": prefix, "metadata": obj, "body": jsonText,})
    } catch (e) {
      // do nothing.
    }
  }
}
