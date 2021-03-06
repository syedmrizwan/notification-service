"""import libraries"""
import asyncio
import json
import logging.config
import os
import uuid
from nats.aio.client import Client as NATS
from stan.aio.client import Client as STAN
from utils import constants

# Constants
NATS_URL = os.environ['NATS_URL']
NATS_CLUSTER_ID = os.environ['NATS_CLUSTER_ID']
NOTIFICATION_HANDLER_CHANNEL_NAME = os.environ['NOTIFICATION_HANDLER_CHANNEL_NAME']

logging.config.fileConfig(constants.LOGGING_FILE_NAME, disable_existing_loggers=False)
logger = logging.getLogger(__name__)
logger.info("Notification Handler has been started.")


async def get_nats_connection():
    """get_nats_connection () -> nats_connection"""
    nats_conn = NATS()
    await nats_conn.connect(NATS_URL)
    return nats_conn


async def handle_requests(nats_conn, subject_name):
    """handle_request ()"""
    sc = STAN()
    await sc.connect(NATS_CLUSTER_ID, str(uuid.uuid1()), nats=nats_conn)
    logger.info("Listening for requests on %s", subject_name)

    async def cb(msg):
        """cb (msg)"""
        notification_data = json.loads(msg.data.decode())
        try:
            logger.info("Notification Handler received %s", notification_data)
            # Handler internal logic by calling SMS or Email Vendor
            await sc.ack(msg)
        except Exception as e:
            logger.exception(e)

    await sc.subscribe(subject_name, manual_acks=True, start_at="last_received", cb=cb, ack_wait=600)


if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    nats_connection = loop.run_until_complete(get_nats_connection())
    loop.run_until_complete(handle_requests(nats_connection, NOTIFICATION_HANDLER_CHANNEL_NAME))
    loop.run_forever()
    loop.close()
