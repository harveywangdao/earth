#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>
#include <string.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <netinet/in.h>
#include <stdint.h>
#include <netdb.h>
#include <time.h>
#include <sys/select.h>
#include <poll.h>
#include <sys/epoll.h>
#include <sys/un.h>
#include <stddef.h>
#include <dbus/dbus.h>

//Methodcall Methodreturn Error Signal
//进程间函数调用以及进程间信号广播
//glib-2.0
//pkg-config --libs dbus-glib-1

/*
gcc -o app dbus.c -ldbus-1
dbus-daemon
dbus-launch
eval `dbus-launch --auto-syntax`
eval `dbus-launch --sh-syntax`
dbus-monitor
dbus-send
*/

void sendsignal()
{
  DBusError error;
  DBusConnection *conn;

  DBusMessage *msg;
  DBusMessageIter args;
  char *sigvalue = "i am sendsignalpp";
  dbus_uint32_t serial = 0;
  int ret;

  sleep(2);

  dbus_error_init(&error);

  conn = dbus_bus_get(DBUS_BUS_SESSION, &error); //DBUS_BUS_SYSTEM DBUS_BUS_STARTER
  if (dbus_error_is_set(&error))
  {
    printf("dbus_bus_get fail, err:%s\n", error.message);
    dbus_error_free(&error);
  }

  if (conn == NULL)
  {
    printf("dbus_bus_get fail\n");
    return;
  }

  ret = dbus_bus_request_name(conn, "dbusname.sendsignal", DBUS_NAME_FLAG_REPLACE_EXISTING, &error);
  if (dbus_error_is_set(&error))
  {
    printf("dbus_bus_request_name fail, err:%s\n", error.message);
    dbus_error_free(&error);
  }

  if (ret != DBUS_REQUEST_NAME_REPLY_PRIMARY_OWNER)
  {
    printf("dbus_bus_request_name fail\n");
    return;
  }

  msg = dbus_message_new_signal("/test/signal/Object", "test.signal.Type", "Test");
  if (msg == NULL)
  {
    printf("dbus_message_new_signal fail\n");
    return;
  }
  dbus_message_iter_init_append(msg, &args);

  if (!dbus_message_iter_append_basic(&args, DBUS_TYPE_STRING, &sigvalue))
  {
    printf("dbus_message_iter_append_basic fail\n");
    return;
  }

  if (!dbus_connection_send(conn, msg, &serial))
  {
    printf("dbus_connection_send fail\n");
    return;
  }

  dbus_connection_flush(conn);
  printf("dbus_connection_send serial = %d\n", serial);

  dbus_message_unref(msg);

  //dbus_connection_close(conn);
}

void recvsignal()
{
  DBusError error;
  DBusConnection *conn;

  DBusMessage *msg;
  DBusMessageIter args;
  char *sigvalue;
  int ret;

  dbus_error_init(&error);

  conn = dbus_bus_get(DBUS_BUS_SESSION, &error); //DBUS_BUS_SYSTEM DBUS_BUS_STARTER
  if (dbus_error_is_set(&error))
  {
    printf("dbus_bus_get fail, err:%s\n", error.message);
    dbus_error_free(&error);
  }

  if (conn == NULL)
  {
    printf("dbus_bus_get fail\n");
    return;
  }

  ret = dbus_bus_request_name(conn, "dbusname.recvsignal", DBUS_NAME_FLAG_REPLACE_EXISTING, &error);
  if (dbus_error_is_set(&error))
  {
    printf("dbus_bus_request_name fail, err:%s\n", error.message);
    dbus_error_free(&error);
  }

  if (ret != DBUS_REQUEST_NAME_REPLY_PRIMARY_OWNER)
  {
    printf("dbus_bus_request_name fail\n");
    return;
  }

  dbus_bus_add_match(conn, "type='signal',interface='test.signal.Type'", &error); // see signals from the given interface
  dbus_connection_flush(conn);
  if (dbus_error_is_set(&error))
  {
    printf("dbus_bus_add_match fail, err:%s\n", error.message);
    dbus_error_free(&error);
  }

  while(1)
  {
    printf("dbus_connection_read_write start\n");
    dbus_connection_read_write(conn, 1111110);
    printf("dbus_connection_read_write end\n");

    msg = dbus_connection_pop_message(conn);
    if (NULL == msg)
    {
      printf("msg is null\n");
      sleep(1);
      continue;
    }

    if (dbus_message_is_signal(msg, "test.signal.Type", "Test"))
    {
      if (!dbus_message_iter_init(msg, &args))
      {
        printf("message has no arguments!\n");
      }
      else if (DBUS_TYPE_STRING != dbus_message_iter_get_arg_type(&args))
      {
        printf("argument is not string!\n");
      }
      else 
      {
        dbus_message_iter_get_basic(&args, &sigvalue);
        printf("get signal with value:%s\n", sigvalue);
        dbus_message_unref(msg);
        return;
      }
    }

    dbus_message_unref(msg);
  }
  
  dbus_connection_close(conn);
}

