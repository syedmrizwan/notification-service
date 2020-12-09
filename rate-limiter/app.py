"""import libraries"""
import asyncio
import json
import logging.config
import os
import uuid
import redis
import time
from nats.aio.client import Client as NATS
from stan.aio.client import Client as STAN
from utils import constants

# Constants
NATS_URL = os.environ['NATS_URL']
NATS_CLUSTER_ID = os.environ['NATS_CLUSTER_ID']
REDIS_HOST = os.environ['REDIS_HOST']
REDIS_PORT = os.environ['REDIS_PORT']
RATE_LIMITER_CHANNEL_NAME = os.environ['RATE_LIMITER_CHANNEL_NAME']
NOTIFICATION_HANDLER_CHANNEL_NAME = os.environ['NOTIFICATION_HANDLER_CHANNEL_NAME']

logging.config.fileConfig(constants.LOGGING_FILE_NAME, disable_existing_loggers=False)
logger = logging.getLogger(__name__)
logger.info("Rate Limiter has been started.")


async def get_nats_connection():
    """get_nats_connection () -> nats_connection"""
    nats_conn = NATS()
    await nats_conn.connect(NATS_URL)
    return nats_conn


async def handle_requests(nats_conn, subject_name):
    """handle_request ()"""
    sc = STAN()
    await sc.connect(NATS_CLUSTER_ID, str(uuid.uuid1()), nats=nats_conn)
    logger.info("Listening for requests on %s subject...", subject_name)

    async def cb(msg):
        """cb (msg)"""
        notification_data = json.loads(msg.data.decode())
        try:
            logger.info("Message received %s ...", notification_data)
            if throttle(notification_data['notification_handler']['name'],
                        notification_data['notification_handler']['rate_per_minute']):
                # fetch data from user service
                notification_data['phone'] = '0321-8550442'
                notification_data['email'] = 'syed.m.rizwan@outlook.com'
                notification_data['device_id'] = 'A4X-2dx'
                await sc.publish(
                    NOTIFICATION_HANDLER_CHANNEL_NAME + "." + notification_data['notification_handler']['name'],
                    json.dumps(notification_data).encode('utf-8'), ack_wait=600)
            await sc.ack(msg)
        except Exception as e:
            logger.exception(e)

    await sc.subscribe(subject_name, manual_acks=True, start_at="last_received", cb=cb, ack_wait=600)


def throttle(handler, rate_per_min):
    logger.info("Message received for throttle %s ... %s", handler, rate_per_min)

    # set the points for this route to 4
    ROUTE_SCORE = 4

    r = redis.Redis(host=REDIS_HOST, port=REDIS_PORT, db=0)

    epoch_ms = int(time.time() * 1000)
    pipe = r.pipeline()

    pipe.zremrangebyscore("%s:per_minute" % handler, 0, epoch_ms - 6000)
    pipe.zadd("%s:per_minute" % handler, {"%d:%d" % (epoch_ms, ROUTE_SCORE): epoch_ms})
    pipe.zrange("%s:per_minute" % handler, 0, -1)
    pipe.expire("%s:per_minute" % handler, 6001)

    res = pipe.execute()

    minute_score = sum(int(i.decode("utf-8").split(':')[-1]) for i in res[2])

    if minute_score > rate_per_min:
        # "DATA Exceeded", status=429
        return False
    return True


if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    nats_connection = loop.run_until_complete(get_nats_connection())
    loop.run_until_complete(handle_requests(nats_connection, RATE_LIMITER_CHANNEL_NAME + ".High"))
    loop.run_until_complete(handle_requests(nats_connection, RATE_LIMITER_CHANNEL_NAME + ".Medium"))
    loop.run_until_complete(handle_requests(nats_connection, RATE_LIMITER_CHANNEL_NAME + ".Low"))
    loop.run_forever()
    loop.close()
