import pika
import random
import time
def on_message_recieved(ch, method, properties, body):
    processing_time=random.randint(1,6)
    print(f"received: {body} will take {processing_time} to process") 
    time.sleep(processing_time)
    ch.basic_ack(delivery_tag=method.delivery_tag)
    print("Finished processing the message")
conn_para=pika.ConnectionParameters('localhost')
conn=pika.BlockingConnection(conn_para)
channel=conn.channel()
channel.queue_declare(queue="letterbox")
channel.basic_qos(prefetch_count=1)
channel.basic_consume(queue='letterbox', on_message_callback=on_message_recieved)
print("Starting Consuming")
channel.start_consuming()