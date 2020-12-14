#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "cge.h"

enum opcode{ NOP, ACC, JMP};

typedef struct instr_t {
	int op;
	int arg;
} instr;

typedef struct prog_t {
	int g_acc;
	int pc;
	int n_instrs;
	int* vis;
	instr** instrs;
} prog;

instr* stoinst(char* s)
{
	if (s == NULL) return NULL;
	char buf[64] = {0};
	int targ = 0;
	instr* i = malloc(sizeof(instr));
	if (sscanf(s, "%s %d\n", buf, &targ) == 2) {
		if (strcmp(buf, "nop") == 0) {
			i->op = NOP; i->arg = targ;
		} else if (strcmp(buf, "acc") == 0) {
			i->op = ACC; i->arg = targ;
		} else if (strcmp(buf, "jmp") == 0) {
			i->op = JMP; i->arg = targ;
		}
		return i;
	}
	return NULL;
}

prog* load_prog(char** lines, int n)
{
	if (lines == NULL) return NULL;
	prog* p = (prog*)malloc(sizeof(prog));
	p->g_acc = 0;
	p->pc = 0;
	p->n_instrs = n;
	p->vis = (int*)calloc(1, sizeof(int)*n);
	p->instrs = (instr**)calloc(1, sizeof(prog*)*n);

	for (int i = 0; i < n; i++) {
		p->vis[i] = 0;
		p->instrs[i] = stoinst(lines[i]);
	}
	return p;
}

void reset_prog(prog* p)
{
	p->g_acc = 0;
	p->pc = 0;
	for (int i = 0; i < p->n_instrs; i++) {
		p->vis[i] = 0;
	}
}

void free_prog(prog* p)
{
	free(p->vis);
	ffree((void**)p->instrs, p->n_instrs);
	free(p);
}

int run_prog_no_rep(prog* p)
{
	if (p == NULL) return 0;
	instr* instr;
	for (p->pc = 0; p->pc < p->n_instrs; p->pc++) {
		instr = p->instrs[p->pc];
		if (p->vis[p->pc] == 1) {
			return 0;
		}
		p->vis[p->pc] = 1;

		switch (instr->op) {
			case NOP:
				break;
			case ACC:
				p->g_acc += instr->arg;
				break;
			case JMP:
				p->pc += instr->arg - 1;
			default:
				break;
		}
	}
	return 1;
}

int* get_nops_jmps(prog* p, int* rn)
{
	int size = 16;
	int* res = (int*)malloc(sizeof(int)*size);
	int r = 0;

	for (int i = 0; i < p->n_instrs; i++) {
		if (p->instrs[i]->op == NOP || p->instrs[i]->op == JMP) {
			if (r == size-1) {
				size <<= 2;
				res = (int*)realloc(res, sizeof(int)*size);
			}
			res[r++] = i;
		}
	}
	*rn = r;
	return res;
}

int main(int argc, char *argv[])
{
	int n = 0;
	char** lines = load_file("./input", 0, &n);

	prog* p = load_prog(lines, n);
	int rn = 0;
	int* njis = get_nops_jmps(p, &rn);

	for (int i = 0; i < rn; i++) {
		if(p->instrs[njis[i]]->op == NOP) {
			p->instrs[njis[i]]->op = JMP;
		} else {
			p->instrs[njis[i]]->op = NOP;
		}

		if (run_prog_no_rep(p)) {
			DUMP("%d", p->g_acc);
			break;
		}

		if(p->instrs[njis[i]]->op == NOP) {
			p->instrs[njis[i]]->op = JMP;
		} else {
			p->instrs[njis[i]]->op = NOP;
		}
		reset_prog(p);
	}
	free(njis);

	free_prog(p);
	ffree((void**)lines, n);
	return 0;
}
