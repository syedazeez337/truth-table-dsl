grammar TruthExpr;

// ---- Parser rules ----
expr
    : '!' expr                     # Not
    | expr '&' expr                # And
    | expr '|' expr                # Or
    | '(' expr ')'                 # Parens
    | TRUE                         # TrueLit
    | FALSE                        # FalseLit
    | IDENT                        # Var
    ;

// ---- Lexer rules ----
TRUE  : [Tt] [Rr] [Uu] [Ee] ;
FALSE : [Ff] [Aa] [Ll] [Ss] [Ee] ;
IDENT : [A-Za-z_] [A-Za-z_0-9]* ;

WS    : [ \t\r\n]+ -> skip ;
LINE_COMMENT : '//' ~[\r\n]* -> skip ;

// Optional: single-quoted strings (not used now, but future-proof)
fragment ESC : '\\' . ;
