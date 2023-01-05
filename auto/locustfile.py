import time
from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(1, 2.5)

    @task
    def request_cpu_eater(self):
        self.client.get("/cpu.eater")