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

int main(int argc, char const *argv[])
{
  time_t now = time(NULL);
  printf("now is %ld %s\n", now, ctime(&now));

  struct tm *localnow = localtime(&now);
  struct tm *gmnow = gmtime(&now);

  printf("localnow : %d-%d-%d %d:%d:%d\n", localnow->tm_year, localnow->tm_mon, localnow->tm_mday, localnow->tm_hour, localnow->tm_min, localnow->tm_sec);
  printf("gmnow : %d-%d-%d %d:%d:%d\n", gmnow->tm_year, gmnow->tm_mon, gmnow->tm_mday, gmnow->tm_hour, gmnow->tm_min, gmnow->tm_sec);

  printf("localnow is %ld : %s\n", mktime(localnow), asctime(localnow));
  printf("gmnow is %ld : %s\n", mktime(gmnow), asctime(gmnow));

  struct timeval tv;
  struct timezone tz;
  gettimeofday(&tv, &tz);

  printf("%ld %ld\n", tv.tv_sec, tv.tv_usec);
  printf("%d %d\n", tz.tz_minuteswest, tz.tz_dsttime);

  printf("sleep start is %ld\n", time(NULL));
  sleep(2);
  printf("sleep stop is %ld\n", time(NULL));

  gettimeofday(&tv, &tz);
  printf("usleep start %ld %ld\n", tv.tv_sec, tv.tv_usec);
  usleep(5000);
  gettimeofday(&tv, &tz);
  printf("usleep stop %ld %ld\n", tv.tv_sec, tv.tv_usec);

  struct timespec req;
  req.tv_sec = 1;
  req.tv_nsec = 5000000;
  struct timespec rem;

  gettimeofday(&tv, &tz);
  printf("nanosleep start %ld %ld\n", tv.tv_sec, tv.tv_usec);
  nanosleep(&req, &rem);
  gettimeofday(&tv, &tz);
  printf("nanosleep stop %ld %ld\n", tv.tv_sec, tv.tv_usec);

  time_t t1,t2;
  t1 = time(NULL);
  sleep(2);
  t2 = time(NULL);

  printf("difftime : %f\n", difftime(t1, t2));

  struct tms startms, stoptms;

  clock_t tms1 = times(&startms);
  printf("tms1 = %ld %ld %ld %ld %ld\n", tms1, startms.tms_utime, startms.tms_stime, startms.tms_cutime, startms.tms_cstime);
  sleep(2);
  clock_t tms2 = times(&stoptms);
  printf("tms2 = %ld %ld %ld %ld %ld\n", tms2, stoptms.tms_utime, stoptms.tms_stime, stoptms.tms_cutime, stoptms.tms_cstime);

  struct timespec tsc;
  clock_gettime(CLOCK_REALTIME, &tsc);
  printf("CLOCK_REALTIME:%ld %ld\n", tsc.tv_sec, tsc.tv_nsec);
  sleep(1);
  clock_gettime(CLOCK_REALTIME, &tsc);
  printf("CLOCK_REALTIME:%ld %ld\n", tsc.tv_sec, tsc.tv_nsec);

  clock_gettime(CLOCK_MONOTONIC, &tsc);
  printf("CLOCK_MONOTONIC:%ld %ld\n", tsc.tv_sec, tsc.tv_nsec);
  sleep(1);
  clock_gettime(CLOCK_MONOTONIC, &tsc);
  printf("CLOCK_MONOTONIC:%ld %ld\n", tsc.tv_sec, tsc.tv_nsec);

  return 0;
}