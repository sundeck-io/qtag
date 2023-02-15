grammar QTag;

// Lexer
BT_IDENTIFIER: '`' ( ~'`' | '``' )* '`';
DQ_IDENTIFIER: '"' ( ~'"' | '""' )* '"';
SQ_LITERAL: '\'' ('\\' . | '\'\'' | ~['\\])* '\'';
DOLLAR_LITERAL: '$$' ( ~'"' | '""' )* '"';

DASH_COMMENT: '--';
OPENC: '/*';
CLOSEC: '*/';
C_COMMENT: '//';
NL: '\n';
CR: '\r';
SPACE: ('\t' | ' ');
CURLY_O: '{';

OTHER: .+?;

// Parser
body:
  (.*? comment)*
  .*?
  EOF;

comment:
  lineComment
  | multilineComment;

lineComment: (C_COMMENT | DASH_COMMENT) commentBody (NL | CR | EOF);

multilineComment: OPENC commentBody CLOSEC;

commentBody:
 SPACE*
 prefix=commentPrefix
 SPACE*
 json=jsonBody;

commentPrefix:
  ~(CURLY_O)*?
  ;

jsonBody:
  CURLY_O .*?
  ;
