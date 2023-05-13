#define YEAR 2015
#define DAY 4
#include "../shared_2015.h"

#define NUM_THREADS 3

pthread_t thread_ids[NUM_THREADS];

// https://stackoverflow.com/a/8389763
static inline void md5_leading_zeros(const char *s, int32_t length, unsigned char *out) {
  MD5_CTX c;
  MD5_Init(&c);
  MD5_Update(&c, s, length);
  MD5_Final(out, &c);
}

uint32_t suffix_1 = 0;
uint32_t suffix_2 = 0;
pthread_mutex_t lock;

void *check_hashes(void *starting_i) {
  uint32_t i = *(uint32_t*)starting_i;

  struct input_file file = get_input_file();
  // Remove trailing newline
  file.bytes[file.len-1] = '\0';
  uint64_t l = file.len+8;
  char *buf = malloc(l);
  unsigned char digest[16] = {0};
  int prefix_len = snprintf(buf, l, "%s", file.bytes);
  uint64_t rl = l-prefix_len;

  while (!suffix_1 || !suffix_2) {
    int x = snprintf(buf+prefix_len, rl, "%d", i);

    md5_leading_zeros(buf, prefix_len+x, digest);

    int num_zeros = (4 * (!digest[0] && !digest[1])) + (digest[2] < 0x10) + (digest[2] == 0);
    if (!suffix_1 && num_zeros >= 5) {
      pthread_mutex_lock(&lock);
      if (!suffix_1 || i < suffix_1) {
        suffix_1 = i;
      }
      pthread_mutex_unlock(&lock);
    }
    if (!suffix_2 && num_zeros >= 6) {
      pthread_mutex_lock(&lock);
      if (!suffix_2 || i < suffix_2) {
        suffix_2 = i;
      }
      pthread_mutex_unlock(&lock);
    }

    i+=NUM_THREADS;
  }

  free(buf);
  free_input_file(&file);
  free(starting_i);
  return NULL;
}

int32_t main(int32_t argc, char **argv) {
  pthread_mutex_init(&lock, NULL); // TODO: check

  for (uint32_t i = 0; i < NUM_THREADS; i++) {
    uint32_t *starting_i = malloc(sizeof(uint32_t));
    *starting_i = i;
    pthread_create(&thread_ids[i], NULL, &check_hashes, (void*)starting_i);
  }

  for (uint32_t i = 0; i < NUM_THREADS; i++) {
    pthread_join(thread_ids[i], NULL);
  }

  part_1("%d", suffix_1);
  part_2("%d", suffix_2);

  pthread_mutex_destroy(&lock);
  return 0;
}
