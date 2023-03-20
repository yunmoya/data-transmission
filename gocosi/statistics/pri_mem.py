import numpy as np
import openpyxl
import pandas
import matplotlib.pyplot as plt
from pandas import DataFrame

plt.rcParams['font.family'] = 'Times New Roman'
plt.rcParams['font.size'] = '10.5'


def round_time_fig():
    func_list = ['ReadEC', 'EndorsementInc', 'StatusAssetExists', 'SetStatus', 'SelectCP', 'UpdatePropotion']
    # arange返回一个数据，range返回一个list
    # arange 用于创建等差数组

    fig = plt.figure()
    # ax1显示y1  ,ax2显示y2
    # y = [7.42, 24.1, 24.81, 24.3, 17.55, 15]
    y = [21.8, 21.15, 21.5, 21.45, 18.7, 18.15]
    for i in range(len(func_list)):
        plt.bar(func_list[i], y[i], color='#1f77b4', width=0.5)
        # plt.text(func_list[i], y[i], "%.f"%y[i], ha='center')

    # plt.xlabel('数据类型', fontdict=font1)
    plt.xticks(rotation=30)
    plt.ylabel('Memory Usage(MB)')
    plt.show()


if __name__ == '__main__':
    round_time_fig()
    # print(parse_latency_xls("gocosi.xlsx", "latency_round_time"))
    # print(round_time_cpu([5, 10, 50, 100, 500, 800, 1000]))
