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
#include <signal.h>
#include <setjmp.h>
#include <sys/signalfd.h>

#define handle_error(msg) \
           do { perror(msg); exit(EXIT_FAILURE); } while (0)

typedef void (*sighandler)(int);

static void handler1(int no)
{
	printf("no = %d\n", no);
	printf("SIGABRT = %d\n", SIGABRT);
	printf("SIGALRM = %d\n", SIGALRM);
	printf("SIGINT = %d\n", SIGINT);
	printf("SIGUSR1 = %d\n", SIGUSR1);
}

void do1()
{
	sighandler sig;
	sig = signal(SIGABRT, handler1);//SIGSTOP SIGKILL
	if (sig == SIG_ERR)
	{
		perror("signal fail");
		return;
	}

	sig = signal(SIGALRM, handler1);
	if (sig == SIG_ERR)
	{
		perror("signal fail");
		return;
	}

	sig = signal(SIGINT, handler1);
	if (sig == SIG_ERR)
	{
		perror("signal fail");
		return;
	}

	sig = signal(SIGUSR1, handler1);
	if (sig == SIG_ERR)
	{
		perror("signal fail");
		return;
	}

	//kill(getpid(), SIGUSR1);
	//raise(SIGUSR1);
	//abort();
	//alarm(2);

	sleep(4);
	printf("sleep end\n");
}

void do2()
{
	sighandler sig;
	sig = signal(SIGUSR1, handler1);
	if (sig == SIG_ERR)
	{
		perror("signal fail");
		return;
	}

	pid_t pid;
	pid = fork();
	if (pid == -1)
	{
		printf("fork fail\n");
		return;
	}
	else if (pid == 0)
	{
		printf("son start, pid = %d, ppid = %d\n", getpid(), getppid());
		sleep(3);
		kill(getppid(), SIGUSR1);
		printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
		exit(0);
	}
	else
	{
		printf("pause start\n");
		pause();
		printf("pause end\n");

		int status = 0;
		int ret = waitpid(pid, &status, 0);
		if (ret == -1)
		{
			printf("son ret = %d, status = %d\n", ret, status);
			return;
		}

		printf("son ret = %d, status = %d\n", ret, status);
	}
}

void do3()
{
	sighandler sig;
	sig = signal(SIGALRM, handler1);
	if (sig == SIG_ERR)
	{
		perror("signal fail");
		return;
	}

	printf("sleep start\n");
	sleep(2);
	printf("sleep end\n");
}

void do4()
{
	int n = 0;
	jmp_buf buf;

	int ret = setjmp(buf);

	n++;
	printf("dddd %d, %d\n",ret, n);

	if (n == 20)
	{
		exit(0);
	}

	longjmp(buf, 1);
}

void do5()
{
	sigset_t mask;
	int sfd;
	struct signalfd_siginfo fdsi;
	ssize_t s;

	sigemptyset(&mask);
	sigaddset(&mask, SIGINT);
	sigaddset(&mask, SIGQUIT);

	/* Block signals so that they aren't handled
	according to their default dispositions */

	if (sigprocmask(SIG_BLOCK, &mask, NULL) == -1)
		handle_error("sigprocmask");

	sfd = signalfd(-1, &mask, 0);
	if (sfd == -1)
		handle_error("signalfd");

	for (;;) {
		s = read(sfd, &fdsi, sizeof(struct signalfd_siginfo));
		if (s != sizeof(struct signalfd_siginfo))
			handle_error("read");

		if (fdsi.ssi_signo == SIGINT) {
			printf("Got SIGINT\n");
		} else if (fdsi.ssi_signo == SIGQUIT) {
			printf("Got SIGQUIT\n");
			exit(EXIT_SUCCESS);
		} else {
			printf("Read unexpected signal\n");
		}
	}
}

void now()
{
	time_t now = time(NULL);
	printf("now is %s\n", ctime(&now));
}

static void sig_alarm(int sig)
{
	printf("alarm: %d\n", sig);
	now();
}

void do6()
{
	signal(SIGALRM, sig_alarm);

	unsigned int n1 = alarm(7);
	printf("%d\n", n1);
	now();

	sleep(3);

	unsigned int n2 = alarm(7);
	printf("%d\n", n2);
	now();

	sleep(100);
	//sleep(100);
	printf("sleep done\n");
	now();
}

void info(const char *str)
{
	time_t now = time(NULL);
	char *timeStr = ctime(&now);
	size_t n = strlen(timeStr);
	if (n >= 1 && timeStr[n-1] == '\n')
	{
		timeStr[n-1] = '\0';
	}
	printf("%s -- %s\n", timeStr, str);
}

static void sig_any(int sig)
{
	char str[20] = {0};
	sprintf(str, "sig no: %d", sig);
	info(str);
}

void do7()
{
	//signal(SIGUSR1, SIG_IGN); // 忽略信号,不影响sleep
	//signal(SIGUSR1, SIG_DFL); // 默认进程终止
	signal(SIGUSR1, sig_any); // 会先执行信号处理函数,然后sleep退出并返回剩余秒数

	info("sleep1 start");
	printf("left: %ds\n", sleep(20));
	info("sleep1 done");

	info("sleep2 start");
	printf("left: %ds\n", sleep(20));
	info("sleep2 done");
}

