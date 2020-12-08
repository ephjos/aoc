#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "../include/cge.h"

#define HT_SIZE 655212
#define HOLDS 512

typedef struct bag_t {
	int count;
	int* counts;
	char** holds;
} bag;

typedef struct bag_ht_t {
	int key_count;
	char** keys;
	bag** bags;
} bag_ht;

bag* init_bag()
{
	bag* b = malloc(sizeof(bag));
	b->count = 0;
	b->counts = malloc(sizeof(int));
	b->holds = malloc(sizeof(char*));
	return b;
}

void add_bag(bag* b, int c, char* s)
{
	b->count++;
	b->counts = realloc(b->counts, sizeof(int)*b->count);
	b->holds = realloc(b->holds, sizeof(char*)*b->count);
	b->counts[b->count-1] = c;
	b->holds[b->count-1] = s;
}

void free_bag(bag* b)
{
	for (int i = 0; i < b->count; i++) {
		free(b->holds[i]);
	}
	free(b->counts);
	free(b->holds);
	free(b);
}

bag_ht* init_bag_ht()
{
	bag_ht* bht = malloc(sizeof(bag_ht));
	bht->key_count = 0;
	bht->keys = calloc(1,sizeof(char*));
	bht->bags = calloc(1,sizeof(bag*)*HT_SIZE);
	return bht;
}

void free_bag_ht(bag_ht* bht)
{
	for (int i = 0; i < bht->key_count; i++) {
		free(bht->keys[i]);
	}
	for (int i = 0; i < HT_SIZE; i++) {
		if (bht->bags[i] != NULL) {
			free_bag(bht->bags[i]);
		}
	}
	free(bht->keys);
	free(bht->bags);
	free(bht);
}

// adler32
unsigned long hash(const char* buf, const int buf_length)
{
	const u_int8_t* buffer = (const u_int8_t*)buf;

	unsigned long s1=1;
	unsigned long s2=0;

	for (int i = 0; i < buf_length; i++) {
		s1 = (s1 + buffer[i]) % 65521;
		s2 = (s2 + s1) % 65521;
	}
	return (s2 << 16) | s1;
}

void insert_bag_ht(bag_ht* bht, char* key, bag* b)
{
	bht->key_count++;
	bht->keys = realloc(bht->keys, bht->key_count*sizeof(char*));
	bht->keys[bht->key_count-1] = key;
	long long h = hash(key, strlen(key)) % HT_SIZE;
	if (bht->bags[h]) ERR("Already set\n");
	bht->bags[h] = b;
}

bag* get_bag_ht(bag_ht* bht, char* key)
{
	long long h = hash(key, strlen(key)) % HT_SIZE;
	if (bht->bags[h]) {
		return bht->bags[h];
	}
	return NULL;
}

char* cat2(const char* a, const char* b)
{
	char* res = calloc(1,sizeof(char)*(strlen(a)+strlen(b)+1));
	strcat(res, a);
	strcat(res, b);
	return res;
}

int required_bags(bag_ht* bht, char* name) {
	bag* b;
	if ((b = get_bag_ht(bht, name)) == NULL) {
		return 0;
	}
	int res = 0;
	for (int i = 0; i < b->count; i++) {
		res += b->counts[i] + (b->counts[i] * required_bags(bht, b->holds[i]));
	}
	return res;
}

int main(int argc, char *argv[])
{
	int n = 0; int m = 0; int l;
	char** lines = load_file("./input", 0, &n);
	char* root_name;
	bag* b;
	bag_ht* bht = init_bag_ht();

	for (int i = 0; i < n; i++) {
		char** toks = split(lines[i], " ", &m);
		root_name = cat2(toks[0], toks[1]);
		if (strcmp(toks[4], "no") == 0) {
			ffree((void**)toks, m);
			free(root_name);
			continue;
		}
		l = 0;
		b = init_bag();
		for (int j = 4; j < m-2; j+=4) {
			add_bag(b, atoi(toks[j]), cat2(toks[j+1], toks[j+2]));
		}
		insert_bag_ht(bht, root_name, b);
		ffree((void**)toks, m);
	}

	DUMP("%d", required_bags(bht, "shinygold"));

	free_bag_ht(bht);
	ffree((void**)lines, n);

	return 0;
}
