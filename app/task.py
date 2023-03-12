from .celery import app
import time
from celery import group, chain as celery_chain, shared_task

@app.task()
def add(x, y):
    print('proces')
    time.sleep(10)
    return x + y


@app.task()
def mul(x, y):
    print("mul")
    return x * y

@app.task()
def deliver(x):
    print('reciever')
    sig = group([add.s(i, i+1).set(queue='test') | mul.s(2).set(queue='test') for i in range(0, x)])
    t = sig.apply_async()
    return t

@app.task()
def xsum(x):
    sig = group([add.s(i, i+1) | mul.s(2) | deliver.s().set(queue='test') for i in range(0, x)])
    t = sig.apply_async()
    return t

    # return sum(numbers)