void reply_to_method_call(DBusMessage *msg, DBusConnection *conn)
{
  DBusMessage *reply;
  DBusMessageIter args;
  int stat = 1;
  dbus_uint32_t level = 21614;
  dbus_uint32_t serial = 0;
  char *param = "";

  if (!dbus_message_iter_init(msg, &args))
  {
    printf("message has no arguments!\n");
  }
  else if (DBUS_TYPE_STRING != dbus_message_iter_get_arg_type(&args))
  {
    printf("argument is not string!\n");
  }
  else
  {
    dbus_message_iter_get_basic(&args, &param);
  }

  printf("method called with %s\n", param);

  reply = dbus_message_new_method_return(msg);
  dbus_message_iter_init_append(reply, &args);
  if (!dbus_message_iter_append_basic(&args, DBUS_TYPE_BOOLEAN, &stat))
  {
    printf("dbus_message_iter_append_basic fail\n");
    return;
  }

  if (!dbus_message_iter_append_basic(&args, DBUS_TYPE_UINT32, &level))
  {
    printf("dbus_message_iter_append_basic fail\n");
    return;
  }

  if (!dbus_connection_send(conn, reply, &serial))
  {
    printf("dbus_connection_send fail\n");
    return;
  }

  dbus_connection_flush(conn);
  dbus_message_unref(reply);
}

void exposemethod()
{
  DBusError error;
  DBusConnection *conn;

  DBusMessage *msg;
  DBusMessageIter args;
  int ret;

  dbus_error_init(&error);

  conn = dbus_bus_get(DBUS_BUS_SESSION, &error); //DBUS_BUS_SYSTEM DBUS_BUS_STARTER
  if (dbus_error_is_set(&error))
  {
    printf("dbus_bus_get fail, err:%s\n", error.message);
    dbus_error_free(&error);
  }

  if (conn == NULL)
  {
    printf("dbus_bus_get fail\n");
    return;
  }

  ret = dbus_bus_request_name(conn, "dbusname.exposemethod", DBUS_NAME_FLAG_REPLACE_EXISTING, &error);
  if (dbus_error_is_set(&error))
  {
    printf("dbus_bus_request_name fail, err:%s\n", error.message);
    dbus_error_free(&error);
  }

  if (ret != DBUS_REQUEST_NAME_REPLY_PRIMARY_OWNER)
  {
    printf("dbus_bus_request_name fail\n");
    return;
  }

  while(1)
  {
    printf("dbus_connection_read_write start\n");
    dbus_connection_read_write(conn, 0);
    printf("dbus_connection_read_write end\n");

    msg = dbus_connection_pop_message(conn);
    if (NULL == msg)
    {
      printf("msg is null\n");
      sleep(1);
      continue;
    }

    if (dbus_message_is_method_call(msg, "test.method.Type", "Method"))
    {
      printf("someone call me\n");
      reply_to_method_call(msg, conn);
    }

    dbus_message_unref(msg);
  }

  dbus_connection_close(conn);
}

