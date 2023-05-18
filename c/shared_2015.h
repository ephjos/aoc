//
// This header file is used by defining the YEAR and DAY macros as the integers
// that correspond to the problem date.
//

#ifndef __SHARED_2015__
#define __SHARED_2015__

#ifndef YEAR
#error YEAR is not defined // Forces compilation to fail if YEAR is not defined
#define YEAR 0 // Makes the rest of this code pass checks in isolation, never
               // hit since this branch exits before this statement
#endif

#ifndef DAY
#error DAY is not defined // Forces compilation to fail if DAY is not defined
#define DAY 0 // Makes the rest of this code pass checks in isolation, never
              // hit since this branch exits before this statement
#endif

// =============================================================================
// Includes
// =============================================================================
#include <assert.h>
#include <ctype.h>
#include <inttypes.h>
#include <limits.h>
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>

// MD5 hashes
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wunknown-pragmas"
#pragma GCC diagnostic ignored "-Wdeprecated-declarations"
#if defined(__APPLE__)
#  define COMMON_DIGEST_FOR_OPENSSL
#  include <CommonCrypto/CommonDigest.h>
#  define SHA1 CC_SHA1
// https://stackoverflow.com/a/8389763
unsigned char *MD5(
    const unsigned char *s, 
    int32_t length, 
    unsigned char *out) {
  MD5_CTX c;
  MD5_Init(&c);
  MD5_Update(&c, s, length);
  MD5_Final(out, &c);
  return NULL;
}
#else
#  include <openssl/md5.h>
#endif
#pragma GCC diagnostic pop

// =============================================================================
// Defines
// =============================================================================

#ifndef NUM_THREADS
#define NUM_THREADS 2
#endif

// =============================================================================
// Macros
// =============================================================================
#define eprintf( format, ... ) \
    fprintf(stderr, format, __VA_ARGS__) \

#define part_1( format, ... ) \
    fprintf(stdout, "%d.%02d.1: " format "\n", YEAR, DAY, __VA_ARGS__) \

#define part_2( format, ... ) \
    fprintf(stdout, "%d.%02d.2: " format "\n", YEAR, DAY, __VA_ARGS__) \

