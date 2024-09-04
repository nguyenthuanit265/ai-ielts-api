package enums

type Channel string
type QueueName string
type ConsumerTag string

const (
	TelegramChannel Channel = "Telegram"
	EmailChannel    Channel = "Email"
	SmsChannel      Channel = "Sms"
)

const (
	TelegramQueue QueueName = "telegram_queue"
	EmailQueue    QueueName = "email_queue"
	SmsQueue      QueueName = "sms_queue"
)

const (
	ConsumerTelegramQueue ConsumerTag = "telegram_queue_tag"
	ConsumerEmailQueue    ConsumerTag = "email_queue_tag"
	ConsumerSmsQueue      ConsumerTag = "sms_queue_tag"
)
