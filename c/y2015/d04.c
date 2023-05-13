#define YEAR 2015
#define DAY 4
#include "../shared_2015.h"

int32_t main(int32_t argc, char **argv) {
  struct input_file file = get_input_file();
  // Remove trailing newline
  file.bytes[file.len-1] = '\0';

  int32_t i = 0;
  int32_t suffix_1 = -1;
  int32_t suffix_2 = -1;
  char buf[64] = {0};

  for (i = 0; i < INT_MAX; i++) {
    snprintf(buf, 64, "%s%d", file.bytes, i);
    char *hash = str2md5(buf, (int32_t)strlen(buf));

    if (
        suffix_1 == -1 &&
        hash[0] == '0' &&
        hash[1] == '0' &&
        hash[2] == '0' &&
        hash[3] == '0' &&
        hash[4] == '0'
        ) {
      suffix_1 = i;
      part_1("%d", suffix_1);
    }

    if (
        suffix_2 == -1 &&
        hash[0] == '0' &&
        hash[1] == '0' &&
        hash[2] == '0' &&
        hash[3] == '0' &&
        hash[4] == '0' &&
        hash[5] == '0'
        ) {
      suffix_2 = i;
      part_2("%d", suffix_2);
    }

    if (suffix_1 != -1 && suffix_2 != -1) {
      break;
    }
  }
  
  if (suffix_1 == -1) {
    eprintf("Could not find answer for part 1, tried %d numbers\n", i);
  }
  
  if (suffix_2 == -1) {
    eprintf("Could not find answer for part 2, tried %d numbers\n", i);
  }

  free_input_file(&file);
  return 0;
}
