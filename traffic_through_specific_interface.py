import socket

import requests
from requests.adapters import HTTPAdapter
from urllib3.poolmanager import PoolManager


class BindToInterfaceAdapter(HTTPAdapter):
    def __init__(self, interface_name, *args, **kwargs):
        self.interface_name = interface_name
        super().__init__(*args, **kwargs)

    def init_poolmanager(self, *args, **kwargs):
        # Подготавливаем null-terminated строку с именем интерфейса
        iface = (self.interface_name + "\0").encode("utf-8")
        kwargs["socket_options"] = [(socket.SOL_SOCKET, socket.SO_BINDTODEVICE, iface)]
        self.poolmanager = PoolManager(*args, **kwargs)


# Creating session and mounting adapter
session = requests.Session()
adapter = BindToInterfaceAdapter("wlan0")  # Proving correct interface name like wlan0
session.mount("http://", adapter)
session.mount("https://", adapter)

# Checking
response = session.get("https://ifconfig.me")
print(response.text.strip())
