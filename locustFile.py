from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(1, 5)
    @task
    def CreateGiftCodeService(self):
        self.client.post("CreateGiftCode",{
            "gift_des": "dec",
            "available_times": 4,
            "valid_period": 100000,
            "creator":"tom-creator",
            "user": "tom",
            "code_type":-2,
            "ContentList":
                {
                    "gold_coins":888,
                    "diamonds":999,
                    "props":777,
                    "heroes":666
                }
        })

    @task
    def GetGiftCodeInfo(self):
        self.client.get("GetGiftCodeInfo?code=G1RWHLQA")

    @task
    def VerifyGiftCode(self):
        self.client.get("VerifyGiftCode?code=72JJ3KL1")