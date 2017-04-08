#!/bin/bash

#mv lex.yy.c lex.yy.bak

sed -e 's/yy_switch_to_buffer/yy_switch_to_buffer_java/g' \
 -e 's/yy_delete_buffer/yy_delete_buffer_java/g' \
 -e 's/yy_flush_buffer/yy_flush_buffer_java/g' \
 -e 's/yy_create_buffer/yy_create_buffer_java/g' \
 -e 's/yyrestart/yyrestart_java/g' \
 -e 's/yylex/yylex_java/g' \
 -e 's/yypush_buffer_state/yypush_buffer_state_java/g' \
 -e 's/yypop_buffer_state/yypop_buffer_state_java/g' \
 -e 's/yy_scan_buffer/yy_scan_buffer_java/g' \
 -e 's/yy_scan_bytes/yy_scan_bytes_java/g' \
 -e 's/yy_scan_string/yy_scan_string_java/g' \
 -e 's/yyget_lineno/yyget_lineno_java/g' \
 -e 's/yyget_in/yyget_in_java/g' \
 -e 's/yyget_out/yyget_out_java/g' \
 -e 's/yyget_leng/yyget_leng_java/g' \
 -e 's/yyget_text/yyget_text_java/g' \
 -e 's/yyset_lineno/yyset_lineno_java/g' \
 -e 's/yyset_in/yyset_in_java/g' \
 -e 's/yyset_out/yyset_out_java/g' \
 -e 's/yyget_debug/yyget_debug_java/g' \
 -e 's/yyset_debug/yyset_debug_java/g' \
 -e 's/yylex_java_destroy/yylex_destroy_java/g' \
 -e 's/yyalloc/yyalloc_java/g' \
 -e 's/yyrealloc/yyrealloc_java/g' \
 -e 's/yyfree/yyfree_java/g' \
 -e 's/yywrap/yywrap_java/g' \
 -e 's/yylineno/yylineno_java/g' \
 -e 's/yyout/yyout_java/g' \
 -e 's/yy_flex_debug/yy_flex_debug_java/g' \
 -e 's/yyin/yyin_java/g' \
 -e 's/yy_load_buffer_state/yy_load_buffer_state_java/g' \
 -e 's/yy_init_buffer/yy_init_buffer_java/g' \
 lex.yy.c.bak > lex.yy.c