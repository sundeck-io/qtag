// Generated from java-escape by ANTLR 4.11.1
// jshint ignore: start
import antlr4 from 'antlr4';
import QTagListener from './QTagListener.js';
const serializedATN = [4,1,13,75,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,
2,5,7,5,2,6,7,6,1,0,5,0,16,8,0,10,0,12,0,19,9,0,1,0,5,0,22,8,0,10,0,12,0,
25,9,0,1,0,5,0,28,8,0,10,0,12,0,31,9,0,1,0,1,0,1,1,1,1,3,1,37,8,1,1,2,1,
2,1,2,1,2,1,3,1,3,1,3,1,3,1,4,5,4,48,8,4,10,4,12,4,51,9,4,1,4,1,4,5,4,55,
8,4,10,4,12,4,58,9,4,1,4,1,4,1,5,5,5,63,8,5,10,5,12,5,66,9,5,1,6,1,6,5,6,
70,8,6,10,6,12,6,73,9,6,1,6,4,17,29,64,71,0,7,0,2,4,6,8,10,12,0,3,2,0,5,
5,8,8,1,1,9,10,1,0,12,12,75,0,23,1,0,0,0,2,36,1,0,0,0,4,38,1,0,0,0,6,42,
1,0,0,0,8,49,1,0,0,0,10,64,1,0,0,0,12,67,1,0,0,0,14,16,9,0,0,0,15,14,1,0,
0,0,16,19,1,0,0,0,17,18,1,0,0,0,17,15,1,0,0,0,18,20,1,0,0,0,19,17,1,0,0,
0,20,22,3,2,1,0,21,17,1,0,0,0,22,25,1,0,0,0,23,21,1,0,0,0,23,24,1,0,0,0,
24,29,1,0,0,0,25,23,1,0,0,0,26,28,9,0,0,0,27,26,1,0,0,0,28,31,1,0,0,0,29,
30,1,0,0,0,29,27,1,0,0,0,30,32,1,0,0,0,31,29,1,0,0,0,32,33,5,0,0,1,33,1,
1,0,0,0,34,37,3,4,2,0,35,37,3,6,3,0,36,34,1,0,0,0,36,35,1,0,0,0,37,3,1,0,
0,0,38,39,7,0,0,0,39,40,3,8,4,0,40,41,7,1,0,0,41,5,1,0,0,0,42,43,5,6,0,0,
43,44,3,8,4,0,44,45,5,7,0,0,45,7,1,0,0,0,46,48,5,11,0,0,47,46,1,0,0,0,48,
51,1,0,0,0,49,47,1,0,0,0,49,50,1,0,0,0,50,52,1,0,0,0,51,49,1,0,0,0,52,56,
3,10,5,0,53,55,5,11,0,0,54,53,1,0,0,0,55,58,1,0,0,0,56,54,1,0,0,0,56,57,
1,0,0,0,57,59,1,0,0,0,58,56,1,0,0,0,59,60,3,12,6,0,60,9,1,0,0,0,61,63,8,
2,0,0,62,61,1,0,0,0,63,66,1,0,0,0,64,65,1,0,0,0,64,62,1,0,0,0,65,11,1,0,
0,0,66,64,1,0,0,0,67,71,5,12,0,0,68,70,9,0,0,0,69,68,1,0,0,0,70,73,1,0,0,
0,71,72,1,0,0,0,71,69,1,0,0,0,72,13,1,0,0,0,73,71,1,0,0,0,8,17,23,29,36,
49,56,64,71];


const atn = new antlr4.atn.ATNDeserializer().deserialize(serializedATN);

const decisionsToDFA = atn.decisionToState.map( (ds, index) => new antlr4.dfa.DFA(ds, index) );

const sharedContextCache = new antlr4.PredictionContextCache();

export default class QTagParser extends antlr4.Parser {

    static grammarFileName = "java-escape";
    static literalNames = [ null, null, null, null, null, "'--'", "'/*'", 
                            "'*/'", "'//'", "'\\n'", "'\\r'", null, "'{'" ];
    static symbolicNames = [ null, "BT_BLOCK", "DQ_BLOCK", "SQ_BLOCK", "DD_BLOCK", 
                             "DASH_COMMENT", "OPENC", "CLOSEC", "C_COMMENT", 
                             "NL", "CR", "SPACE", "CURLY_O", "OTHER" ];
    static ruleNames = [ "body", "comment", "lineComment", "multilineComment", 
                         "commentBody", "commentPrefix", "jsonBody" ];

