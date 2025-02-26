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
#include <sys/timerfd.h>
#include <signal.h>
#include <setjmp.h>
#include <sys/signalfd.h>
#include <sys/eventfd.h>

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
	//jmp_buf buf;
	sigjmp_buf buf;

	//int ret = setjmp(buf);
	int ret = sigsetjmp(buf, 1);

	n++;
	printf("dddd %d, %d\n",ret, n);

	if (n == 20)
	{
		exit(0);
	}

	//longjmp(buf, 1);
	siglongjmp(buf, 1);
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

	printf("sival_ptr = %p; ", si->si_value.sival_ptr);
	printf("*sival_ptr = %ld\n", (long) *tidp);

	or = timer_getoverrun(*tidp);
	if (or == -1)
		handle_error("timer_getoverrun");
	else
		printf("overrun count = %d\n", or);
}

static void handler(int sig, siginfo_t *si, void *uc)
{
	printf("caught signal %d\n", sig);
	print_siginfo(si);
	//signal(sig, SIG_IGN);
}

static void sa_sigaction_func(int sig, siginfo_t *sinfo, void *ucontext)
{
	psiginfo(sinfo, "msg");
	char str[128] = {0};
	sprintf(str, "sig no: %d, si_signo: %d, si_errno: %d, si_code: %d, si_pid: %d, si_value.sival_int: %d, si_int: %d", 
		sig, sinfo->si_signo, sinfo->si_errno, sinfo->si_code, sinfo->si_pid, sinfo->si_value.sival_int, sinfo->si_int);
	info(str);
}

// gcc -o app signal.c -lrt
// /proc/[pid]/timers
void do12()
{
	struct sigevent sev;
	timer_t timerid;
	struct itimerspec its;

	struct sigaction sa;
	sa.sa_flags = SA_SIGINFO;
	sa.sa_sigaction = handler;
	sigemptyset(&sa.sa_mask);
	if (sigaction(SIGUSR1, &sa, NULL) == -1)
		handle_error("sigaction");

	sigset_t mask;
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

	its.it_value.tv_sec = 5;
	its.it_value.tv_nsec = 5;
	its.it_interval.tv_sec = 5;
	its.it_interval.tv_nsec = 5;

	if (timer_settime(timerid, 0, &its, NULL) == -1)
		handle_error("timer_settime");

	struct itimerspec curr_value;
	if (timer_gettime(timerid, &curr_value) == -1)
		handle_error("timer_gettime");
	printf("%ld %ld %ld %ld\n", curr_value.it_value.tv_sec, curr_value.it_value.tv_nsec, curr_value.it_interval.tv_sec, curr_value.it_interval.tv_nsec);

	info("sleep start");
	sleep(6);
	info("sleep done");

	if (sigprocmask(SIG_UNBLOCK, &mask, NULL) == -1)
		handle_error("sigprocmask");
	info("sigprocmask reset done");

	while(1);
}

void do13()
{
	signal(SIGUSR1, sig_any);
	signal(SIGUSR2, sig_any);

	sigset_t newset, oldset;
	sigemptyset(&newset);
	sigaddset(&newset, SIGUSR1);
	sigaddset(&newset, SIGUSR2);
	if (sigprocmask(SIG_BLOCK, &newset, &oldset) == -1)
		handle_error("sigprocmask fail");

	info("sleep1 start");
	sleep(30);
	info("sleep1 done");

	sigset_t pendset;
	if (sigpending(&pendset) == -1)
		handle_error("sigpending fail");

	int ret = sigismember(&pendset, SIGUSR1);
	if (ret == 1)
	{
		printf("SIGUSR1 is pending\n");
	}
	else if (ret == -1)
	{
		handle_error("sigismember fail");
	}

	ret = sigismember(&pendset, SIGUSR2);
	if (ret == 1)
	{
		printf("SIGUSR2 is pending\n");
	}
	else if (ret == -1)
	{
		handle_error("sigismember fail");
	}

	info("sigprocmask reset start");
	if (sigprocmask(SIG_SETMASK, &oldset, NULL) == -1)
		handle_error("sigprocmask fail");
	info("sigprocmask reset done");

	while(1);
}

void do14()
{
	struct sigaction act;
	sigemptyset(&act.sa_mask);
	//act.sa_handler = sig_any;
	//act.sa_flags = SA_RESTART;
	act.sa_sigaction = sa_sigaction_func;
	act.sa_flags = SA_SIGINFO;
	act.sa_restorer = NULL;
	if (sigaction(SIGUSR1, &act, NULL) == -1)
		handle_error("sigaction fail");

	while(1);
}

void do15()
{
	signal(SIGUSR1, sig_any);
	signal(SIGUSR2, sig_any);

	sigset_t set;
	sigemptyset(&set);
	sigaddset(&set, SIGUSR1);

	info("sigsuspend start");
	sigsuspend(&set);
	info("sigsuspend done");
}

void do16(int argc, char const *argv[])
{
	if (argc != 2)
	{
		fprintf(stderr, "use error\n");
		exit(EXIT_FAILURE);
	}

	pid_t pid = atoi(argv[1]);
	union sigval value;
	value.sival_int = 44;

	info("sigqueue start");
	if (sigqueue(pid, SIGUSR1, value) == -1)
		handle_error("sigqueue fail");
	info("sigqueue done");
}

