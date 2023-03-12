from celery import Celery

app = Celery('app',
             broker='amqp://localhost:5672',
             backend='redis://localhost:6379/0',
             include=['app.task'])

# Optional configuration, see the application user guide.
app.conf.update(
    result_expires=3600,
)
app.conf.task_protocol = 1
if __name__ == '__main__':
    app.start()