void callmethod()
{
  DBusError error;
  DBusConnection *conn;

  DBusMessage *msg;
  DBusMessageIter args;
  char *sigvalue = "i am callmethod";
  DBusPendingCall *pending = NULL;
  int ret;

  sleep(5);
  dbus_error_init(&error);

  conn = dbus_bus_get(DBUS_BUS_SESSION, &error); //DBUS_BUS_SYSTEM DBUS_BUS_STARTER
  if (dbus_error_is_set(&error))
  {
    printf("dbus_bus_get fail, err:%s\n", error.message);
    dbus_error_free(&error);
  }

  if (conn == NULL)
  {
    printf("dbus_bus_get fail\n");
    return;
  }

  ret = dbus_bus_request_name(conn, "dbusname.callmethod", DBUS_NAME_FLAG_REPLACE_EXISTING, &error);
  if (dbus_error_is_set(&error))
  {
    printf("dbus_bus_request_name fail, err:%s\n", error.message);
    dbus_error_free(&error);
  }

  if (ret != DBUS_REQUEST_NAME_REPLY_PRIMARY_OWNER)
  {
    printf("dbus_bus_request_name fail\n");
    return;
  }

  msg = dbus_message_new_method_call("dbusname.exposemethod", // target for the method call
                                     "/test/method/Object", // object to call on
                                     "test.method.Type", // interface to call on
                                     "Method"); // method name
  if (msg == NULL)
  {
    printf("dbus_message_new_method_call fail\n");
    return;
  }

  dbus_message_iter_init_append(msg, &args);
  if (!dbus_message_iter_append_basic(&args, DBUS_TYPE_STRING, &sigvalue))
  {
    printf("dbus_message_iter_append_basic fail\n");
    return;
  }

  if (!dbus_connection_send_with_reply(conn, msg, &pending, 100000))
  {
    // -1 is default timeout
    printf("dbus_connection_send_with_reply fail\n");
    return;
  }

  if (NULL == pending)
  {
    printf("dbus_connection_send_with_reply fail\n");
    return;
  }

  printf("call method send success\n");

  dbus_connection_flush(conn);
  dbus_message_unref(msg);

  int stat = 0;
  dbus_uint32_t level = 0;

  printf("call method wait replying1\n");
  dbus_pending_call_block(pending);

  printf("call method wait replying2\n");
  msg = dbus_pending_call_steal_reply(pending);
  if (msg == NULL)
  {
    printf("dbus_pending_call_steal_reply fail\n");
    return;
  }

  dbus_pending_call_unref(pending);

  if (!dbus_message_iter_init(msg, &args))
  {
      printf("message has no arguments!\n");
  }
  else if (DBUS_TYPE_BOOLEAN != dbus_message_iter_get_arg_type(&args))
  {
      printf("argument is not boolean!\n");
  }
  else
  {
      dbus_message_iter_get_basic(&args, &stat);
  }

  if (!dbus_message_iter_next(&args))
  {
      printf("message has too few arguments!\n");
  }
  else if (DBUS_TYPE_UINT32 != dbus_message_iter_get_arg_type(&args))
  {
      printf("argument is not int!\n");
  }
  else
  {
      dbus_message_iter_get_basic(&args, &level);
  }

  printf("get reply: %d, %d\n", stat, level);

  dbus_message_unref(msg);

  //dbus_connection_close(conn);
}

void do1()
{
  int ret;
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

    //sendsignal();
    callmethod();

    printf("son end, pid = %d, ppid = %d\n", getpid(), getppid());
    exit(0);
  }
  else
  {
    //recvsignal();
    exposemethod();

    /*dbus_bus_get
    dbus_message_new_signal
    dbus_message_iter_init_append
    dbus_connection_send
    dbus_message_unref

    dbus_bus_get
    dbus_bus_add_match
    dbus_connection_pop_message
    dbus_message_is_signal

    dbus_message_new_method_call  
    dbus_message_iter_init_append  
    dbus_connection_send_with_reply  
    dbus_pending_call_block  
    dbus_pending_call_steal_reply  
    dbus_message_iter_init  


    dbus_bus_add_match  
    dbus_connection_pop_message  
    dbus_message_is_method_call 
    dbus_message_iter_init  
    dbus_message_new_method_return  
    dbus_message_iter_init_append  
    dbus_connection_send  */

    int status = 0;
    ret = waitpid(pid, &status, 0);
    if (ret == -1)
    {
      printf("son ret = %d, status = %d\n", ret, status);
      return;
    }

    printf("son ret = %d, status = %d\n", ret, status);
  }
}

int main(int argc, char const *argv[])
{
  do1();

  return 0;
}