    constructor(input) {
        super(input);
        this._interp = new antlr4.atn.ParserATNSimulator(this, atn, decisionsToDFA, sharedContextCache);
        this.ruleNames = QTagParser.ruleNames;
        this.literalNames = QTagParser.literalNames;
        this.symbolicNames = QTagParser.symbolicNames;
    }

    get atn() {
        return atn;
    }



	body() {
	    let localctx = new BodyContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 0, QTagParser.RULE_body);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 23;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,1,this._ctx)
	        while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1) {
	                this.state = 17;
	                this._errHandler.sync(this);
	                var _alt = this._interp.adaptivePredict(this._input,0,this._ctx)
	                while(_alt!=1 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	                    if(_alt===1+1) {
	                        this.state = 14;
	                        this.matchWildcard(); 
	                    }
	                    this.state = 19;
	                    this._errHandler.sync(this);
	                    _alt = this._interp.adaptivePredict(this._input,0,this._ctx);
	                }

	                this.state = 20;
	                this.comment(); 
	            }
	            this.state = 25;
	            this._errHandler.sync(this);
	            _alt = this._interp.adaptivePredict(this._input,1,this._ctx);
	        }

	        this.state = 29;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,2,this._ctx)
	        while(_alt!=1 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1+1) {
	                this.state = 26;
	                this.matchWildcard(); 
	            }
	            this.state = 31;
	            this._errHandler.sync(this);
	            _alt = this._interp.adaptivePredict(this._input,2,this._ctx);
	        }

	        this.state = 32;
	        this.match(QTagParser.EOF);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	comment() {
	    let localctx = new CommentContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 2, QTagParser.RULE_comment);
	    try {
	        this.state = 36;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 5:
	        case 8:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 34;
	            this.lineComment();
	            break;
	        case 6:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 35;
	            this.multilineComment();
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	lineComment() {
	    let localctx = new LineCommentContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 4, QTagParser.RULE_lineComment);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 38;
	        _la = this._input.LA(1);
	        if(!(_la===5 || _la===8)) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	        this.state = 39;
	        this.commentBody();
	        this.state = 40;
	        _la = this._input.LA(1);
	        if(!(((((_la - -1)) & ~0x1f) == 0 && ((1 << (_la - -1)) & 3073) !== 0))) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	multilineComment() {
	    let localctx = new MultilineCommentContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 6, QTagParser.RULE_multilineComment);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 42;
	        this.match(QTagParser.OPENC);
	        this.state = 43;
	        this.commentBody();
	        this.state = 44;
	        this.match(QTagParser.CLOSEC);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	commentBody() {
	    let localctx = new CommentBodyContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 8, QTagParser.RULE_commentBody);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 49;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,4,this._ctx)
	        while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1) {
	                this.state = 46;
	                this.match(QTagParser.SPACE); 
	            }
	            this.state = 51;
	            this._errHandler.sync(this);
	            _alt = this._interp.adaptivePredict(this._input,4,this._ctx);
	        }

	        this.state = 52;
	        localctx.prefix = this.commentPrefix();
	        this.state = 56;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===11) {
	            this.state = 53;
	            this.match(QTagParser.SPACE);
	            this.state = 58;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 59;
	        localctx.json = this.jsonBody();
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	commentPrefix() {
	    let localctx = new CommentPrefixContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 10, QTagParser.RULE_commentPrefix);
	    var _la = 0; // Token type
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 64;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,6,this._ctx)
	        while(_alt!=1 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1+1) {
	                this.state = 61;
	                _la = this._input.LA(1);
	                if(_la<=0 || _la===12) {
	                this._errHandler.recoverInline(this);
	                }
	                else {
	                	this._errHandler.reportMatch(this);
	                    this.consume();
	                } 
	            }
	            this.state = 66;
	            this._errHandler.sync(this);
	            _alt = this._interp.adaptivePredict(this._input,6,this._ctx);
	        }

	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	jsonBody() {
	    let localctx = new JsonBodyContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 12, QTagParser.RULE_jsonBody);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 67;
	        this.match(QTagParser.CURLY_O);
	        this.state = 71;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,7,this._ctx)
	        while(_alt!=1 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1+1) {
	                this.state = 68;
	                this.matchWildcard(); 
	            }
	            this.state = 73;
	            this._errHandler.sync(this);
	            _alt = this._interp.adaptivePredict(this._input,7,this._ctx);
	        }

	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}


}

QTagParser.EOF = antlr4.Token.EOF;
QTagParser.BT_BLOCK = 1;
QTagParser.DQ_BLOCK = 2;
QTagParser.SQ_BLOCK = 3;
QTagParser.DD_BLOCK = 4;
QTagParser.DASH_COMMENT = 5;
QTagParser.OPENC = 6;
QTagParser.CLOSEC = 7;
QTagParser.C_COMMENT = 8;
QTagParser.NL = 9;
QTagParser.CR = 10;
QTagParser.SPACE = 11;
QTagParser.CURLY_O = 12;
QTagParser.OTHER = 13;

