#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(int argc, char const *argv[])
{
	printf("Hello world\n");

	int way = 4;

	if (way == 1)
	{
		printf("exit\n");
		exit(0);
	}
	else if (way == 2)
	{
		printf("_exit\n");
		_exit(0);
	}
	else if (way == 3)
	{
		printf("abort\n");
		abort();
	}
	else if (way == 4)
	{
		printf("_Exit\n");
		_Exit(0);
	}

	printf("return\n");
	return 0;
}