#define DDUMP( fmt, ... ) \
    fprintf(stderr, "%s:%d:%s(): " fmt "\n", __FILE__, \
        __LINE__, __func__, ##__VA_ARGS__); \

#define DUMP( fmt, val ) \
    printf("%s = " fmt "\n", #val, val);

#undef MAX
#define MAX( a, b ) ((a) > (b) ? (a) : (b))

#undef MIN
#define MIN( a, b ) ((a) > (b) ? (b) : (a))

#undef ABS
#define ABS( x ) (((x) < 0) ? (-x) : (x))

#define IS_BETWEEN( x, a, b ) \
    ((unsigned char)((x) >= (a) && (x) <= (b)))

#define XIS_BETWEEN( x, a, b ) \
    ((unsigned char)((x) > (a) && (x) < (b)))

// =============================================================================
// List template
// =============================================================================

#define LIST_INITIAL_CAP 1024

#define list_free(list) free((list)->arr)
#define list_size(list) ((list)->len)
#define list_is_empty(list) ((list)->len == 0)

#define DEFINE_LIST_FOR(T, ...)                                                \
  typedef struct T##_list_t T##_list_t;                                        \
  typedef struct T##_list_t {                                                  \
    uint32_t len;                                                              \
    uint32_t cap;                                                              \
    T *arr;                                                                    \
  } T##_list_t;                                                                \
                                                                               \
  T##_list_t T##_list_create() {                                               \
    return (T##_list_t) {                                                      \
      .len = 0,                                                                \
      .cap = LIST_INITIAL_CAP,                                                 \
      .arr = (T*)calloc(sizeof(T), LIST_INITIAL_CAP),                          \
    };                                                                         \
  }                                                                            \
                                                                               \
  void T##_list_add(T##_list_t *list, T item) {                                \
    list->arr[list->len++] = item;                                             \
                                                                               \
    if (list->len == list->cap) {                                              \
      assert(!(list->cap & 0xF0000000));                                       \
      list->cap <<= 1;                                                         \
      list->arr = (T*)realloc(list->arr, sizeof(T)*list->cap);                 \
    }                                                                          \
  }                                                                            \
                                                                               \
  void T##_list_insert(T##_list_t *list, uint32_t index, T item) {             \
    if (index <= list->len) {                                                  \
      list->len++;                                                             \
      if (list->len == list->cap) {                                            \
        assert(!(list->cap & 0xF0000000));                                     \
        list->cap <<= 1;                                                       \
        list->arr = (T*)realloc(list->arr, sizeof(T)*list->cap);               \
      }                                                                        \
                                                                               \
      memmove(&list->arr[index+1],                                             \
          &list->arr[index],                                                   \
          (list->len-index-1)*sizeof(T));                                      \
      list->arr[index] = item;                                                 \
      return;                                                                  \
    }                                                                          \
                                                                               \
    fprintf(stderr,                                                            \
        "Cannot insert at index %u outside list length %u\n",                  \
        index,                                                                 \
        list->len);                                                            \
    exit(1);                                                                   \
  }                                                                            \
                                                                               \
  void T##_list_set(T##_list_t *list, uint32_t index, T item) {                \
    if (index < list->len) {                                                   \
      list->arr[index] = item;                                                 \
      return;                                                                  \
    }                                                                          \
                                                                               \
    fprintf(stderr,                                                            \
        "Cannot set index %u outside list length %u\n",                        \
        index,                                                                 \
        list->len);                                                            \
    exit(1);                                                                   \
  }                                                                            \
                                                                               \
  T T##_list_get(T##_list_t *list, uint32_t index) {                           \
    if (index < list->len) {                                                   \
      return list->arr[index];                                                 \
    }                                                                          \
                                                                               \
    fprintf(stderr,                                                            \
        "Cannot get index %u outside list length %u\n",                        \
        index,                                                                 \
        list->len);                                                            \
    exit(1);                                                                   \
  }                                                                            \
                                                                               \
  int32_t T##_list_index_of(T##_list_t *list, T needle) {                      \
    for (uint32_t i = 0; i < list->len; i++) {                                 \
      if (compare_##T(&list->arr[i], &needle) == 0) {                          \
        return i;                                                              \
      }                                                                        \
    }                                                                          \
                                                                               \
    return -1;                                                                 \
  }                                                                            \
                                                                               \
  void T##_list_remove(T##_list_t *list, uint32_t index) {                     \
    if (index < list->len) {                                                   \
      list->len--;                                                             \
                                                                               \
      memmove(&list->arr[index],                                               \
          &list->arr[index+1],                                                 \
          (list->len-index)*sizeof(T));                                        \
      return;                                                                  \
    }                                                                          \
                                                                               \
    fprintf(stderr,                                                            \
        "Cannot remove index %u outside list length %u\n",                     \
        index,                                                                 \
        list->len);                                                            \
    exit(1);                                                                   \
  }                                                                            \
                                                                               \
  void T##_list_delete(T##_list_t *list, T needle) {                           \
    int32_t index = T##_list_index_of(list, needle);                           \
    if (index == -1) {                                                         \
      fprintf(stderr, "Cannot remove item\n");                                 \
      exit(1);                                                                 \
    }                                                                          \
                                                                               \
    T##_list_remove(list, index);                                              \
  }                                                                            \

// =============================================================================
// Hash Table template
// =============================================================================

#define HT_INITIAL_CAP_SHIFT 10
#define HT_MAX_LOAD 0.65
#define HT_MIN_LOAD_FACTOR 0.25
#define ht_free(ht) free((ht)->arr)

// djb2: http://www.cse.yorku.ca/~oz/hash.html
uint64_t hash_string(const char *str) {
  uint64_t h = 5381;
  int c;

  while ((c = *str++)) {
    h = ((h << 5) + h) + c;
  }

  return h;
}

// https://en.wikipedia.org/wiki/Hash_function
// https://probablydance.com/2018/06/16/fibonacci-hashing-the-optimization-that-the-world-forgot-or-a-better-alternative-to-integer-modulo/
#define A_16 40503u
#define A_32 2654435769u
#define A_48 173961102589771lu
#define A_64 11400714819323198485llu
uint64_t fibonacci_slot(uint64_t hash, uint64_t bit_slots) {
  hash ^= hash >> (64-bit_slots); 
  return (A_64*hash) >> (64 - bit_slots);
}

#define DEFINE_HT_FOR(T, ...)                                                  \
typedef struct he_##T##_t he_##T##_t;                                          \
typedef struct he_##T##_t {                                                    \
  uint8_t used;                                                                \
  uint8_t rem;                                                                 \
  uint64_t key;                                                                \
  uint64_t hash;                                                               \
  T value;                                                                     \
} he_##T##_t;                                                                  \
                                                                               \
typedef struct ht_##T##_t ht_##T##_t;                                          \
typedef struct ht_##T##_t {                                                    \
  uint64_t cnt;                                                                \
  uint64_t rem;                                                                \
  uint64_t cap;                                                                \
  uint64_t bit;                                                                \
  double alpha;                                                                \
  double alpha_max;                                                            \
  he_##T##_t *arr;                                                             \
} ht_##T##_t;                                                                  \
                                                                               \
