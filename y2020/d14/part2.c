#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "cge.h"
#include "clist.h"

#define MEM_ISIZE 8192
#define MASK_LEN 36

void print_bin(ulong x)
{
	//ulong l = sizeof(x)*8;
	ulong l = 36;
	for (long i = l-1; i >= 0; i--) {
		printf("%lu", (x >> i) & 1);
	}
	printf("\n");
}

int r = 0;
void make_masks(char* mask, int l, int xs, char** res)
{
	//printf("%s\n", mask);
	if (xs == 0) {
		strcpy(res[r++], mask);
	}

	for (int i = l; i < MASK_LEN; i++) {
		if (mask[i] == 'X') {
			mask[i] = '0'; make_masks(mask, i+1, xs-1, res);
			mask[i] = '1'; make_masks(mask, i+1, xs-1, res);
			mask[i] = 'X';
		}
	}
}

ulong parse_mask(char* buf)
{
	ulong res = 0;

	for(int i = 0; i < MASK_LEN; i++) {
		res = (res | (buf[i] == '1')) << (i != MASK_LEN-1);
	}

	return res;
}

char** parse_all(char* line, ulong addr, int* rk)
{
	char buf[48] = {0};
	sscanf(line, "mask = %s\n", buf);
	int k = 0;
	for (int i = 0; i < MASK_LEN; i++) {
		k += buf[i] == 'X';
		if ((addr >> (MASK_LEN-i-1)) & 1 && buf[i] != 'X') buf[i] = '1';
	}
	*rk = (1<<k);

	char** bufs = calloc(1, sizeof(char*)*(*rk));
	for (int i = 0; i < *rk; i++) {
		bufs[i] = calloc(1, sizeof(char)*MASK_LEN+1);
	}

	r = 0;
	make_masks(buf, 0, k, bufs);
	return bufs;

	/*
	ulong* res = malloc(sizeof(ulong)*(*rk));
	for (int i = 0; i < *rk; i++) {
		res[i] = parse_mask(bufs[i]);
	}

	FFREE(bufs, *rk);
	return res;
	*/
}

// TODO: IMPLEMENT HASH TABLE WITH BUCKETS

#define HT_SIZE 32768
#define HOLDS 512

typedef struct __entry entry;
typedef struct __entry {
	char* k;
	ulong v;
	entry* next;
} entry;

typedef struct __ht {
	long key_count;
	char** keys;
	entry** items;
} ht;

entry* init_entry()
{
	entry* e = malloc(sizeof(entry));
	e->k = NULL;
	e->v = 0;
	e->next = NULL;
	return e;
}

void free_entry(entry* e)
{
	free(e->k);
	free(e);
}

ht* init_ht()
{
	ht* d = malloc(sizeof(ht));
	d->key_count = 0;
	d->keys = calloc(1,sizeof(char*));
	d->items = calloc(1,sizeof(entry*)*HT_SIZE);
	return d;
}

void free_ht(ht* d)
{
	for (int i = 0; i < d->key_count; i++) {
		free(d->keys[i]);
	}
	for (int i = 0; i < HT_SIZE; i++) {
		if (d->items[i] != NULL) {
			free_entry(d->items[i]);
		}
	}
	free(d->keys);
	free(d->items);
	free(d);
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

void insert_wht(ht* d, char* key, ulong val)
{
	entry* e = init_entry();
	e->k = key; e->v = val;

	long long h = hash(key, strlen(key)) % HT_SIZE;
	//long long h = parse_mask(key) % HT_SIZE;
	if (!d->items[h]) {
		d->items[h] = e;
		d->key_count++;
		d->keys = realloc(d->keys, d->key_count*sizeof(char*));
		d->keys[d->key_count-1] = key;
		return;
	}

	entry* prev = NULL;
	entry* curr = d->items[h];
	while (curr) {
		if (strcmp(curr->k, key) == 0) {
			curr->v = val;
			return;
		}
		prev = curr;
		curr = curr->next;
	}
	prev->next = e;
	d->key_count++;
	d->keys = realloc(d->keys, d->key_count*sizeof(char*));
	d->keys[d->key_count-1] = key;
}

entry* get_ht(ht* d, char* key)
{
	long long h = hash(key, strlen(key)) % HT_SIZE;
	if (d->items[h]) {
		entry* curr = d->items[h];
		while (curr) {
			if (strcmp(curr->k, key) == 0) {
				return curr;
			}
			curr = curr->next;
		}
	}
	return NULL;
}

int main(int argc, char *argv[])
{
	int ln = 0;
	char** lines;

	if (argc == 3 &&
			(strcmp(argv[1], "-f") == 0 ||
			 strcmp(argv[1], "--file") == 0)) {
		lines = load_file(argv[2], 0, &ln);
	} else {
		lines = load_file("./input", 0, &ln);
	}

	//ulong_list* mem = new_ulong_list(MEM_ISIZE);
	ht* mem = init_ht();
	char* mask;
	int k;
	char** masks = NULL;
	ulong addr, val;

	for (int i = 0; i < ln; i++) {
		if (lines[i][1] == 'a') { // mask
			mask = lines[i];
		} else { // mem
			sscanf(lines[i], "mem[%ld] = %ld\n", &addr, &val);
			masks = parse_all(mask, addr, &k);
			for (int i = 0; i < k; i++) {
				//printf("%lu\n", masks[i]);
				//print_bin(masks[i]);
				//set_ulong_list(mem, masks[i], val);
				insert_wht(mem, masks[i], val);
			}
			free(masks);
		}
	}

	//print_ulong_list(mem);

	ulong sum = 0;
	//for (int i = 0; i < mem->size; i++) {
		//sum += get_ulong_list(mem, i);
	//}

	entry* e = NULL;
	for (int i = 0; i < mem->key_count; i++) {
		e = get_ht(mem, mem->keys[i]);
		sum += e->v;
	}

	DUMP("%lu", sum);

	//free_ulong_list(mem);
	FFREE(lines, ln);
	return 0;
}
