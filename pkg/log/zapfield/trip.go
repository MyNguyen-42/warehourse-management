package zapfield

import (
	"go.uber.org/zap"
	"strings"
)

// TripCode return trip_code field
func TripCode(value string) zap.Field {
	return zap.String("trip_code", value)
}

// TripItem return trip_item field
func TripItem(value string) zap.Field {
	return zap.String("trip_item", value)
}

// DriverId return driver_id field
func DriverId(value string) zap.Field {
	return zap.String("driver_id", value)
}

// HubId return hub_id field
func HubId(value string) zap.Field {
	return zap.String("hub_id", value)
}

func CombinedCode(value string) zap.Field {
	return zap.String("combined_code", value)
}

// OrderCodes return order_codes field
func OrderCodes(value []string) zap.Field {
	return zap.String("order_codes", strings.Join(value, ","))
}

// OrderCode return order_code field
func OrderCode(value string) zap.Field {
	return zap.String("order_code", value)
}

func UserId(value string) zap.Field {
	return zap.String("user_id", value)
}
