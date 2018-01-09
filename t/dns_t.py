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
    f=open(sys.argv[1])
    l = f.read().splitlines()
    f.close()
    t = []
    num =  sys.argv[2]
    for j in xrange(int(num)):
        print "* " * 30
        print "[%s]" % str(j)
        for i in l:
            t.append(threading.Thread(target=get_dns, args=(i,)))
        for tx in t:
            tx.start()
        print "* " * 30
