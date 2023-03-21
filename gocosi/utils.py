import socket

import toml


def get_host_ip():
    s = None
    _host = ''
    try:
        s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        s.connect(("8.8.8.8", 80))
        _host = s.getsockname()[0]
    finally:
        s.close()
    return _host


def read_from_toml(file: str):
    with open(file, 'r') as f:
        info = toml.load(f)
        print(f"read from {file} successfully")
    return info


def write_to_toml(info: dict, file: str):
    with open(file, 'w') as f:
        toml.dump(info, f)
        print(f"write to {file} successfully")