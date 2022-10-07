package zapfield

import "go.uber.org/zap"

func WorkerTask(value string) zap.Field {
	return zap.String("worker_task", value)
}

func WorkerConsumer(value string) zap.Field {
	return zap.String("worker_consumer", value)
}

func WorkerQueue(value string) zap.Field {
	return zap.String("worker_queue", value)
}

func WorkerTopic(value string) zap.Field {
	return zap.String("worker_topic", value)
}