QTagParser.RULE_body = 0;
QTagParser.RULE_comment = 1;
QTagParser.RULE_lineComment = 2;
QTagParser.RULE_multilineComment = 3;
QTagParser.RULE_commentBody = 4;
QTagParser.RULE_commentPrefix = 5;
QTagParser.RULE_jsonBody = 6;

class BodyContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = QTagParser.RULE_body;
    }

	EOF() {
	    return this.getToken(QTagParser.EOF, 0);
	};

	comment = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(CommentContext);
	    } else {
	        return this.getTypedRuleContext(CommentContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof QTagListener ) {
	        listener.enterBody(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof QTagListener ) {
	        listener.exitBody(this);
		}
	}


}



class CommentContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = QTagParser.RULE_comment;
    }

	lineComment() {
	    return this.getTypedRuleContext(LineCommentContext,0);
	};

	multilineComment() {
	    return this.getTypedRuleContext(MultilineCommentContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof QTagListener ) {
	        listener.enterComment(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof QTagListener ) {
	        listener.exitComment(this);
		}
	}


}



class LineCommentContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = QTagParser.RULE_lineComment;
    }

	commentBody() {
	    return this.getTypedRuleContext(CommentBodyContext,0);
	};

	C_COMMENT() {
	    return this.getToken(QTagParser.C_COMMENT, 0);
	};

	DASH_COMMENT() {
	    return this.getToken(QTagParser.DASH_COMMENT, 0);
	};

	NL() {
	    return this.getToken(QTagParser.NL, 0);
	};

	CR() {
	    return this.getToken(QTagParser.CR, 0);
	};

	EOF() {
	    return this.getToken(QTagParser.EOF, 0);
	};

	enterRule(listener) {
	    if(listener instanceof QTagListener ) {
	        listener.enterLineComment(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof QTagListener ) {
	        listener.exitLineComment(this);
		}
	}


}



class MultilineCommentContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = QTagParser.RULE_multilineComment;
    }

	OPENC() {
	    return this.getToken(QTagParser.OPENC, 0);
	};

	commentBody() {
	    return this.getTypedRuleContext(CommentBodyContext,0);
	};

	CLOSEC() {
	    return this.getToken(QTagParser.CLOSEC, 0);
	};

	enterRule(listener) {
	    if(listener instanceof QTagListener ) {
	        listener.enterMultilineComment(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof QTagListener ) {
	        listener.exitMultilineComment(this);
		}
	}


}



class CommentBodyContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = QTagParser.RULE_commentBody;
        this.prefix = null; // CommentPrefixContext
        this.json = null; // JsonBodyContext
    }

	commentPrefix() {
	    return this.getTypedRuleContext(CommentPrefixContext,0);
	};

	jsonBody() {
	    return this.getTypedRuleContext(JsonBodyContext,0);
	};

	SPACE = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(QTagParser.SPACE);
	    } else {
	        return this.getToken(QTagParser.SPACE, i);
	    }
	};


	enterRule(listener) {
	    if(listener instanceof QTagListener ) {
	        listener.enterCommentBody(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof QTagListener ) {
	        listener.exitCommentBody(this);
		}
	}


}



class CommentPrefixContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = QTagParser.RULE_commentPrefix;
    }

	CURLY_O = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(QTagParser.CURLY_O);
	    } else {
	        return this.getToken(QTagParser.CURLY_O, i);
	    }
	};


	enterRule(listener) {
	    if(listener instanceof QTagListener ) {
	        listener.enterCommentPrefix(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof QTagListener ) {
	        listener.exitCommentPrefix(this);
		}
	}


}



class JsonBodyContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = QTagParser.RULE_jsonBody;
    }

	CURLY_O() {
	    return this.getToken(QTagParser.CURLY_O, 0);
	};

	enterRule(listener) {
	    if(listener instanceof QTagListener ) {
	        listener.enterJsonBody(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof QTagListener ) {
	        listener.exitJsonBody(this);
		}
	}


}




QTagParser.BodyContext = BodyContext; 
QTagParser.CommentContext = CommentContext; 
QTagParser.LineCommentContext = LineCommentContext; 
QTagParser.MultilineCommentContext = MultilineCommentContext; 
QTagParser.CommentBodyContext = CommentBodyContext; 
QTagParser.CommentPrefixContext = CommentPrefixContext; 
QTagParser.JsonBodyContext = JsonBodyContext; 
