import pika
import time
import random

# Connection parameters
conn_para = pika.ConnectionParameters('localhost')
conn = pika.BlockingConnection(conn_para)
channel = conn.channel()

# Declare the queue
channel.queue_declare(queue='letterbox')

# Initialize message ID
messageId = 1

print("Starting to send messages...")
try:
    while True:
        # Create the message
        mess = f"sending messageId: {messageId}"

        # Publish the message
        channel.basic_publish(exchange='', routing_key='letterbox', body=mess)
        print(f"Sent message: {mess}")

        # Increment the message ID
        messageId += 1

        # Random delay between 1 and 4 seconds
        time.sleep(random.randint(1, 4))

except KeyboardInterrupt:
    print("\nStopping message production...")

finally:
    # Close the connection
    conn.close()
    print("Connection closed.")