ht_##T##_t ht_##T##_create() {                                                 \
  return (ht_##T##_t) {                                                        \
    .cnt = 0,                                                                  \
    .cap = 1 << HT_INITIAL_CAP_SHIFT,                                          \
    .bit = HT_INITIAL_CAP_SHIFT,                                               \
    .arr = (he_##T##_t*)calloc(1 << HT_INITIAL_CAP_SHIFT, sizeof(he_##T##_t)), \
  };                                                                           \
}                                                                              \
                                                                               \
void ht_##T##_put_core(ht_##T##_t *ht, uint64_t key, T value) {                \
  uint64_t slot = fibonacci_slot(key, ht->bit);                                \
  uint64_t i = slot;                                                           \
                                                                               \
  do {                                                                         \
    if (!ht->arr[i].used ||                                                    \
        ht->arr[i].rem ||                                                      \
        compare_##T(&ht->arr[i].key, &key) == 0) {                             \
      ht->arr[i] = (he_##T##_t) {                                              \
        .used = 1,                                                             \
        .key = key,                                                            \
        .hash = slot,                                                          \
        .value = value,                                                        \
      };                                                                       \
      ht->cnt++;                                                               \
      ht->alpha = (double)(ht->cnt+ht->rem) / ht->cap;                         \
      ht->alpha_max = MAX(ht->alpha_max, ht->alpha);                           \
                                                                               \
      return;                                                                  \
    }                                                                          \
  } while ((i = (i+1) % ht->cap) != slot);                                     \
                                                                               \
  fprintf(stderr, "Table full   \n");                                          \
  exit(1);                                                                     \
}                                                                              \
                                                                               \
static inline void ht_##T##_resize(ht_##T##_t *ht) {                           \
  if (ht->alpha > HT_MAX_LOAD) {                                               \
    ht->bit += 1;                                                              \
  } else if (ht->alpha < (ht->alpha_max/4)) {                                  \
    ht->bit -= 1;                                                              \
  } else {                                                                     \
    return;                                                                    \
  }                                                                            \
                                                                               \
  uint64_t old_cap = ht->cap;                                                  \
  uint64_t old_cnt = ht->cnt;                                                  \
  he_##T##_t *old_arr = (he_##T##_t*)calloc(old_cap, sizeof(he_##T##_t));      \
  memcpy(old_arr, ht->arr, ht->cap*sizeof(he_##T##_t));                        \
                                                                               \
  ht->cap = 1 << ht->bit;                                                      \
  ht->arr = (he_##T##_t*)calloc(ht->cap, sizeof(he_##T##_t));                  \
                                                                               \
  for (uint64_t i = 0; i < old_cap; i++) {                                     \
    if (old_arr[i].used) {                                                     \
      ht_##T##_put_core(ht, old_arr[i].key, old_arr[i].value);                 \
    }                                                                          \
  }                                                                            \
                                                                               \
  ht->alpha = (double)(ht->cnt+ht->rem) / ht->cap;                             \
  ht->alpha_max = MAX(ht->alpha_max, ht->alpha);                               \
  ht->cnt = old_cnt;                                                           \
  ht->rem = 0;                                                                 \
                                                                               \
  free(old_arr);                                                               \
}                                                                              \
                                                                               \
void ht_##T##_put(ht_##T##_t *ht, uint64_t key, T value) {                     \
  ht_##T##_put_core(ht, key, value);                                           \
  ht_##T##_resize(ht);                                                         \
}                                                                              \
                                                                               \
T *ht_##T##_get(ht_##T##_t *ht, uint64_t key) {                                \
  uint64_t slot = fibonacci_slot(key, ht->bit);                                \
  uint64_t i = slot;                                                           \
                                                                               \
  do {                                                                         \
    if (ht->arr[i].used && compare_##T(&ht->arr[i].key, &key) == 0) {          \
      return &ht->arr[i].value;                                                \
    }                                                                          \
  } while ((i = (i+1) % ht->cap) != slot);                                     \
                                                                               \
  return NULL;                                                                 \
}                                                                              \
                                                                               \
void ht_##T##_remove(ht_##T##_t *ht, uint64_t key) {                           \
  uint64_t slot = fibonacci_slot(key, ht->bit);                                \
  uint64_t i = slot;                                                           \
                                                                               \
  do {                                                                         \
    if (ht->arr[i].used && compare_##T(&ht->arr[i].key, &key) == 0) {          \
      ht->arr[i].rem = 1;                                                      \
      ht->cnt--;                                                               \
      ht->rem++;                                                               \
      return;                                                                  \
    }                                                                          \
  } while ((i = (i+1) % ht->cap) != slot);                                     \
                                                                               \
  fprintf(stderr, "Cannot remove key %" PRIu64 ", not found   \n", key);       \
  exit(1);                                                                     \
}                                                                              \

