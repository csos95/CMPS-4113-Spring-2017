#!/bin/bash

#mv lex.yy.c > lex.yy.bak

sed -e 's/yy_switch_to_buffer/yy_switch_to_buffer_cpp/g' \
 -e 's/yy_delete_buffer/yy_delete_buffer_cpp/g' \
 -e 's/yy_flush_buffer/yy_flush_buffer_cpp/g' \
 -e 's/yy_create_buffer/yy_create_buffer_cpp/g' \
 -e 's/yyrestart/yyrestart_cpp/g' \
 -e 's/yylex/yylex_cpp/g' \
 -e 's/yypush_buffer_state/yypush_buffer_state_cpp/g' \
 -e 's/yypop_buffer_state/yypop_buffer_state_cpp/g' \
 -e 's/yy_scan_buffer/yy_scan_buffer_cpp/g' \
 -e 's/yy_scan_bytes/yy_scan_bytes_cpp/g' \
 -e 's/yy_scan_string/yy_scan_string_cpp/g' \
 -e 's/yyget_lineno/yyget_lineno_cpp/g' \
 -e 's/yyget_in/yyget_in_cpp/g' \
 -e 's/yyget_out/yyget_out_cpp/g' \
 -e 's/yyget_leng/yyget_leng_cpp/g' \
 -e 's/yyget_text/yyget_text_cpp/g' \
 -e 's/yyset_lineno/yyset_lineno_cpp/g' \
 -e 's/yyset_in/yyset_in_cpp/g' \
 -e 's/yyset_out/yyset_out_cpp/g' \
 -e 's/yyget_debug/yyget_debug_cpp/g' \
 -e 's/yyset_debug/yyset_debug_cpp/g' \
 -e 's/yylex_cpp_destroy/yylex_destroy_cpp/g' \
 -e 's/yyalloc/yyalloc_cpp/g' \
 -e 's/yyrealloc/yyrealloc_cpp/g' \
 -e 's/yyfree/yyfree_cpp/g' \
 -e 's/yywrap/yywrap_cpp/g' \
 -e 's/yylineno/yylineno_cpp/g' \
 -e 's/yyout/yyout_cpp/g' \
 -e 's/yy_flex_debug/yy_flex_debug_cpp/g' \
 -e 's/yyin/yyin_cpp/g' \
 -e 's/yy_load_buffer_state/yy_load_buffer_state_cpp/g' \
 -e 's/yy_init_buffer/yy_init_buffer_cpp/g' \
 lex.yy.c.bak > lex.yy.c