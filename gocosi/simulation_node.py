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
    parser.add_argument('-f', '--fnum', default=0, type=int, help='failing nodes num')
    args = parser.parse_args()
    node_num = args.num
    round_num = args.round
    failing_node_num = args.fnum
    if node_num < 4:
        print("the number of nodes is at least 4")
    if failing_node_num >= node_num:
        print("the number of failing nodes should be smaller than node num")

    for node_num in [3, 10, 20, 30, 40, 50]:
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
        neighbours = 0
        try:
            for port in ports:
                p = subprocess.Popen(["python", "gocosi.py", "-p", str(port), "-s", str(node_num)])
                pList.append(p)
            time.sleep(60)

            for i in range(failing_node_num):
                index = len(ports) - 1
                ports.remove(ports[index])
                host = hosts[index]
                hosts.remove(hosts[index])
                pList[index].kill()
                pList.remove(pList[index])
                print(f"remove node {host}")
            time.sleep(10)

            for i in range(round_num):
                port = random.choice(ports)
                address = "localhost:" + str(port)
                client.sig_msg(address, f"hi{i}")
                time.sleep(10)


            for host in hosts:
                response2 = client.get_info(host)
                latency_list.extend(response2.latency)
                round_time = response2.round_time
                neighbours = response2.neighbours
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
                data = node_num, failing_node_num, neighbours, round_num, round_time, \
                       latency_avg, latency_max, latency_min
                xls_utils.wirte_to_xls_neighbours(data)
            else:
                print("latency list is empty")
            time.sleep(10)
            print("====================================end====================================")