void do17()
{
	signal(SIGUSR1, sig_any);
	signal(SIGUSR2, sig_any);

	sigset_t set;
	sigemptyset(&set);
	sigaddset(&set, SIGUSR1);
	int sig;

	info("sigwait start");
	//if (sigwait(&set, &sig) != 0)
	//	handle_error("sigwait fail");
	//printf("sig: %d\n", sig);

	siginfo_t sinfo;
	struct timespec ts;
	ts.tv_sec = 20;
	ts.tv_nsec = 3;
	//if (sigwaitinfo(&set, &sinfo) == -1)
	if (sigtimedwait(&set, &sinfo, &ts) == -1)
		handle_error("sigwaitinfo fail");
	printf("si_signo: %d, si_errno: %d, si_code: %d, si_pid: %d, si_value.sival_int: %d, si_int: %d\n", 
		sinfo.si_signo, sinfo.si_errno, sinfo.si_code, sinfo.si_pid, sinfo.si_value.sival_int, sinfo.si_int);

	info("sigwait done");
}

void do18()
{
	signal(SIGUSR1, sig_any);

	int tfd = timerfd_create(CLOCK_REALTIME, TFD_CLOEXEC);
	if (tfd == -1)
		handle_error("timerfd_create fail");

	struct itimerspec its;
	its.it_value.tv_sec = 10;
	its.it_value.tv_nsec = 0;
	its.it_interval.tv_sec = 10;
	its.it_interval.tv_nsec = 0;
	if (timerfd_settime(tfd, 0, &its, NULL) == -1)
		handle_error("timerfd_settime fail");

	while(1)
	{
		uint64_t dd;
		info("read start");
		int n = read(tfd, &dd, sizeof(dd));
		info("read done");
		printf("%d, %ld\n", n, dd);

		struct itimerspec its2;
		if (timerfd_gettime(tfd, &its2) == -1)
			handle_error("timerfd_gettime fail");
		printf("%ld %ld %ld %ld\n", its2.it_value.tv_sec, its2.it_value.tv_nsec, its2.it_interval.tv_sec, its2.it_interval.tv_nsec);
	}
}

void do19()
{
	int efd = eventfd(0, 0);

	switch (fork()) 
	{
		case 0: //子进程
		{
			uint64_t u1 = 0;
			for (int i = 0; i < 5; ++i)
			{
				int s = write(efd, &u1, sizeof(uint64_t));
				if (s != sizeof(uint64_t))
					handle_error("write");
				printf("child write: %ld, size: %d\n", u1, s);
				//u1++;
			}
			exit(EXIT_SUCCESS);
		}

		default: //父进程
		{
			//sleep(2); //先休眠2秒，等待子进程写完数据

			while(1)
			{
				uint64_t u2;
				info("read start");
				int s = read(efd, &u2, sizeof(uint64_t));
				info("read done");
				if (s != sizeof(uint64_t))
					handle_error("read");
				printf("parent read %ld\n", u2);
			}

			exit(EXIT_SUCCESS);		
		}

		case -1:
			handle_error("fork");
	}
}

void do20()
{
	psignal(SIGUSR1, "msg"); // msg: User defined signal 1
	printf("%s\n", strsignal(SIGUSR1));

	printf("%ld\n", sysconf(_SC_SIGQUEUE_MAX));
}

void do21()
{
	signal(SIGABRT, sig_any);

	abort();

	while(1);
}

static void sig_any2(int sig)
{
	char str[20] = {0};
	char str2[128] = {0};
	sprintf(str, "sig no: %d", sig);
	info(str);

	time_t start = time(NULL);
	while(1)
	{
		if (time(NULL) > start+5)
		{
			sprintf(str2, "sig_any2 exit, %s", str);
			info(str2);
			printf("\n");
			break;
		}
	}
}

void do22()
{
	signal(SIGUSR1, sig_any2);
	signal(SIGUSR2, sig_any2);
	while(1);
}

static void sa_sigaction_func2(int sig, siginfo_t *sinfo, void *ucontext)
{
	int a = 9;
	printf("%p\n", &a);
}

void do23()
{
	stack_t ss;
	ss.ss_sp = malloc(SIGSTKSZ);
	if (ss.ss_sp == NULL)
		handle_error("malloc");

	printf("%p\n", ss.ss_sp);
	printf("SIGSTKSZ: %d\n", SIGSTKSZ);

	ss.ss_size = SIGSTKSZ;
	ss.ss_flags = 0;
	if (sigaltstack(&ss, NULL) == -1)
		handle_error("sigaltstack");

	struct sigaction act;
	sigemptyset(&act.sa_mask);
	act.sa_sigaction = sa_sigaction_func2;
	act.sa_flags = SA_ONSTACK;
	if (sigaction(SIGUSR1, &act, NULL) == -1)
		handle_error("sigaction fail");

	while(1);
}

int main(int argc, char const *argv[])
{
	printf("pid: %d\n", getpid());
	//do16(argc, argv);
	do23();
	return 0;
}
