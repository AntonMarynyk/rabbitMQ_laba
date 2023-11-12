package utils

func Producer() {
	validateProducerInput()

	conn, ch := init_connection_channel()
	defer conn.Close()
	defer ch.Close()

	ctx, cancel := declareContext()
	defer cancel()

	publishMessage(ctx, ch)
}
