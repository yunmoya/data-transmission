import numpy as np
import openpyxl
import pandas
import matplotlib.pyplot as plt
from pandas import DataFrame

plt.rcParams['font.family'] = 'Times New Roman'
plt.rcParams['font.size'] = '10.5'


def parse_top_txt(filename):
    """
    解析top指令生成的txt top指令为：top -b -d 1 -n 200 | grep python  > psr101000.txt
    :param filename: top文件名
    :return: 平均CPU占用率和内存占用率（已去除simulation文件的影响）
    """
    with open(filename, encoding="utf-8") as fo:
        print(f"read {filename}")
        datas = dict()
        while True:
            line = fo.readline()
            if not line:
                break
            line = ' '.join(line.split())
            items = line.split(" ")
            if len(items) < 9:
                continue
            if items[0] in datas:
                datas[items[0]].get("CPU").append(float(items[8]))
                datas[items[0]].get("MEM").append(float(items[9]))
            else:
                per = {
                    "CPU": [float(items[8])],
                    "MEM": [float(items[9])]
                }
                datas[items[0]] = per

        min_pid_cpu = 200
        min_pid = None
        for pid in datas:
            data = datas[pid]
            cpu_list = data['CPU']
            mem_list = data['MEM']
            avg_cpu = np.mean(cpu_list)
            min_cpu = np.min(cpu_list)
            max_cpu = np.max(cpu_list)
            print(f"pid: {pid}")
            # print(cpu_list)
            print(f"cpu performance: {(avg_cpu, min_cpu, max_cpu)}")
            avg_mem = np.mean(mem_list)
            min_mem = np.min(mem_list)
            max_mem = np.max(mem_list)
            print(f"memory performance: {(avg_mem, min_mem, max_mem)}")
            datas[pid] = {
                "CPU": avg_cpu,
                "MEM": avg_mem
            }
            if avg_cpu < min_pid_cpu:
                min_pid_cpu = avg_cpu
                min_pid = pid
        print(f"simulation pid: {min_pid}")
        avg_cpu_list = []
        avg_mem_list = []
        for pid in datas:
            # 去除simulation进程的影响
            if pid == min_pid:
                pass
            else:
                avg_cpu_list.append(datas[pid].get('CPU'))
                avg_mem_list.append(datas[pid].get('MEM'))
        return np.mean(avg_cpu_list), np.mean(avg_mem_list)


def round_time_cpu(round_times):
    """
    top数据分析，输出csv文件返回不同节点不同t的CPU使用率
    :return: CPU使用率 例如：{3: [6.058799171842651, 5.709756097560976, 5.64792899408284, 5.50078895463511, 3.355945419103314, 3.233911368015414, 4.468131868131868], 10: [14.914911242603551, 14.509771428571431, 12.542443181818184, 12.059766081871347, 3.9168825991628458, 3.232994652406417, 3.1752659574468085], 20: [7.330449999999999, 7.885925, 18.411315789473687, 15.367922077922078, 4.135336787564766, 3.386928934010153, 3.13215]}
    """
    cpu_dict = dict()
    mem_dict = dict()
    for nn in [3, 10, 20, 30]:
        cpu_list = []
        mem_list = []
        for round_time in round_times:
            file_name = f"psr{nn}{round_time}.txt"
            avg_cpu, avg_mem = parse_top_txt(file_name)
            cpu_list.append(avg_cpu)
            mem_list.append(avg_mem)
        cpu_dict[nn] = cpu_list
        mem_dict[nn] = mem_list
    cpu_df = pandas.DataFrame(cpu_dict, index=round_times)
    mem_df = pandas.DataFrame(mem_dict, index=round_times)
    print(cpu_df)
    print(mem_df)
    cpu_df.to_csv("cpu_round_time.csv")
    return cpu_dict


def round_time_fig():
    round_times = [50, 100, 150, 200, 250, 300, 350, 400, 450, 500]
    # round_times = [50, 100, 150]
    cpu_dict = round_time_cpu(round_times)
    color = {3: 'r', 10: 'b', 20: 'g', 30: 'c'}
    # arange返回一个数据，range返回一个list
    # arange 用于创建等差数组

    fig = plt.figure()
    y1 = round_times
    # ax1显示y1  ,ax2显示y2
    ax1 = fig.subplots()
    for nn in cpu_dict:
        color_nn = color[nn]
        y1 = cpu_dict[nn]
        if nn == 3:
            nn = 4
        ax1.plot(round_times, y1, f'{color_nn}-', marker='.', label=f'N={nn}')

    ax1.set_xlabel('Gossip Period(ms)')
    ax1.set_ylabel('CPU Usage(%)')

    plt.xticks(round_times)
    plt.legend()
    plt.grid()
    plt.show()


if __name__ == '__main__':
    round_time_fig()
    # print(parse_latency_xls("gocosi.xlsx", "latency_round_time"))
    # print(round_time_cpu([5, 10, 50, 100, 500, 800, 1000]))
