#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "cge.h"

typedef struct __range range;
typedef struct __range {
	int l, h;
} range;

typedef struct __rule rule;
typedef struct __rule {
	range rs[2];
	char* label;
} rule;

int ticket_fits_rule(rule r, int* ticket, int n)
{
	for (int i = 0; i < n; i++) {
		if (IS_BETWEEN(ticket[i], r.rs[0].l, r.rs[0].h) ||
				IS_BETWEEN(ticket[i], r.rs[1].l, r.rs[1].h))
			return 1;
	}
	return 0;
}

#define ON_RANGE( r, x ) ((r.l <= x) && (x <= r.h))
int on_ranges(range* ranges, int n, int x)
{
	int res = 0;
	for (int i = 0; i < n; i++) {
		res |= ON_RANGE(ranges[i], x);
	}
	return res;
}

void print2a(int** arr, int n)
{
	for (int i = 0; i < n; i++) {
		for (int j = 0; j < n; j++) {
			printf("%d ", arr[i][j]);
		}
		printf("\n");
	}
}

int suma(int* arr, int n)
{
	int sum = 0;
	for (int i = 0; i < n; i++) {
		sum += arr[i];
	}
	return sum;
}

int get1ia(int* arr, int n)
{
	for (int i = 0; i < n; i++) {
		if (arr[i] == 1) return i;
	}
	return -1;
}

int sum2a(int** arr, int n)
{
	int sum = 0;
	for (int i = 0; i < n; i++) {
		for (int j = 0; j < n; j++) {
			sum += arr[i][j];
		}
	}
	return sum;
}

int main(int argc, char *argv[])
{
	int n = 0;
	char** lines;

	if (argc == 3 &&
			(strcmp(argv[1], "-f") == 0 ||
			 strcmp(argv[1], "--file") == 0)) {
		lines = load_file(argv[2], 0, &n);
	} else {
		lines = load_file("./input", 0, &n);
	}

	// Inits
	rule* rules;
	int* my_ticket;
	int** nearby_tickets;
	int** valid_tickets;
	char** toks;
	int i = 0, j, m, n_rules, n_tickets, n_valid = 0;

	// Count rules
	while (i < n && strcmp(lines[i], "") != 0) {
		i++;
	}
	n_rules = i;
	rules = calloc(1, sizeof(rule)*n_rules);
	int l1, h1, l2, h2;
	char buf[64] = {0};
	// Get rules
	for (int j = 0; j < n_rules; j++) {
		sscanf(lines[j], "%[a-z ]: %d-%d or %d-%d\n", buf, &l1, &h1, &l2, &h2);
		rules[j].label = strdup(buf);
		rules[j].rs[0] = (range){ .l = l1, .h = h1 };
		rules[j].rs[1] = (range){ .l = l2, .h = h2 };
	}

	// Read my ticket
	i+=2; // skip whitespace and text
	toks = split(lines[i], ",", &m);
	my_ticket = sstoi(toks, m);
	FFREE(toks, m);

	// Read my nearby tickets
	i+=3; // skip whitespace and text
	j = 0;
	n_tickets = n-i;
	nearby_tickets = calloc(1, sizeof(int*)*n_tickets);
	while (i < n && strcmp(lines[i], "") != 0) {
		toks = split(lines[i], ",", &m);
		nearby_tickets[j] = sstoi(toks, m);
		FFREE(toks, m);
		i++; j++;
	}

	// Remove bad tickets
	int res = 0;
	valid_tickets = calloc(n_tickets, sizeof(int*));
	for (int k = 0; k < n_tickets; k++) {
		for (int j = 0; j < m; j++) {
			res = 0;
			for (int l = 0; l < n_rules; l++) {
				res |= IS_BETWEEN(nearby_tickets[k][j], rules[l].rs[0].l, rules[l].rs[0].h) ||
					IS_BETWEEN(nearby_tickets[k][j], rules[l].rs[1].l, rules[l].rs[1].h);
			}
			if (!res) break;
		}
		if (res) {
			valid_tickets[n_valid] = malloc(sizeof(int)*m);
			memcpy(valid_tickets[n_valid], nearby_tickets[k], sizeof(int)*m);
			n_valid++;
		}
	}
	FFREE(nearby_tickets, n_tickets);
	valid_tickets = realloc(valid_tickets, n_valid*sizeof(int)*m);

	int** rules2cols = calloc(m, sizeof(int*));
	for (int i = 0; i < m; i++) {
		rules2cols[i] = calloc(m, sizeof(int));
	}

	for (int i = 0; i < n_rules; i++) {
		for (int j = 0; j < m; j++) {
			res = 1;
			for (int k = 0; k < n_valid; k++) {
				res &= IS_BETWEEN(valid_tickets[k][j], rules[i].rs[0].l, rules[i].rs[0].h) ||
					IS_BETWEEN(valid_tickets[k][j], rules[i].rs[1].l, rules[i].rs[1].h);
			}
			rules2cols[i][j] = res;
		}
	}

	int k = 0;
	while(sum2a(rules2cols, m) != m) {
		for (int i = 0; i < m; i++) {
			if (suma(rules2cols[i], m) == 1) {
				k = get1ia(rules2cols[i], m);
				for (int j = 0; j < m; j++) {
					if (j == i) continue;
					rules2cols[j][k] = 0;
				}
			}
		}
		//print2a(rules2cols, m);
		//printf("\n");
	}

	long mul = 1;
	for (int i = 0; i < 6; i++) {
		mul *= my_ticket[get1ia(rules2cols[i], m)];
	}

	DUMP("%ld", mul);

	// Frees
	FFREE(rules2cols, m);
	FFREE(valid_tickets, n_valid);
	free(my_ticket);

	for (int i = 0; i < n_rules; i++) {
		free(rules[i].label);
	}
	free(rules);
	FFREE(lines, n);
	return 0;
}