// =============================================================================
// Number XMacros
// =============================================================================

// Lists all of the single-word name number types
#define NUMBER_TYPES(F) \
  F(char, _) \
  F(short, _) \
  F(int, _) \
  F(long, _) \
  F(float, _) \
  F(double, _) \
  F(uint8_t, _) \
  F(uint16_t, _) \
  F(uint32_t, _) \
  F(uint64_t, _) \
  F(int8_t, _) \
  F(int16_t, _) \
  F(int32_t, _) \
  F(int64_t, _) \

// Define a qsort compatible comparator for each number type
#define NUMBER_COMPARATOR(T, ...)                                              \
  int compare_##T(const void *a, const void *b) {                              \
    T cast_a = *((const T*)a);                                                 \
    T cast_b = *((const T*)b);                                                 \
    return (cast_a > cast_b) - (cast_a < cast_b);                              \
  }                                                                            \

NUMBER_TYPES(NUMBER_COMPARATOR)

  // Define a simple qsort wrapper for each number type
#define NUMBER_QSORT(T, ...)                                                   \
    void qsort_##T(T *arr, uint32_t nmemb) {                                   \
      qsort(arr, nmemb, sizeof(T), compare_##T);                               \
    }                                                                          \

NUMBER_TYPES(NUMBER_QSORT)

NUMBER_TYPES(DEFINE_LIST_FOR)

NUMBER_TYPES(DEFINE_HT_FOR)

// =============================================================================
// Enum macros
// =============================================================================
/*
   Allows creating a enum with enum to string mapping:

   #define COLORS(N, V) \
     N(Colors) \
     V(Red) \
     V(Blue) \
     V(Green) \
   
   DEFINE_ENUM(COLORS);
*/

#define E_SHOW_NAME(N) N
#define E_SHOW_NAME_COUNT(N) N##Count
#define E_SHOW_NAME_STRINGS(N) N##Strings
#define E_SHOW_VALUE(V) V,
#define E_VALUE_STRING(V) #V,
#define E_NOOP(X)
#define DEFINE_ENUM(E) \
enum E(E_SHOW_NAME, E_NOOP) { \
  E(E_NOOP, E_SHOW_VALUE) \
  E(E_SHOW_NAME_COUNT, E_NOOP), \
}; \
\
const char *E(E_SHOW_NAME_STRINGS,E_NOOP)[E(E_SHOW_NAME_COUNT, E_NOOP)] = { \
  E(E_NOOP, E_VALUE_STRING) \
}\
  

// =============================================================================
// Input File
// =============================================================================
#define INPUT_FILENAME_BUF_SIZE 24
  struct input_file {
    uint64_t len;
    char filename[INPUT_FILENAME_BUF_SIZE];
    char *bytes;
  };


void free_input_file(struct input_file *i);
inline void free_input_file(struct input_file *i) {
  free(i->bytes);
}

// Load an input file for a given day
struct input_file get_input_file();
inline struct input_file get_input_file() {
  struct input_file i;

  // Build the file path
  int32_t n = snprintf(
      i.filename, 
      INPUT_FILENAME_BUF_SIZE, 
      "../../inputs/y%d/d%02d", 
      YEAR, 
      DAY);
  if (n < 0) {
    eprintf("Could not write filename with year %d and day %d\n", YEAR, DAY);
    exit(1);
  } else if (n > INPUT_FILENAME_BUF_SIZE) {
    eprintf("Filename buffer was too small, had %d but needed %d", 
        INPUT_FILENAME_BUF_SIZE, 
        n);
    exit(1);
  }

  // Open the file for reading
  FILE *fp = fopen(i.filename, "r");
  if (fp == NULL) {
    eprintf("Could not open file \"%s\" for reading\n", i.filename);
    exit(1);
  }

  // Get the total size of the file (in bytes)
  fseek(fp, 0, SEEK_END);
  i.len = ftell(fp);
  fseek(fp, 0, SEEK_SET);

  // Load file into memory
  i.bytes = (char*)malloc(i.len+1);
  fread(i.bytes, i.len, 1, fp);
  i.bytes[i.len] = '\0';

  // Cleanup
  fclose(fp);

  return i;
}

// =============================================================================
// Points
// =============================================================================
struct point {
  int32_t x;
  int32_t y;
};

#define POINT_EQUALS(a, b) (a.x == b.x && a.y == b.y)

#endif
