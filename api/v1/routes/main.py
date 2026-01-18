import logging
import os
import sys
from analytics_worker import config, worker

def main():
    # Initialize the logger
    logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(name)s - %(levelname)s - %(message)s')
    logger = logging.getLogger(__name__)

    try:
        # Load the configuration
        configuration = config.load_configuration()

        # Create a worker instance
        analytics_worker = worker.AnalyticsWorker(configuration)

        # Start the worker
        analytics_worker.start()
    except Exception as e:
        logger.error("An error occurred: %s", e)
        sys.exit(1)

if __name__ == "__main__":
    main()