#include <stdlib.h>
#define YEAR 2015
#define DAY 2

#include "../shared_2015.h"

enum TokenType {
  TOKEN_NUM,
  TOKEN_X,
  TOKEN_NEWLINE
};

struct Token {
  enum TokenType type;
  union {
    uint32_t value;
  };
};

int main(const int argc, const char *argv[]) {
  struct input_file file = get_input_file();

  uint32_t cap = 1024;
  uint32_t len = 0;
  struct Token *tokens = malloc(sizeof(struct Token) * cap);

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
      tokens[len++] = (struct Token){
        .type = TOKEN_NUM,
          .value = atoi(num_buf),
      };
    } else if (c == 'x') {
      tokens[len++] = (struct Token){
        .type = TOKEN_X,
      };
      i++;
    } else if (c == '\n') {
      tokens[len++] = (struct Token){
        .type = TOKEN_NEWLINE,
      };
      i++;
    } else {
      i++;
    }

    if (len == cap) {
      cap <<= 2;
      tokens = realloc(tokens, sizeof(struct Token) * cap);
    }
  }

  uint32_t total_paper = 0;
  uint32_t total_ribbon = 0;
  i = 0;
  while (i < len) {
    struct Token token = tokens[i++];
    if (token.type == TOKEN_NUM) {
      uint32_t l = token.value;

      token = tokens[i++];
      assert(token.type == TOKEN_X);

      token = tokens[i++];
      assert(token.type == TOKEN_NUM);
      uint32_t w = token.value;

      token = tokens[i++];
      assert(token.type == TOKEN_X);

      token = tokens[i++];
      assert(token.type == TOKEN_NUM);
      uint32_t h = token.value;

      token = tokens[i++];
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

  free(tokens);
  free_input_file(&file);
  return 0;
}
