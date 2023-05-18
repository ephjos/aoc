#define YEAR 2015
#define DAY 4
#include "../shared_2015.h"

#define DIGITS_INT_MAX 10

typedef struct thread_args_t {
  uint32_t start;
  uint32_t end;
  uint32_t suffix_5;
  uint32_t suffix_6;
} thread_args_t;


struct input_file file;

void *check_hashes(void *args) {
  thread_args_t *thread_args = (thread_args_t*)args;

  uint64_t l = file.len+DIGITS_INT_MAX;
  char *buf = malloc(l);
  unsigned char digest[16] = {0};
  int prefix_len = snprintf(buf, l, "%s", file.bytes);
  uint64_t rl = l-prefix_len;

  for (uint32_t i = thread_args->start; i < thread_args->end; i++) {
    int x = snprintf(buf+prefix_len, rl, "%d", i);
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wdeprecated-declarations"
    MD5((unsigned char*)buf, prefix_len+x, digest);
#pragma GCC diagnostic pop

    int num_zeros = (4 * (!digest[0] && !digest[1])) + 
                    (digest[2] < 0x10) + 
                    (digest[2] == 0);

    if (num_zeros >= 5) {
      thread_args->suffix_5 = MIN(thread_args->suffix_5, i);
    }

    if (num_zeros >= 6) {
      thread_args->suffix_6 = MIN(thread_args->suffix_6, i);
    }

    if (thread_args->suffix_5 != INT_MAX && thread_args->suffix_6 != INT_MAX)  {
      break;
    }
  }

  free(buf);
  return NULL;
}

int32_t main(int32_t argc, char **argv) {
  file = get_input_file();

  // Remove trailing newline
  file.bytes[file.len-1] = '\0';

  // Size of each block and last value to check
  const uint32_t LAST_SUFFIX = INT_MAX;
  const uint32_t BLOCK_SIZE = 131072;
                                  
  // Threads
  pthread_t thread_ids[NUM_THREADS];
  thread_args_t thread_args[NUM_THREADS] = { 0 };

  // Bounds on thread group
  uint32_t thread_group_start = 0;
  uint32_t thread_group_end = MIN(BLOCK_SIZE*NUM_THREADS, LAST_SUFFIX);

  // Accumulators for found values
  uint32_t final_suffix_5 = INT_MAX;
  uint32_t final_suffix_6 = INT_MAX;

  do { 
    // Kick off a thread for each block 
    for (uint32_t i = 0; i < NUM_THREADS; i++) {
      thread_args[i].start = thread_group_start+(i*BLOCK_SIZE);
      thread_args[i].end = thread_group_start+((i+1)*BLOCK_SIZE);
      thread_args[i].suffix_5 = INT_MAX;
      thread_args[i].suffix_6 = INT_MAX;
      pthread_create(&thread_ids[i], NULL, &check_hashes, (void*)&thread_args[i]);
    }

    // Gather threads, and take their result if it is better than the 
    // current value
    for (uint32_t i = 0; i < NUM_THREADS; i++) {
      pthread_join(thread_ids[i], NULL);
      final_suffix_5 = MIN(final_suffix_5, thread_args[i].suffix_5);
      final_suffix_6 = MIN(final_suffix_6, thread_args[i].suffix_6);
    }

    // If we find both values, we are done!
    if (final_suffix_5 != INT_MAX && final_suffix_6 != INT_MAX) {
      break;
    }

    // Otherwise, shift to the next groups of blocks
    thread_group_start += BLOCK_SIZE*NUM_THREADS;
    thread_group_end += BLOCK_SIZE*NUM_THREADS;
  } while (thread_group_end < LAST_SUFFIX);


  part_1("%d", final_suffix_5);
  part_2("%d", final_suffix_6);

  free_input_file(&file);
  return 0;
}
