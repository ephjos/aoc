#include <ctype.h>
#define YEAR 2015
#define DAY 5

#include "../shared_2015.h"

#define A_OFFSET 97

const uint8_t IS_VOWEL[26] = { 1,0,0,0,1,0,0,0,1,0,0,0,0,0,1,0,0,0,0,0,1,0,0,0,0,0 };

struct part_1_state {
  uint32_t nice_count;
  uint32_t vowel_count;
  uint8_t has_double;
  uint8_t has_illegal;
};

struct part_2_state {
  uint32_t nice_count;
  uint8_t has_double_pair;
  uint8_t has_spaced_pair;
  uint8_t pairs[26][26];
};

int main(const int argc, const char *argv[]) {
  struct input_file file = get_input_file();

  uint32_t i = 0;
  char prev_prev_c = '\0';
  char prev_c = '\0';
  char c = '\0';

  struct part_1_state state_1 = { 0,0,0,0 };
  struct part_2_state state_2 = {
    .nice_count = 0,
    .has_double_pair = 0,
    .has_spaced_pair = 0,
    .pairs = {0},
  };

  while (i < file.len && (c = file.bytes[i++]) != EOF) {
    if (c == '\n') {
      state_1.nice_count += (state_1.vowel_count>=3) && 
        (state_1.has_double) && 
        (!state_1.has_illegal);
      state_1.vowel_count = 0;
      state_1.has_double = 0;
      state_1.has_illegal = 0;

      state_2.nice_count += (state_2.has_double_pair && 
          state_2.has_spaced_pair);
      state_2 = (struct part_2_state){
        .nice_count = state_2.nice_count,
          .has_double_pair = 0,
          .has_spaced_pair = 0,
          .pairs = {0},
      };

      prev_c = '\0';
      prev_prev_c = '\0';
      continue;
    } else {
      uint8_t n = c-A_OFFSET;
      state_1.vowel_count += IS_VOWEL[n];
      state_1.has_double = MAX(state_1.has_double, c == prev_c);
      state_1.has_illegal = MAX(state_1.has_illegal, prev_c == 'a' && c == 'b');
      state_1.has_illegal = MAX(state_1.has_illegal, prev_c == 'c' && c == 'd');
      state_1.has_illegal = MAX(state_1.has_illegal, prev_c == 'p' && c == 'q');
      state_1.has_illegal = MAX(state_1.has_illegal, prev_c == 'x' && c == 'y');

      if (isalpha(prev_c)) {
        uint8_t prev_n = prev_c-A_OFFSET;
        state_2.pairs[prev_n][n] += 1;

        if ((c != prev_c) || (c != prev_prev_c) || (state_2.pairs[prev_n][n] > 2)) {
          state_2.has_double_pair = MAX(state_2.has_double_pair, 
              state_2.pairs[prev_n][n]>1);
        }
        state_2.has_spaced_pair = MAX(state_2.has_spaced_pair, 
            c == prev_prev_c);
      }
    }

    prev_prev_c = prev_c;
    prev_c = c;
  }

  part_1("%d", state_1.nice_count);
  part_2("%d", state_2.nice_count);

  free_input_file(&file);
  return 0;
}
