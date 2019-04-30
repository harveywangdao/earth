#include <stdio.h>
#include <stdlib.h>
#include <sys/select.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <string.h>

char *host_name ="127.0.0.1";
int port = 8000;

int main(int argc, char** argv)
{
    char buf[8192];
    char message[256];
    int socket_descriptor;
    struct sockaddr_in pin;

    /*
    * hostent记录主机的信息，包括主机名、别名、地址类型、地址长度和地址列表
    * struct hostent {
    　　       char *h_name;地址的正式名称
    　　       char **h_aliases;空字节-地址的预备名称的指针
    　　       int h_addrtype;地址类型; 通常是AF_INET。
    　       　int h_length;地址的比特长度。
    　       　char **h_addr_list;零字节-主机网络地址指针。网络字节顺序。
    　　    };
    　　   #define h_addr h_addr_list[0] //h_addr_list中的第一地址
    */
    struct hostent *server_host_name;

    char *str = "A default test string";

    if (argc < 2)//运行程序时送给main函数到命令行参数个数
    {
        printf("Usage:test \"Any test string\"\n");
        printf("we will send a default test string. \n");
    }
    else
    {
        str =argv[1];
    }

    if ((server_host_name = gethostbyname(host_name)) == 0)
    {
        perror("Error resolving local host \n");
        exit(1);
    }

    memset(&pin, 0, sizeof(pin));
    pin.sin_family = AF_INET;
    pin.sin_addr.s_addr = htonl(INADDR_ANY);
    pin.sin_addr.s_addr = ((struct in_addr *)(server_host_name->h_addr))->s_addr;

    if (inet_pton(AF_INET, host_name, &pin.sin_addr) != 1)
    {
        perror("inet_pton\n");
        exit(1);
    }

    //printf("pin.sin_addr.s_addr = %d.\n", pin.sin_addr.s_addr);

    pin.sin_port = htons(port);

    if ((socket_descriptor = socket(AF_INET, SOCK_STREAM, 0)) == -1)
    {
        perror("Error opening socket \n");
        exit(1);
    }

    if (connect(socket_descriptor, (void *)&pin, sizeof(pin))==-1)
    {
        perror("Error connecting to socket \n"); ////
        exit(1);
    }

    printf("Sending message:[%s]\n", str);

    if (send(socket_descriptor, str, strlen(str),0) == -1)
    {
        perror("Error in send\n");
        exit(1);
    }

    if (recv(socket_descriptor, buf, 8192, 0) == -1)
    {
        perror("Error in receiving response from server \n");
        exit(1);
    }

    printf("Response from server:[%s]\n",buf);

    close(socket_descriptor);

    return (EXIT_SUCCESS);
}