# -*- coding: utf-8 -*-
import requests

"""
JSON RPC client using Python
"""
def rpc_call():
    url = 'http://localhost:1234/rpc'
    payload = {
	    'id': 1,
	    'method': 'Arith.Multiply',
	    'params': [{'A': 2, 'B': 3}]
	}
    r = requests.post(url, json=payload)
    print r.text


if __name__ == '__main__':
    rpc_call()