import socket
import sys, threading

def get_dns(host):
    try:
        print  socket.gethostbyname(host)
    except socket.gaierror:
        print  "%s socket.gaierror" % str(i)

def list_hosts(list):
    for i in list:
        get_dns(i)
if __name__=="__main__":
    get_dns("www.baidu.com")