void do8()
{
	//signal(SIGUSR1, SIG_IGN); // 忽略信号,不影响pause
	//signal(SIGUSR1, SIG_DFL); // 默认进程终止
	signal(SIGUSR1, sig_any); // 会先执行信号处理函数,然后sleep退出并返回剩余秒数

	info("pause start");
	pause();
	info("pause done");
}

void do9()
{
	//signal(SIGUSR1, SIG_IGN); // 忽略信号,不影响pause
	//signal(SIGUSR1, SIG_DFL); // 默认进程终止
	signal(SIGUSR1, sig_any); // 会先执行信号处理函数,然后sleep退出并返回剩余秒数

	info("pause start");
	pause();
	info("pause done");
}

void do10()
{
	//signal(SIGALRM, sig_any);
	//signal(SIGVTALRM, sig_any);
	signal(SIGPROF, sig_any);

	struct itimerval it;
	it.it_interval.tv_sec = 2;
	it.it_interval.tv_usec = 0;
	it.it_value.tv_sec = 5;
	it.it_value.tv_usec = 0;

	//int which = ITIMER_REAL;
	//int which = ITIMER_VIRTUAL;
	int which = ITIMER_PROF;

	setitimer(which, &it, NULL);

	info("setitimer start");

	struct itimerval oit;
	getitimer(which, &oit);
	printf("%ld\n", oit.it_interval.tv_sec);
	printf("%ld\n", oit.it_interval.tv_usec);
	printf("%ld\n", oit.it_value.tv_sec);
	printf("%ld\n", oit.it_value.tv_usec);

	while(1);
}

void do11()
{
	struct timespec ts;
	int ret1 = clock_getres(CLOCK_REALTIME, &ts);
	printf("clock_getres ret: %d\n", ret1);
	printf("clock_getres tv_sec: %ld\n", ts.tv_sec);
	printf("clock_getres tv_nsec: %ld\n", ts.tv_nsec);

	int ret2 = clock_gettime(CLOCK_REALTIME, &ts);
	printf("clock_gettime ret: %d\n", ret2);
	printf("clock_gettime tv_sec: %ld\n", ts.tv_sec);
	printf("clock_gettime tv_nsec: %ld\n", ts.tv_nsec);

	ts.tv_sec++;
	int ret3 = clock_settime(CLOCK_REALTIME, &ts);
	printf("clock_settime ret: %d\n", ret3);
	perror("clock_settime fail");
}

       
static void print_siginfo(siginfo_t *si)
{
	timer_t *tidp;
	int or;

	tidp = si->si_value.sival_ptr;

	printf("    sival_ptr = %p; ", si->si_value.sival_ptr);
	printf("    *sival_ptr = 0x%lx\n", (long) *tidp);

	or = timer_getoverrun(*tidp);
	if (or == -1)
		handle_error("timer_getoverrun");
	else
		printf("    overrun count = %d\n", or);
}

static void handler(int sig, siginfo_t *si, void *uc)
{
	printf("caught signal %d\n", sig);
	print_siginfo(si);
	signal(sig, SIG_IGN);
}

// gcc -o app signal.c -lrt
// /proc/[pid]/timers
void do12()
{
	struct sigevent sev;
	timer_t timerid;
	struct itimerspec its;
	long long freq_nanosecs = 100;
	sigset_t mask;
	struct sigaction sa;

	sa.sa_flags = SA_SIGINFO;
	sa.sa_sigaction = handler;
	sigemptyset(&sa.sa_mask);
	if (sigaction(SIGUSR1, &sa, NULL) == -1)
		handle_error("sigaction");

	sigemptyset(&mask);
	sigaddset(&mask, SIGUSR1);
	if (sigprocmask(SIG_SETMASK, &mask, NULL) == -1)
		handle_error("sigprocmask");

	sev.sigev_notify = SIGEV_SIGNAL;
	sev.sigev_signo = SIGUSR1;
	sev.sigev_value.sival_ptr = &timerid;
	if (timer_create(CLOCK_REALTIME, &sev, &timerid) == -1)
		handle_error("timer_create");

	printf("timer_create timerid: %ld\n", (long)timerid);

	its.it_value.tv_sec = freq_nanosecs / 1000000000;
	its.it_value.tv_nsec = freq_nanosecs % 1000000000;
	its.it_interval.tv_sec = its.it_value.tv_sec;
	its.it_interval.tv_nsec = its.it_value.tv_nsec;

	if (timer_settime(timerid, 0, &its, NULL) == -1)
		handle_error("timer_settime");

	while(1);
}

/*
struct sigaction {
	void     (*sa_handler)(int);
	void     (*sa_sigaction)(int, siginfo_t *, void *);
	sigset_t   sa_mask;
	int        sa_flags;
	void     (*sa_restorer)(void);
};
*/
void do13()
{
	sigaction();
}

int main(int argc, char const *argv[])
{
	do13();
	return 0;
}
