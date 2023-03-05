#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2021/6/2 4:08 下午
# @Author  : CrissChan
# @Site    : https://blog.csdn.net/crisschan
# @File    : load_test_grpc.py
# @Software: 这是调用gRPC的Locust脚本
import random
import sys
import grpc
import inspect
import time
import gevent
from locust.contrib.fasthttp import FastHttpUser
from locust import task, events, constant
from locust.runners import STATE_STOPPING, STATE_STOPPED, STATE_CLEANUP, WorkerRunner
import gocosi_rpc_client as client


def stopwatch(func):
    def wrapper(*args, **kwargs):

        previous_frame = inspect.currentframe().f_back
        _, _, task_name, _, _ = inspect.getframeinfo(previous_frame)
        start = time.time()
        result = None
        try:
            result = func(*args, **kwargs)
        except Exception as e:
            total = int((time.time() - start) * 1000)
            events.request_failure.fire(request_type="TYPE",
                                        name=task_name,
                                        response_time=total,
                                        response_length=0,
                                        exception=e)
        else:
            total = int((time.time() - start) * 1000)
            events.request_success.fire(request_type="TYPE",
                                        name=task_name,
                                        response_time=total,
                                        response_length=0)
        return result

    return wrapper


def get_msg():
    index = random.randint(0, 10)
    return f"hi {index}"


def get_address():
    index = random.randint(30000, 30003)
    return f"localhost:{index}"


class GRPCMyLocust(FastHttpUser):
    host = 'http://127.0.0.1:50000'
    wait_time = constant(0)

    def on_start(self):
        pass

    def on_stop(self):
        pass

    @task
    @stopwatch
    def grpc_new_msg(self):
        start_time = time.time()
        address = get_address()
        msg = get_msg()
        client.sig_msg(address, msg)
        end_time = time.time()
        delay = end_time - start_time
        if delay < 1:
            time.sleep(1-delay)

    # @task
    # @stopwatch
    # def grpc_info(self):
    #     address = get_address()
    #     resp = client.get_info(address)
    #     print(resp.latency)
