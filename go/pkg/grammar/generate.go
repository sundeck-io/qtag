package grammar

//go:generate wget -nc https://www.antlr.org/download/antlr-4.11.1-complete.jar
//go:generate -command antlr java -Xmx500M -cp "./antlr-4.11.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool
//go:generate antlr -Dlanguage=Go -package generated ./QTag.g4 -Werror -o ./generated/
//go:generate antlr -Dlanguage=JavaScript ./QTag.g4 -Werror -o ../../../js/src/generated/
