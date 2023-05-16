#include <ctype.h>
#include <string.h>
#define YEAR 2015
#define DAY 6

#include "../shared_2015.h"

enum TokenType {
  TOKEN_ON,
  TOKEN_OFF,
  TOKEN_TOGGLE,
  TOKEN_POINT,
  TOKEN_NEWLINE
};

struct Token {
  enum TokenType type;
  union {
    struct point value;
  };
};

int main(const int argc, const char *argv[]) {
  struct input_file file = get_input_file();

  uint32_t cap = 1024;
  uint32_t len = 0;
  struct Token *tokens = malloc(sizeof(struct Token) * cap);

  const uint32_t BUF_SIZE = 12;
  char buf[BUF_SIZE];
  #define CLEAR_BUF(buf) for (uint32_t b = 0; b < BUF_SIZE; b++) { buf[b] = '\0'; } buf_i = 0

  uint32_t buf_i = 0;

  uint32_t i = 0;
  char c;
  while (i < file.len) {
    c = file.bytes[i];
    if (isalpha(c)) {
      while (isalpha(c)) {
        c = file.bytes[i++];
        buf[buf_i++] = c;
      }

      i--;
      buf[--buf_i] = '\0';

      if (strcmp(buf, "on") == 0) {
        tokens[len++] = (struct Token){
          .type = TOKEN_ON,
        };
      } else if (strcmp(buf, "off") == 0) {
        tokens[len++] = (struct Token){
          .type = TOKEN_OFF,
        };
      } else if (strcmp(buf, "toggle") == 0) {
        tokens[len++] = (struct Token){
          .type = TOKEN_TOGGLE,
        };
      }

      CLEAR_BUF(buf);
    } else if (isdigit(c)) {
      struct point p = {0,0};
      while (isdigit(c)) {
        c = file.bytes[i++];
        buf[buf_i++] = c;
      }

      buf[--buf_i] = '\0';
      p.x = atoi(buf);

      CLEAR_BUF(buf);

      c = file.bytes[i];
      while (isdigit(c)) {
        c = file.bytes[i++];
        buf[buf_i++] = c;
      }

      i--;
      buf[--buf_i] = '\0';
      p.y = atoi(buf);

      CLEAR_BUF(buf);

      tokens[len++] = (struct Token){
        .type = TOKEN_POINT,
          .value = p,
      };
      
    } else if (c == '\n') {
      tokens[len++] = (struct Token){
        .type = TOKEN_NEWLINE,
      };
      i++;
    } else if (isspace(c)) {
      i++;
    } else if (c == ',') {
      i++;
    }

    if (len == cap) {
      cap <<= 2;
      tokens = realloc(tokens, sizeof(struct Token) * cap);
    }
  }

  i = 0;
  uint8_t op = 0;
  struct point ps[2];
  uint8_t pi = 0;

  const uint32_t D = 1000;
  uint32_t lights[D][D];
  for (uint32_t y = 0; y < D; y++) {
    for (uint32_t x = 0; x < D; x++) {
      lights[y][x] = 0;
    }
  }

  uint32_t lights_2[D][D];
  for (uint32_t y = 0; y < D; y++) {
    for (uint32_t x = 0; x < D; x++) {
      lights_2[y][x] = 0;
    }
  }

  while (i < len) {
    struct Token t = tokens[i];

    if (t.type == TOKEN_NEWLINE) {
      for (uint32_t y = (uint32_t)ps[0].y; y <= (uint32_t)ps[1].y; y++) {
        for (uint32_t x = (uint32_t)ps[0].x; x <= (uint32_t)ps[1].x; x++) {
          switch (op) {
            case TOKEN_ON:
              lights[y][x] = 1;
              lights_2[y][x] += 1;
              break;
            case TOKEN_OFF:
              lights[y][x] = 0;
              if (lights_2[y][x] > 0) {
                lights_2[y][x] -= 1;
              }
              break;
            case TOKEN_TOGGLE:
              lights[y][x] = !lights[y][x];
              lights_2[y][x] += 2;
              break;
            default:
              __builtin_unreachable();
          }
        }
      }

      op = 0;
      pi = 0;
    } else if (t.type == TOKEN_POINT) {
      ps[pi++] = t.value;
    } else {
      op = t.type;
    }

    i++;
  }

  uint64_t total_lit = 0;
  uint64_t total_brightness = 0;
  for (uint32_t y = 0; y < D; y++) {
    for (uint32_t x = 0; x < D; x++) {
      total_lit += lights[y][x];
      total_brightness += lights_2[y][x];
    }
  }

  part_1("%"PRIu64, total_lit);
  part_2("%"PRIu64, total_brightness);

  free(tokens);
  free_input_file(&file);
  return 0;
}
