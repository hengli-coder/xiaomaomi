# -*- coding: UTF-8 -*-
__author__ = 'bbjiang'

import md5
import urllib
import httplib
from kafka import KafkaConsumer

from thrift.transport import TSocket,TTransport
from thrift.protocol import TBinaryProtocol
from hbase import Hbase
import happybase

# thrift默认端口是9090
socket = TSocket.TSocket('172.16.154.6', 10299)
socket.setTimeout(50000)

transport = TTransport.TBufferedTransport(socket)
protocol = TBinaryProtocol.TBinaryProtocol(transport)

client = Hbase.Client(protocol)
socket.open()


src = 'mysun:meta:21169:1.0'
m1 = md5.new()
m1.update(src.encode(encoding='utf-8'))
m = (m1.hexdigest())
print m


row = "%s#%s" % (m, src)
print row
#print client.getColumnDescriptors("msp_osp30_bin")
print client.get('msp_osp30_bin', row, 'f:bin')

httplib.HTTPConnection("172.16.154.205")
httplib.HTTP


#consumer = KafkaConsumer('osp_publish', bootstrap_servers = ['172.16.154.205:9092', '172.16.154.251:9092', '172.16.154.206:9092'])

#for msg in consumer:
#    print msg