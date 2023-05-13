#include <stdio.h>
#define YEAR 2015
#define DAY 4
#include "../shared_2015.h"
  
unsigned char digest[16];

// https://stackoverflow.com/a/8389763
static inline void md5_leading_zeros(const char *s, int32_t length, unsigned char *out) {
  MD5_CTX c;
  MD5_Init(&c);
  MD5_Update(&c, s, length);
  MD5_Final(out, &c);
}

int32_t main(int32_t argc, char **argv) {
  struct input_file file = get_input_file();
  // Remove trailing newline
  file.bytes[file.len-1] = '\0';

  int32_t i = 0;
  int32_t suffix_1 = 0;
  int32_t suffix_2 = 0;
  uint64_t l = file.len+8;
  char *buf = malloc(l);
  int prefix_len = snprintf(buf, l, "%s", file.bytes);
  uint64_t rl = l-prefix_len;

  while (!suffix_1 || !suffix_2) {
    int x = snprintf(buf+prefix_len, rl, "%d", i);

    md5_leading_zeros(buf, prefix_len+x, digest);

    int num_zeros = (4 * (!digest[0] && !digest[1])) + (digest[2] < 0x10) + (digest[2] == 0);
    if (!suffix_1 && num_zeros >= 5) {
        suffix_1 = i;
        part_1("%d", suffix_1);
    }
    if (!suffix_2 && num_zeros >= 6) {
        suffix_2 = i;
        part_2("%d", suffix_2);
    }

    i++;
  }

  if (suffix_1 == -1) {
    eprintf("Could not find answer for part 1, tried %d numbers\n", i);
  }
  
  if (suffix_2 == -1) {
    eprintf("Could not find answer for part 2, tried %d numbers\n", i);
  }

  free(buf);
  free_input_file(&file);
  return 0;
}
