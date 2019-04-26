#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <string.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <errno.h>
#include <time.h>
#include <sys/time.h>
#include <sys/times.h>
#include <wchar.h>

void do1()
{
	char buf[] = "123456789";

	FILE *pf = fmemopen(buf, strlen(buf), "r+");
	if (pf == NULL)
	{
		printf("fmemopen fail\n");
		return;
	}

	char ch;
	while((ch = fgetc(pf)) != EOF)
	{
		printf("%c\n", ch);
	}

	fclose(pf);
}

/*
fwide

setbuf
setvbuf

freopen
fdopen

fprintf
fscanf

fread
fwrite

ferror
feof
clearerr

ungetc

fileno

tmpnam
tmpfile

fmemopen
open_memstream
open_wmemstream
*/

void do2()
{
	int ret;
	FILE *fp = fopen("file.txt", "w+");
	if (fp == NULL)
	{
		printf("fopen fail\n");
		return;
	}

	putc('1', fp);
	fputc('2', fp);
	fputc('3', fp);
	fputc('4', fp);

	rewind(fp);
	int ch;
	ch = getc(fp);
	if (ch == EOF)
	{
		printf("getc EOF\n");
	}
	else
	{
		printf("getc:%c\n", ch);
	}

	while(!feof(fp))
	{
		ch = fgetc(fp);
		if (ch == EOF)
		{
			printf("fgetc EOF\n");
		}
		else
		{
			printf("fgetc:%c\n", ch);
		}
	}

	fputs("ABCD\nEFGH", fp);

	char buf[] = "123456";
	rewind(fp);
	char *p = fgets(buf, sizeof(buf), fp);
	if (p == NULL)
	{
		printf("fgets fail\n");
		return;
	}
	printf("1 fgets buf:%s\n", buf);

	p = fgets(buf, sizeof(buf), fp);
	if (p == NULL)
	{
		printf("fgets fail\n");
		return;
	}
	printf("2 fgets buf:%s\n", buf);//fgets include '\n'

	p = fgets(buf, sizeof(buf), fp);
	if (p == NULL)
	{
		printf("fgets fail\n");
		return;
	}
	printf("3 fgets buf:%s\n", buf);

	printf("pos = %ld\n", ftell(fp));
	fseek(fp, 1, SEEK_SET);
	printf("after fseek pos = %ld\n", ftell(fp));

	fpos_t pt;
	fgetpos(fp, &pt);
	printf("fgetpos:%ld\n", pt.__pos);

	pt.__pos = 2;
	fsetpos(fp, &pt);
	fgetpos(fp, &pt);
	printf("after fsetpos fgetpos:%ld\n", pt.__pos);

	fflush(fp);
	fclose(fp);
	remove("file.txt");
}

void do3()
{
	printf("input char:");
	int ch = getchar();//getc(stdin);
	printf("getchar:%c\n", ch);

	printf("putchar:");
	putchar(ch);//putc(c, stdout);
	printf("\n");

	//gets is dangerous
	printf("puts:");
	puts("0123456789");
	puts("01234567\n89");
}

void do4()
{
	printf("input char:");
	int ch = getchar();//include '\n'
	printf("getchar:%c\n", ch);

	//fflush(stdin);//useless in linux
	while((ch = getchar()) != EOF)
	{
		printf("getchar:%c\n", ch);
	}

	printf("getchar end\n");
}

void do5()
{
	printf("111222");
	fflush(stdout);
	sleep(2);
	printf("111222\n");
}

int main(int argc, char const *argv[])
{
	//do1();
	//do2();
	//do3();
	//do4();
	do5();

	return 0;
}