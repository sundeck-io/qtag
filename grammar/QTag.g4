grammar QTag;

// Lexer
BT_BLOCK: '`' ( ~'`' | '``' )* '`';
DQ_BLOCK: '"' ( ~'"' | '""' )* '"';
SQ_BLOCK: '\'' ('\\' . | '\'\'' | ~['\\])* '\'';
DD_BLOCK: '$$' .*? '$$';

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


