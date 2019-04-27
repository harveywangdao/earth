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

ferror
feof
clearerr




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

void do6()
{
	//stdout
	FILE *outfp = freopen("out.txt", "w+", stdout);
	if (outfp == NULL)
	{
		printf("freopen fail\n");
		return;
	}

	printf("testfreopen\n");
	fclose(outfp);

	freopen("/dev/tty", "w", stdout);

	FILE *fp = fopen("out.txt", "r");
	if (fp == NULL)
	{
		printf("fopen fail\n");
		return;
	}

	char buf[32];
	char *p = fgets(buf, sizeof(buf), fp);
	if (p == NULL)
	{
		printf("fgets fail\n");
		return;
	}
	printf("fgets buf:%s\n", buf);
	fclose(fp);

	//stdin
	FILE *infp = freopen("out.txt", "r", stdin);
	if (infp == NULL)
	{
		printf("freopen fail\n");
		return;
	}

	memset(buf, 0, sizeof(buf));
	scanf("%s", buf);
	printf("scanf:%s\n", buf);

	int ch;
	while((ch = getchar()) != EOF)
	{
		printf("getchar:%c\n", ch);
	}

	fclose(infp);

	freopen("/dev/tty", "r", stdin);

	printf("input char:");
	ch = getchar();
	printf("getchar:%c\n", ch);

	remove("out.txt");
}

void do7()
{
	int fd = open("file.txt", O_RDWR|O_CREAT|O_TRUNC, 0666);
	if (fd == -1)
	{
	  perror("open fail");
	  return;
	}

	char *str = "test fdopen";
	int nb = write(fd, str, strlen(str));
	if (nb == -1)
	{
	  perror("write fail");
	  return;
	}

	FILE *fp = fdopen(fd, "r");
	if (fp == NULL)
	{
		printf("fdopen fail\n");
		return;
	}

	rewind(fp);
	char buf[32];
	char *p = fgets(buf, sizeof(buf), fp);
	if (p == NULL)
	{
		printf("fgets fail\n");
		return;
	}
	printf("fgets buf:%s\n", buf);

	fclose(fp);
  close(fd);

  remove("file.txt");
}

void do8()
{
	FILE *fp = fopen("file.txt", "w+");
	if (fp == NULL)
	{
		printf("fopen fail\n");
		return;
	}

	fprintf(fp, "%s", "fprintftest123");
	
	rewind(fp);
	char buf[32]={0};
	fscanf(fp, "%s", buf);
	printf("fscanf:%s\n", buf);

	rewind(fp);
	char *str = "fwrite test";
	int n = fwrite(str, sizeof(str[0]), strlen(str), fp);
	printf("fwrite n:%d\n", n);

	rewind(fp);
	n = fread(buf, sizeof(buf[0]), sizeof(buf)-1, fp);
	buf[n] = '\0';
	printf("fread n:%d:%s\n", n, buf);

	int fd = fileno(fp);
	lseek(fd, 0, SEEK_SET);
	str = "fileno";
	int nb = write(fd, str, strlen(str));
	if (nb == -1)
	{
	  perror("write fail");
	  return;
	}

	lseek(fd, 0, SEEK_SET);
	memset(buf, 0, sizeof(buf));
	nb = read(fd, buf, sizeof(buf));
	if (nb == -1)
	{
	  perror("read fail");
	  return;
	}
	printf("read %d:%s\n", nb, buf);

	fclose(fp);
	remove("file.txt");
}

void do9()
{
	FILE *fp = fopen("file.txt", "w+");
	if (fp == NULL)
	{
		printf("fopen fail\n");
		return;
	}

	fprintf(fp, "%s", "123456");

	rewind(fp);
	int ch = getc(fp);
	if (ch == EOF)
	{
		printf("getc EOF\n");
	}
	else
	{
		printf("getc:%c\n", ch);
	}

	ch = ungetc('A', fp);
	if (ch == EOF)
	{
		printf("ungetc EOF\n");
	}
	else
	{
		printf("ungetc:%c\n", ch);
	}
	ungetc('B', fp);
	ungetc('C', fp);
	ungetc('D', fp);

	while((ch = getc(fp)) != EOF)
	{
		printf("getc:%c\n", ch);
	}

	rewind(fp);
	char buf[32]={0};
	int n = fread(buf, sizeof(buf[0]), sizeof(buf)-1, fp);
	buf[n] = '\0';
	printf("fread n:%d:%s\n", n, buf);

	fclose(fp);
	remove("file.txt");
}

void do10()
{
	char name[L_tmpnam];

	for (int i = 0; i < 3; ++i)
	{
		tmpnam(name);//tmpnam is dangerous
		printf("tmpnam:%s\n", name);
	}

	char *p = tempnam("/home/thomas/golang/src/earth/paper", "aaa_");//tempnam is dangerous
	printf("tempnam:%s\n", p);

	FILE *fp = tmpfile();
	fprintf(fp, "%s", "123456");

	rewind(fp);
	char buf[32]={0};
	int n = fread(buf, sizeof(buf[0]), sizeof(buf)-1, fp);
	buf[n] = '\0';
	printf("fread n:%d:%s\n", n, buf);
	fclose(fp);

	char dir[] = "helloa_XXXXXX";
	p = mkdtemp(dir);
	if (p == NULL)
	{
		perror("mkdtemp fail");
		return;
	}
	printf("%s\n%s\n", dir, p);

	remove(dir);

	char tpl[] = "hellos_XXXXXX";
	int fd = mkstemp(tpl);
	if (fd == -1)
	{
		perror("mkstemp fail\n");
		return;
	}

	char *str = "mkstemp test";
	int nb = write(fd, str, strlen(str));
	if (nb == -1)
	{
	  perror("write fail");
	  return;
	}
	
	lseek(fd, 0, SEEK_SET);
	memset(buf, 0, sizeof(buf));
	nb = read(fd, buf, sizeof(buf));
	if (nb == -1)
	{
	  perror("read fail");
	  return;
	}
	printf("read %d:%s\n", nb, buf);
	close(fd);
	unlink(tpl);
}

int main(int argc, char const *argv[])
{
	//do1();
	//do2();
	//do3();
	//do4();
	//do5();
	//do6();
	//do7();
	//do8();
	//do9();
	do10();

	return 0;
}