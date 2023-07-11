import matplotlib.pyplot as plt

plt.rcParams['font.family'] = 'Times New Roman'
plt.rcParams['font.size'] = '12'
plt.rcParams['figure.figsize'] = (8.0, 6.0)  # 单位是inches

def node_time_fig():
    node_nums = [4, 10, 20, 30, 40, 50]

    color = {'gocosi': 'r', 'blscosi': 'b'}
    # arange返回一个数据，range返回一个list
    # arange 用于创建等差数组

    fig = plt.figure()
    # ax1显示y1  ,ax2显示y2
    ax1 = fig.subplots()
    gocosi_normal = [0.645321655, 1.043728733, 1.162276497, 1.195117044, 1.246071625, 1.36917716]
    gocosi_failing = [0.702014, 1.933499, 2.799725, 2.975602, 3.337802, 3.389581]
    blscosi_normal = [0.011547, 0.020596, 0.040082, 0.05019, 0.062895, 0.067557]
    blscosi_failing = [10.009992, 10.01263, 8.349161, 5.571029667, 5.02049275, 6.0231808]

    ax1.plot(node_nums, gocosi_normal, 'r-', marker='.', label='GoCosi')
    ax1.plot(node_nums, gocosi_failing, 'r-', marker='x', label='GoCosi with 1/3 failures')
    ax1.plot(node_nums, blscosi_normal, 'b--', marker='.', label='BLSCosi')
    ax1.plot(node_nums, blscosi_failing, 'b--', marker='x', label='BLSCosi with 1/3 failures')

    ax1.set_xlabel('Number of Nodes')
    ax1.set_ylabel('Latency(s)')

    plt.xticks(node_nums)
    plt.legend()
    plt.grid()
    plt.savefig('./output/gocosiCp.eps', dpi=300)  # eps文件，用于LaTeX
    plt.show()


if __name__ == '__main__':
    node_time_fig()