#!/bin/bash

#mv lex.yy.c > lex.yy.bak

sed -e 's/yy_switch_to_buffer/yy_switch_to_buffer_c/g' \
 -e 's/yy_delete_buffer/yy_delete_buffer_c/g' \
 -e 's/yy_flush_buffer/yy_flush_buffer_c/g' \
 -e 's/yy_create_buffer/yy_create_buffer_c/g' \
 -e 's/yyrestart/yyrestart_c/g' \
 -e 's/yylex/yylex_c/g' \
 -e 's/yypush_buffer_state/yypush_buffer_state_c/g' \
 -e 's/yypop_buffer_state/yypop_buffer_state_c/g' \
 -e 's/yy_scan_buffer/yy_scan_buffer_c/g' \
 -e 's/yy_scan_bytes/yy_scan_bytes_c/g' \
 -e 's/yy_scan_string/yy_scan_string_c/g' \
 -e 's/yyget_lineno/yyget_lineno_c/g' \
 -e 's/yyget_in/yyget_in_c/g' \
 -e 's/yyget_out/yyget_out_c/g' \
 -e 's/yyget_leng/yyget_leng_c/g' \
 -e 's/yyget_text/yyget_text_c/g' \
 -e 's/yyset_lineno/yyset_lineno_c/g' \
 -e 's/yyset_in/yyset_in_c/g' \
 -e 's/yyset_out/yyset_out_c/g' \
 -e 's/yyget_debug/yyget_debug_c/g' \
 -e 's/yyset_debug/yyset_debug_c/g' \
 -e 's/yylex_c_destroy/yylex_destroy_c/g' \
 -e 's/yyalloc/yyalloc_c/g' \
 -e 's/yyrealloc/yyrealloc_c/g' \
 -e 's/yyfree/yyfree_c/g' \
 -e 's/yywrap/yywrap_c/g' \
 -e 's/yylineno/yylineno_c/g' \
 -e 's/yyout/yyout_c/g' \
 -e 's/yy_flex_debug/yy_flex_debug_c/g' \
 -e 's/yyin/yyin_c/g' \
 -e 's/yy_load_buffer_state/yy_load_buffer_state_c/g' \
 -e 's/yy_init_buffer/yy_init_buffer_c/g' \
 -e 's/yytext/yytext_c/g' \
 -e 's/yyleng/yyleng_c/g' \
 lex.yy.c.bak > lex.yy.c