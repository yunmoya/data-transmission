import subprocess
import threading
import time
import traceback
import random

import numpy

import xls_utils
import sys

sys.path.append('mygrpc')
import gocosi_rpc_client as client


def send_msg(index):
    port = random.choice(ports)
    address = "10.1.0.133:" + str(port)
    client.sig_msg(address, f"hi{index}")


if __name__ == '__main__':
    from argparse import ArgumentParser

    parser = ArgumentParser()
    parser.add_argument('-n', '--num', default=4, type=int, help='nodes num')
    parser.add_argument('-r', '--round', default=1, type=int, help='round number')
    parser.add_argument('-p', '--pnum', default=4, type=int, help='parallel requests num')
    parser.add_argument('-f', '--fnum', default=0, type=int, help='failing nodes num')
    args = parser.parse_args()
    node_num = args.num
    round_num = args.round
    parallel_num = args.pnum
    failing_node_num = args.fnum
    if node_num < 3:
        print("the number of nodes is at least 3")
    if failing_node_num >= node_num:
        print("the number of failing nodes should be smaller than node num")

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
        time.sleep(60)

        for i in range(parallel_num):
            threading.Thread(target=send_msg, args=(i,)).start()

        for i in range(failing_node_num):
            index = random.randint(0, len(ports) - 1)
            ports.remove(ports[index])
            hosts.remove(hosts[index])
            pList[index].kill()
            pList.remove(pList[index])
        time.sleep(60)

        for host in hosts:
            response2 = client.get_info(host)
            latency_list.extend(response2.latency)
            round_time = response2.round_time
    except Exception:
        traceback.print_exc()
    finally:
        for p in pList:
            p.kill()
        print("all hosts are killed")
        if len(latency_list) > 0:
            latency_avg = numpy.mean(latency_list)
            latency_max = numpy.max(latency_list)
            latency_min = numpy.min(latency_list)
            print(f'latency_avg: {latency_avg}')
            print(f'latency_max: {latency_max}')
            print(f'latency_min: {latency_min}')
            data = node_num, failing_node_num, round_num, round_time, \
                   parallel_num, latency_avg, latency_max, latency_min
            xls_utils.wirte_to_xls_p(data)
        else:
            print("latency list is empty")