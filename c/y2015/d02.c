#define YEAR 2015
#define DAY 2

#include "../shared_2015.h"

enum TokenType {
  TOKEN_NUM,
  TOKEN_X,
  TOKEN_NEWLINE
};

typedef struct Token {
  enum TokenType type;
  union {
    uint32_t value;
  };
} Token;

int compare_Token(Token *a, Token *b) {
  return (a->type > b->type) - (a->type < b->type);
}

DEFINE_LIST_FOR(Token, _)

int main(const int argc, const char *argv[]) {
  struct input_file file = get_input_file();

  Token_list_t tokens = Token_list_create();

  uint32_t i = 0;
  while (i < file.len){
    char c = file.bytes[i];
    if (isdigit(c)) {
      uint32_t buf_i = 0;
      char num_buf[12];
      while (isdigit(c)) {
        c = file.bytes[i++];
        num_buf[buf_i++] = c;
      }
      i--;

      Token_list_add(
          &tokens, 
          (Token){
            .type = TOKEN_NUM,
              .value = atoi(num_buf),
          }
      );
    } else if (c == 'x') {
      Token_list_add(
          &tokens, 
          (Token){
            .type = TOKEN_X,
          }
      );
      i++;
    } else if (c == '\n') {
      Token_list_add(
          &tokens, 
          (Token){
            .type = TOKEN_NEWLINE,
          }
      );
      i++;
    } else {
      i++;
    }
  }

  uint32_t total_paper = 0;
  uint32_t total_ribbon = 0;
  i = 0;
  while (i < list_size(&tokens)) {
    Token token = Token_list_get(&tokens, i++);
    if (token.type == TOKEN_NUM) {
      uint32_t l = token.value;

      token = Token_list_get(&tokens, i++);
      assert(token.type == TOKEN_X);

      token = Token_list_get(&tokens, i++);
      assert(token.type == TOKEN_NUM);
      uint32_t w = token.value;

      token = Token_list_get(&tokens, i++);
      assert(token.type == TOKEN_X);

      token = Token_list_get(&tokens, i++);
      assert(token.type == TOKEN_NUM);
      uint32_t h = token.value;

      token = Token_list_get(&tokens, i++);
      assert(token.type == TOKEN_NEWLINE);

      uint32_t area_0 = 2*l*w;
      uint32_t area_1 = 2*w*h;
      uint32_t area_2 = 2*h*l;

      uint32_t dims[3] = { w, l, h };
      qsort_uint32_t(dims, 3);

      total_paper += area_0 + area_1 + area_2 + (dims[0]*dims[1]);

      total_ribbon += (dims[0] * dims[1] * dims[2]) + ((2*dims[0])+(2*dims[1]));
    }
  }

  part_1("%d", total_paper);
  part_2("%d", total_ribbon);

  list_free(&tokens);
  free_input_file(&file);
  return 0;
}
