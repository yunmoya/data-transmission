import subprocess
import time
import traceback
import random

import numpy

import xls_utils
import sys

sys.path.append('mygrpc')
import gocosi_rpc_client as client

if __name__ == '__main__':
    from argparse import ArgumentParser

    parser = ArgumentParser()
    parser.add_argument('-n', '--num', default=4, type=int, help='nodes num')
    parser.add_argument('-r', '--round', default=1, type=int, help='round number')
    args = parser.parse_args()
    node_num = args.num
    round_num = args.round
    failing_node_num = args.fnum
    if node_num < 4:
        print("the number of nodes is at least 4")
    ports = []
    hosts = []
    for i in range(node_num):
        port = 30000 + i
        ports.append(port)
        hosts.append("localhost:" + str(port))
    
    pList = []
    latency_list = []
    round_time = 0
    stage1_time = 0
    try:
        for port in ports:
            p = subprocess.Popen(["python", "gocosi.py", "-p", str(port), "-s", str(node_num)])
            pList.append(p)
        time.sleep(3600)

    except Exception:
        traceback.print_exc()
    finally:
        for p in pList:
            p.kill()
        print("all hosts are killed")
